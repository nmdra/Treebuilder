package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

var dryRun bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "treebuilder <structure-file>",
	Short: "Quickly scaffold a directory structure from a text file",
	Long: `A CLI tool that reads a tree-formatted text file and creates the corresponding 
	directories and empty files on your disk.

	Example input file:
	myproject/
	├── main.go
	└── pkg/
	└── utils.go`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		buildStructure(filePath)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Preview what will be created without modifying the disk")
}

// buildStructure contains the core directory generation logic
func buildStructure(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var stack []string
	firstLine := true

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		if firstLine {
			root := strings.TrimSuffix(strings.TrimSpace(line), "/")
			stack = []string{root}

			if dryRun {
				fmt.Println("[DRY RUN] dir :", root)
			} else {
				if err := os.MkdirAll(root, os.ModePerm); err != nil {
					log.Fatalf("Failed to create root dir: %v", err)
				}
				fmt.Println("dir :", root)
			}

			firstLine = false
			continue
		}

		level, name := parseLine(line)
		if name == "" {
			continue
		}

		isDir := strings.HasSuffix(name, "/")
		name = strings.TrimSuffix(name, "/")

		depth := level + 1
		if depth > len(stack) {
			depth = len(stack)
		}

		stack = stack[:depth]
		stack = append(stack, name)

		fullPath := filepath.Join(stack...)

		if isDir {
			if dryRun {
				fmt.Println("[DRY RUN] dir :", fullPath)
			} else {
				if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
					log.Fatalf("Failed to create dir %s: %v", fullPath, err)
				}
				fmt.Println("dir :", fullPath)
			}
		} else {
			if dryRun {
				fmt.Println("[DRY RUN] file:", fullPath)
			} else {
				if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
					log.Fatalf("Failed to create parent dir for %s: %v", fullPath, err)
				}

				f, err := os.Create(fullPath)
				if err != nil {
					log.Fatalf("Failed to create file %s: %v", fullPath, err)
				}
				f.Close()

				fmt.Println("file:", fullPath)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
}

func parseLine(line string) (int, string) {
	idx1 := strings.Index(line, "├──")
	idx2 := strings.Index(line, "└──")

	idx := idx1
	if idx2 != -1 && (idx == -1 || idx2 < idx) {
		idx = idx2
	}

	if idx == -1 {
		return 0, strings.TrimSpace(line)
	}

	indentStr := line[:idx]
	level := utf8.RuneCountInString(indentStr) / 4

	namePart := line[idx:]
	namePart = strings.TrimPrefix(namePart, "├── ")
	namePart = strings.TrimPrefix(namePart, "└── ")
	namePart = strings.TrimPrefix(namePart, "├──")
	namePart = strings.TrimPrefix(namePart, "└──")

	return level, strings.TrimSpace(namePart)
}
