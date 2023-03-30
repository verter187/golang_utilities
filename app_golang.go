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

const refString = "test" //"Вася Петя, Вася! Леня Петя Женя Игорь"

func test() {
	words := strings.Fields(refString)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
	fmt.Println(Dedupe(words))
}

func main() {
	fmt.Println(Dedupe([]string{"a", "b", "a", "c"})) // [a b c]
	test()

	// для более мелких файлов
	fContent, err := ioutil.ReadFile("./temp/anna1.txt")
	if err != nil {
		panic(err)
	}

	replace_marks_str := "!,.?:;\"-[]()123456789"
	replace_marks := strings.SplitAfter(replace_marks_str, "")
	content := string(fContent)
	for i := 0; i < len(replace_marks); i++ {
		content = strings.ReplaceAll(content, replace_marks[i], "")
	}

	words := strings.Fields(content)

	fmt.Println(len(words))
	fmt.Println(len(Dedupe(words)))

}
