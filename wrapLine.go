package main

import (
	"log"
	"strings"

	stackLib "github.com/golang-collections/collections/stack"
)

func wrapLine(text string) string {
	arr := strings.Split(text, "\n")
	lines := len(arr)
	stack := stackLib.New()

	for lineNo := 0; lineNo < lines; {
		i := 0
		line := arr[lineNo]
		end := len(line)
		content := []byte(line)
		isOverLength := false
		if len(line) > 115 {
			isOverLength = true
		}
		inComment := false
		for {
			if i >= end {
				break
			}
			// log.Printf("content:%v, i:%v", content, i)
			currentChar := content[i]
			// 判定是不是在注释中
			if currentChar == '/' && i != 0 {
				if content[i-1] == '/' {
					inComment = true
				}
			}

			if currentChar == '`' || currentChar == '"' {
				if stack.Len() != 0 && stack.Peek() == currentChar {
					stack.Pop()
				} else {
					stack.Push(currentChar)
				}
			}

			stackLen := stack.Len()
			if stackLen > 0 || inComment {
				//在字符串中啥也不干
			} else {
				if isOverLength && (currentChar == ',' || currentChar == '(') {
					log.Printf("%v\n", stack)
					// 插入一个\n
					splitArr := make([]string, 0)
					splitArr = append(splitArr, string(content[0:i+1]))
					splitArr = append(splitArr, string(content[i+1:]))
					arr[lineNo] = splitArr[0]
					arr = InsertSlice(arr, lineNo+1, splitArr[1])
					lines++
					break
				}
			}
			i++
		}
		lineNo++
	}

	return strings.Join(arr, "\n")
}

// Stack ...
type Stack struct {
	inner []byte
}

// Push ...
func (s *Stack) Push(v byte) {
	s.inner = append(s.inner, v) // Push
}

// Pop ...
func (s *Stack) Pop() byte {
	n := len(s.inner) - 1 // Top element
	v := s.inner[n]
	s.inner = s.inner[:n] // Pop
	return v
}

// Peek ...
func (s *Stack) Peek() byte {
	n := len(s.inner) - 1 // Top element
	v := s.inner[n]
	return v
}

// Len ...
func (s *Stack) Len() int {
	return len(s.inner)
}
