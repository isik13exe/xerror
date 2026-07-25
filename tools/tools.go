package tools

import (
	"os"
	"path/filepath"
	"strings"
)

func Logo() string {
	data, err := os.ReadFile("./tools/logo.ascii")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func Lang(path string) string {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".go":
		return "go"
	case ".py":
		return "python"
	case ".js", ".mjs", ".cjs":
		return "javascript"
	case ".ts", ".tsx":
		return "typescript"
	case ".rs":
		return "rust"
	case ".c", ".h":
		return "c"
	case ".cpp", ".cc", ".hpp", ".cxx":
		return "cpp"
	case ".java":
		return "java"
	case ".md":
		return "markdown"
	case ".json":
		return "json"
	case ".yaml", ".yml":
		return "yaml"
	case ".html", ".htm":
		return "html"
	case ".css":
		return "css"
	case ".sh", ".bash":
		return "bash"
	case ".sql":
		return "sql"
	default:
		return "" // no highlighting
	}
}
