package bootstrap

import (
	"haven/internal/adapter/repository"
	"haven/internal/adapter/rest"
)

func RegisterDatabase() {
	appContainer.RegisterService("database", new(repository.Gorm))
}

func RegisterCache() {
	//appContainer.RegisterService("cache", new(repository.Cache))
}

func RegisterRest() {
	appContainer.RegisterService("fiber", new(rest.Fiber))
}

func RegisterToggleService() {
}

func RegisterRepository() {
	//appContainer.RegisterService("domainRepository", new(gorm.DomainRepository))
}
