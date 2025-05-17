package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

const (
	baseURL = "https://github.com/github/gitignore"
	rawURL  = "https://raw.githubusercontent.com/github/gitignore/main/"
	fileExt = ".gitignore"
)

func main() {
	list := flag.Bool("list", false, "List all available templates from github/gitignore")
	output := flag.String("o", ".gitignore", "Output file name")
	flag.Parse()

	if *list {
		templates, err := fetchAvailableTemplates()
		if err != nil {
			fmt.Println("Error fetching template list:", err)
			os.Exit(1)
		}
		fmt.Println("📄 Available templates from github/gitignore:\n")
		for _, name := range templates {
			fmt.Println(" -", name)
		}
		return
	}

	if flag.NArg() < 1 {
		fmt.Println("Usage: gitignore <template>[,<template>...]")
		fmt.Println("Use --list to see all available templates.")
		os.Exit(1)
	}

	templateNames := strings.Split(flag.Arg(0), ",")
	var combined strings.Builder

	for _, name := range templateNames {
		template := strings.TrimSpace(name)
		url := rawURL + escapeTemplateName(template) + fileExt

		resp, err := http.Get(url)
		if err != nil || resp.StatusCode != 200 {
			fmt.Printf("⚠️  Failed to fetch: %s\n", url)
			continue
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		combined.WriteString("# ===== " + template + " =====\n")
		combined.Write(body)
		combined.WriteString("\n\n")
	}

	err := os.WriteFile(*output, []byte(combined.String()), 0644)
	if err != nil {
		fmt.Println("❌ Failed to write file:", err)
		os.Exit(1)
	}

	fmt.Printf("✅ %s generated.\n", *output)
}

func fetchAvailableTemplates() ([]string, error) {
	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)
	var templates []string
	rx := regexp.MustCompile(`(.+)\.gitignore`)

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		t := z.Token()
		if t.Type == html.StartTagToken && t.Data == "a" {
			for _, attr := range t.Attr {
				if attr.Key == "title" && strings.HasSuffix(attr.Val, ".gitignore") {
					if match := rx.FindStringSubmatch(attr.Val); match != nil {
						templates = append(templates, match[1])
					}
				}
			}
		}
	}
	return templates, nil
}

func escapeTemplateName(name string) string {
	// Capitalize first letter, handle spaces and special cases (e.g., "C++")
	safe := strings.ReplaceAll(name, " ", "%20")
	switch strings.ToLower(name) {
	case "c++":
		return "C%2B%2B"
	case "c":
		return "C"
	case "csharp":
		return "VisualStudio"
	default:
		return strings.Title(safe)
	}
}
