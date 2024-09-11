// challenge:write a pogram that reverses a string without using built in functions

package main

import "fmt"

func main() {
	s := "My name is nyagooh and i'm made up of great things"
	fmt.Println(ReverseString(s))
	fmt.Println(Split(s))
}
func ReverseString(word string) string {
	s := Split(word)
	str := ""
	for i := len(s) - 1; i >= 0; i-- {
		str += s[i] + " "
	}
	return str
}
func Split(s string) []string {
	// implement your own logic here
	word := ""
	result := []string{}
	for _, char := range s {
		if char == ' ' && word != ""{
			result = append(result, word)
			word = ""
	}else if char != ' '{
		word += string(char)
	}
}
	if word != "" {
		result = append(result, word)
	}
	return result
}
