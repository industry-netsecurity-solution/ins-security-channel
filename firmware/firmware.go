package firmware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/industry-netsecurity-solution/ins-security-channel/ins"
	"github.com/industry-netsecurity-solution/ins-security-channel/logger"
	"github.com/industry-netsecurity-solution/ins-security-channel/request"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"
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

func CheckUpdate(Conf ins.FirmwareConfigurations) (*bytes.Buffer, error) {

	// 기존의 Config file을 로딩한다.
	var buf *bytes.Buffer = nil
	var err error = nil

	if buf, err = ins.ReadFile(Conf.ConfigFilepath); err != nil {
		return nil, err
	}

	// 기존의 Config file로 firmware 업데이트가 있는지 검사한다.
	checkParam := request.NewRequestParam()
	checkParam.SetQuerypath("firmware/check")
	//checkParam.querypath = "code/json/sample1.json"
	checkParam.SetHeader("Content-Type", "application/json")
	checkParam.SetHeader("Accept", "application/json")
	checkParam.SetData(buf.Bytes())

	results := new(ins.SMap)
	httpRequest := request.HttpRequest{Conf.Http}
	if _, e := httpRequest.DoRequest(checkParam, func(resp *resty.Response) error {
		status := resp.StatusCode()
		if status != 200 && status != 201 && status != 202 && status != 205 {
			return errors.New(fmt.Sprintf("%d %s - %s", status, http.StatusText(status), resp.Request.URL))
		}

		contents := resp.Body()
		err := json.Unmarshal(contents, results)
		if err != nil {
			return err
		}

		logger.Println("Received firmware configuration:", string(contents))

		return nil
	}); e != nil {
		return nil, e
	}

	if v := results.Get("type"); v == nil {
		err = errors.New("firmware/check, unknown response")
		return nil, err
	} else {
		if strings.Compare(strings.ToLower(v.(string)), "success") != 0 {
			if message := results.Get("message"); message == nil {
				err = errors.New("unknown error")
				return nil, err
			} else {
				err = errors.New(message.(string))
				return nil, err
			}
		}
	}

	if v := results.Get("result"); v == nil {
		// result == null 이면 동일한 버전임
		return nil, nil
	} else {
		contents := v.(map[string]interface{})
		if cfg, e := json.MarshalIndent(contents, "", "    "); e == nil {
			firmwareConfig := new(bytes.Buffer)
			firmwareConfig.Write(cfg)
			return firmwareConfig, nil
		} else {
			return nil, e
		}
	}

	// Never Reached
	return nil, nil
}

func DownloadFirmware(Conf ins.FirmwareConfigurations, firmwareConfig *bytes.Buffer) error {

	var err error = nil
	//fmt.Println(string(firmwareConfig.Bytes()))

	// 임시 파일로 저장한다.
	uuidfirmware := uuid.New()
	tempfirmname := fmt.Sprintf("%s/%s", Conf.DownlaodFilepath, uuidfirmware.String())

	firmwareParam := request.NewRequestParam()
	firmwareParam.SetQuerypath("firmware/download")
	//firmwareParam.querypath: "code/json/sample2.json",
	firmwareParam.SetHeader("Content-Type", "application/json")
	firmwareParam.SetHeader("Accept", "application/octet-stream; charset=binary")
	firmwareParam.SetData(firmwareConfig.Bytes())
	firmwareParam.SetFilename(tempfirmname)

	var mediatype string
	var params map[string]string
	httpRequest := request.HttpRequest{Conf.Http}
	if _, e := httpRequest.DoRequest(firmwareParam, func(resp *resty.Response) error {
		headers := resp.Header()
		disposition := headers.Get("Content-Disposition")
		mediatype , params , err = mime.ParseMediaType(disposition)

		return nil
	}); e != nil {
		return e
	}

	firmwareInfo := new(ins.SMap)
	if err = json.Unmarshal(firmwareConfig.Bytes(), firmwareInfo); err != nil {
		_ = os.Remove(tempfirmname)
		return err
	}

	if v := firmwareInfo.Get("hash"); v != nil {
		if hash, err := ins.SumSHA256(tempfirmname); err != nil {
			_ = os.Remove(tempfirmname)
			return err
		} else {
			if strings.EqualFold(v.(string), hash) == false {
				return errors.New(fmt.Sprintf("mismatch hash: %s", hash))
			}
		}
	}

	uuidconfig := uuid.New()
	tempconfname := fmt.Sprintf("%s/%s", path.Dir(Conf.ConfigFilepath), uuidconfig.String())

	if err = ioutil.WriteFile(tempconfname, firmwareConfig.Bytes(), 0644); err != nil {
		_ = os.Remove(tempfirmname)
		return err
	}

	var newpath string
	if filename, ok :=  params["filename"]; ok {
		newpath = fmt.Sprintf("%s/%s", Conf.DownlaodFilepath, filename)
	} else {
		newpath = fmt.Sprintf("%s/firmware.bin", Conf.DownlaodFilepath)
	}


	if _, err := ins.Copy(tempfirmname, newpath); err != nil {
		_ = os.Remove(tempfirmname)
		logger.Println("Remove:", tempfirmname)

		_ = os.Remove(tempconfname)
		logger.Println("Remove:", tempconfname)
		return err;
	}
	_ = os.Remove(tempfirmname)
	logger.Println("Firmmware file downlaod:", newpath)

	if _, err := ins.Copy(tempconfname, Conf.ConfigFilepath); err != nil {
		_ = os.Remove(newpath)
		logger.Println("Remove:", newpath)

		_ = os.Remove(tempconfname)
		logger.Println("Remove:", tempconfname)
		return err;
	}
	_ = os.Remove(tempconfname)
	logger.Println("Firmmware config downlaod:", Conf.ConfigFilepath)

	return nil
}


func UpdateFirmware(Conf ins.FirmwareConfigurations) error {

	var firmwareConfig *bytes.Buffer
	var err error

	if Conf.Enable == false {
		return nil
	}

	if firmwareConfig, err = CheckUpdate(Conf); err == nil {
		if firmwareConfig == nil {
			// 신규 버전 없음
			return nil
		}

		logger.Println("Try download:", string(firmwareConfig.Bytes()))
		err = DownloadFirmware(Conf, firmwareConfig)
	}
	return err
}
