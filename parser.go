package codefile

import "strings"

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

// ScoreLines calculates scores for each programming language based on the given lines
// TODO : make this more efficient , like detamining position of the keyword
// in the file then weight it by position, also add more keywords
func ScoreLines(lines []string, languages []Language) map[string]int {
	languageScores := make(map[string]int)

	for _, line := range lines {
		for _, lang := range languages {
			for _, keyword := range lang.Keywords {
				if strings.Contains(line, keyword.Pattern) {
					languageScores[lang.Name] += keyword.Weight
				}
			}
		}
	}

	return languageScores
}
