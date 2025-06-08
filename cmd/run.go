package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func dockerAvailable() bool {
	cmd := exec.Command("docker", "info")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

var runCmd = &cobra.Command{
	Use:   "run <image_name>",
	Short: "Run a Docker container from a loaded image",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !dockerAvailable() {
			fmt.Println("Docker is not running or not installed. Please start Docker and try again.")
			return
		}

		imageName := args[0]

		fmt.Printf("Running container from image '%s'...\n", imageName)
		runCmd := exec.Command("docker", "run", "-d", "--name", imageName, imageName)
		runOut, err := runCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to run container: %s\n", string(runOut))
			return
		}
		fmt.Printf("Container started with ID: %s\n", strings.TrimSpace(string(runOut)))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
