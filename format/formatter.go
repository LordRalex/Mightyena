package format

import "regexp"

var BBCodeBold = regexp.MustCompile(`\[/?b\]`)
var BBCodeUnderline = regexp.MustCompile(`\[/?u\]`)
var BBCodeItalic = regexp.MustCompile(`\[/?i\]`)

const IrcBold = string(0x02)
const IrcItalics = string(0x1D)
const IrcUnderline = string(0x1F)
const IrcPlain = string(0x0F)

func ParseFromBBCode(input string) string {
	result := input

	result = BBCodeBold.ReplaceAllString(result, IrcBold)
	result = BBCodeUnderline.ReplaceAllString(result, IrcUnderline)
	result = BBCodeItalic.ReplaceAllString(result, IrcItalics)

	return result
}
