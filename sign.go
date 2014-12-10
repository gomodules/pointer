package aws

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strings"
	"time"
)

func sign(service, region string, creds Credentials, req *http.Request) {
	req.Header.Set("host", req.Host) // host header must be included as a signed header
	payloadHash := payloadHash(req)
	req.Header.Set("x-amz-content-sha256", payloadHash)

	t := requestTime(req)
	creq := canonicalRequest(req, payloadHash)
	sts := stringToSign(t, creq, region, service)
	signature := signature(t, sts, region, service, creds.SecretAccessKey())
	auth := authorization(req.Header, t, region, service, creds.AccessKeyID(), signature)
	req.Header.Set("Authorization", auth)
	if s := creds.SecurityToken(); s != "" {
		req.Header.Set("X-Amz-Security-Token", s)
	}
	return
}

const (
	iso8601BasicFormat      = "20060102T150405Z"
	iso8601BasicFormatShort = "20060102"
)

func requestTime(req *http.Request) time.Time {
	// Get "x-amz-date" header
	date := req.Header.Get("x-amz-date")

	// Attempt to parse as ISO8601BasicFormat
	t, err := time.Parse(iso8601BasicFormat, date)
	if err == nil {
		return t
	}

	// Attempt to parse as http.TimeFormat
	t, err = time.Parse(http.TimeFormat, date)
	if err == nil {
		req.Header.Set("x-amz-date", t.Format(iso8601BasicFormat))
		return t
	}

	// Get "date" header
	date = req.Header.Get("date")

	// Attempt to parse as http.TimeFormat
	t, err = time.Parse(http.TimeFormat, date)
	if err == nil {
		return t
	}

	// Create a current time header to be used
	t = time.Now().UTC()
	req.Header.Set("x-amz-date", t.Format(iso8601BasicFormat))
	return t
}

func canonicalRequest(req *http.Request, pHash string) string {
	if pHash == "" {
		pHash = payloadHash(req)
	}
	c := new(bytes.Buffer)
	fmt.Fprintf(c, "%s\n", req.Method)
	fmt.Fprintf(c, "%s\n", canonicalURI(req.URL))
	fmt.Fprintf(c, "%s\n", canonicalQueryString(req.URL))
	fmt.Fprintf(c, "%s\n\n", canonicalHeaders(req.Header))
	fmt.Fprintf(c, "%s\n", signedHeaders(req.Header))
	fmt.Fprintf(c, "%s", pHash)
	return c.String()
}

func canonicalURI(u *url.URL) string {
	u = &url.URL{Path: u.Path}
	canonicalPath := u.String()
	slash := strings.HasSuffix(canonicalPath, "/")
	canonicalPath = path.Clean(canonicalPath)
	if canonicalPath != "/" && slash {
		canonicalPath += "/"
	}

	return canonicalPath
}

func canonicalQueryString(u *url.URL) string {
	var a []string
	for k, vs := range u.Query() {
		k = url.QueryEscape(k)
		for _, v := range vs {
			if v == "" {
				a = append(a, k+"=")
			} else {
				v = url.QueryEscape(v)
				a = append(a, k+"="+v)
			}
		}
	}
	sort.Strings(a)
	return strings.Join(a, "&")
}

func canonicalHeaders(h http.Header) string {
	i, a := 0, make([]string, len(h))
	for k, v := range h {
		for j, w := range v {
			v[j] = strings.Trim(w, " ")
		}
		sort.Strings(v)
		a[i] = strings.ToLower(k) + ":" + strings.Join(v, ",")
		i++
	}
	sort.Strings(a)
	return strings.Join(a, "\n")
}

func signedHeaders(h http.Header) string {
	i, a := 0, make([]string, len(h))
	for k, _ := range h {
		a[i] = strings.ToLower(k)
		i++
	}
	sort.Strings(a)
	return strings.Join(a, ";")
}

func payloadHash(req *http.Request) string {
	var b []byte
	if req.Body == nil {
		b = []byte("")
	} else {
		var err error
		b, err = ioutil.ReadAll(req.Body)
		if err != nil {
			// TODO: I REALLY DON'T LIKE THIS PANIC!!!!
			panic(err)
		}
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	return hash(string(b))
}

func stringToSign(t time.Time, creq, region, service string) string {
	w := new(bytes.Buffer)
	fmt.Fprint(w, "AWS4-HMAC-SHA256\n")
	fmt.Fprintf(w, "%s\n", t.Format(iso8601BasicFormat))
	fmt.Fprintf(w, "%s\n", credentialScope(t, region, service))
	fmt.Fprintf(w, "%s", hash(creq))
	return w.String()
}

func credentialScope(t time.Time, region, service string) string {
	return fmt.Sprintf("%s/%s/%s/aws4_request", t.Format(iso8601BasicFormatShort), region, service)
}

func signature(t time.Time, sts, region, service, secret string) string {
	h := mac(derivedKey(t, region, service, secret), []byte(sts))
	return fmt.Sprintf("%x", h)
}

func derivedKey(t time.Time, region, service, secret string) []byte {
	h := mac([]byte("AWS4"+secret), []byte(t.Format(iso8601BasicFormatShort)))
	h = mac(h, []byte(region))
	h = mac(h, []byte(service))
	h = mac(h, []byte("aws4_request"))
	return h
}

func authorization(header http.Header, t time.Time, region, service, key, signature string) string {
	w := new(bytes.Buffer)
	fmt.Fprint(w, "AWS4-HMAC-SHA256 ")
	fmt.Fprintf(w, "Credential=%s/%s, ", key, credentialScope(t, region, service))
	fmt.Fprintf(w, "SignedHeaders=%s, ", signedHeaders(header))
	fmt.Fprintf(w, "Signature=%s", signature)
	return w.String()
}

func hash(in string) string {
	h := sha256.New()
	fmt.Fprintf(h, "%s", in)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func mac(key, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}
