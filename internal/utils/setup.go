package utils

import (
	"os"
)

func VerifyUploadsDir(path string) error {
	entries, err := os.ReadDir(".")
	if err != nil {
		return err
	}
	
	// look for dir and handle if found
	for _, e := range entries {
		if e.IsDir() && e.Name() == path {
			println("found uploads dir")
			return nil
		}
	}

	// create new if not found
	err = os.Mkdir("uploads", 0755)
	if err != nil {
		return err
	}
	println("created uploads dir")

	return nil
}
