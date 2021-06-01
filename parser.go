package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

// NeedCommentLine 这一行需要加注释
type NeedCommentLine struct {
	OriginLineNo int
	Name         string
}

// AddCommentToText ...
func AddCommentToText(text string) string {
	originTextLines := strings.Split(text, "\n")

	// 需要加注释的一行已经有注释或者在注释中，处理这种情况
	inCommentMap := make(map[int]bool)
	isInComment := false
	for i, line := range originTextLines {
		if isInComment {
			inCommentMap[i] = true
		}
		trimLine := strings.Trim(line, " ")
		if strings.HasPrefix(trimLine, "//") {
			log.Printf("line:%v is in comment", i)
			inCommentMap[i] = true
		}
		if strings.HasPrefix(trimLine, "/*") {
			// 进入大注释
			inCommentMap[i] = true
			isInComment = true
		}
		if strings.HasSuffix(trimLine, "*/") {
			inCommentMap[i] = true
			isInComment = false
		}
	}

	// 获取所有需要加注释的line
	needCommentLines := ParseFunc(text)
	needCommentLines = append(needCommentLines, ParseStruct(text)...)
	needCommentLines = append(needCommentLines, ParseType(text)...)
	sort.Slice(needCommentLines, func(i, j int) bool {
		return needCommentLines[i].OriginLineNo < needCommentLines[j].OriginLineNo
	})

	offset := 0
	for _, c := range needCommentLines {
		willInsertLine := c.OriginLineNo + offset
		log.Printf("willInsertLine:%v", willInsertLine)

		// 虽然它满足加注释的类型，但是在原文件中，它已经有注释了
		_, ok := inCommentMap[c.OriginLineNo-1]
		if !ok {
			comment := fmt.Sprintf("// %v ...", c.Name)
			originTextLines = InsertSlice(originTextLines, willInsertLine, comment)
			offset++
		}
	}

	return strings.Join(originTextLines, "\n")
}

// InsertSlice ...
func InsertSlice(ss []string, index int, inserted string) []string {
	rear := append([]string{}, ss[index:]...)
	ss = append(ss[0:index], inserted)
	ss = append(ss, rear...)
	return ss
}

// InsertSliceByte ...
func InsertSliceByte(ss []byte, index int, inserted byte) []byte {
	rear := append([]byte{}, ss[index:]...)
	ss = append(ss[0:index], inserted)
	ss = append(ss, rear...)
	return ss
}
