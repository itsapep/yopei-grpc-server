package manager

import "github.com/itsapep/yopei-grpc/server/service"

type ServiceManager interface {
	YopeiService() *service.YopeiService
}

type serviceManager struct {
	yopeiService *service.YopeiService
}

// YopeiService implements ServiceManager
func (s *serviceManager) YopeiService() *service.YopeiService {
	return s.yopeiService
}

func NewServiceManager(repoManager RepositoryManager) ServiceManager {
	return &serviceManager{
		yopeiService: service.NewYopeiService(repoManager.YopeiRepository()),
	}
}
