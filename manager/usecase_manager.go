package manager

import "menu-manage/usecase"

type UseCaseManager interface {
	MenuUseCase() usecase.MenuUseCase
	CustomerTableUseCase() usecase.CustomerTableUseCase
}

type usecaseManager struct {
	repo RepoManager
}

func (uc *usecaseManager) MenuUseCase() usecase.MenuUseCase {
	return usecase.NewMenuUseCase(uc.repo.MenuRepo())
}

func (uc *usecaseManager) CustomerTableUseCase() usecase.CustomerTableUseCase {
	return usecase.NewCustomerTableUseCase(uc.repo.TableTransactionRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &usecaseManager{
		repo: repoManager,
	}
}
