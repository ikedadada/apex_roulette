package usecase

import (
	"apex_roulette/domain/model"
	"apex_roulette/domain/repository"
	"errors"
	"math/rand"
)

type Roulette interface {
	Start() (RouletteUsecaseOutput, error)
}

type roulette struct {
	charactorRepository repository.CharactorRepository
	weaponRepository    repository.WeaponRepository
}

func NewRoulette(c repository.CharactorRepository, w repository.WeaponRepository) Roulette {
	return &roulette{
		charactorRepository: c,
		weaponRepository:    w,
	}
}

type RouletteUsecaseOutput struct {
	PlayersSelectionStatus [3]PlayerSelectionStatus
}

type PlayerSelectionStatus struct {
	Charactor model.Charactor
	Weapons   [2]model.Weapon
}

var ErrCharactorMustBeMoreThanThree = errors.New("charactor must be more than three")

func (r *roulette) Start() (RouletteUsecaseOutput, error) {
	charactors, err := r.charactorRepository.FindAll()
	if err != nil {
		return RouletteUsecaseOutput{}, err
	}
	if len(charactors) < 3 {
		return RouletteUsecaseOutput{}, ErrCharactorMustBeMoreThanThree
	}
	weapons, err := r.weaponRepository.FindOnlyCanGetFromFields()
	if err != nil {
		return RouletteUsecaseOutput{}, err
	}

	players := [3]PlayerSelectionStatus{}
	pickedCharactorNames := []string{}
	pickedWeaponTypes := []string{}

	for i := range players {
		players[i].Charactor = pickUniqueCharactor(charactors, &pickedCharactorNames)
		for c := 0; c < 2; c++ {
			players[i].Weapons[c] = pickUniqueWeapon(weapons, &pickedWeaponTypes)
		}
	}

	return RouletteUsecaseOutput{
		PlayersSelectionStatus: players,
	}, nil
}

func pickUniqueCharactor(charactors []model.Charactor, pickedNames *[]string) model.Charactor {
	for {
		charactor := charactors[rand.Intn(len(charactors))]
		if !contains(*pickedNames, charactor.Name) {
			*pickedNames = append(*pickedNames, charactor.Name)
			return charactor
		}
	}
}

func pickUniqueWeapon(weapons []model.Weapon, pickedTypes *[]string) model.Weapon {
	for {
		weapon := weapons[rand.Intn(len(weapons))]
		if !contains(*pickedTypes, weapon.Type) {
			*pickedTypes = append(*pickedTypes, weapon.Type)
			return weapon
		}
	}
}

func contains(array []string, target string) bool {
	for _, v := range array {
		if v == target {
			return true
		}
	}
	return false
}
