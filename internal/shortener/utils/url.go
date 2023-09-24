package utils

import (
	"regexp"
)

type Blacklist map[string]struct{}

func IsValidUrl(str string, blacklist Blacklist) (bool, error) {
	r, err := regexp.Match(`^(?:https?:\/\/)?(?:\w+[.])+[a-zA-Z]+(?:.*)*$`, []byte(str))
	if err != nil {
		return false, err
	}
	_, blacklisted := blacklist[str]
	return r && !blacklisted, nil
}

func Shorten(str string, shortener func([]byte) string, length int) string {
	short := shortener([]byte(str))
	return short[:min(len(short), length)]
}
