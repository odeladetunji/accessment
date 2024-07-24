package repository

type RepositoryIntf interface {
}

type Repo struct {
	CommitRepo RepositoryIntf
	RepoDetail RepositoryIntf
}

func NewRepo(commitRepo RepositoryIntf, repoDetail RepositoryIntf) *Repo {
	return &Repo{
		CommitRepo: commitRepo,
		RepoDetail: repoDetail,
	}
}

func (rep *Repo) NewCommitRepo() *CommitRepo {
	return &CommitRepo{}
}

func (rep *Repo) NewRepoDetailRepo() *RepoDetailRepo {
	return &RepoDetailRepo{}
}
