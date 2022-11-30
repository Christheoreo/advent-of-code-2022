package filereader

import (
	"strings"
	"testing"
)

func TestGetDataFolder(t *testing.T) {
	path, err := getDataFolderPath()

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if !strings.Contains(path, "/data") {
		t.Logf("incorrect path - %s", path)
		t.Fail()
	}
}

func TestReadFileInDataFolder(t *testing.T) {
	file, err := readFileInDataFolder("testFile.txt")

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if !strings.Contains(file.Name(), "testFile") {
		t.Logf("incorrect file - %v", file)
		t.Fail()
	}
}

func TestReadFileToStringArray(t *testing.T) {
	content, err := ReadFileToStringArray("testFile.txt")

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if len(content) != 3 {
		t.Logf("length of content should be 3, currently is %d", len(content))
		t.Fail()
	}

}

func TestReadNonExistantFileToStringArray(t *testing.T) {
	content, err := ReadFileToStringArray("fileThatDoesNotExist.txt")

	if err == nil {
		t.Log("error should not be nil", err)
		t.Fail()
	}

	if len(content) > 0 {
		t.Log("content should be empty")
		t.Fail()
	}

}
