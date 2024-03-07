package productrepository

import (
	"github.com/stephan-lopes/golang-clean-arch/adapter/postgres"
	"github.com/stephan-lopes/golang-clean-arch/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

func New(db postgres.PoolInterface) domain.ProductRepository {
	return &repository{
		db: db,
	}
}
