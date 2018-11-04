package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error: status code ", resp.StatusCode)
		return
	}

	encode := determineEncoding(resp.Body)

	all, err := ioutil.ReadAll(
		transform.NewReader(resp.Body, encode.NewDecoder()))
	if err != nil {
		panic(err)
	}
	//fmt.Printf("contents are %s\n", all)
	printCitiesList(all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, e := bufio.NewReader(r).Peek(1024)
	if e != nil {
		panic(e)
	}
	encode, _, _ := charset.DetermineEncoding(bytes, "")
	return encode
}

func printCitiesList(contents []byte) {
	exp := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]+)</a>`)
	matches := exp.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("%s\n", m[1])
		fmt.Printf("%s\n", m[2])
	}
	fmt.Printf("%d matches was found.", len(matches))
}
