package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Dedupe[T comparable](arr []T) []T {
	m, uniq := make(map[T]struct{}), make([]T, 0, len(arr))
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v], uniq = struct{}{}, append(uniq, v)
		}
	}
	return uniq
}

func main() {
	// для более мелких файлов
	filename := "./temp/anna1.txt" //os.Args[1] //"./temp/anna1.txt"
	fContent, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	replace_marks_str := "!,.?:;\"-[]()123456789"
	replace_marks := strings.SplitAfter(replace_marks_str, "")
	content := string(fContent)

	for _, mark := range replace_marks {
		content = strings.ReplaceAll(content, mark, "")
	}

	words := strings.Fields(content)

	fmt.Println(len(words))
	fmt.Println(len(Dedupe(words)))
}
