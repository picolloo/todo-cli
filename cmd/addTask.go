package cmd

import (
	"fmt"

	"github.com/picolloo/todo-cli/models"
	"github.com/picolloo/todo-cli/services"
	"github.com/picolloo/todo-cli/storage"
	"github.com/spf13/cobra"
)

var description string
var addTaskCmd = &cobra.Command{

	Use:   "add-task",
	Short: "Creates a new task.",
	Long:  "Creates a new task.",
	Run: func(cmd *cobra.Command, args []string) {
		db := storage.NewDBConnection()
		taskService := services.NewTaskService(db)
		task := models.NewTask(description, models.WAITING)
		err := taskService.SaveTask(task)
		if err != nil {
			fmt.Printf("%+v", task)
			fmt.Println(err.Error())
		}
	},
}

func init() {
	addTaskCmd.Flags().StringVarP(&description, "description", "d", "", "The description of the task")
	addTaskCmd.MarkFlagRequired("description")
}
