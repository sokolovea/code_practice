package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type dirAndLevel struct {
	directories []os.FileInfo
	path        string
	level       int
}

func printDirectoryOrFile(parOut io.Writer, fileName string, level int, isLast bool) {
	var stringPath strings.Builder
	var levelRepresentation string
	if isLast {
		levelRepresentation = "└───"
	} else {
		levelRepresentation = "├───"
	}
	for i := 1; i < level; i++ {
		stringPath.WriteString("|\t")
	}
	stringPath.WriteString(levelRepresentation)
	stringPath.WriteString(fileName)
	stringPath.WriteRune('\n')
	parOut.Write([]byte(stringPath.String()))
}

func dirTree(parOut io.Writer, parPath string, parPrintFiles bool) error {

	var dirFiles []os.FileInfo

	var dirFilesStack []dirAndLevel = make([]dirAndLevel, 0, 10)

	var currentDirectoryLevel int = 0

	for {
		currentDirectory, error := os.Open(parPath)
		if error != nil {
			return error
		}

		error = currentDirectory.Chdir()
		if error != nil {
			return error
		}

		dirFiles, error = currentDirectory.Readdir(0)
		if error != nil {
			return error
		}

		sort.Slice(dirFiles, func(i, j int) bool {
			return dirFiles[i].Name() < dirFiles[j].Name()
		})

		var dirAndLevel dirAndLevel

		dirAndLevel.directories = dirFiles
		dirAndLevel.level = currentDirectoryLevel + 1
		dirAndLevel.path = parPath

		if len(dirAndLevel.directories) == 0 {
			continue
		}

		dirFilesStack = append(dirFilesStack, dirAndLevel)

		if len(dirFilesStack) == 0 {
			fmt.Printf("Debug: EXIT!")
			break
		}

		currentFilesLevel := dirFilesStack[len(dirFilesStack)-1]
		dirFilesStack = dirFilesStack[:len(dirFilesStack)-1]

		var currentFilesLevelLength = len(currentFilesLevel.directories)

		if currentFilesLevelLength > 0 {
			currentFilesLevel.directories, parPath = currentFilesLevel.directories[:len(currentFilesLevel.directories)-1], currentFilesLevel.directories[len(currentFilesLevel.directories)-1].Name()
			parPath = dirAndLevel.path + string(os.PathSeparator) + parPath
			printDirectoryOrFile(parOut, parPath, currentFilesLevel.level, len(currentFilesLevel.directories) == 0)
			dirFilesStack = append(dirFilesStack, currentFilesLevel)
		}
	}
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
