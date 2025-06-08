package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func dockerAvailable() bool {
	cmd := exec.Command("docker", "info")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

var pullCmd = &cobra.Command{
	Use:   "pull <image_name>",
	Short: "Import a local Docker image from tar and run it",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !dockerAvailable() {
			fmt.Println("Docker is not running or not installed. Please start Docker and try again.")
			return
		}

		imageName := args[0]
		tarFile := fmt.Sprintf("%s.tar", imageName)

		fmt.Printf("Loading Docker image from %s...\n", tarFile)
		loadCmd := exec.Command("docker", "load", "-i", tarFile)
		loadOut, err := loadCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to load image: %s\n", string(loadOut))
			return
		}
		fmt.Printf("Image loaded: %s\n", string(loadOut))

		// docker run:

		// fmt.Printf("Running container from image '%s'...\n", imageName)
		// runCmd := exec.Command("docker", "run", "-d", "--name", imageName, imageName)
		// runOut, err := runCmd.CombinedOutput()
		// if err != nil {
		// 	fmt.Printf("Failed to run container: %s\n", string(runOut))
		// 	return
		// }
		// fmt.Printf("Container started with ID: %s\n", strings.TrimSpace(string(runOut)))
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
