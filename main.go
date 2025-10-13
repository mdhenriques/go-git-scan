package main

import (
    "fmt"
    "github.com/mdhenriques/go-git-scan/scanner"
)

func main() {
    fmt.Println("Credential Scanner v0.1")
    
    // Create a scanner
    s := scanner.NewScanner()
    
    // Scan current directory
    findings, err := s.ScanDirectory(".")
    
    if err != nil {
        fmt.Printf("Error scanning: %v\n", err)
        return
    }
    
    // Print results
    if len(findings) == 0 {
        fmt.Println("✓ No credentials found!")
    } else {
        fmt.Printf("⚠ Found %d potential credentials:\n\n", len(findings))
        
        for _, finding := range findings {
            fmt.Printf("[%s] %s\n", finding.Severity, finding.PatternName)
            fmt.Printf("  File: %s:%d\n", finding.FilePath, finding.LineNumber)
            fmt.Printf("  Line: %s\n\n", finding.LineContent)
        }
    }
}