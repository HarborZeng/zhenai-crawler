package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36"

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	//<-rateLimiter
	request, e := http.NewRequest(http.MethodGet, url, nil)
	if e != nil {
		panic(e)
	}
	request.Header.Add("User-Agent", UserAgent)

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong http status: %s", resp.Status)
	}

	bodyReader := bufio.NewReader(resp.Body)
	encode := determineEncoding(bodyReader)

	return ioutil.ReadAll(
		transform.NewReader(bodyReader, encode.NewDecoder()))
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, e := r.Peek(1024)
	if e != nil {
		log.Printf("Fetch err: wrong encoding %v", e)
		return unicode.UTF8
	}
	encode, _, _ := charset.DetermineEncoding(bytes, "")
	return encode
}
