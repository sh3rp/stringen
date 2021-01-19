package stringen

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"log"
)

func EncodeBase64(r *bufio.Reader) string {
	var data string
	nBytes, nChunks := int64(0), int64(0)
	buf := make([]byte, 0, 4*1024)

	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		nChunks++
		nBytes += int64(len(buf))
		data += base64.StdEncoding.EncodeToString(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
	return data
}

func DecodeBase64(r *bufio.Reader) string {
	var data string
	nBytes, nChunks := int64(0), int64(0)
	buf := make([]byte, 0, 4*1024)

	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		nChunks++
		nBytes += int64(len(buf))
		data += string(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
	bytes, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		fmt.Println("Error base64 decoding")
		return ""
	}
	return string(bytes)
}
