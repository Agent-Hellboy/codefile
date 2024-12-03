package codefile

import (
	"os"
	"testing"
)

func TestDetectCodeFileType(t *testing.T) {
	// Setup: Create temporary test files with representative content
	testFiles := map[string]string{
		"main.py":       "def hello():\n    print('Hello, World!')\nimport os\n",
		"main.go":       "package main\nfunc main() {\n    println(\"Hello, World!\")\n}",
		"main.cpp":      "#include <iostream>\nint main() {\n    std::cout << \"Hello, World!\";\n    return 0;\n}",
		"script.sh":     "#!/bin/bash\necho 'Hello, World!'\n",
		"unknown.txt":   "Some random text\nWith no specific language features\n",
		"java_sample":   "class HelloWorld {\n    public static void main(String[] args) {\n        System.out.println(\"Hello, World!\");\n    }\n}",
		"javascript.js": "function greet() {\n    console.log('Hello, World!');\n}\n",
	}

	// Create files
	for name, content := range testFiles {
		err := os.WriteFile(name, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", name, err)
		}
	}

	// Cleanup after tests
	defer func() {
		for name := range testFiles {
			_ = os.Remove(name)
		}
	}()

	// Test cases
	testCases := []struct {
		fileName string
		expected string
	}{
		{"main.py", "Python"},
		{"main.go", "Go"},
		{"main.cpp", "C++"},
		{"script.sh", "Shell"},
		{"unknown.txt", ""},
		{"java_sample", "Java"},
		{"javascript.js", "JavaScript"},
	}

	for _, tc := range testCases {
		t.Run(tc.fileName, func(t *testing.T) {
			result, ok := DetectCodeFileType(tc.fileName)
			if ok && result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
			if !ok && tc.expected != "" {
				t.Errorf("Expected %s, but detection failed", tc.expected)
			}
		})
	}
}
