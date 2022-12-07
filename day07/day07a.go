package day07

import (
	"strconv"
	"strings"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

type FolderA struct {
	Children map[string]*FolderA
	Parent   *FolderA
	Size     int
}

func SolveFirstA(filename string) int {
	folder := helper(filename)
	// this will update the runningTotal value.
	return getRunningTotal(folder, 0)
}
func SolveSecondA(filename string) int {
	const capacity = 70000000
	const spaceNeeded = 30000000

	folder := helper(filename)

	totalSize := getSizeOfThisFolderA(folder)
	deleteSize := spaceNeeded - (capacity - totalSize)

	return getSmallerFolderA(folder, deleteSize, 0)
}

// this will sort all the text in to tree logic to sort through recursively
func helper(filename string) *FolderA {
	data, _ := filereader.ReadFileToStringArray(filename)
	var currentFolder *FolderA
	for _, val := range data {

		if strings.Contains(val, "$ ls") {
			continue
		}

		if strings.Contains(val, "$ cd") {
			if currentFolder == nil {
				currentFolder = &FolderA{
					Children: map[string]*FolderA{},
					Size:     0,
					Parent:   nil,
				}
			} else if val[5:] == ".." {
				// going up a directory - set currentFolder to the parent
				currentFolder = currentFolder.Parent
			} else {
				// Find the correct child by path and set the current folder
				currentFolder = currentFolder.Children[val[5:]]
			}
			continue
		}

		if strings.Contains(val, "dir") {
			newFolder := &FolderA{
				Children: map[string]*FolderA{},
				Size:     0,
				Parent:   currentFolder,
			}
			// create a folder ready to cd in to later on
			currentFolder.Children[val[4:]] = newFolder
			continue
		}

		numString := strings.Split(val, " ")[0]

		num, _ := strconv.Atoi(numString)

		currentFolder.Size += num
	}

	for currentFolder.Parent != nil {
		currentFolder = currentFolder.Parent
	}
	return currentFolder
}

// wrapper methods for ease of use

func getRunningTotal(folder *FolderA, runningTotal int) int {
	_, rt := getSizeARecur(folder, runningTotal)
	return rt
}
func getSizeOfThisFolderA(folder *FolderA) int {
	var size int = folder.Size
	for _, val := range folder.Children {
		s, _ := getSizeARecur(val, 0)
		size += s
	}
	return size
}

// returns size of the folder, and the smallest folder needed
func getSmallerFolderA(folder *FolderA, sizeNeeded int, smallestFolder int) int {
	_, sf := getSmallerFolderARecur(folder, sizeNeeded, smallestFolder)
	return sf
}

// Methods below are called by wrapper methods above for readability. (think we lose some performance though)

func getSizeARecur(folder *FolderA, runningTotal int) (int, int) {
	var size int = folder.Size
	var rt = runningTotal
	var s int
	for _, val := range folder.Children {
		s, rt = getSizeARecur(val, rt)
		size += s
	}
	if size <= 100000 {
		rt += size
	}
	return size, rt
}

// returns size of the folder, and the smallest folder needed
func getSmallerFolderARecur(folder *FolderA, sizeNeeded int, smallestFolder int) (int, int) {
	var size int = folder.Size
	var sf, s int = smallestFolder, 0
	for _, val := range folder.Children {
		s, sf = getSmallerFolderARecur(val, sizeNeeded, sf)
		size += s
	}
	if size > sizeNeeded && (sf == 0 || size < sf) {
		sf = size
	}
	return size, sf
}
