package main

import "os"

func ReadImage(filePath string) bool {
	_, err := os.Stat(filePath)

	return err == nil
}
