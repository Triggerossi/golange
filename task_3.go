package main

import (
	"fmt"
	"strings"
)

func encryptWord(word string) string {
	if len(word) <= 1 {
		return word
	}

	firstChar := string(word[0])

	rest := word[1:]
	reversed := reverseString(rest)

	return firstChar + reversed
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func encryptPhrase(phrase string) string {
	words := strings.Fields(phrase)

	encryptedWords := make([]string, len(words))
	for i, word := range words {
		encryptedWords[i] = encryptWord(word)
	}

	return strings.Join(encryptedWords, " ")
}

func main() {
	testPhrases := []string{
		"Pepe Schnele is a legend",
		"Hello World",
		"Go is awesome",
		"I love programming",
		"Crypto culture research",
		"A",
		"",
	}

	for i, phrase := range testPhrases {
		fmt.Printf("тест #%d:\n", i+1)
		fmt.Printf("исходная фраза:  \"%s\"\n", phrase)

		if phrase == "" {
			fmt.Println("зашифрованная фраза: \"\" (пустая строка)")
		} else {
			encrypted := encryptPhrase(phrase)
			fmt.Printf("зашифрованная фраза: \"%s\"\n", encrypted)
		}

		if i == 0 {
			fmt.Println("Разбивка на слова:", strings.Fields(phrase))
			words := strings.Fields(phrase)
			for _, word := range words {
				fmt.Printf("  '%s' -> '%s'\n", word, encryptWord(word))
			}
		}

	}

	examples := []struct {
		word     string
		expected string
	}{
		{"Pepe", "Pepe"},
		{"Schnele", "Selenhc"},
		{"is", "is"},
		{"a", "a"},
	}

	for _, example := range examples {
		result := encryptWord(example.word)
		fmt.Printf("%s '%s' -> '%s'\n",
			example.word, result, example.expected)
	}
}
