package filereader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// go up a level until we find the data folder
// only works in we have 1 folder and only 1 folder called data
func getDataFolderPath() (string, error) {
	path, _ := os.Getwd()
	index := 0
	for {
		_, err := os.Stat(fmt.Sprintf("%s/data", path))
		if !os.IsNotExist(err) {
			break
		}
		path = fmt.Sprintf("%s/..", path)
		if index > 5 {
			return "", errors.New("could not find folder")
		}
	}
	return fmt.Sprintf("%s/data", path), nil
}

func readFileInDataFolder(filename string) (*os.File, error) {
	dataFolder, err := getDataFolderPath()
	if err != nil {
		return nil, err
	}
	filePath := fmt.Sprintf("%s/%s", dataFolder, filename)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("%s - %s", err.Error(), filePath)
	}
	return file, nil
}

func ReadFileToStringArray(filepath string) ([]string, error) {

	file, err := readFileInDataFolder(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
