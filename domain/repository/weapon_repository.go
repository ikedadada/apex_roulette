package repository

import "apex_roulette/domain/model"

type WeaponRepository interface {
	FindOnlyCanGetFromFields() ([]model.Weapon, error)
}
