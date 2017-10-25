package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

// DICOMResult DICOMResult
type DICOMResult struct {
	SOPInstanceUID                 string `json:"SOP Instance UID"`
	FileMetaInformationGroupLength int    `json:"File Meta Information Group Length"`
	PatientName                    string `json:"Patients Name"`
	PatientSex                     string `json:"Patients Sex"`
	PatientBirthDate               string `json:"Patient's Birth Date"`
	BodyPartExamined               string `json:"Body Part Examined"`
	Modality                       string `json:"Modality"`
	InstitutionName                string `json:"Institution Name"`
	Score                          string `json:"(暂无)"`
	FilePath                       string `json:"File Path"`
}

// Query body提交二进制数据
func Query(data []byte) ([]byte, error) {
	url := "http://127.0.0.1:9200/dicom_test/_search?pretty"
	body := bytes.NewReader(data)
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Printf("http.NewRequest,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	request.Header.Set("content-type", "application/json; charset=UTF-8")
	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("http.Do failed,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("http.Do failed,[err=%s][url=%s]", err, url)
	}
	return b, err
}

// ProcessResult 处理ES返回的字符串（可选择返回字段）
func ProcessResult(fields string, data []byte) ([]byte, error) {
	resultMaps := make([]map[string]string, 0)
	strs := strings.Split(fields, ",")
	hits := gjson.Get(string(data), "hits.hits.#._source")
	for _, hit := range hits.Array() {
		hitMap := hit.Map()
		newMap := make(map[string]string)
		for _, s := range strs {
			newMap[s] = hitMap[s].String()
		}
		resultMaps = append(resultMaps, newMap)
	}
	return json.Marshal(resultMaps)
}

// Save 保存DICOM信息到ES
func Save(data []byte) ([]byte, error) {
	var obj DICOMResult
	json.Unmarshal(data, &obj) // 解析

	url := "http://127.0.0.1:9200/dicom_test/external/" + obj.SOPInstanceUID

	body := bytes.NewReader(data)
	request, err := http.NewRequest("PUT", url, body)
	if err != nil {
		fmt.Printf("http.NewRequest,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	request.Header.Set("content-type", "application/json; charset=UTF-8")
	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("http.Do failed,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("http.Do failed,[err=%s][url=%s]", err, url)
	}
	return b, err
}
