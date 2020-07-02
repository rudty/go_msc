package main

import (
	"io"
	"net/http"
	"os"
)

func pipeStream(istream io.Reader, ostream io.Writer) error {
	var buf [8094]byte
	for {
		recvLen, err := istream.Read(buf[:])
		switch err {
		case io.EOF:
			return nil
		default:
			return err
		case nil:
		}
		_, err = ostream.Write(buf[:recvLen])
		if err != nil {
			return err
		}
	}
}
func StaticFileDownload(url string, savePath string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return err
	}

	f, err := os.Create(savePath)
	if err != nil {
		return err
	}

	if err := pipeStream(res.Body, f); err != nil {
		return err
	}
	return nil
}

func main() {
	StaticFileDownload("https://blog.golang.org/gopher/header.jpg", "test.jpg")
}
