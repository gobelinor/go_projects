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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a task",
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
		WriteFileWithoutLine(records, n)
		fmt.Println("Tâche supprimée.")
	},
}

func WriteFileWithoutLine(records [][]string, n int) {
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
	k := 0
	for _, record := range records {
		if record[0] == strconv.Itoa(n) {
			continue
		}
		record[0] = strconv.Itoa(k)
		err := writer.Write(record)
		if err != nil {
			log.Fatal(err)
		}
		k++
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
