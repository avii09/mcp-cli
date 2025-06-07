package model

type Repository struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Run struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
	Port    int      `json:"port"`
}

type MCP struct {
	Name        string     `json:"name"`
	Version     string     `json:"version"`
	Description string     `json:"description"`
	Author      string     `json:"author"`
	License     string     `json:"license"`
	Keywords    []string   `json:"keywords"`
	Repository  Repository `json:"repository"`
	Run         Run        `json:"run"`
}
