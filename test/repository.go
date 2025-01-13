package test

import (
	"nexus-ai/model"
	"nexus-ai/repository"
)

func TestUserRepository() {
	userRepo := repository.NewUserRepository(model.GetDB())
	userRepo.Benchmark(10)
}

func TestUserGroupRepository() {
	userGroupRepo := repository.NewUserGroupRepository(model.GetDB())
	userGroupRepo.Benchmark(10)
}
