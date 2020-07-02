package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// 보통 큰 이미지는 하나에 3MB정도 함
// 그래서 4MB 할당해서 복사없이 한번에 할수있게함
// 수정, 똥컴 배려로 낮춤
const datomiDefaultReadSize = 262144

func newReadSlice(size int) []byte {
	return make([]byte, size)
}

func growSlice(b []byte) []byte {
	c := cap(b)
	n := newReadSlice(c * 2)
	copy(n, b)
	return n
}

func readFromNetworkStream(r io.Reader) (buf []byte, err error) {
	buf = newReadSlice(datomiDefaultReadSize)
	readIndex := 0
	for {
		len, e := r.Read(buf[readIndex:])
		if e == io.EOF {
			return
		}

		if len < 0 {
			return nil, errors.New("read < 0")
		}

		if e != nil {
			return nil, e
		}

		readIndex += len
		if readIndex == cap(buf) {
			buf = growSlice(buf)
		}
	}
}

//httpGet 실제로 http 요청해서 download함
func httpGet(url string, referer string) []byte {
	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err, "make request error continue")
			continue
		}

		if len(referer) > 0 {
			req.Header.Add("referer", referer)
			req.Header.Add("DNT", "1")
		}
		client := http.Client{
			Timeout: time.Duration(60 * time.Minute),
		}

		res, err := client.Do(req)

		if err != nil {
			fmt.Println(err, "http connection error continue")
			continue
		}
		defer res.Body.Close()

		data, err := readFromNetworkStream(res.Body)

		if err != nil {
			fmt.Println(err, "http read error continue")
			continue
		}

		return data
	}
}
