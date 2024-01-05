package database

import (
	"apex_roulette/domain/model"
	"apex_roulette/domain/repository"
	"encoding/json"
	"fmt"
	"os"
)

type charactorRepository struct {
	data []model.Charactor
}

func NewCharactorRepository() repository.CharactorRepository {

	dataPath := GetDataDirPath()

	filePath := fmt.Sprintf("%s/charactors.json", dataPath)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var charactors []model.Charactor
	if err := json.NewDecoder(file).Decode(&charactors); err != nil {
		panic(err)
	}
	return &charactorRepository{
		data: charactors,
	}
}

func (c *charactorRepository) FindAll() ([]model.Charactor, error) {
	return c.data, nil
}
