package augeuHttp

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

func GetRequest(target string, header Header, proxy string) (string, error) {

	req, err := http.NewRequest(Get, target, nil)
	if err != nil {
		return "", fmt.Errorf("http.NewRequest -> %v", err)
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	var proxyUrl *url.URL
	if proxy != "" {
		proxyUrl, err = url.Parse(proxy)
		if err != nil {
			return "", fmt.Errorf("url.Parse -> %v", err)
		}
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("http.RequestWithJson -> %v", err)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String(), nil
}
