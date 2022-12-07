package day07

import (
	"strconv"
	"strings"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

type Folder struct {
	Children []*Folder
	Parent   *Folder
	Size     int
	Path     string
}

var runningTotal int
var smallestFolder int

func SolveFirst(filename string) int {
	runningTotal = 0
	data, _ := filereader.ReadFileToStringArray(filename)
	var currentFolder *Folder
	for _, val := range data {

		if strings.Contains(val, "$ ls") {
			continue
		}

		if strings.Contains(val, "$ cd") {
			slice := strings.Split(val, " ")
			folderName := slice[len(slice)-1]
			if currentFolder == nil {
				currentFolder = &Folder{
					Children: make([]*Folder, 0),
					Size:     0,
					Parent:   nil,
					Path:     folderName,
				}
			} else if val[5:] == ".." {
				// going up a directory - set currentFolder to the parent
				currentFolder = currentFolder.Parent
			} else {
				// Find the correct child by path and set the current folder
				for _, folder := range currentFolder.Children {
					if folder.Path == val[5:] {
						currentFolder = folder
						break
					}
				}
			}
			continue
		}

		if strings.Contains(val, "dir") {
			slice := strings.Split(val, " ")
			folderName := slice[len(slice)-1]
			newFolder := &Folder{
				Children: make([]*Folder, 0),
				Size:     0,
				Parent:   currentFolder,
				Path:     folderName,
			}
			// create a folder ready to cd in to later on
			currentFolder.Children = append(currentFolder.Children, newFolder)
			continue
		}

		numString := strings.Split(val, " ")[0]

		num, _ := strconv.Atoi(numString)

		currentFolder.Size += num
	}

	for currentFolder.Path != "/" {
		currentFolder = currentFolder.Parent
	}
	// this will update the runningTotal value.
	getSize(currentFolder)
	return runningTotal
}
func SolveSecond(filename string) int {
	smallestFolder = 0
	data, _ := filereader.ReadFileToStringArray(filename)
	var currentFolder *Folder
	for _, val := range data {

		if strings.Contains(val, "$ ls") {
			continue
		}

		if strings.Contains(val, "$ cd") {
			slice := strings.Split(val, " ")
			folderName := slice[len(slice)-1]
			if currentFolder == nil {
				currentFolder = &Folder{
					Children: make([]*Folder, 0),
					Size:     0,
					Parent:   nil,
					Path:     folderName,
				}
			} else if val[5:] == ".." {
				// going up a directory - set currentFolder to the parent
				currentFolder = currentFolder.Parent
			} else {
				// Find the correct child by path and set the current folder
				for _, folder := range currentFolder.Children {
					if folder.Path == val[5:] {
						currentFolder = folder
						break
					}
				}
			}
			continue
		}

		if strings.Contains(val, "dir") {
			slice := strings.Split(val, " ")
			folderName := slice[len(slice)-1]
			newFolder := &Folder{
				Children: make([]*Folder, 0),
				Size:     0,
				Parent:   currentFolder,
				Path:     folderName,
			}
			// create a folder ready to cd in to later on
			currentFolder.Children = append(currentFolder.Children, newFolder)
			continue
		}

		numString := strings.Split(val, " ")[0]

		num, _ := strconv.Atoi(numString)

		currentFolder.Size += num
	}

	for currentFolder.Path != "/" {
		currentFolder = currentFolder.Parent
	}

	const capacity = 70000000
	const spaceNeeded = 30000000
	totalSize := getSizeOfThisFolder(currentFolder)
	deleteSize := spaceNeeded - (capacity - totalSize)
	getSmallerFolder(currentFolder, deleteSize)
	return smallestFolder
}
func getSize(folder *Folder) int {
	var size int = folder.Size
	for _, val := range folder.Children {
		s := getSize(val)
		size += s
	}
	if size <= 100000 {
		runningTotal += size
	}
	return size
}
func getSizeOfThisFolder(folder *Folder) int {
	var size int = folder.Size
	for _, val := range folder.Children {
		s := getSize(val)
		size += s
	}
	return size
}
func getSmallerFolder(folder *Folder, sizeNeeded int) int {
	var size int = folder.Size
	for _, val := range folder.Children {
		s := getSmallerFolder(val, sizeNeeded)
		size += s
	}
	if size > sizeNeeded && (smallestFolder == 0 || size < smallestFolder) {
		smallestFolder = size
	}
	return size
}
