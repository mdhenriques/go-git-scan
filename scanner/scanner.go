package scanner

import (
	"bufio"
	"os"
	"path/filepath"
	"github.com/mdhenriques/go-git-scan/patterns"
)

type Finding struct {
	FilePath   string
	LineNumber int
	LineContent string
	PatternName string
	Severity string
}

type Scanner struct {
	patterns []patterns.Pattern
}

func NewScanner() *Scanner {
	return &Scanner{
		patterns: patterns.GetPatterns(),
	}
}

func (s *Scanner) ScanDirectory(rootPath string) ([]Finding, error) {
	var findings []Finding

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fileFindings, err := s.scanFile(path)
		if err != nil {
			return nil
		}

		findings = append(findings, fileFindings...)

		return nil
	})

	return findings, err
}

func (s *Scanner) scanFile(filePath string) ([]Finding, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var findings []Finding
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		for _, pattern := range s.patterns {
			if pattern.Regex.MatchString(line) {
				findings = append(findings, Finding{
					FilePath:	filePath,
					LineNumber: lineNumber,
					LineContent: line,
					PatternName: pattern.Name,
					Severity:    pattern.Severity,
				})
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return findings, nil
}