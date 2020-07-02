package main

import (
	"encoding/binary"
	"fmt"
	"net/http"
	"time"
)

func httpGetBinary(url string) {
	for {
		client := http.Client{
			Timeout: time.Duration(60 * time.Second),
		}

		res, err := client.Get(url)

		if err != nil {
			fmt.Println(err, "http connection error continue")
			continue
		}
		defer res.Body.Close()

		var v int32
		var cnt = 0
		for true {
			if err := binary.Read(res.Body, binary.BigEndian, &v); err != nil {
				break
			}

		}
		fmt.Println(cnt)
		return

	}
}
