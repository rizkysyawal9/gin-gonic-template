package manager

import "travelezat-dev/usecase"

type UseCaseManager interface {
	MenuUseCase() usecase.MenuUseCase
}

type usecaseManager struct {
	repo RepoManager
}

func (uc *usecaseManager) MenuUseCase() usecase.MenuUseCase {
	return usecase.NewMenuUseCase(uc.repo.MenuRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &usecaseManager{
		repo: repoManager,
	}
}
