package request

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/industry-netsecurity-solution/ins-security-channel/ins"
	"net/http"
	"strings"
)

type RequestParam struct {
	method    string
	querypath string
	headers   map[string]string
	data      interface{}
	filename  string
}

func NewRequestParam() *RequestParam {
	param := new(RequestParam)
	param.method = http.MethodPost
	param.headers = make(map[string]string)
	param.data = nil

	return param
}

func (v *RequestParam) SetMethod(method string) {
	v.method = method
}

func (v *RequestParam) SetQuerypath(querypath string) {
	v.querypath = querypath
}

func (v *RequestParam) SetHeader(key, value string) {
	v.headers[key] = value
}

func (v *RequestParam) SetData(data interface{}) {
	v.data = data
}

func (v *RequestParam) SetFilename(filename string) {
	v.filename = filename
}

type HttpRequest struct {
	ins.HttpConfigurations
}

func (v HttpRequest) DoRequest(param *RequestParam, handler func(resp *resty.Response) error) (int, error) {
	if 0 < len(v.Authorization) {
		encoded := base64.StdEncoding.EncodeToString([]byte(v.Authorization))
		param.headers["Authorization"] = fmt.Sprintf("Bearer %s", encoded)
	}

	u, err := v.Url(param.querypath)
	if err != nil {
		return -1, err
	}
	httpurl := u.String()

	client := resty.New()
	if strings.EqualFold(u.Scheme, "https") {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.SetTransport(tr)
	}
	client.SetCloseConnection(true)

	req := client.R().
		SetHeaders(param.headers).
		SetBody(param.data)
	if 0 < len(param.filename) {
		req.SetOutput(param.filename)
	}
	var resp *resty.Response

	if param.method == http.MethodGet {
		if resp, err = req.Get(httpurl); err != nil {
			return -1, err
		}
	} else if param.method == http.MethodPost {
		if resp, err = req.Post(httpurl); err != nil {
			return -1, err
		}
	} else if param.method == http.MethodPut {
		if resp, err = req.Put(httpurl); err != nil {
			return -1, err
		}
	} else if param.method == http.MethodDelete {
		if resp, err = req.Delete(httpurl); err != nil {
			return -1, err
		}
	} else if param.method == http.MethodHead {
		if resp, err = req.Head(httpurl); err != nil {
			return -1, err
		}
	} else if param.method == http.MethodOptions {
		if resp, err = req.Options(httpurl); err != nil {
			return -1, err
		}
	} else if param.method == http.MethodPatch {
		if resp, err = req.Patch(httpurl); err != nil {
			return -1, err
		}
	} else {
		return -1, errors.New(http.ErrNotSupported.Error())
	}
	defer resp.RawBody().Close()
	status := resp.StatusCode()

	if handler != nil {
		if err := handler(resp); err != nil {
			return status, err
		}
	}

	return status, nil
}

