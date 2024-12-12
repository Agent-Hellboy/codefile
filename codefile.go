package codefile

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
	lines, err := ScanFile(filePath, 50) // Scanning up to 50 lines
	if err != nil {
		return "", false
	}

	scores := ScoreLines(lines, languages)

	var bestMatch string
	highestScore := 0
	for lang, score := range scores {
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
