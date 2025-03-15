/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete [task ID]",
	Short: "Complete a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !FileExists() {
			fmt.Println("Aucune tâche enregistrée.")
			return
		}
		records := ReadFile()
		if len(records) <= 1 {
			fmt.Println("Aucune tâche enregistrée.")
			return
		}
		taskID := args[0]
		n, _ := strconv.Atoi(taskID)
		if len(records) <= n {
			fmt.Println("Aucune tâche avec cet ID.")
			return
		}
		for i, record := range records {
			if i == n {
				record[2] = "true"
				break
			}
		}
		WriteFile(records)
		fmt.Println("Tâche complétée.")
	},
}

// write the task list to the file
// open the file in write mode replacing the content

func WriteFile(records [][]string) {
	// delete the file
	err := os.Remove(GetFilePath())
	if err != nil {
		log.Fatal(err)
	}
	// create a new file
	file, err := os.OpenFile(GetFilePath(), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
