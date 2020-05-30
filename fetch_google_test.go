package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func fetchGoogle(timeout time.Duration, t *testing.T) string {
	r, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), timeout)
	defer cancelFunc()

	g := r.WithContext(timeoutRequest)

	res, err := http.DefaultClient.Do(g)

	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	return string(body)
}

func TestFetchGoogleOK(t *testing.T) {
	body := fetchGoogle(10*time.Second, t)
	fmt.Println(body)
}

func TestFetchGoogleFail(t *testing.T) {
	body := fetchGoogle(1*time.Nanosecond, t)
	fmt.Println(body)
}
