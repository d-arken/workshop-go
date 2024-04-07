package pokemon

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (r Repository) BatchCreate(pokemons []Model, batchSize int) error {
	return r.DB.CreateInBatches(pokemons, batchSize).Error
}

func (r Repository) GetAll(pokemons *[]Model) error {
	return r.DB.Find(pokemons).Error
}
