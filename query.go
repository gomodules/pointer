package aws

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// QueryClient is the underlying client for Query APIs.
type QueryClient struct {
	Context    Context
	Client     *http.Client
	Endpoint   string
	APIVersion string
}

// Do sends an HTTP request and returns an HTTP response, following policy
// (e.g. redirects, cookies, auth) as configured on the client.
func (c *QueryClient) Do(op, method, uri string, req, resp interface{}) error {
	body := url.Values{"Action": {op}, "Version": {c.APIVersion}}
	if err := c.loadValues(body, req, ""); err != nil {
		return err
	}

	httpReq, err := http.NewRequest(method, c.Endpoint+uri, strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Set("User-Agent", "aws-go")
	if err := c.Context.sign(httpReq); err != nil {
		return err
	}

	httpResp, err := c.Client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		var err queryErrorResponse
		if err := xml.NewDecoder(httpResp.Body).Decode(&err); err != nil {
			return err
		}
		return err.Err()
	}

	if resp != nil {
		return xml.NewDecoder(httpResp.Body).Decode(resp)
	}
	return nil
}

type queryErrorResponse struct {
	XMLName   xml.Name `xml:"ErrorResponse"`
	Type      string   `xml:"Error>Type"`
	Code      string   `xml:"Error>Code"`
	Message   string   `xml:"Error>Message"`
	RequestID string   `xml:"RequestId"`
}

func (e queryErrorResponse) Err() error {
	return APIError{
		Type:      e.Type,
		Code:      e.Code,
		Message:   e.Message,
		RequestID: e.RequestID,
	}
}

func (c *QueryClient) loadValues(v url.Values, i interface{}, prefix string) error {
	value := reflect.ValueOf(i)

	// follow any pointers
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			if prefix == "" {
				prefix = strconv.Itoa(i + 1)
			} else {
				prefix = prefix + ".member." + strconv.Itoa(i+1)
			}
			if err := c.loadValues(v, value.Index(i).Interface(), prefix); err != nil {
				return err
			}
		}
		return nil
	}

	return c.loadStruct(v, value, prefix)
}

func (c *QueryClient) loadStruct(v url.Values, value reflect.Value, prefix string) error {
	if !value.IsValid() {
		return nil
	}

	t := value.Type()
	for i := 0; i < value.NumField(); i++ {
		value := value.Field(i)
		paths := strings.Split(t.Field(i).Tag.Get("xml"), ">")
		name := paths[0]

		if name == "" {
			name = t.Field(i).Name
		}
		if prefix != "" {
			name = prefix + "." + name
		}
		switch casted := value.Interface().(type) {
		case StringValue:
			if casted != nil {
				v.Set(name, *casted)
			}
		case BooleanValue:
			if casted != nil {
				v.Set(name, strconv.FormatBool(*casted))
			}
		case LongValue:
			if casted != nil {
				v.Set(name, strconv.FormatInt(*casted, 10))
			}
		case IntegerValue:
			if casted != nil {
				v.Set(name, strconv.Itoa(*casted))
			}
		case DoubleValue:
			if casted != nil {
				v.Set(name, strconv.FormatFloat(*casted, 'f', -1, 64))
			}
		case FloatValue:
			if casted != nil {
				v.Set(name, strconv.FormatFloat(float64(*casted), 'f', -1, 32))
			}
		case []string:
			if len(casted) != 0 {
				for i, val := range casted {
					v.Set(fmt.Sprintf("%s.member.%d", name, i+1), val)
				}
			}
		default:
			if err := c.loadValues(v, value.Interface(), name); err != nil {
				return err
			}
		}
	}
	return nil
}
