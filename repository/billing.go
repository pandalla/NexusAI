package repository

import (
	"fmt"
	"math/rand"
	"nexus-ai/common"
	dto "nexus-ai/dto/model"
	"nexus-ai/model"
	"nexus-ai/utils"
	"time"

	"gorm.io/gorm"
)

// BillingRepository 账单仓储接口
type BillingRepository interface {
	Create(billing *dto.Billing) error
	Update(billing *dto.Billing) error
	Delete(billingID string) error
	GetByID(billingID string) (*dto.Billing, error)
	GetByBillingNo(billingNo string) (*dto.Billing, error)
	List(page, pageSize int) ([]*dto.Billing, int64, error)
	ListByUser(userID string, page, pageSize int) ([]*dto.Billing, int64, error)
	ListByStatus(status int8, page, pageSize int) ([]*dto.Billing, int64, error)
	ListByType(billingType string, page, pageSize int) ([]*dto.Billing, int64, error)
	ListByCycle(billingCycle string, page, pageSize int) ([]*dto.Billing, int64, error)
	Benchmark(count int) error
}

type billingRepository struct {
	db *gorm.DB
}

// NewBillingRepository 创建账单仓储实例
func NewBillingRepository(db *gorm.DB) BillingRepository {
	return &billingRepository{db: db}
}

// convertToDTO 将数据模型转换为DTO
func (r *billingRepository) convertToDTO(model *model.Billing) *dto.Billing {
	if model == nil {
		return nil
	}

	var billingMoney dto.BillingMoney
	var billingDetails dto.BillingDetails
	var quotaDetails dto.QuotaDetails

	if err := model.BillingMoney.ToStruct(&billingMoney); err != nil {
		utils.SysError("解析账单金额失败:" + err.Error())
	}

	if err := model.BillingDetails.ToStruct(&billingDetails); err != nil {
		utils.SysError("解析账单明细失败:" + err.Error())
	}

	if err := model.QuotaDetails.ToStruct(&quotaDetails); err != nil {
		utils.SysError("解析配额明细失败:" + err.Error())
	}

	var deletedAt *utils.MySQLTime
	if model.DeletedAt.Valid {
		t := utils.MySQLTime(model.DeletedAt.Time)
		deletedAt = &t
	}

	return &dto.Billing{
		BillingID:       model.BillingID,
		UserID:          model.UserID,
		BillingNo:       model.BillingNo,
		BillingType:     model.BillingType,
		BillingCycle:    model.BillingCycle,
		Remark:          model.Remark,
		BillingCurrency: model.BillingCurrency,
		BillingStatus:   model.BillingStatus,
		BillingMoney:    billingMoney,
		BillingDetails:  billingDetails,
		QuotaDetails:    quotaDetails,
		PayTime:         model.PayTime,
		StartTime:       model.StartTime,
		EndTime:         model.EndTime,
		DueTime:         model.DueTime,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
		DeletedAt:       deletedAt,
	}
}

// convertToModel 将DTO转换为数据模型
func (r *billingRepository) convertToModel(dto *dto.Billing) (*model.Billing, error) {
	if dto == nil {
		return nil, nil
	}

	billingMoneyJSON, err := common.FromStruct(dto.BillingMoney)
	if err != nil {
		return nil, fmt.Errorf("转换账单金额失败: %w", err)
	}

	billingDetailsJSON, err := common.FromStruct(dto.BillingDetails)
	if err != nil {
		return nil, fmt.Errorf("转换账单明细失败: %w", err)
	}

	quotaDetailsJSON, err := common.FromStruct(dto.QuotaDetails)
	if err != nil {
		return nil, fmt.Errorf("转换配额明细失败: %w", err)
	}

	var deletedAt gorm.DeletedAt
	if dto.DeletedAt != nil {
		deletedAt.Time = time.Time(*dto.DeletedAt)
		deletedAt.Valid = true
	}

	return &model.Billing{
		BillingID:       dto.BillingID,
		UserID:          dto.UserID,
		BillingNo:       dto.BillingNo,
		BillingType:     dto.BillingType,
		BillingCycle:    dto.BillingCycle,
		Remark:          dto.Remark,
		BillingCurrency: dto.BillingCurrency,
		BillingStatus:   dto.BillingStatus,
		BillingMoney:    billingMoneyJSON,
		BillingDetails:  billingDetailsJSON,
		QuotaDetails:    quotaDetailsJSON,
		PayTime:         dto.PayTime,
		StartTime:       dto.StartTime,
		EndTime:         dto.EndTime,
		DueTime:         dto.DueTime,
		CreatedAt:       dto.CreatedAt,
		UpdatedAt:       dto.UpdatedAt,
		DeletedAt:       deletedAt,
	}, nil
}

// Create 创建账单记录
func (r *billingRepository) Create(billing *dto.Billing) error {
	model, err := r.convertToModel(billing)
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

// Update 更新账单记录
func (r *billingRepository) Update(billing *dto.Billing) error {
	modelData, err := r.convertToModel(billing)
	if err != nil {
		return err
	}
	return r.db.Model(&model.Billing{}).Where("billing_id = ?", billing.BillingID).Updates(modelData).Error
}

// Delete 删除账单记录
func (r *billingRepository) Delete(billingID string) error {
	return r.db.Delete(&model.Billing{}, "billing_id = ?", billingID).Error
}

// GetByID 根据ID获取账单记录
func (r *billingRepository) GetByID(billingID string) (*dto.Billing, error) {
	var billing model.Billing
	if err := r.db.Where("billing_id = ?", billingID).First(&billing).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&billing), nil
}

// GetByBillingNo 根据账单编号获取账单记录
func (r *billingRepository) GetByBillingNo(billingNo string) (*dto.Billing, error) {
	var billing model.Billing
	if err := r.db.Where("billing_no = ?", billingNo).First(&billing).Error; err != nil {
		return nil, err
	}
	return r.convertToDTO(&billing), nil
}

// List 获取账单记录列表
func (r *billingRepository) List(page, pageSize int) ([]*dto.Billing, int64, error) {
	var total int64
	var billings []model.Billing

	offset := (page - 1) * pageSize

	if err := r.db.Model(&model.Billing{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Offset(offset).Limit(pageSize).Find(&billings).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.Billing, len(billings))
	for i, b := range billings {
		dtoList[i] = r.convertToDTO(&b)
	}

	return dtoList, total, nil
}

// ListByUser 获取用户的账单记录列表
func (r *billingRepository) ListByUser(userID string, page, pageSize int) ([]*dto.Billing, int64, error) {
	var total int64
	var billings []model.Billing

	offset := (page - 1) * pageSize
	query := r.db.Model(&model.Billing{}).Where("user_id = ?", userID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&billings).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.Billing, len(billings))
	for i, b := range billings {
		dtoList[i] = r.convertToDTO(&b)
	}

	return dtoList, total, nil
}

// ListByStatus 获取指定状态的账单记录列表
func (r *billingRepository) ListByStatus(status int8, page, pageSize int) ([]*dto.Billing, int64, error) {
	var total int64
	var billings []model.Billing

	offset := (page - 1) * pageSize
	query := r.db.Model(&model.Billing{}).Where("billing_status = ?", status)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&billings).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.Billing, len(billings))
	for i, b := range billings {
		dtoList[i] = r.convertToDTO(&b)
	}

	return dtoList, total, nil
}

// ListByType 获取指定类型的账单记录列表
func (r *billingRepository) ListByType(billingType string, page, pageSize int) ([]*dto.Billing, int64, error) {
	var total int64
	var billings []model.Billing

	offset := (page - 1) * pageSize
	query := r.db.Model(&model.Billing{}).Where("billing_type = ?", billingType)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&billings).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.Billing, len(billings))
	for i, b := range billings {
		dtoList[i] = r.convertToDTO(&b)
	}

	return dtoList, total, nil
}

// ListByCycle 获取指定周期的账单记录列表
func (r *billingRepository) ListByCycle(billingCycle string, page, pageSize int) ([]*dto.Billing, int64, error) {
	var total int64
	var billings []model.Billing

	offset := (page - 1) * pageSize
	query := r.db.Model(&model.Billing{}).Where("billing_cycle = ?", billingCycle)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&billings).Error; err != nil {
		return nil, 0, err
	}

	dtoList := make([]*dto.Billing, len(billings))
	for i, b := range billings {
		dtoList[i] = r.convertToDTO(&b)
	}

	return dtoList, total, nil
}

// Benchmark 执行基准测试
func (r *billingRepository) Benchmark(count int) error {
	utils.SysInfo("开始执行账单记录基准测试...")
	startTime := time.Now()

	for i := 0; i < count; i++ {
		testBilling := &dto.Billing{
			BillingID:       utils.GenerateRandomUUID(12),
			UserID:          utils.GenerateRandomUUID(12),
			BillingNo:       fmt.Sprintf("BILL_%s", utils.GenerateRandomString(16)),
			BillingType:     []string{"consumption", "refund"}[rand.Intn(2)],
			BillingCycle:    []string{"once", "daily", "weekly", "monthly"}[rand.Intn(4)],
			BillingCurrency: []string{"USD", "CNY", "HKD", "EUR"}[rand.Intn(4)],
			BillingStatus:   int8(rand.Intn(4) + 1),
			StartTime:       utils.MySQLTime(time.Now()),
			EndTime:         utils.MySQLTime(time.Now().AddDate(0, 0, rand.Intn(30)+1)),
			BillingMoney: dto.BillingMoney{
				TotalAmount:               float64(rand.Intn(1000000)) / 100,
				DiscountAmount:            float64(rand.Intn(10000)) / 100,
				ActualAmount:              float64(rand.Intn(1000000)) / 100,
				PaidAmount:                float64(rand.Intn(1000000)) / 100,
				UnpaidAmount:              float64(rand.Intn(1000000)) / 100,
				PaymentPlatformCommission: float64(rand.Intn(1000000)) / 100,
				PaymentPlatformFee:        float64(rand.Intn(1000000)) / 100,
				RefundAmount:              float64(rand.Intn(1000000)) / 100,
				OverdueAmount:             float64(rand.Intn(1000000)) / 100,
				OverduePenalty:            float64(rand.Intn(1000000)) / 100,
			},
			BillingDetails: dto.BillingDetails{
				ModelUsage:      map[string]float64{"gpt-3.5-turbo": float64(rand.Intn(100000)) / 100},
				TextUsage:       map[string]float64{"gpt-3.5-turbo": float64(rand.Intn(100000)) / 100},
				ImageUsage:      map[string]float64{"gpt-3.5-turbo": float64(rand.Intn(100000)) / 100},
				AudioUsage:      map[string]float64{"gpt-3.5-turbo": float64(rand.Intn(100000)) / 100},
				VideoUsage:      map[string]float64{"gpt-3.5-turbo": float64(rand.Intn(100000)) / 100},
				MultiMediaUsage: map[string]float64{"gpt-3.5-turbo": float64(rand.Intn(100000)) / 100},
				OtherUsage:      map[string]float64{"gpt-3.5-turbo": float64(rand.Intn(100000)) / 100},
			},
			QuotaDetails: dto.QuotaDetails{
				QuotaID:         utils.GenerateRandomUUID(12),
				QuotaType:       "payment",
				QuotaAmount:     float64(rand.Intn(1000000)) / 100,
				QuotaUsage:      float64(rand.Intn(100000)) / 100,
				QuotaFrozen:     float64(rand.Intn(100000)) / 100,
				QuotaRemaining:  float64(rand.Intn(100000)) / 100,
				QuotaExpired:    float64(rand.Intn(100000)) / 100,
				QuotaExpireTime: utils.MySQLTime(time.Now().AddDate(0, 0, rand.Intn(30)+1)),
			},
		}

		testBilling.BillingMoney.ActualAmount = testBilling.BillingMoney.TotalAmount - testBilling.BillingMoney.DiscountAmount
		testBilling.DueTime = utils.MySQLTime(time.Time(testBilling.EndTime).AddDate(0, 0, 15))

		// 创建
		if err := r.Create(testBilling); err != nil {
			utils.SysError("创建账单记录失败: " + err.Error())
			return err
		}

		// 获取创建后的记录
		createdBilling, err := r.GetByID(testBilling.BillingID)
		if err != nil {
			utils.SysError("获取创建的账单记录失败: " + err.Error())
			return err
		}

		// 更新
		createdBilling.BillingStatus = int8(rand.Intn(4) + 1)
		if err := r.Update(createdBilling); err != nil {
			utils.SysError("更新账单记录失败: " + err.Error())
			return err
		}

		// 删除
		if err := r.Delete(createdBilling.BillingID); err != nil {
			utils.SysError("删除账单记录失败: " + err.Error())
			return err
		}
	}

	duration := time.Since(startTime)
	utils.SysInfo("基准测试完成，总耗时: " + duration.String() + ", 平均每组操作耗时: " + (duration / time.Duration(count)).String())
	return nil
}
