package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// FetchResponse 요청에 대한 응답을 반환합니다
type FetchResponse struct {
	body []byte
}

// String Convert the response body to a string.
func (f *FetchResponse) String() string {
	if f == nil || f.body == nil {
		return ""
	}
	return string(f.body)
}

// JSON Convert the response body to a JSON.
func (f *FetchResponse) JSON(v interface{}) error {
	if f == nil || f.body == nil {
		return errors.New("FetchResponse nil body")
	}
	if err := json.Unmarshal(f.body, v); err != nil {
		return err
	}
	return nil
}

// JSONMap Convert the response body to a JSON map.
func (f *FetchResponse) JSONMap() map[string]interface{} {
	m := make(map[string]interface{}, 10)
	f.JSON(&m)
	return m
}

// FetchBody http get 요청으로 body 만을 가져옵니다.
func FetchBody(url string) (*FetchResponse, error) {
	return FetchBodyWithContext(context.Background(), url)
}

// FetchBodyWithContext http get 요청으로 body 만을 가져옵니다.
func FetchBodyWithContext(ctx context.Context, url string) (*FetchResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return FetchBodyWithRequest(req.WithContext(ctx))
}

// FetchBodyWithRequest request 를 통해 요청해서 body 만을 가져옵니다.
func FetchBodyWithRequest(req *http.Request) (*FetchResponse, error) {
	res, err := http.DefaultClient.Do(req)
	if res != nil {
		defer res.Body.Close()
		defer io.Copy(ioutil.Discard, res.Body)
	}
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &FetchResponse{
		body: buf,
	}, nil
}
