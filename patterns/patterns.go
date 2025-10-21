package patterns

import "regexp"

type Pattern struct {
	Name        string
	Regex       *regexp.Regexp
	Description string
	Severity    string
}

// Compiled once when package loads
var defaultPatterns = []Pattern{
	{
		Name:        "AWS Access Key",
		Regex:       regexp.MustCompile(`AKIA[0-9A-Z]{16}`),
		Description: "AWS Access Key ID",
		Severity:    "HIGH",
	},
	{
		Name:        "GitHub Token",
		Regex:       regexp.MustCompile(`gh[pousr]_[A-Za-z0-9]{36,}`),
		Description: "GitHub Personal Access Token",
		Severity:    "HIGH",
	},
	{
		Name:        "Private Key",
		Regex:       regexp.MustCompile(`-----BEGIN\s+(?:RSA|DSA|EC|OPENSSH)?\s*PRIVATE KEY-----`),
		Description: "Private SSH/SSL Key",
		Severity:    "CRITICAL",
	},
	{
		Name:        "Generic Password",
		Regex:       regexp.MustCompile(`(?i)password\s*=\s*["'][^"']{3,}["']`),
		Description: "Hardcoded password in code",
		Severity:    "MEDIUM",
	},
}

// GetPatterns returns the default pattern set
func GetPatterns() []Pattern {
	return defaultPatterns
}
