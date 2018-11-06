package common

import (
	"net/http"
)

const ftUrl = "https://sc.ftqq.com/SCU6693T4108bacae69e4d52fbb8dd73f048bb185b6d687b94350.send"

func ReportError(title string, err error) {
	http.Get(ftUrl + "?text=" + title + "&desc=" + err.Error())
}

func ReportMessage(title string) {
	http.Get(ftUrl + "?text=" + title)
}
