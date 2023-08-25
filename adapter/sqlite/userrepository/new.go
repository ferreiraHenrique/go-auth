package userrepository

import (
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/core/domain"
)

type repository struct {
	db sqlite.PoolInterface
}

func New(db sqlite.PoolInterface) domain.UserRepository {
	return &repository{
		db: db,
	}
}
