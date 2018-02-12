package main

import (
	"flag"
	"fmt"
	"unicode/utf8"
)

const (
	headerStart  = "＿"
	headerPiece  = "人"
	headerEnd    = "＿"
	bodyStart    = "＞"
	bodySpace    = "　"
	bodyEnd      = "＜"
	footerStart  = "￣"
	footerPiece1 = "Y"
	footerPiece2 = "^"
	footerEnd    = "￣"
)

func main() {
	help := flag.Bool("h", false, "output help message")
	input := flag.String("t", "suddenly-d", "input text be suddenly-death of ASCII art")
	flag.Parse()

	if *help == true {
		flag.PrintDefaults()
		return
	}

	creater := creater{Input: *input}
	creater.assemble()
	creater.output()
}

type creater struct {
	Input  string
	Header string
	Body   string
	Footer string
}

func (c *creater) assemble() {
	asciiCount, multiCount := 0, 0
	for _, t := range c.Input {
		if isASCII(string(t)) {
			asciiCount++
		} else {
			multiCount++
		}
	}

	bodyLength := (asciiCount / 2) + multiCount + 2

	headerBody, footerBody := "", ""
	for i := 0; i < bodyLength; i++ {
		headerBody += headerPiece
	}
	c.Header = headerStart + headerBody + headerEnd

	for i := 0; i < bodyLength; i++ {
		footerPiece := footerPiece1 + footerPiece2
		footerBody += footerPiece
	}
	footerBody += footerPiece1
	c.Footer = footerStart + footerBody + footerEnd
	c.Body = bodyStart + bodySpace + c.Input + bodySpace + bodyEnd
}

func (c *creater) output() {
	fmt.Println(c.Header)
	fmt.Println(c.Body)
	fmt.Println(c.Footer)
}

func isASCII(text string) bool {
	return utf8.ValidString(text) && utf8.RuneCountInString(text) == len(text)
}
