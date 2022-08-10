package manager

import "github.com/itsapep/yopei-grpc/server/repository"

type RepositoryManager interface {
	YopeiRepository() repository.YopeiRepository
}

type repositoryManager struct {
	yopeiRepo repository.YopeiRepository
}

// YopeiRepository implements RepositoryManager
func (r *repositoryManager) YopeiRepository() repository.YopeiRepository {
	return r.yopeiRepo
}

func NewRepositoryManager() RepositoryManager {
	repo := new(repositoryManager)
	repo.yopeiRepo = repository.NewYopeiRepository()
	return repo
}
