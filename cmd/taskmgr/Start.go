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

	// 3. register it
	rootCmd.AddCommand(addCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
