/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if !FileExists() {
			fmt.Println("Aucune tâche enregistrée.")
			return
		}
		records := ReadFile()
		// Vérifier si la liste est vide (hormis l'en-tête)
		if len(records) <= 1 {
			fmt.Println("Aucune tâche enregistrée.")
			return
		}
		w := tabwriter.NewWriter(os.Stdout, 4, 0, 2, ' ', 0)
		defer w.Flush()
		fmt.Println("\n📋 LISTE DES TÂCHES 📋")
		fmt.Println("───────────────────────────────────────────────────────")
		fmt.Fprintln(w, "ID\tTask\tStatus\tCreated")
		for _, record := range records[1:] {
			// Remplacer "true" par ✅ et "false" par ❌
			status := ""
			if record[2] == "true" {
				status = "✅"
			} else {
				status = "❌"
			}
			layout := "2006-01-02 15:04:05"
			date, _ := time.ParseInLocation(layout, record[3], time.Local)
			elapsed := timediff.TimeDiff(date)
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", record[0], record[1], status, elapsed)
		}
		w.Flush()
		// fmt.Println("\n")
		// fmt.Println("───────────────────────────────────────────────────────\n")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
