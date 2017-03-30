package main

import (
	"fmt"
	"testing"
)

func BenchmarkSubDistance(b *testing.B) {
	wordHashes := make(map[string]map[string]struct{})
	for _, word := range wordsToTest {
		wordHashes[word] = hashWord(word)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, searchWord := range searchWords {
			searchWordHash := hashWord(searchWord)
			for _, word := range wordsToTest {
				compareHash(searchWordHash, wordHashes[word])
			}
		}
	}
}

func BenchmarkSubDistance2(b *testing.B) {
	wordHashes := make(map[string]map[string]struct{})
	for _, word := range wordsToTest {
		wordHashes[word] = hashWord(word)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, searchWord := range searchWords {
			searchWordHash := hashWord(searchWord)
			for _, word := range wordsToTest {
				foo := wordHashes[word]
				compareHashIfBetter(&searchWordHash, &foo, 400, len(word)+len(searchWord))
			}
		}
	}
}

func BenchmarkFindBestWord(b *testing.B) {
	wordHashes := make(map[string]map[string]struct{})
	for _, word := range wordsToTest {
		wordHashes[word] = hashWord(word)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, searchWord := range searchWords {
			findBestWordWithWordHash(searchWord, wordsToTest, wordHashes)
		}
	}
}

func BenchmarkFindBestWord2(b *testing.B) {
	wordHashes := make(map[string]map[string]struct{})
	for _, word := range wordsToTest {
		wordHashes[word] = hashWord(word)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, searchWord := range searchWords {
			findBestWordWithWordHashIfBetter(searchWord, wordsToTest, wordHashes)
		}
	}
}

func TestGeneral(t *testing.T) {
	wordHashes := make(map[string]map[string]struct{})
	for _, word := range wordsToTest {
		wordHashes[word] = hashWord(word)
	}
	fmt.Println("\nNew algorithm:\n")
	for _, searchWord := range searchWords {
		fmt.Printf("'%s'\tmatched\t'%s'\n", searchWord, findBestWordWithWordHashIfBetter(searchWord, wordsToTest, wordHashes))
	}
}
