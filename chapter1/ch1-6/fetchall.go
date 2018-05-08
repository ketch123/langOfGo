package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //ゴルーチンの開始
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //chチャネルから受信
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

/*
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //chチャネルへ送信
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() //資源をリークさせない
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2fs  %7d  %s", secs, nbytes, url)
}
*/

//練習問題 1.10, 1.11
func fetch(url string, ch chan<- string) {
	var b bytes.Buffer
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //chチャネルへ送信
		return
	}

	nbytes, err := io.Copy(&b, resp.Body)
	resp.Body.Close() //資源をリークさせない
	file, err := os.Create("./output/fetchall-out")
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	file.WriteString(b.String())
	fmt.Printf("file name is %s\n", file.Name())
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2fs  %7d  %s", secs, nbytes, url)
}
