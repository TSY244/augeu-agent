package augeuHttp

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

func PostRequestWithJson(target string, header Header, body string) (string, error) {

	payload := strings.NewReader(body)

	req, err := http.NewRequest(Post, target, payload)
	if err != nil {
		return "", fmt.Errorf("http.NewRequest -> %v", err)
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	// 设置json 头
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("http.RequestWithJson -> %v", err)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String(), nil
}

func PostWithFilePayload(target string, payload io.Reader, writer *multipart.Writer) (string, error) {
	fmt.Println("into PostWithFilePayload")
	client := &http.Client{}
	req, err := http.NewRequest(Post, target, payload)
	if err != nil {
		return "", fmt.Errorf("http.NewRequest -> %v", err)
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Connection", "keep-alive")
	//req.Header.Add("Content-Type", "multipart/form-data; boundary=--------------------------105291221491684600198529")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("http.RequestWithJson -> %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("io.ReadAll -> %v", err)
	}

	return string(body), nil
}
