package link

import (
	"go/adv-demo2/db"
)

type LinkRepositoryDeps struct {
	Database *db.Db
}

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(db *db.Db) *LinkRepository {
	return &LinkRepository{Database: db}
}

func (repo *LinkRepository) Create(link *Link) error {
	repo.Database
}
