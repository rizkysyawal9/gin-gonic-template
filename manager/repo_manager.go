package manager

import "travelezat-dev/repo"

type RepoManager interface {
	MenuRepo() repo.MenuRepository
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) MenuRepo() repo.MenuRepository {
	return repo.NewMenuRepository(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{infra: infra}
}
