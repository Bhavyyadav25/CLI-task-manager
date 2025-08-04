package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Bhavyyadav25/CLI-task-manager/internal/store"
	"github.com/Bhavyyadav25/CLI-task-manager/internal/usecase"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "taskmgr",
	Short: "CLI task manager",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	repo := store.NewFileRepo("internal/store/user_data.json")
	svc := usecase.NewTaskService(repo)

	addCmd := &cobra.Command{
		Use:   "add [description]",
		Short: "Create a new task",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			desc := strings.Join(args, " ")
			task, err := svc.AddTask(desc)
			if err != nil {
				return err
			}
			fmt.Printf("✔️  Added task %s: %s\n", task.ID, task.Description)
			return nil
		},
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := svc.ListTasks()
			if err != nil {
				return err
			}
			if len(tasks) == 0 {
				fmt.Println("No tasks found.")
				return nil
			}
			fmt.Printf("%-4s  %-36s  %-30s  %-30s  %-30s\n", "No.", "ID", "Description", "Created At", "Updated At")
			fmt.Println(strings.Repeat("-", 110))
			for i, task := range tasks {
				fmt.Printf(
					"%02d.   %-36s  %-30s  %-30s  %-30s\n",
					i+1,
					task.ID,
					task.Description,
					task.CreatedAt.Format("02-01-2006 15:04"),
					task.UpdateAt.Format("02-01-2006 15:04"),
				)
			}
			return nil
		},
	}

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
