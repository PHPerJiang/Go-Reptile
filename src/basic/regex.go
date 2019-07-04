package basic

import (
	"fmt"
	"regexp"
)

func StringRegex(txt string) {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	math  := re.FindAllStringSubmatch(txt,-1)
	fmt.Println(math)
}
