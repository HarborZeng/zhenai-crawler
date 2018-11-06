package parser

import (
	"testing"
	"zhenai-crawler/crawler/fetcher"
)

func TestParseCityList(t *testing.T) {
	contents, _ := fetcher.Fetch(CityListURL)

	result := ParseCityList(contents)

	const resultSize = 470

	if len(result.Requests) != resultSize {
		t.Errorf("There should be %d requests, but %d instead.",
			resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("There should be %d items rather than %d",
			resultSize, len(result.Items))
	}
}
