package service

import (
	Repository "accessment.com/microservice/db/repository"
)

type ServiceModuleIntf interface {
}

type ServiceModule struct {
	RepoService ServiceModuleIntf
}

func NewRepoService() *RepoService {
	return &RepoService{}
}

func NewServiceModule(repoService ServiceModuleIntf) *ServiceModule {
	return &ServiceModule{
		RepoService: repoService,
	}
}

var (
	DbRepo = Repository.NewRepo(
		&Repository.Repo{},
		&Repository.Repo{},
	)

	CommitRepo = *DbRepo.NewCommitRepo()
	RepoDetail = *DbRepo.NewRepoDetailRepo()
	Reposervice = NewRepoService()
)
