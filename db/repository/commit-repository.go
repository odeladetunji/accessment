package repository

import (
	Entity "accessment.com/microservice/db/entity"
	"gorm.io/gorm"
)

type CommitRepository interface {
	Store(commit Entity.Commit) error
	GetCommit(repoName string) ([]Entity.Commit, error)
	GetCommitInSha(shaList []string) ([]Entity.Commit, error)
	StoreList(commits []Entity.Commit) error
}

type CommitRepo struct {
	dB *gorm.DB
}

func (com *CommitRepo) Store(commit Entity.Commit) error {
	err := com.dB.Create(&commit).Error
	return err
}

func (com *CommitRepo) StoreList(commits []Entity.Commit) error {
	err := com.dB.Create(&commits).Error
	return err
}

func (com *CommitRepo) GetCommit(repoName string) ([]Entity.Commit, error) {
	var commits []Entity.Commit
	err := com.dB.Where(&Entity.Commit{RepoName: repoName}).Find(&commits).Error
	return commits, err
}

func (com *CommitRepo) GetCommitInSha(shaList []string) ([]Entity.Commit, error) {
	var commits []Entity.Commit
	err := com.dB.Where("commits.sha IN ?", shaList).Find(&commits).Error
	return commits, err
}
