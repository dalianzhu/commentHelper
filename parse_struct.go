package main

import (
    "log"
    "regexp"
    "strings"
)

// ParseStruct ...
func ParseStruct(text string) []*WillCommentLine {
    re, err := regexp.Compile(`^type\s+([A-Z][a-z0-9A-Z]*)\s+struct\s*{|^type\s+([A-Z][a-z0-9A-Z]+)\s+interface\s*{`)
    if err != nil {
        log.Println(err)
        return []*WillCommentLine{}
    }
    lines := strings.Split(text, "\n")
    willComments := filterRegex(re, lines)
    return willComments
}
