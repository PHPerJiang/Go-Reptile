package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	content, err := ioutil.ReadFile("test.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(content)
	const resultSize  = 470
	if len(result.Requests) != resultSize {
		t.Errorf("Result should have %d requests, but had %d requests now",resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("Result should have %d items, but had %d items now", resultSize, len(result.Items))
	}
}
