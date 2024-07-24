package repository

import (
	Entity "accessment.com/microservice/db/entity"
	"gorm.io/gorm"
)

type RepoDetailRepository interface {
	Store(repoDetail Entity.RepoDetail) error
	GetByName(name string) (Entity.RepoDetail, error)
	GetAll() ([]Entity.RepoDetail, error)
}

type RepoDetailRepo struct {
	dB *gorm.DB
}

func (rep *RepoDetailRepo) Store(repoDetail Entity.RepoDetail) error {
	err := rep.dB.Create(&repoDetail).Error
	return err
}

func (rep *RepoDetailRepo) GetByName(name string) (Entity.RepoDetail, error) {
	var repoDetail Entity.RepoDetail
	err := rep.dB.Where(&Entity.RepoDetail{Name: name}).Find(&repoDetail).Error
	return repoDetail, err
}

func (rep *RepoDetailRepo) GetAll() ([]Entity.RepoDetail, error) {
	var repoDetail []Entity.RepoDetail
	err := rep.dB.Find(&repoDetail).Error
	return repoDetail, err
}
