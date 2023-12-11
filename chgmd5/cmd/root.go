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
	Use:   "chgmd5",
	Short: "A small tool to change a file's MD5 hash",
	Long:  `A small tool to change a file's MD5 hash`,
	Run: func(cmd *cobra.Command, args []string) {
		appendStr, err := cmd.Flags().GetString("append")
		if err != nil {
			fmt.Println("Invalid append flags: ", err)
			return
		}

		chgmd5(args, appendStr)
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

