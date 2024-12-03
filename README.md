# codefile

`codefile` is a Go library for detecting the programming language of a given file. It uses content-based detection with weighted keyword matching, ensuring robust and accurate identification, even for files without extensions.

---

[![Go Reference](https://pkg.go.dev/badge/github.com/Agent-Hellboy/codefile.svg)](https://pkg.go.dev/github.com/Agent-Hellboy/codefile)
[![Go Report Card](https://goreportcard.com/badge/github.com/Agent-Hellboy/codefile)](https://goreportcard.com/report/github.com/Agent-Hellboy/codefile)
[![codecov](https://codecov.io/gh/Agent-Hellboy/codefile/branch/main/graph/badge.svg)](https://codecov.io/gh/Agent-Hellboy/codefile)

## Features

- **Content-Based Detection**: 
  - Detects programming languages by inspecting file content for unique constructs and patterns.
- **Weighted Scoring**: 
  - Each language feature is assigned a weight to improve detection accuracy.
- **Efficient Scanning**:
  - Only inspects the first 20 lines of a file for optimal performance.

---

## Installation

Install the package using `go get`:

```bash
go get github.com/Agent-Hellboy/codefile
```

## Usage
Basic Language Detection
Detect the programming language of a file:

```go
package main

import (
	"fmt"
	"github.com/Agent-Hellboy/codefile"
)

func main() {
	filePath := "example.py"
	language, ok := codefile.DetectCodeFileType(filePath)
	if ok {
		fmt.Printf("The language of the file is: %s\n", language)
	} else {
		fmt.Println("Language could not be detected.")
	}
}

The language of the file is: Go
```

### Supported Languages
The library supports the following programming languages out of the box:

- Python
- Go
- C++
- Java
- JavaScript
- TypeScript
- Shell
