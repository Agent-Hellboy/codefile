package codefile

import (
	"bufio"
	"os"
	"strings"
)

// Keyword represents a content pattern with an associated weight
type Keyword struct {
	Pattern string
	Weight  int
}

// Language represents a programming language with associated keywords
type Language struct {
	Name     string
	Keywords []Keyword
}

// Predefined language rules
var languages = []Language{
	{
		Name: "Python",
		Keywords: []Keyword{
			{"def ", 3}, {"class ", 3}, {"import ", 2}, {"print(", 2},
			{"self.", 1}, {"__init__", 1}, {"from ", 2},
		},
	},
	{
		Name: "Go",
		Keywords: []Keyword{
			{"package ", 3}, {"func ", 3}, {"import ", 2}, {"var ", 1},
		},
	},
	{
		Name: "C++",
		Keywords: []Keyword{
			{"#include", 3}, {"int main", 3}, {"using namespace", 2},
			{"class ", 2}, {"cout", 1}, {"std::", 1},
		},
	},
	{
		Name: "Java",
		Keywords: []Keyword{
			{"class ", 3}, {"public static void main", 3}, {"import ", 2},
			{"System.out.println", 2}, {"new ", 1},
		},
	},
	{
		Name: "JavaScript",
		Keywords: []Keyword{
			{"function ", 3}, {"const ", 2}, {"let ", 2}, {"var ", 1},
			{"console.log", 2},
		},
	},
	{
		Name: "Shell",
		Keywords: []Keyword{
			{"#!/bin/bash", 3}, {"#!/usr/bin/env bash", 3}, {"echo ", 2},
		},
	},
}

// DetectCodeFileType detects the programming language of a file
func DetectCodeFileType(filePath string) (string, bool) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	languageScores := make(map[string]int)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineCount++

		for _, lang := range languages {
			for _, keyword := range lang.Keywords {
				if strings.Contains(line, keyword.Pattern) {
					languageScores[lang.Name] += keyword.Weight
				}
			}
		}

		if lineCount > 20 {
			break
		}
	}

	var bestMatch string
	highestScore := 0
	for lang, score := range languageScores {
		if score > highestScore {
			highestScore = score
			bestMatch = lang
		}
	}

	if highestScore > 3 {
		return bestMatch, true
	}

	return "", false
}
