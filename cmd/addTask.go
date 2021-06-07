package cmd

import (
	"github.com/spf13/cobra"
)

var description string
var addTaskCmd = &cobra.Command{

	Use:   "add-task",
	Short: "zzzzzz",
	Long:  "Creates a new task.",
	Run: func(cmd *cobra.Command, args []string) {
		task := NewTask(description, WAITING)
		Tasks = append(Tasks, task)
	},
}

func init() {
	addTaskCmd.Flags().StringVarP(&description, "description", "d", "", "The description of the task")
	addTaskCmd.MarkFlagRequired("description")
}
