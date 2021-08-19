package firmware

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/industry-netsecurity-solution/ins-security-channel/ins"
	"net/http"
)

// POST http://{host}:9092/firmware/check
// Content-type: application/json
// Accept: application/json

// POST http://{host}:9092/firmware/download
// Content-type: application/json
// Accept: application/octet-stream

//{
//"gwid": "i-am-center-gw-0001",
//"gwtype": "집중게이트웨이",
//"manufacturer": "(주)에프아이시스",
//"model": "BA000000-000001",
//"version": "2.1",
//"hash": null,
//"date": null
//}

// Response POST http://{host}:9092/firmware/check
//{
//	"message": "firmware download Version Check success.",
//	"type": "SUCCESS",
//	"result": {
//		"gwid": "i-am-center-gw-0001",
//		"gwtype": "집중게이트웨이",
//		"model": "BA000000-000001",
//		"manufacturer": "(주)에프아이시스",
//		"version": "2.4",
//		"hash": "ec3826548263f8eab4179f3d95543d5adf664e1bd6bf6af0779ab3d33fb4998a",
//		"date": "2021-08-18T06:40:22.000+0000“,
//	},
//	"respDate": "2021-08-19T14:42:17"
//}

type FIrmware struct {
	GwId              string   `json:"gwid"`
	GwType            string   `json:"gwtype"`
	Model             string   `json:"model"`
	Manufacturer      string   `json:"manufacturer"`
	Version           string   `json:"version"`
	Date              string   `json:"date,omitempty"`
	hash              string   `json:"hash,omitempty"`
}

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

func CheckUpdate(Conf ins.FirmwareConfigurations) error {

	// 기존의 Config file을 로딩한다.
	var buf *bytes.Buffer = nil
	var err error = nil
	if buf, err = ins.ReadFile(Conf.ConfigFilepath); err != nil {
		return err
	}

	// 기존의 Config file로 firmware 업데이트가 있는지 검사한다.
	checkParam := NewRequestParam()
	checkParam.SetQuerypath("firmware/check")
	//checkParam.querypath = "code/json/sample1.json"
	checkParam.SetHeader("Content-Type", "application/json")
	checkParam.SetHeader("Accept", "application/json")
	checkParam.SetData(buf.Bytes())

	firmwareConfig := bytes.Buffer{}
	request := HttpRequest{Conf.Http}
	if r, e := request.DoRequest(checkParam, func(resp *resty.Response) error {
		status := resp.StatusCode()
		if status != 200 && status != 201 && status != 202 && status != 205 {
			return errors.New(fmt.Sprintf("%d %s - %s", status, http.StatusText(status), resp.Request.URL))
		}

		contents := resp.Body()
		firmwareConfig.Write(contents)
		fmt.Println(string(contents))
		return nil
	}); e != nil {
		return e
	} else {
		fmt.Println(r)
	}

	// 임시 파일로 저장한다.
	UUID := uuid.New()

	firmwareParam := NewRequestParam()
	firmwareParam.SetQuerypath("firmware/download")
	//firmwareParam.querypath: "code/json/sample2.json",
	firmwareParam.SetHeader("Content-Type", "application/json")
	firmwareParam.SetHeader("Accept", "application/octet-stream; charset=binary")
	firmwareParam.SetData(firmwareConfig.Bytes())
	firmwareParam.SetFilename(UUID.String())

	if r, e := request.DoRequest(firmwareParam, nil); e != nil {
		return e
	} else {
		fmt.Println(r)
	}

	// TODO: firmwareConfig는 파일로 저장
	// TODO: 임시 저장된 Firmware 파일

	return nil
}