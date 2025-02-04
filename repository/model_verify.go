package repository

import (
	"fmt"
	dto "nexus-ai/dto/model"
	"nexus-ai/model"
)

func ModelVerify(modelName string) (*dto.Model, error) {
	modelRepo := NewModelRepository(model.GetDB())
	model, err := modelRepo.GetByName(modelName)
	if err != nil { // 模型不存在
		return nil, fmt.Errorf("%s not found", modelName)
	}
	if model.Status == 0 { // 0: 禁用
		return nil, fmt.Errorf("%s is not available now", modelName)
	}
	return model, nil
}
