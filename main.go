package main

import (
	"flag"
	"fmt"
    "strings"

	"github.com/mdhenriques/go-git-scan/scanner"
)

func main() {

	scanPath := flag.String("path", ".", "Directory path to scan")

	flag.Parse()

	fmt.Println("Credential Scanner v0.1")
	fmt.Printf("Scanning directory %s\n\n", *scanPath)

	// Create a scanner
	s := scanner.NewScanner()

	// Scan current directory
	findings, err := s.ScanDirectory(*scanPath)

	if err != nil {
		fmt.Printf("Error scanning: %v\n", err)
		return
	}

	// Print results
    printResults(findings)
}

func printResults(findings []scanner.Finding) {
	if len(findings) == 0 {
		fmt.Println("No credentials found!")
		return
	}

	fmt.Printf("Found %d potential credentials:\n\n", len(findings))

	critical := 0
	high := 0
	medium := 0

	for _, finding := range findings {
		switch finding.Severity {
		case "CRITICAL":
			critical++
		case "HIGH":
			high++
		case "MEDIUM":
			medium++

		}
	}

    if critical > 0 {
        fmt.Printf("CRITICAL: %d\n", critical)
    }
    if high > 0 {
        fmt.Printf("HIGH: %d\n", high)
    }
    if medium > 0 {
        fmt.Printf(" MEDIUM: %d\n", medium)
    }

    fmt.Println("\nDetails:")
    fmt.Println(strings.Repeat("-", 80))

	for i, finding := range findings {
		fmt.Printf("\n[%d] %s\n", i+1, finding.PatternName)
		fmt.Printf("    Severity: %s\n", finding.Severity)
		fmt.Printf("    Location: %s:%d\n", finding.FilePath, finding.LineNumber)
		fmt.Printf("    Line: %s\n", finding.LineContent)
	}

	fmt.Println("\n" + strings.Repeat("─", 80))
	fmt.Printf("\n💡 Tip: Review these findings and remove sensitive data from your code.\n")
}
