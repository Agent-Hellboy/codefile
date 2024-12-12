package codefile

import (
	"io/ioutil"
	"os"
	"testing"
)

// Helper function to create a temporary file with sample content
func createTempFile(t *testing.T, content string) string {
	t.Helper()
	tmpFile, err := ioutil.TempFile("", "codefile_test_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer tmpFile.Close()

	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	return tmpFile.Name()
}

// TestDetectCodeFileType tests the language detection functionality
func TestDetectCodeFileType(t *testing.T) {
	tests := []struct {
		name         string
		content      string
		expectedLang string
		expectedOk   bool
	}{
		{
			name: "Python file",
			content: `
				def hello_world():
					print("Hello, world!")
				class Test:
					def __init__(self):
						pass
			`,
			expectedLang: "Python",
			expectedOk:   true,
		},
		{
			name: "Go file",
			content: `
				package main

				import "fmt"

				func main() {
					fmt.Println("Hello, world!")
				}
			`,
			expectedLang: "Go",
			expectedOk:   true,
		},
		{
			name: "C++ file",
			content: `
				#include <iostream>
				using namespace std;

				int main() {
					cout << "Hello, world!" << endl;
					return 0;
				}
			`,
			expectedLang: "C++",
			expectedOk:   true,
		},
		{
			name: "JavaScript file",
			content: `
				function helloWorld() {
					console.log("Hello, world!");
				}
				const x = 10;
			`,
			expectedLang: "JavaScript",
			expectedOk:   true,
		},
		{
			name:         "Empty file",
			content:      "",
			expectedLang: "",
			expectedOk:   false,
		},
		{
			name: "Ambiguous content",
			content: `
				# A comment that matches multiple languages
				print("Hello, ambiguous world!")
			`,
			expectedLang: "",
			expectedOk:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile := createTempFile(t, tt.content)
			defer os.Remove(tmpFile)

			lang, ok := DetectCodeFileType(tmpFile)
			if lang != tt.expectedLang || ok != tt.expectedOk {
				t.Errorf("DetectCodeFileType(%s) = (%q, %v), want (%q, %v)",
					tt.name, lang, ok, tt.expectedLang, tt.expectedOk)
			}
		})
	}
}
