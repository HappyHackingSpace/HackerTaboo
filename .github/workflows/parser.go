package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Keyword struct {
	Keyword             string   `json:"keyword"`
	Category            string   `json:"category"`
	TabooWords          []string `json:"taboo_words"`
	SourceLanguage      string   `json:"source_language"`
	DestinationLanguage string   `json:"destination_language"`
	Description         string   `json:"description"`
}

func extractValue(issueBody, label string) string {
	re := regexp.MustCompile(fmt.Sprintf(`(?m)^### %s\s*\n([^#]+)`, label))
	match := re.FindStringSubmatch(issueBody)
	if len(match) > 1 {
		return strings.TrimSpace(match[1])
	}
	return ""
}

func extractList(issueBody, label string) []string {
	raw := extractValue(issueBody, label)
	if raw == "" {
		return []string{}
	}
	return strings.Split(raw, ", ")
}

func isDuplicate(entry Keyword, existing []Keyword) bool {
	for _, k := range existing {
		if k.Keyword == entry.Keyword &&
			k.Category == entry.Category &&
			strings.Join(k.TabooWords, ",") == strings.Join(entry.TabooWords, ",") &&
			k.SourceLanguage == entry.SourceLanguage &&
			k.DestinationLanguage == entry.DestinationLanguage &&
			k.Description == entry.Description {
			return true
		}
	}
	return false
}

func main() {
	fileBody, err := os.ReadFile("issue_body.txt")
	if err != nil {
		fmt.Println("Error reading issue body:", err)
		os.Exit(1)
	}
	issueBody := string(fileBody)

	data := Keyword{
		Keyword:             extractValue(issueBody, "Keyword"),
		Category:            extractValue(issueBody, "Category"),
		TabooWords:          extractList(issueBody, "Taboo Words"),
		SourceLanguage:      extractValue(issueBody, "Source Language"),
		DestinationLanguage: extractValue(issueBody, "Destination Language"),
		Description:         extractValue(issueBody, "Description"),
	}

	if data.Keyword == "" || data.SourceLanguage == "" || data.DestinationLanguage == "" {
		fmt.Println("Incomplete data found. Skipping update.")
		return
	}

	file, err := os.OpenFile("cards.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var jsonData map[string][]Keyword
	fileData, _ := io.ReadAll(file)
	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &jsonData)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			os.Exit(1)
		}
	} else {
		jsonData = map[string][]Keyword{"keywords": {}}
	}

	if !isDuplicate(data, jsonData["keywords"]) {
		jsonData["keywords"] = append(jsonData["keywords"], data)
		newFileData, _ := json.MarshalIndent(jsonData, "", "  ")
		os.WriteFile("cards.json", newFileData, 0644)
	} else {
		fmt.Println("Duplicate entry detected. Skipping update.")
	}
}
