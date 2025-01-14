package test

import (
	"nexus-ai/model"
	"nexus-ai/repository"
)

func TestRepository(num int) {
	TestUserRepository(num)
	TestUserGroupRepository(num)
	TestTokenRepository(num)
}

func TestUserRepository(num int) {
	userRepo := repository.NewUserRepository(model.GetDB())
	userRepo.Benchmark(num)
}

func TestUserGroupRepository(num int) {
	userGroupRepo := repository.NewUserGroupRepository(model.GetDB())
	userGroupRepo.Benchmark(num)
}

func TestTokenRepository(num int) {
	tokenRepo := repository.NewTokenRepository(model.GetDB())
	tokenRepo.Benchmark(num)
}
