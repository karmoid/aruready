package main

import "fmt"

type (
	LineContext struct {
		cmd string
		tag []string
	}
)

func (l LineContext) Tag() []string {
	fmt.Println("hello")
	return l.tag
}

func (l LineContext) GetTag(tag int) string {
	fmt.Println("hello")
	return ""
}

func (l LineContext) FindTag(tag string) int {
	fmt.Println("hello")
	return -1
}

func CliAnalyze(cmd string) (LineContext, error) {
	return LineContext{}, nil
}
