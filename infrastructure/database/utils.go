package database

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetDataDirPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if filepath.Base(dir) == "database" {
		dir = filepath.Dir(filepath.Dir(dir))
	} else if filepath.Base(dir) == "usecase" {
		dir = filepath.Dir(filepath.Dir(dir))
	}
	return fmt.Sprintf("%s/infrastructure/database/data", dir)
}
