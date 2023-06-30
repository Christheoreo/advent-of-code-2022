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

const max = 100000

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

func SolveFirst(data string) int {

	// directories := make([]byte, 1)
	// directories[0] = '/'
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

	// answer := 0

	for storage.Parent != nil {
		storage = storage.Parent
	}

	_, answer := countX(storage)
	return answer
}

func countX(s *StorageStructure) (totalSize int, answer int) {

	if len(s.Children) > 0 {
		for _, c := range s.Children {
			newTotalSize, newAnswer := countX(c)
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
	return 0
}
