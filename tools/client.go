package tools

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
		log.Println("http.NewRequest,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	request.Header.Set("content-type", "application/json; charset=UTF-8")
	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		log.Println("http.Do failed,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("http.Do failed,[err=%s][url=%s]", err, url)
	}
	return b, err
}

// Save Save
func Save(data []byte) ([]byte, error) {
	var obj DICOMResult
	json.Unmarshal(data, &obj) // 解析

	url := "http://127.0.0.1:9200/dicom_test/external/" + obj.SOPInstanceUID

	body := bytes.NewReader(data)
	request, err := http.NewRequest("PUT", url, body)
	if err != nil {
		log.Println("http.NewRequest,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	request.Header.Set("content-type", "application/json; charset=UTF-8")
	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		log.Println("http.Do failed,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("http.Do failed,[err=%s][url=%s]", err, url)
	}
	return b, err
}
