package main

import (
	"os"
	"testing"
)

func TestImageDownload(t *testing.T) {
	d := Downloader{}
	d.imageDownload("http://www.google.com", "test.html")
	err := os.Remove("test.html")
	if err != nil {
		t.Error(err)
	}
}
