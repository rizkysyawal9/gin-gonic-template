package manager

import "menu-manage/repo"

type RepoManager interface {
	MenuRepo() repo.MenuRepository
	TableTransactionRepo() repo.TableTransactionRepository
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) MenuRepo() repo.MenuRepository {
	return repo.NewMenuRepository(r.infra.SqlDb())
}

func (r *repoManager) TableTransactionRepo() repo.TableTransactionRepository {
	return repo.NewTableTransactionRepository(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{infra: infra}
}
