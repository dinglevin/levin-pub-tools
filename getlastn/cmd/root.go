package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "getlastn",
	Short: "A small tool to get last n bytes from a file",
	Long:  `A small tool to get last n bytes from a file`,
	Run: func(cmd *cobra.Command, args []string) {
		num, err := cmd.Flags().GetInt64("num")
		if err != nil {
			fmt.Println("Invalid number flags: ", err)
			return
		}

		GetLastNFrom(args, num)
		return
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Int64P("num", "n", 9, "Number of last bytes to read")
}