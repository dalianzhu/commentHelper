package main

import (
	"log"
	"regexp"
	"strings"
)

// ParseStruct ...
func ParseStruct(text string) []*NeedCommentLine {
	re, err := regexp.Compile(`^type\s+([a-z0-9A-Z]*)\s+struct\s*{|^type\s+([a-z0-9A-Z]+)\s+interface\s*{`)
	if err != nil {
		log.Println(err)
		return []*NeedCommentLine{}
	}
	lines := strings.Split(text, "\n")
	willComments := filterRegex(re, lines)
	return willComments
}
