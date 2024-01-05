package database

import (
	"apex_roulette/domain/model"
	"apex_roulette/domain/repository"
	"encoding/json"
	"fmt"
	"os"
)

type weaponRepository struct {
	data []model.Weapon
}

func NewWeaponRepository() repository.WeaponRepository {

	dataPath := GetDataDirPath()

	filePath := fmt.Sprintf("%s/weapons.json", dataPath)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var weapons []model.Weapon
	if err := json.NewDecoder(file).Decode(&weapons); err != nil {
		panic(err)
	}
	return &weaponRepository{
		data: weapons,
	}
}

func (w *weaponRepository) FindOnlyCanGetFromFields() ([]model.Weapon, error) {
	var weapons []model.Weapon
	for _, weapon := range w.data {
		if weapon.IsGetFromField() {
			weapons = append(weapons, weapon)
		}
	}
	return weapons, nil
}
