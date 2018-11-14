package reporter

import (
	"net/http"
)

const ftUrl = "https://sc.ftqq.com/SCU6693T4108bacae69e4d52fbb8dd73f048bb185b6d687b94350.send"

func ReportError(title string, err error) {
	go func() {
		http.Get(ftUrl + "?text=" + title + "&desp=" + err.Error())
	}()
}

func ReportMessage(title string) {
	go func() {
		http.Get(ftUrl + "?text=" + title)
	}()
}
