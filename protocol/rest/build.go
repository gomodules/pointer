package rest

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

func Build(r *aws.Request) {
	if r.ParamsFilled() {
		v := reflect.ValueOf(r.Params).Elem()
		buildLocationElements(r, v)
		buildBody(r, v)
	}
}

func buildLocationElements(r *aws.Request, v reflect.Value) {
	query := r.HTTPRequest.URL.Query()

	for i := 0; i < v.NumField(); i++ {
		m := v.Field(i)
		if n := v.Type().Field(i).Name; n[0:1] == strings.ToLower(n[0:1]) {
			continue
		}

		if m.IsValid() {
			field := v.Type().Field(i)
			name := field.Tag.Get("locationName")
			if name == "" {
				name = field.Name
			}
			if m.Kind() == reflect.Ptr {
				m = m.Elem()
			}

			switch field.Tag.Get("location") {
			case "header":
				buildHeader(r, m, name)
			case "uri":
				buildURI(r, m, name)
			case "querystring":
				buildQueryString(r, m, name, query)
			}
		}
		if r.Error != nil {
			return
		}
	}

	r.HTTPRequest.URL.RawQuery = query.Encode()
	updatePath(r.HTTPRequest.URL, r.HTTPRequest.URL.Path)
}

func buildBody(r *aws.Request, v reflect.Value) {
	if field, ok := v.Type().FieldByName("SDKShapeTraits"); ok {
		if payloadName := field.Tag.Get("payload"); payloadName != "" {
			payload := v.FieldByName(payloadName)
			switch reader := payload.Interface().(type) {
			case io.ReadSeeker:
				r.SetReaderBody(reader)
			case []byte:
				r.SetBufferBody(reader)
			case string:
				r.SetBufferBody([]byte(reader))
			default:
				r.Error = fmt.Errorf("Unknown payload type %s", payload.Type())
			}
		}
	}
}

func buildHeader(r *aws.Request, v reflect.Value, name string) {
	str, err := convertType(v)
	if err != nil {
		r.Error = err
	} else if str != nil {
		r.HTTPRequest.Header.Add(name, *str)
	}
}

func buildURI(r *aws.Request, v reflect.Value, name string) {
	value, err := convertType(v)
	if err != nil {
		r.Error = err
	} else if value != nil {
		uri := r.HTTPRequest.URL.Path
		uri = strings.Replace(uri, "{"+name+"}", escapePath(*value, true), -1)
		uri = strings.Replace(uri, "{"+name+"+}", escapePath(*value, false), -1)
		r.HTTPRequest.URL.Path = uri
	}
}

func buildQueryString(r *aws.Request, v reflect.Value, name string, query url.Values) {
	str, err := convertType(v)
	if err != nil {
		r.Error = err
	} else if str != nil {
		query.Set(name, *str)
	}
}

func updatePath(url *url.URL, urlPath string) {
	scheme, query := url.Scheme, url.RawQuery

	// clean up path
	urlPath = path.Clean(urlPath)

	// get formatted URL minus scheme so we can build this into Opaque
	url.Scheme, url.Path, url.RawQuery = "", "", ""
	s := url.String()
	url.Scheme = scheme

	// build opaque URI
	url.Opaque = s + urlPath
	if query != "" {
		url.Opaque += "?" + query
	}
}

// Whether the byte value can be sent without escaping in AWS URLs
var noEscape [256]bool
var noEscapeInitialized = false

// initialise noEscape
func initNoEscape() {
	for i := range noEscape {
		// Amazon expects every character except these escaped
		noEscape[i] = (i >= 'A' && i <= 'Z') ||
			(i >= 'a' && i <= 'z') ||
			(i >= '0' && i <= '9') ||
			i == '-' ||
			i == '.' ||
			i == '_' ||
			i == '~'
	}
}

// escapePath escapes part of a URL path in Amazon style
func escapePath(path string, encodeSep bool) string {
	if !noEscapeInitialized {
		initNoEscape()
		noEscapeInitialized = true
	}

	var buf bytes.Buffer
	for i := 0; i < len(path); i++ {
		c := path[i]
		if noEscape[c] || (c == '/' && !encodeSep) {
			buf.WriteByte(c)
		} else {
			buf.WriteByte('%')
			buf.WriteString(strings.ToUpper(strconv.FormatUint(uint64(c), 16)))
		}
	}
	return buf.String()
}

func convertType(v reflect.Value) (*string, error) {
	if !v.IsValid() {
		return nil, nil
	}

	var str string
	switch value := v.Interface().(type) {
	case string:
		str = value
	case []byte:
		str = base64.StdEncoding.EncodeToString(value)
	case bool:
		str = strconv.FormatBool(value)
	case int64:
		str = strconv.FormatInt(value, 10)
	case int:
		str = strconv.Itoa(value)
	case float64:
		str = strconv.FormatFloat(value, 'f', -1, 64)
	case float32:
		str = strconv.FormatFloat(float64(value), 'f', -1, 32)
	case time.Time:
		const ISO8601UTC = "2006-01-02T15:04:05Z"
		str = value.UTC().Format(ISO8601UTC)
	default:
		err := fmt.Errorf("Unsupported value for param %v (%s)", v.Interface(), v.Type().Name())
		return nil, err
	}
	return &str, nil
}
