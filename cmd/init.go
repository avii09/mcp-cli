/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"bufio"
	"cli/model"
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var yesFlag bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new mcp.json file",
	Run: func(cmd *cobra.Command, args []string) {
		var mcp model.MCP
		if yesFlag {
			cwd, err := os.Getwd()
			projectName := "my-project"
			if err == nil {
				projectName = filepath.Base(cwd)
			}
			mcp = model.MCP{
				Name:        projectName,
				Version:     "1.0.0",
				Description: "",
				Author:      "",
				License:     "",
				Keywords:    []string{},
				Repository: model.Repository{
					Type: "git",
					URL:  "",
				},
				Run: model.Run{
					Command: "",
					Args:    []string{""},
					Port:    5050,
				},
			}
		} else {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Project name: ")
			mcp.Name = readLine(reader)

			fmt.Print("Version: ")
			mcp.Version = readLine(reader)

			fmt.Print("Description: ")
			mcp.Description = readLine(reader)

			fmt.Print("Author: ")
			mcp.Author = readLine(reader)

			fmt.Print("License: ")
			mcp.License = readLine(reader)

			fmt.Print("Keywords: ")
			keywords := readLine(reader)
			mcp.Keywords = strings.Split(keywords, ",")

			fmt.Print("Repository type: ")
			mcp.Repository.Type = readLine(reader)

			fmt.Print("Repository URL: ")
			mcp.Repository.URL = readLine(reader)

			fmt.Print("Run command: ")
			mcp.Run.Command = readLine(reader)

			fmt.Print("Run arguments: ")
			argsStr := readLine(reader)
			mcp.Run.Args = strings.Split(argsStr, ",")

			fmt.Print("Port: ")
			fmt.Scanf("%d\n", &mcp.Run.Port)
		}

		file, err := os.Create("mcp.json")
		if err != nil {
			fmt.Println("Error creating mcp.json:", err)
			return
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(mcp); err != nil {
			fmt.Println("Error writing mcp.json:", err)
			return
		}

		fmt.Println("✅ mcp.json created successfully!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&yesFlag, "yes", "y", false, "Use default values")
}

func readLine(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
