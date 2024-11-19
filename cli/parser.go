package cli

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Parser *cobra.Command

func FileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func ParseFile(cmd *cobra.Command, args []string) error {

	filePath, err := cmd.Flags().GetString("file")
	if err != nil {
		return fmt.Errorf("[!] ERROR : l'argument file n'a pas été trouvé")
	}

	// Remplacer les antislashs par des slashs dans le chemin du fichier
	filePath = strings.Replace(filePath, "\\", "/", -1)

	fmt.Println("Chemin du fichier:", filePath)

	if !FileExist(filePath) {
		fmt.Println("[!] ERROR : fichier introuvable")
		return nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("[!] ERROR : erreur lors de l'ouverture du fichier")
		return nil
	}
	defer file.Close()

	lecteur := csv.NewReader(file)

	_, err = lecteur.Read() // Ignorer la première ligne
	if err != nil {
		fmt.Println("[!] ERROR : erreur lors de la lecture de la première ligne du fichier CSV")
		return nil
	}

	lineCount := 0

	// Compter le nombre de lignes
	for {
		_, err := lecteur.Read()
		if err != nil {
			break
		}
		lineCount++
	}

	fmt.Println("Nombre de lignes : ", lineCount)

	// Réinitialiser le fichier pour une nouvelle lecture
	file.Seek(0, 0)
	_, err = lecteur.Read() // Ignorer la première ligne à nouveau
	if err != nil {
		fmt.Println("[!] ERROR : erreur lors de la lecture de la première ligne du fichier CSV")
		return nil
	}

	// Ouvrir le fichier de sortie
	outputFile, err := os.Create("resultats.txt")
	if err != nil {
		fmt.Println("[!] ERROR : erreur lors de la création du fichier de sortie")
		return nil
	}
	defer outputFile.Close()

	// Lire les numéros dans le CSV
	var numbers []string

	for {
		record, err := lecteur.Read()
		if err != nil {
			break
		}

		// Récupérer le numéro de la première colonne
		str := record[0]
		num := strings.Split(str, ";")[0]
		numbers = append(numbers, num)
	}

	// Générer la chaîne formatée en fonction du nombre de numéros
	var formatted string

	switch lineCount {
	case 5:
		formatted = "s/^(?!(42)).*$/1${R:[0-4]1}/"
		for i, num := range numbers {
			formatted += fmt.Sprintf(";s/^%d/%s/", 10+i, num)
		}
		formatted += "/;s/[+]//;"
	case 10:
		formatted = "s/^(?!(42)).*$/1${R:[0-9]1}/"
		for i, num := range numbers {
			formatted += fmt.Sprintf(";s/^%d/%s/", 10+i, num)
		}
		formatted += "/;s/[+]//;"
	case 26:
		formatted = "s/.*/${R:[a-z]1}/"
		for i, num := range numbers {
			formatted += fmt.Sprintf(";s/^%c/%s/", 'a'+rune(i), num)
		}
		formatted += "/"
	default:
		return fmt.Errorf("[!] ERROR : Nombre de lignes non pris en charge")
	}

	// Écrire la chaîne formatée dans le fichier de sortie
	outputFile.WriteString(formatted + "\n")

	fmt.Println("Traitement terminé. Résultats enregistrés dans resultats.txt.")
	return nil
}

func init() {
	Parser = &cobra.Command{
		Use:   "parser",
		Short: "Parse phone numbers from excel file",
		RunE:  ParseFile,
	}

	Parser.Flags().StringP("file", "f", "", "Path to excel file")
	Parser.MarkFlagRequired("file")

	rootCmd.AddCommand(Parser)
}
