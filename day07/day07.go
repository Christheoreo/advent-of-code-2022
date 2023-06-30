package day07

import (
	"strconv"
	"strings"
)

type StorageStructure struct {
	SizeAtThisLevel int
	Children        map[string]*StorageStructure
	Parent          *StorageStructure
}

func (s *StorageStructure) GetTotalSize() int {
	size := s.SizeAtThisLevel
	for _, c := range s.Children {
		size += c.GetTotalSize()
	}
	return size
}

func (s *StorageStructure) AddChildFolder(folderName string) {
	_, ok := s.Children[folderName]

	if !ok {
		s.Children[folderName] = &StorageStructure{
			Parent:          s,
			Children:        make(map[string]*StorageStructure),
			SizeAtThisLevel: 0,
		}
	}
}

func createStorage(data string) *StorageStructure {
	storage := &StorageStructure{
		SizeAtThisLevel: 0,
		Children:        make(map[string]*StorageStructure),
	}

	for lineIndex, line := range strings.Split(data, "\n") {
		if lineIndex == 0 {
			continue
		}

		if line == "$ ls" {
			continue
		}

		if strings.Index(line, "dir") == 0 {
			storage.AddChildFolder(line[4:])
			continue
		}

		if strings.Index(line, "$ cd ") == 0 {
			if line == "$ cd .." {
				storage = storage.Parent
				continue
			}
			folder := line[5:]
			storage = storage.Children[folder]
			continue
		}

		// else it is a file
		s := strings.Split(line, " ")
		numberS := s[0]
		number, _ := strconv.Atoi(numberS)
		storage.SizeAtThisLevel += number
	}

	// Go to the top level Dir
	for storage.Parent != nil {
		storage = storage.Parent
	}
	return storage
}

func SolveFirst(data string) int {
	const max = 100000
	storage := createStorage(data)

	_, answer := countX(storage, max)
	return answer
}

func countX(s *StorageStructure, max int) (totalSize int, answer int) {

	if len(s.Children) > 0 {
		for _, c := range s.Children {
			newTotalSize, newAnswer := countX(c, max)
			totalSize += newTotalSize
			answer += newAnswer
		}
	} else {
		if s.SizeAtThisLevel <= max {
			answer = s.SizeAtThisLevel
			return s.SizeAtThisLevel, answer
		}
	}
	totalSize += s.SizeAtThisLevel

	if totalSize <= max {
		answer += totalSize
	}

	return totalSize, answer

}
func SolveSecond(data string) int {
	const totalSize = 70000000
	const amountNeeded = 30000000
	storage := createStorage(data)

	storageTotal := countTotalSize(storage)

	spaceAvailable := totalSize - storageTotal

	toDelete := amountNeeded - spaceAvailable

	_, answer := countY(storage, toDelete, storageTotal)
	return answer
}

func countTotalSize(s *StorageStructure) int {
	size := s.SizeAtThisLevel
	if len(s.Children) > 0 {
		for _, c := range s.Children {
			size += countTotalSize(c)
		}
	}
	return size
}

func countY(s *StorageStructure, spaceNeeded int, currentBestFit int) (totalSize int, currentBestFitNew int) {
	currentBestFitNew = currentBestFit
	if len(s.Children) > 0 {
		for _, c := range s.Children {
			newTotalSize, cb := countY(c, spaceNeeded, currentBestFitNew)
			currentBestFitNew = cb
			totalSize += newTotalSize
		}
	}

	totalSize += s.SizeAtThisLevel

	if totalSize >= spaceNeeded && totalSize < currentBestFitNew {
		currentBestFitNew = totalSize
	}

	return
}
