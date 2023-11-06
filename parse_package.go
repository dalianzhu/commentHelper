package main

import (
	"log"
	"regexp"
	"strings"
)

// ParsePackage ...
func ParsePackage(text string) []*NeedCommentLine {
	re, err := regexp.Compile(`^package\s+([a-z0-9A-Z]+)`)
	if err != nil {
		log.Println(err)
		return []*NeedCommentLine{}
	}
	lines := strings.Split(text, "\n")
	willComments := filterRegex(re, lines, "Package ")
	return willComments
}
