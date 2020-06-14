package main

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
)

// FetchBody http get 요청으로 body 만을 가져옵니다.
func FetchBody(url string) (string, error) {
	return FetchBodyWithContext(context.Background(), url)
}

// FetchBodyWithContext http get 요청으로 body 만을 가져옵니다.
func FetchBodyWithContext(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	return FetchBodyWithHTTPRequest(req.WithContext(ctx))
}

// FetchBodyWithHTTPRequest request 를 통해 요청해서 body 만을 가져옵니다.
func FetchBodyWithHTTPRequest(req *http.Request) (string, error) {
	res, err := http.DefaultClient.Do(req)
	if res != nil {
		defer res.Body.Close()
		defer io.Copy(ioutil.Discard, res.Body)
	}
	if err != nil {
		return "", err
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
