package utils

import "os"

func CreateFile(filePath string, content []byte) error {
	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	file.Write(content)

	defer file.Close()
	return nil
}
