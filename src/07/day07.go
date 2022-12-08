package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type directory struct {
	parent         *directory
	name           string
	files          []file
	subDirectories []*directory
	totalSize      int
}

type file struct {
	filename string
	size     int
}

func main() {

	data := parseFile("../../input/day07.txt")

	var files []file
	var directories []directory

	for _, line := range data {

		firstCharacter := line[0:1]

		// Step 1 : Create lists for directories and for files
		if lineIsFile(firstCharacter) {

			fileNameSize := strings.Split(line, " ")

			size, _ := strconv.Atoi(fileNameSize[0])
			var filename string = fileNameSize[1]

			newFile := file{filename: filename, size: size}

			// TODO : Move these to a seperate function
			found := false
			for _, file := range files {
				if file.filename == filename {
					found = true
					break
				}
			}
			if !found {
				files = append(files, newFile)
			}

			// I am not proud of this either (I am kinda proud)
		} else if lineIsDirectory(firstCharacter) {
			directoryName := strings.Split(line, " ")

			newDirectory := directory{name: directoryName[1]}

			// TODO : Move these to a seperate function
			found := false
			for _, directory := range directories {

				if directory.name == directoryName[1] {
					found = true
					break
				}
			}
			if !found {
				directories = append(directories, newDirectory)
			}
		}

	}

	// Step 1.9 : Put the root directory in the list.
	directories = append(directories, directory{name: "/"})

	// Step 2 : Traverse again and try to build the file tree
	var currentDirectory *directory

	for _, line := range data {

		firstCharacter := line[0:1]
		if lineIsDirectory(firstCharacter) {

			var directoryName string = line[4:]

			if foundDirectory, ok := findDirectory(directoryName, directories); ok {

				found := false
				for _, directory := range currentDirectory.subDirectories {
					if directory.name == foundDirectory.name {
						found = true
						break
					}
				}

				// TODO : Move these to a seperate function
				if !found && foundDirectory.name != currentDirectory.name {
					currentDirectory.subDirectories = append(currentDirectory.subDirectories, foundDirectory)
					foundDirectory.parent = currentDirectory
				}

			} else {
				panic("Could not find directory when trying to add sub-directory")
			}

		} else if lineIsFile(firstCharacter) {

			fileNameSize := strings.Split(line, " ")
			var fileName string = fileNameSize[1]

			if foundFile, ok := findFile(fileName, files); ok {

				// TODO : Move these to a seperate function
				found := false
				for _, file := range currentDirectory.files {
					if file.filename == foundFile.filename {
						found = true
						break
					}
				}
				if !found {
					currentDirectory.files = append(currentDirectory.files, *foundFile)
				}

			} else {
				panic("Could not find file when trying to add file to directory")
			}

		} else if lineIsChangeDirectory(line[0:4]) {

			if line[5:] == ".." {

				if parentDirectory, ok := findDirectory(currentDirectory.parent.name, directories); ok {
					currentDirectory = parentDirectory
				} else {
					panic("Could not find directory when cd ..")
				}
			} else {
				var targetDirectoryName string = line[5:]

				if targetDirectory, ok := findDirectory(targetDirectoryName, directories); ok {
					currentDirectory = targetDirectory
				} else {
					panic("Could not find directory when trying to access new directory")
				}
			}

		}

	}

	// Finally : Start calculating from the root directory :
	rootDirectory, _ := findDirectory("/", directories)
	calculateTotalSizeIncludingSubDirectories(rootDirectory)

	var goalSum int = 0
	for _, directory := range directories {

		if directory.totalSize < 100000 {
			goalSum += directory.totalSize
		}
	}

	fmt.Printf("goalSum: %v\n", goalSum)

}

// Recursion baby
func calculateTotalSizeIncludingSubDirectories(directory *directory) int {

	// Files under directory
	for _, file := range directory.files {
		directory.totalSize += file.size
	}

	if len(directory.subDirectories) > 0 {

		for _, subDirectory := range directory.subDirectories {
			directory.totalSize += calculateTotalSizeIncludingSubDirectories(subDirectory)
		}
	}

	println("Finished calculating for", directory.name, "with ", len(directory.subDirectories), "subdirectories")

	return directory.totalSize
}

func findFile(fileName string, files []file) (*file, bool) {

	for i := 0; i < len(files); i++ {
		if files[i].filename == fileName {
			return &files[i], true
		}
	}

	return nil, false
}

func findDirectory(directoryName string, directories []directory) (*directory, bool) {

	for i := 0; i < len(directories); i++ {
		if directories[i].name == directoryName {
			return &directories[i], true
		}
	}

	return nil, false
}

func lineIsFile(character string) bool {
	firstCharacter := []rune(character)
	return unicode.IsNumber(firstCharacter[0])
}

func lineIsDirectory(character string) bool {
	firstCharacter := []rune(character)
	return firstCharacter[0] == []rune("d")[0]
}

func lineIsChangeDirectory(characters string) bool {

	return characters == "$ cd"
}

func parseFile(filename string) []string {

	data, _ := os.ReadFile(filename)
	stringData := string(data)
	return strings.Split(stringData, "\r\n")
}
