/*
Copyright © 2023 dinglevin <dinglevin@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "appendstr",
	Short: "A small tool to append strings to a file, so it can change its MD5 hash",
	Long:  `A small tool to append strings to a file, so it can change its MD5 hash`,
	Run: func(cmd *cobra.Command, args []string) {
		appendContent, err := cmd.Flags().GetString("append")
		if err != nil {
			fmt.Println("Invalid append flags: ", err)
			return
		}

		AppendStrTo(args, appendContent)
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
	rootCmd.Flags().StringP("append", "a", "@@LEVIN@@", "Append string to the file")
}

