package util

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
)

func GenerateApiSign(maps map[string]string) string {
	var keys []string
	for k, _ := range maps {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	newsMap := make(map[string]string)
	for _, k := range keys {
		//textQuoted := strconv.QuoteToASCII(maps[k])
		//textUnquoted := textQuoted[1 : len(textQuoted)-1]
		newsMap[k] = maps[k]
	}

	fmt.Println(newsMap)
	marshal, _ := json.Marshal(newsMap)
	fmt.Println("marshal", string(marshal)+newsMap["secret"])
	newSig := md5.Sum([]byte(fmt.Sprintf("%x", md5.Sum([]byte(string(marshal)+newsMap["secret"])))))
	fmt.Println("newSig", fmt.Sprintf("%x", newSig))
	return fmt.Sprintf("%x", newSig)

}

var host = "http://qymeta_api.newmin.cn/api/"

func PostJsonRequest(uri string, params map[string]string) interface{} {
	fmt.Println("params:", params)
	path := host + uri
	bytesData := &bytes.Buffer{}
	writer := multipart.NewWriter(bytesData)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	_ = writer.Close()
	req, err := http.NewRequest("POST", path, bytesData)
	fmt.Println(writer.FormDataContentType())
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		return err.Error()
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)

}
func GetJsonRequest(uri string, paramMap map[string]string) interface{} {
	fmt.Println("params:", paramMap)
	path := host + uri
	params := url.Values{}
	var parseURL *url.URL
	parseURL, err := url.Parse(path)
	if err != nil {
		return err.Error()
	}
	if paramMap != nil {
		for k, v := range paramMap {
			params.Set(k, v)
		}
	}
	parseURL.RawQuery = params.Encode()
	fmt.Println(parseURL.String())
	req, err := http.NewRequest("GET", parseURL.String(), nil)
	req.Header.Set("Content-Type", "multipart/form-data")
	if err != nil {
		return err.Error()
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}
