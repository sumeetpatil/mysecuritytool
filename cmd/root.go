package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mysecuritytool",
	Short: "My Security Tool",
	Long:  `My Security Tool contains mutiple security tools like fuzzing`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
