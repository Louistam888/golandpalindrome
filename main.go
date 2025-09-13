package main

import (
	"fmt"
	"unicode"
)

func main() {
input := "Madam, I'm Adam"

normalized := ""
//remove non letter characters and convert to lowercase
for _, r := range input {
	if unicode.IsLetter(r) {
		normalized += string(unicode.ToLower(r))
	}
}


var isPalindrome bool = true

length := len(normalized)

for i:=0; i<length/2; i++ {
	if normalized[i] != normalized[length-1-i] {
		isPalindrome = false
	}
}

if isPalindrome {
	fmt.Println("It is a palinedrome")
} else {
	fmt.Println("It is not a palindrome")
}


}
