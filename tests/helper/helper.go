package helper

import (
	"crypto/tls"
	"net/http"
	"time"
)

func HttpWaitForStatus(url string, maxRetry int, interval int, expectedCode int) bool {
	for i := 0; i < maxRetry; i++ {
		transporter := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: transporter}
		resp, err := client.Get(url)
		if err != nil {
			time.Sleep(time.Duration(interval) * time.Second)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode == expectedCode {
			return true
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
	return false
}
