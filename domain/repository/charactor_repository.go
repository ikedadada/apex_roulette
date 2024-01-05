package repository

import "apex_roulette/domain/model"

type CharactorRepository interface {
	FindAll() ([]model.Charactor, error)
}
