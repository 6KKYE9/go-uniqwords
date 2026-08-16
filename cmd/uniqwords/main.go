package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"uniqwords"
)

func main() {
	countMode := false
	for _, a := range os.Args[1:] {
		if a == "-count" {
			countMode = true
		}
	}
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	var sb strings.Builder
	for sc.Scan() {
		sb.WriteString(sc.Text())
		sb.WriteString(" ")
	}
	text := strings.TrimSpace(sb.String())
	if countMode {
		for _, wc := range uniqwords.Count(text) {
			fmt.Printf("%s %d\n", wc.Word, wc.Count)
		}
	} else {
		for _, w := range uniqwords.Unique(text) {
			fmt.Println(w)
		}
	}
}
