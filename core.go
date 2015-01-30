package aws

import (
	"bytes"
	"container/list"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"reflect"
	"regexp"
	"time"
)

func SendHandler(r *Request) {
	r.HTTPResponse, r.Error = r.Service.Config.HTTPClient.Do(r.HTTPRequest)
}

type Operation struct {
	Name        string
	HTTPMethod  string
	HTTPPath    string
	InPayload   string
	OutPayload  string
	Required    []string
	QueryParams []string
	URIParams   []string
	InHeaders   []string
	OutHeaders  []string
}

type Service struct {
	Config       *Config
	Handlers     Handlers
	ManualSend   bool
	ServiceName  string
	APIVersion   string
	Endpoint     string
	JSONVersion  string
	TargetPrefix string
}

var schemeRE = regexp.MustCompile("^([^:]+)://")

func (s *Service) Initialize() {
	if s.Config.HTTPClient == nil {
		s.Config.HTTPClient = http.DefaultClient
	}

	s.Handlers.Send.PushBack(SendHandler)
	s.buildEndpoint()
}

type Request struct {
	*Service
	Time         time.Time
	Operation    *Operation
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Body         io.ReadSeeker
	Params       interface{}
	Error        error
	Data         interface{}
	RequestID    string
	Debug        uint
}

func NewRequest(service *Service, operation *Operation, params interface{}, data interface{}) *Request {
	method := operation.HTTPMethod
	if method == "" {
		method = "POST"
	}
	path := service.Endpoint + operation.HTTPPath

	httpReq, _ := http.NewRequest(method, path, nil)
	if httpReq.URL.Path == "" {
		httpReq.URL.Path = "/"
	}

	r := &Request{
		Service:     service,
		Time:        time.Now(),
		Operation:   operation,
		HTTPRequest: httpReq,
		Body:        nil,
		Params:      params,
		Error:       nil,
		Data:        data,
		Debug:       service.Config.LogLevel,
	}

	r.AddDebugHandlers()

	return r
}

func (s *Service) buildEndpoint() {
	if s.Config.Endpoint != "" {
		s.Endpoint = s.Config.Endpoint
	} else {
		s.Endpoint = s.endpointForRegion()
	}

	if !schemeRE.MatchString(s.Endpoint) {
		scheme := "https"
		if s.Config.DisableSSL {
			scheme = "http"
		}
		s.Endpoint = scheme + "://" + s.Endpoint
	}
}

func (r *Request) AddDebugHandlers() {
	if r.Debug == 0 {
		return
	}

	r.Handlers.Sign.PushBack(func(r *Request) {
		dumpedBody, _ := httputil.DumpRequest(r.HTTPRequest, true)

		fmt.Printf("=> [%s] %s.%s(%+v)\n", r.Time,
			r.Service.ServiceName, r.Operation.Name, r.Params)
		fmt.Printf("---[ REQUEST  ]--------------------------------------\n")
		fmt.Printf("%s\n", string(dumpedBody))
		fmt.Printf("-----------------------------------------------------\n\n")
	})
	r.Handlers.Send.PushBack(func(r *Request) {
		defer r.HTTPResponse.Body.Close()
		dumpedBody, _ := httputil.DumpResponse(r.HTTPResponse, true)

		fmt.Printf("---[ RESPONSE ]--------------------------------------\n")
		fmt.Printf("%s\n", string(dumpedBody))
		fmt.Printf("-----------------------------------------------------\n\n")
	})
}

func (r Request) ParamsFilled() bool {
	return reflect.ValueOf(r.Params).Elem().IsValid()
}

func (r Request) DataFilled() bool {
	return reflect.ValueOf(r.Data).Elem().IsValid()
}

func (r *Request) SetBufferBody(buf []byte) {
	r.SetReaderBody(bytes.NewReader(buf))
}

func (r *Request) SetReaderBody(reader io.ReadSeeker) {
	r.HTTPRequest.Body = ioutil.NopCloser(reader)
	r.Body = reader
}

func (r *Request) Send() error {
	r.Error = nil

	r.Handlers.Build.Run(r)
	if r.Error != nil {
		return r.Error
	}

	r.Handlers.Sign.Run(r)
	if r.Error != nil {
		return r.Error
	}

	r.Handlers.Send.Run(r)
	if r.Error != nil {
		return r.Error
	}

	r.Handlers.ValidateResponse.Run(r)
	if r.Error != nil {
		return r.Error
	}

	r.Handlers.Unmarshal.Run(r)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

type Handlers struct {
	Build            HandlerList
	Sign             HandlerList
	Send             HandlerList
	ValidateResponse HandlerList
	Unmarshal        HandlerList
}

type HandlerList struct {
	list.List
}

func (l *HandlerList) Run(r *Request) {
	for e := l.Front(); e != nil; e = e.Next() {
		h := e.Value.(func(*Request))
		h(r)
	}
}
