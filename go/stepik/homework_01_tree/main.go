package main

import (
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

func printDirectoryOrFile(parOut io.Writer, fileName string, level int, isLastFile bool, isParentDirLast bool) {
	var stringPath strings.Builder
	var levelRepresentation string

	if isLastFile {
		levelRepresentation = "└───"
	} else {
		levelRepresentation = "├───"
	}

	var stringSeparator string = "|\t"

	if isParentDirLast {
		for i := 0; i < level-1; i++ {
			stringPath.WriteString(stringSeparator)
		}
		if level != 0 {
			stringPath.WriteString("\t")
		}
	} else {
		for i := 0; i < level; i++ {
			stringPath.WriteString(stringSeparator)
		}
	}

	stringPath.WriteString(levelRepresentation)
	stringPath.WriteString(fileName)
	stringPath.WriteRune('\n')
	parOut.Write([]byte(stringPath.String()))
}

func dirTreeInner(parOut io.Writer, parDirFiles []os.DirEntry, parBasePath string, parPrintFiles bool, parLevel int) error {
	var currentFilesLevelLength = len(parDirFiles)
	if currentFilesLevelLength > 0 {
		var parCurrentFile os.DirEntry
		parDirFiles, parCurrentFile = parDirFiles[:len(parDirFiles)-1], parDirFiles[len(parDirFiles)-1]
		sort.Slice(parDirFiles, func(i, j int) bool {
			return parDirFiles[i].Name() > parDirFiles[j].Name()
		})
		printDirectoryOrFile(parOut, parCurrentFile.Name(), parLevel, len(parDirFiles) == 0, currentFilesLevelLength == 1)
		currentDirectoryPath := parBasePath + string(os.PathSeparator) + parCurrentFile.Name()

		if parCurrentFile.IsDir() {
			currentDirectory, error := os.Open(currentDirectoryPath)
			if error != nil {
				return error
			}

			subDirFiles, error := currentDirectory.ReadDir(0)

			if error != nil {
				return error
			}

			dirTreeInner(parOut, subDirFiles, currentDirectoryPath, parPrintFiles, parLevel+1)
		}
		dirTreeInner(parOut, parDirFiles, parBasePath, parPrintFiles, parLevel)
	}
	return nil
}

func dirTree(parOut io.Writer, parPath string, parPrintFiles bool) error {

	var dirFiles []os.DirEntry

	currentDirectory, error := os.Open(parPath)
	if error != nil {
		return error
	}

	error = currentDirectory.Chdir()
	if error != nil {
		return error
	}

	dirFiles, error = currentDirectory.ReadDir(0)
	if error != nil {
		return error
	}

	dirTreeInner(parOut, dirFiles, parPath, parPrintFiles, 0)

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
