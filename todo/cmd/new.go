package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

const (
	ErrorTaskEmpty  = TaskError("Le nom de la tâche ne peut pas être vide")
	ErrorTaskExists = TaskError("La tâche existe déjà")
	taskAdded       = "✅ Tâche ajoutée: "
)

type TaskError string

func (e TaskError) Error() string {
	return string(e)
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [task name]",
	Short: "Create a new task in the todo list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		err := CreateFileIfNotExist()
		if err != nil {
			fmt.Println("Erreur lors de la création du fichier: ", err)	
		}
		err = WriteTaskToCsv(name)
		if err != nil {
			fmt.Println("Erreur lors de l'ajout de la tâche: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Get argument for the name/description of the task

}

func GetFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Erreur lors de la récupération du répertoire utilisateur: ", err)
	}
	dir := filepath.Join(homeDir, ".todo")      // ~/.todo
	filePath := filepath.Join(dir, "tasks.csv") // ~/.todo/tasks.csv
	// Crée le répertoire ~/.todo s'il n'existe pas
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			log.Fatal("Erreur lors de la création du répertoire: ", err)
		}
	}
	return filePath
}

func openFile() *os.File {
	file, err := os.Open(GetFilePath())
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture du fichier: ", err)
	}
	return file
}

func ReadFile() (records [][]string) {
	file := openFile()
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Erreur lors de la lecture du fichier: ", err)
	}
	return records
}

func AddLineToCsvFile(file *os.File, record []string) error {
	writer := csv.NewWriter(file)
	defer writer.Flush()
	if err := writer.Write(record); err != nil {
		return err
	}
	return nil
}

// Vérifie si un fichier existe
func FileExists() bool {
	_, err := os.Stat(GetFilePath())
	return !os.IsNotExist(err)
}

// Crée le fichier task.csv s'il n'existe pas
func CreateFileIfNotExist() error {
	if !FileExists() {
		file, err := os.Create(GetFilePath())
		if err != nil {
			log.Fatal("Erreur lors de la création du fichier: ", err)
		}
		defer file.Close()
		header := []string{"ID", "taskName", "done", "createdAt"}
		err = AddLineToCsvFile(file, header)
		if err != nil {
			return err
		}
	}
	return nil
}

// Vérifie si une tâche existe déjà
func assertTaskAlreadyExists(taskName string) bool {
	records := ReadFile()
	for _, record := range records[1:] {
		if len(record) > 1 && record[1] == taskName {
			return true
		}
	}
	return false
}

// Compte le nombre de tâches
func countTasks() int {
	records := ReadFile()
	return len(records) - 1
}

// Écrit une tâche dans un fichier CSV
func WriteTaskToCsv(taskName string) error {
	if taskName == "" {
		return ErrorTaskEmpty
	}
	if assertTaskAlreadyExists(taskName) {
		return ErrorTaskExists
	}
	// Ouvre le fichier en mode append
	file, err := os.OpenFile(GetFilePath(), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture du fichier: ", err)
	}
	defer file.Close()
	// Ajouter une nouvelle tâche
	id := strconv.Itoa(countTasks() + 1)
	record := []string{id, taskName, "false", time.Now().In(time.Local).Format("2006-01-02 15:04:05")}
	err = AddLineToCsvFile(file, record)
	if err != nil {
		return err
	}
	fmt.Println(taskAdded + taskName)
	return nil
}
