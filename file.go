package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func Save(store map[string]string) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("failed to get current dir: ", err)
	}

	path := filepath.Join(pwd, "data.csv")
	// now just overwrite for making delete easy otherwise we need Find func and delete line by line 
	// now just overwrite whole file
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("failed to save: ", err)
	}

	defer file.Close()

	for key, value := range store {
		_, err := fmt.Fprintln(file,fmt.Sprintf("%s,%s",key,value))

		if err != nil {
			fmt.Println("failed to write: ", err)
		}
	}

}
