package helpers

import (
	"net/http"
	"log"
	"errors"
	"io/ioutil"
	"os"
	"fmt"
	"time"
	"path"
	"bytes"
	"strconv"
	"io"
)

func GetBodyFromUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("respone was not 200 for " + url)
	}

	bodyBytes, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		return nil, errors.New("failed to recieve response from url")
	}
	return bodyBytes, nil
}


func PrintDownloadPercent(done chan int64, path string, total int64) {

	var stop bool = false

	for {
		select {
		case <-done:
			stop = true
		default:

			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			fi, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}

			size := fi.Size()

			if size == 0 {
				size = 1
			}

			var percent float64 = float64(size) / float64(total) * 100
			fmt.Print("\r")
			fmt.Printf("%.0f", percent)
			fmt.Print("%  ")
		}

		if stop {
			fmt.Print("\r")
			break
		}

		time.Sleep(time.Second)
	}
}

func DownloadFile(url string, dest string) string {

	file := path.Base(url)

	log.Printf("Downloading file %s from %s\n", file, url)

	var buffer bytes.Buffer
	buffer.WriteString(dest)
	buffer.WriteString("/")
	buffer.WriteString(file)

	start := time.Now()

	out, err := os.Create(buffer.String())

	if err != nil {
		fmt.Println(buffer.String())
		panic(err)
	}

	defer out.Close()

	headResp, err := http.Head(url)

	if err != nil {
		panic(err)
	}

	defer headResp.Body.Close()

	size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))

	if err != nil {
		panic(err)
	}

	done := make(chan int64)

	go PrintDownloadPercent(done, buffer.String(), int64(size))

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)

	if err != nil {
		panic(err)
	}

	done <- n

	elapsed := time.Since(start)
	log.Printf("Download completed in %s", elapsed)
	return buffer.String()
}
