package day07

import (
	"fmt"
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

		if strings.Contains(val, "$ cd") {
			slice := strings.Split(val, " ")
			folderName := slice[len(slice)-1]
			if currentFolder == nil {
				fmt.Printf("Creating new folder called %s\n", folderName)
				currentFolder = &Folder{
					Children: make([]*Folder, 0),
					Size:     0,
					Parent:   nil,
					Path:     folderName,
				}
			} else if val[5:] == ".." {
				fmt.Printf("Gone up a directory - currently in %s , goingto %s \n", currentFolder.Path, currentFolder.Parent.Path)
				currentFolder = currentFolder.Parent
			} else {
				for _, folder := range currentFolder.Children {
					if folder.Path == val[5:] {
						fmt.Printf("Gone down a directory - currently in %s , going down to %s \n", currentFolder.Path, folder.Path)
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
			fmt.Printf("Creating a sub folder in %s called %s\n", currentFolder.Path, newFolder.Path)
			currentFolder.Children = append(currentFolder.Children, newFolder)
			continue
		}

		if strings.Contains(val, "$ ls") {
			continue
		}

		numString := strings.Split(val, " ")[0]

		num, err := strconv.Atoi(numString)

		if err != nil {
			panic(err)
		}

		// fmt.Println(num)

		currentFolder.Size += num
	}

	for currentFolder.Path != "/" {
		currentFolder = currentFolder.Parent
	}
	recursivelyPrint(currentFolder)
	getSize(currentFolder)
	// for currentFolder.Parent != nil {
	// 	currentFolder = currentFolder.Parent
	// }

	// fmt.Println(len(currentFolder.Children))
	// fmt.Println(len(currentFolder.Path))
	return runningTotal
}

func recursivelyPrint(folder *Folder) {
	fmt.Printf("Folder %s has a size of %d\n", folder.Path, folder.Size)
	for _, val := range folder.Children {
		recursivelyPrint(val)
	}
}

func getSize(folder *Folder) int {
	var size int = folder.Size
	for _, val := range folder.Children {
		s := getSize(val)
		size += s
	}

	fmt.Printf("Folder %s has a size of %d\n", folder.Path, size)
	if size <= 100000 {
		fmt.Printf("Adding %d to the total (%d) to make %d\n", size, runningTotal, runningTotal+size)
		runningTotal += size
	}
	return size
}
func SolveSecond(filename string) int {

	smallestFolder = 0
	data, _ := filereader.ReadFileToStringArray(filename)
	var currentFolder *Folder
	for _, val := range data {

		if strings.Contains(val, "$ cd") {
			slice := strings.Split(val, " ")
			folderName := slice[len(slice)-1]
			if currentFolder == nil {
				fmt.Printf("Creating new folder called %s\n", folderName)
				currentFolder = &Folder{
					Children: make([]*Folder, 0),
					Size:     0,
					Parent:   nil,
					Path:     folderName,
				}
			} else if val[5:] == ".." {
				fmt.Printf("Gone up a directory - currently in %s , goingto %s \n", currentFolder.Path, currentFolder.Parent.Path)
				currentFolder = currentFolder.Parent
			} else {
				for _, folder := range currentFolder.Children {
					if folder.Path == val[5:] {
						fmt.Printf("Gone down a directory - currently in %s , going down to %s \n", currentFolder.Path, folder.Path)
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
			fmt.Printf("Creating a sub folder in %s called %s\n", currentFolder.Path, newFolder.Path)
			currentFolder.Children = append(currentFolder.Children, newFolder)
			continue
		}

		if strings.Contains(val, "$ ls") {
			continue
		}

		numString := strings.Split(val, " ")[0]

		num, err := strconv.Atoi(numString)

		if err != nil {
			panic(err)
		}

		// fmt.Println(num)

		currentFolder.Size += num
	}

	for currentFolder.Path != "/" {
		currentFolder = currentFolder.Parent
	}
	recursivelyPrint(currentFolder)

	const capacity = 70000000
	const spaceNeeded = 30000000
	totalSize := getSizeOfThisFolder(currentFolder)

	deleteSize := spaceNeeded - (capacity - totalSize)
	getSizeTwo(currentFolder, deleteSize)

	fmt.Printf("capacity - total size = %d\n", capacity-totalSize)
	fmt.Printf("delete size = %d\n", deleteSize)

	// fmt.Printf("Total size = %d\n")
	// for currentFolder.Parent != nil {
	// 	currentFolder = currentFolder.Parent
	// }

	// fmt.Println(len(currentFolder.Children))
	// fmt.Println(len(currentFolder.Path))
	return smallestFolder
}
func getSizeOfThisFolder(folder *Folder) int {
	var size int = folder.Size
	for _, val := range folder.Children {
		s := getSize(val)
		size += s
	}
	return size
}
func getSizeTwo(folder *Folder, sizeNeeded int) int {
	var size int = folder.Size
	for _, val := range folder.Children {
		s := getSizeTwo(val, sizeNeeded)
		size += s
	}

	fmt.Printf("Folder %s has a size of %d\n", folder.Path, size)
	if size > sizeNeeded && (smallestFolder == 0 || size < smallestFolder) {
		fmt.Printf("Adding %d to the total (%d) to make %d\n", size, runningTotal, runningTotal+size)
		smallestFolder = size
	}
	return size
}
