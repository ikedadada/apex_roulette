package model

type Weapon struct {
	Name          string
	Type          string
	IsCarePackage bool
	IsCraft       bool
}

func (w *Weapon) IsGetFromField() bool {
	return !w.IsCarePackage && !w.IsCraft
}
