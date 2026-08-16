package uniqwords

import "strings"

// Count 按出现顺序统计词频（词用空白切分），返回保序去重后的列表，每个带计数。
func Count(text string) []WordCount {
	seen := map[string]int{}
	order := []string{}
	for _, w := range strings.Fields(text) {
		if _, ok := seen[w]; !ok {
			order = append(order, w)
		}
		seen[w]++
	}
	out := make([]WordCount, 0, len(order))
	for _, w := range order {
		out = append(out, WordCount{Word: w, Count: seen[w]})
	}
	return out
}

// Unique 返回保序去重后的词列表（计数都 >=1，重复词只留第一次出现的位置）。
func Unique(text string) []string {
	seen := map[string]bool{}
	var out []string
	for _, w := range strings.Fields(text) {
		if !seen[w] {
			seen[w] = true
			out = append(out, w)
		}
	}
	return out
}

type WordCount struct {
	Word  string
	Count int
}
