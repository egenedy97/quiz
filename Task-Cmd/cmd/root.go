package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Task",
	Short: "Task is a CLI task manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task called")
	},
}
