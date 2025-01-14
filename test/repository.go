package test

import (
	"nexus-ai/model"
	"nexus-ai/repository"
)

func TestRepository(num int) {
	TestUserRepository(num)
	TestUserGroupRepository(num)
	TestTokenRepository(num)
	TestModelRepository(num)
	TestChannelRepository(num)
	TestModelGroupRepository(num)
	TestChannelGroupRepository(num)
	TestUsageRepository(num)
	TestPaymentRepository(num)
	TestQuotaRepository(num)
	TestBillingRepository(num)
	TestMessageSaveRepository(num)
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

func TestModelRepository(num int) {
	modelRepo := repository.NewModelRepository(model.GetDB())
	modelRepo.Benchmark(num)
}

func TestChannelRepository(num int) {
	channelRepo := repository.NewChannelRepository(model.GetDB())
	channelRepo.Benchmark(num)
}

func TestModelGroupRepository(num int) {
	modelGroupRepo := repository.NewModelGroupRepository(model.GetDB())
	modelGroupRepo.Benchmark(num)
}

func TestChannelGroupRepository(num int) {
	channelGroupRepo := repository.NewChannelGroupRepository(model.GetDB())
	channelGroupRepo.Benchmark(num)
}

func TestUsageRepository(num int) {
	usageRepo := repository.NewUsageRepository(model.GetDB())
	usageRepo.Benchmark(num)
}

func TestPaymentRepository(num int) {
	paymentRepo := repository.NewPaymentRepository(model.GetDB())
	paymentRepo.Benchmark(num)
}

func TestQuotaRepository(num int) {
	quotaRepo := repository.NewQuotaRepository(model.GetDB())
	quotaRepo.Benchmark(num)
}

func TestBillingRepository(num int) {
	billingRepo := repository.NewBillingRepository(model.GetDB())
	billingRepo.Benchmark(num)
}

func TestMessageSaveRepository(num int) {
	messageSaveRepo := repository.NewMessageSaveRepository(model.GetDB())
	messageSaveRepo.Benchmark(num)
}
