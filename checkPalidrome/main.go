//challenge:verify given number or string is a palidrome (reads same forwad as backward)

package main

import "fmt"

func main() {
	input := "A man"
	fmt.Println(IsPalindrome(input))
}
func IsPalindrome(input string) bool {
	str := Reverse(input)
word1 :=Split(str)
word2 := Split(input)
var ok bool
if len(word2) > 1 {
	for i:=0;i<len(word2);i++ {
		for j:=len(word1)-1;j>=0;j-- {
			if word2[i] != word1[j] {
                ok = false
            } else {
				ok = true
			}
		}
	}
} else {
	for i:=0;i<len(input);i++ {
		for j:=len(str)-1;j>=0;j--{
			if string(input[i]) != string(str[j]) {
				ok=false
            }else{
				ok=true
			}
		}
}
}
return ok
}

func Reverse(input string) string {
	word := Split(input)
	str := ""
	if len(word) > 1 {
		for i := len(word) - 1; i >= 0; i-- {
			if word[i] == word[0] {
				str += word[i]
			} else {
				str+=word[i]+" "
			}
		}
	} else {
		for _, char := range word {
			for i := len(char) - 1; i >= 0; i-- {
				str += string(char[i])
			}
		}
	}
	return str
}
func Split(s string) []string {
	word := ""
	result := []string{}

	for _, char := range s {
		if char != ' ' {
			word += string(char)
		}
		if char == ' ' && word != "" {
			result = append(result, word)
			word = ""
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
