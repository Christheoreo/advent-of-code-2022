package day07

import (
	"fmt"
	"strings"

	"github.com/christheoreo/advent-of-code-2022/internal/filereader"
)

type Folder struct {
	Folders []Folder
	Size    int
}

// func SolveFirst(filename string) int {
// 	data, _ := filereader.ReadFileToStringArray(filename)
// 	path := make([]string, 0)
// 	for _, val := range data {
// 		if val == "$ cd /" {
// 			continue
// 		}
// 		if strings.Contains(val, "$ cd") {
// 			slice := strings.Split(val, " ")
// 			folderName := slice[len(slice)-1]
// 			if folderName == ".." {
// 				path = path[:len(path)-1]
// 			} else {
// 				path = append(path, folderName)
// 			}

// 			fmt.Println(path)
// 		}
// 	}
// 	fmt.Println(folders)
// 	return 0
// }

func SolveFirst(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	folders := make([]Folder, 0)
	folder := Folder{
		Folders: []Folder{},
		Size:    0,
	}
	for _, val := range data {
		if val == "$ cd /" {
			continue
		}
		if strings.Contains(val, "$ cd") {
			slice := strings.Split(val, " ")
			folderName := slice[len(slice)-1]

			newFolder := Folder{
				Folders: []Folder{},
				Size:    0,
			}

			
			continue
		}
		fmt.Println(folders)
		return 0
	}
func SolveSecond(filename string) int {
	data, _ := filereader.ReadFileToStringArray(filename)
	for _, val := range data {
		fmt.Println(val)
	}
	return 0
}
