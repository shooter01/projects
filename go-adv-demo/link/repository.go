package link

import (
	"fmt"
	"go/adv-demo/pkg/db"
)

type LinkRepositoryDeps struct {
	Database *db.Db
}

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	fmt.Println(&LinkRepository{
		Database: database,
	})
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil

}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil

	// if result.Error != nil {
	// 	return nil, result.Error
	// }
	// return link, nil

}
