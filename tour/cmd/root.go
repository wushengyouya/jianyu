package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

// 初始化命令
func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(sqlCmd)
}
