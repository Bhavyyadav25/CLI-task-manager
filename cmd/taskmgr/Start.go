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
	pdfE := usecase.NewPDFExporter(repo)

	addCmd := &cobra.Command{
		Use:   "add [description]",
		Short: "Create a new task",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			desc := strings.Join(args, " ")
			task, err := svc.AddTask(desc)
			if err != nil {
				fmt.Fprintf(os.Stderr, "⚠️  No task added: %s\n", err.Error())
				return nil
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
				fmt.Fprintf(os.Stderr, "⚠️  No task found")
				return nil
			}
			fmt.Printf("%-4s  %-36s  %-30s  %-30s  %-30s  %-30s\n", "No.", "ID", "Description", "Done", "Created At", "Updated At")
			fmt.Println(strings.Repeat("-", 110))
			for i, task := range tasks {
				fmt.Printf(
					"%02d.   %-36s  %-30s  %-30v  %-30s  %-30s\n",
					i+1,
					task.ID,
					task.Description,
					task.Done,
					task.CreatedAt.Format("02-01-2006 15:04"),
					task.UpdateAt.Format("02-01-2006 15:04"),
				)
			}
			return nil
		},
	}

	updateCmd := &cobra.Command{
		Use:   "done [description]",
		Short: "Mark the task done",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			desc := strings.Join(args, " ")
			task, err := svc.MarkDone(desc)
			if err != nil {
				fmt.Fprintf(os.Stderr, "⚠️  Cannot mark task done: %s\n", err.Error())
				return nil
			}
			fmt.Printf(
				"%-4s  %-36s  %-30s  %-6s  %-16s  %-16s\n",
				"No.", "ID", "Description", "Done", "Created At", "Updated At",
			)
			fmt.Println(strings.Repeat("-", 112))

			fmt.Printf(
				"01.   %-36s  %-30s  %-6v  %-16s  %-16s\n",
				task.ID,
				task.Description,
				task.Done,
				task.CreatedAt.Format("02-01-2006 15:04"),
				task.UpdateAt.Format("02-01-2006 15:04"),
			)
			return nil
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "del [description]",
		Short: "Deleting a task",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			desc := strings.Join(args, " ")
			err := svc.DeleteTask(desc)
			if err != nil {
				fmt.Fprintf(os.Stderr, "⚠️  No task deleted: %s\n", err.Error())
				return nil
			}
			fmt.Printf("✔️  Deleted task with id: %s\n", desc)
			return nil
		},
	}

	pdfCmd := &cobra.Command{
		Use:   "export [description]",
		Short: "Export task in pdf format",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			desc := strings.Join(args, " ")
			err := pdfE.Export(desc)
			if err != nil {
				fmt.Fprintf(os.Stderr, "⚠️  No task exported: %s\n", err.Error())
				return nil
			}
			fmt.Printf("✔️  Exported task: %s\n", desc)
			return nil
		},
	}

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(pdfCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
