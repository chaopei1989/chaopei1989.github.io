package main

import "fmt"

var (
	dict = []string{
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
)

func letterCombinationsRecursive(digits string, i int, curr string, res []string) []string {
	i++
	if i >= len(digits) {
		res = append(res, curr)
		return res
	}
	dictStr := dict[int32(digits[i]-'2')]
	for k := 0; k < len(dictStr); k++ {
		res = letterCombinationsRecursive(digits, i, curr + dictStr[k:k+1], res)
	}
	return res
}

// digits: "2345"
func letterCombinations(digits string) []string {
	res := make([]string, 0)
	dictStr := dict[int32(digits[0]-'2')]
	for i := 0; i < len(dictStr); i++ {
		res = letterCombinationsRecursive(digits, 0, dictStr[i:i+1], res)
	}
	return res
}

func main() {
	fmt.Println(letterCombinations("23"))
}
