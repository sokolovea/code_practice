package main

import (
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type dirAndLevel struct {
	directories []os.FileInfo
	path        string
	level       int
}

func printDirectoryOrFile(parOut io.Writer, fileName string, isCurrentFileDir bool, isLastFile bool, parentDirsLast []bool, fileSizeBytes int64) {
	var stringPath strings.Builder
	var levelRepresentation string

	if isLastFile {
		levelRepresentation = "└───"
	} else {
		levelRepresentation = "├───"
	}

	for _, isLastDirectory := range parentDirsLast {
		if isLastDirectory {
			stringPath.WriteString(" ")
		} else {
			stringPath.WriteString("|")
		}
		stringPath.WriteString("   ")
	}

	stringPath.WriteString(levelRepresentation)
	stringPath.WriteString(fileName)
	if isCurrentFileDir {
		stringPath.WriteString(" (")
		if fileSizeBytes == 0 {
			stringPath.WriteString("empty")
		} else {
			stringPath.WriteString(strconv.FormatInt(fileSizeBytes, 10))
			stringPath.WriteRune('b')
		}
		stringPath.WriteRune(')')
	}
	stringPath.WriteRune('\n')
	parOut.Write([]byte(stringPath.String()))
}

func dirTreeInner(parOut io.Writer, parDirFiles []os.DirEntry, parBasePath string, parPrintFiles bool, parentDirsLast []bool) error {
	var currentFilesLevelLength = len(parDirFiles)
	if currentFilesLevelLength > 0 {
		var parCurrentFile os.DirEntry
		sort.Slice(parDirFiles, func(i, j int) bool {
			return parDirFiles[i].Name() > parDirFiles[j].Name()
		})
		parDirFiles, parCurrentFile = parDirFiles[:len(parDirFiles)-1], parDirFiles[len(parDirFiles)-1]

		currentFilePath := parBasePath + string(os.PathSeparator) + parCurrentFile.Name()

		fileInfo, error := parCurrentFile.Info()
		var fileSizeBytesInfo int64
		if error != nil {
			fileSizeBytesInfo = 0
		} else {
			fileSizeBytesInfo = fileInfo.Size()
		}

		isCurrentFileDir := parCurrentFile.IsDir()
		if isCurrentFileDir || parPrintFiles {
			printDirectoryOrFile(parOut, parCurrentFile.Name(), isCurrentFileDir, len(parDirFiles) == 0, parentDirsLast, fileSizeBytesInfo)
		}
		if isCurrentFileDir {

			currentFile, error := os.Open(currentFilePath)
			if error != nil {
				return error
			}

			subDirFiles, error := currentFile.ReadDir(0)

			if error != nil {
				return error
			}
			if currentFilesLevelLength == 1 {
				dirTreeInner(parOut, subDirFiles, currentFilePath, parPrintFiles, append(parentDirsLast, true))
			} else {
				dirTreeInner(parOut, subDirFiles, currentFilePath, parPrintFiles, append(parentDirsLast, false))
			}
		}
		dirTreeInner(parOut, parDirFiles, parBasePath, parPrintFiles, parentDirsLast)
	}
	return nil
}

func dirTree(parOut io.Writer, parPath string, parPrintFiles bool) error {

	var dirFiles []os.DirEntry

	currentDirectory, error := os.Open(parPath)
	if error != nil {
		return error
	}

	dirFiles, error = currentDirectory.ReadDir(0)
	if error != nil {
		return error
	}

	var parentDirsLast []bool = make([]bool, 0, 10000)
	dirTreeInner(parOut, dirFiles, parPath, parPrintFiles, parentDirsLast)

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
