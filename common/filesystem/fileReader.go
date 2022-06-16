package filesystem

import (
	"bufio"
	"log"
	"os"
)

type FileService struct {
}

func ProvideFileService() *FileService {
	return &FileService{}
}

func (file *FileService) ReadLines(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Println("Invalid file", path)
	}

	defer f.Close()

	results := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		results = append(results, scanner.Text())
	}

	return results
}
