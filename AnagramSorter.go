package main

import (
"fmt"
"sort"
"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, word := range strs {
		chars := strings.Split(word, "")
		sort.Strings(chars)                   // sort letters in word
		key := strings.Join(chars, "")        // key like "aet" for "eat"

		anagramMap[key] = append(anagramMap[key], word)
	}

	result := [][]string{}
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func main(){
	var choice string
	strs := []string{"eat", "bat", "tea", "tab"}
	fmt.Println("Anagram Grouping Service")
	for _, strings := range groupAnagrams(strs) {
		fmt.Println(strings)
	}

	fmt.Println("Would you like to add more strings? (yes/no)")
	fmt.Scanln(&choice)
	lowerChoice := strings.ToLower(choice)

	for lowerChoice == "yes" {
		var newStr string
		fmt.Println("Enter a string to add:")
		fmt.Scanln(&newStr)
		strs = append(strs, newStr)
		fmt.Println("Would you like to add another string? (yes/no)")
		fmt.Scanln(&choice)
		lowerChoice = strings.ToLower(choice)

	}
	if lowerChoice == "no" {
		grouped := groupAnagrams(strs)

		for _, group := range grouped {
			fmt.Println(group)
		}
	}

}

