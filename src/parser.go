package main

import (
	"errors"
	"mvdan.cc/xurls"
)

func ParseUrlsFromText(text string) ([]string, error) {
	foundUrls := xurls.Strict().FindAllString(text, -1)

	if len(foundUrls) == 0 {
		return nil, errors.New("URLs not found in message.")
	}

	return foundUrls, nil
}
