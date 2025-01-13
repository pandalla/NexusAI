package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// =================== 通用键操作 ===================
// 常用操作:
// [x] Keys   - 获取匹配模式的所有键
// [x] Exists - 检查键是否存在
// [x] Delete - 删除键
// [x] Expire - 设置过期时间
// [x] TTL    - 获取键的剩余过期时间
// [x] Rename - 重命名键
// [x] Type   - 获取键的数据类型

// Keys 获取匹配模式的所有键
func Keys(ctx context.Context, pattern string) ([]string, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	return client.Keys(ctx, pattern).Result()
}

// Exists 检查键是否存在
func Exists(ctx context.Context, key string) (bool, error) {
	client := GetRedisClient()
	if client == nil {
		return false, fmt.Errorf("redis client is nil")
	}
	n, err := client.Exists(ctx, key).Result()
	return n > 0, err
}

// Delete 删除键
func Delete(ctx context.Context, keys ...string) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.Del(ctx, keys...).Err()
}

// Expire 设置过期时间
func Expire(ctx context.Context, key string, expiration time.Duration) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.Expire(ctx, key, expiration).Err()
}

// TTL 获取键的剩余过期时间
func TTL(ctx context.Context, key string) (time.Duration, error) {
	client := GetRedisClient()
	if client == nil {
		return 0, fmt.Errorf("redis client is nil")
	}
	return client.TTL(ctx, key).Result()
}

// Rename 重命名键
func Rename(ctx context.Context, key, newKey string) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.Rename(ctx, key, newKey).Err()
}

// Type 获取键的数据类型
func Type(ctx context.Context, key string) (string, error) {
	client := GetRedisClient()
	if client == nil {
		return "", fmt.Errorf("redis client is nil")
	}
	return client.Type(ctx, key).Result()
}

// =================== String 操作 ===================
// 常用操作:
// [x] Set - 设置键值对
// [x] Get - 获取键值
// [x] SetObj - 将对象序列化后存储到Redis
// [x] GetObj - 获取并解析为指定类型
// [x] MGet/MSet - 批量操作
// [x] Incr/IncrBy - 递增操作
// [x] Decr/DecrBy - 递减操作
// [x] SetNX - 键不存在时设置
// [x] GetSet - 设置新值并返回旧值
// [x] Append - 追加字符串
// [x] SetEX - 设置值和过期时间

// Set 设置键值对
func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.Set(ctx, key, value, expiration).Err()
}

// Get 获取键值
func Get(ctx context.Context, key string) (string, error) {
	client := GetRedisClient()
	if client == nil {
		return "", fmt.Errorf("redis client is nil")
	}
	return client.Get(ctx, key).Result()
}

// SetObj 将对象序列化后存储到Redis
func SetObj(ctx context.Context, key string, obj interface{}, expiration time.Duration) error {
	// 将对象序列化为JSON字节数组
	bytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("marshal object failed: %v", err)
	}

	// 调用Set存储序列化后的数据
	return Set(ctx, key, string(bytes), expiration)
}

// GetObj 获取并解析为指定类型
func GetObj(ctx context.Context, key string, obj interface{}) error {
	value, err := Get(ctx, key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(value), obj)
}

// MGet 批量获取键值
func MGet(ctx context.Context, keys ...string) ([]interface{}, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	return client.MGet(ctx, keys...).Result()
}

// MSet 批量设置键值对
func MSet(ctx context.Context, values ...interface{}) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.MSet(ctx, values...).Err()
}

// Incr 递增
func Incr(ctx context.Context, key string) (int64, error) {
	client := GetRedisClient()
	if client == nil {
		return 0, fmt.Errorf("redis client is nil")
	}
	return client.Incr(ctx, key).Result()
}

// IncrBy 按指定值递增
func IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	client := GetRedisClient()
	if client == nil {
		return 0, fmt.Errorf("redis client is nil")
	}
	return client.IncrBy(ctx, key, value).Result()
}

// Decr 递减
func Decr(ctx context.Context, key string) (int64, error) {
	client := GetRedisClient()
	if client == nil {
		return 0, fmt.Errorf("redis client is nil")
	}
	return client.Decr(ctx, key).Result()
}

// DecrBy 按指定值递减
func DecrBy(ctx context.Context, key string, value int64) (int64, error) {
	client := GetRedisClient()
	if client == nil {
		return 0, fmt.Errorf("redis client is nil")
	}
	return client.DecrBy(ctx, key, value).Result()
}

// SetNX 键不存在时设置
func SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	client := GetRedisClient()
	if client == nil {
		return false, fmt.Errorf("redis client is nil")
	}
	return client.SetNX(ctx, key, value, expiration).Result()
}

// GetSet 设置新值并返回旧值
func GetSet(ctx context.Context, key string, value interface{}) (string, error) {
	client := GetRedisClient()
	if client == nil {
		return "", fmt.Errorf("redis client is nil")
	}
	return client.GetSet(ctx, key, value).Result()
}

// Append 追加字符串
func Append(ctx context.Context, key, value string) (int64, error) {
	client := GetRedisClient()
	if client == nil {
		return 0, fmt.Errorf("redis client is nil")
	}
	return client.Append(ctx, key, value).Result()
}

// SetEX 设置值和过期时间
func SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return Set(ctx, key, value, expiration)
}

// =================== List 操作 ===================
// 常用操作:
// [x] LPush/RPush - 左/右端推入
// [x] LPop/RPop - 左/右端弹出
// [x] BLPop/BRPop - 阻塞式弹出
// [x] LLen - 获取长度
// [x] LRange - 获取指定范围元素
// [x] LRem - 删除指定元素
// [x] LIndex - 获取指定位置元素
// [x] LSet - 设置指定位置元素
// [x] LTrim - 保留指定范围的元素
// [x] LInsert - 在指定位置插入元素

// LPush 从列表左端推入元素
func LPush(ctx context.Context, key string, values ...interface{}) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.LPush(ctx, key, values...).Err()
}

// RPush 从列表右端推入元素
func RPush(ctx context.Context, key string, values ...interface{}) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.RPush(ctx, key, values...).Err()
}

// LPop 从列表左端弹出元素
func LPop(ctx context.Context, key string) (string, error) {
	client := GetRedisClient()
	if client == nil {
		return "", fmt.Errorf("redis client is nil")
	}
	return client.LPop(ctx, key).Result()
}

// RPop 从列表右端弹出元素
func RPop(ctx context.Context, key string) (string, error) {
	client := GetRedisClient()
	if client == nil {
		return "", fmt.Errorf("redis client is nil")
	}
	return client.RPop(ctx, key).Result()
}

// BLPop 阻塞式从列表左端弹出元素
func BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	return client.BLPop(ctx, timeout, keys...).Result()
}

// BRPop 阻塞式从列表右端弹出元素
func BRPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	return client.BRPop(ctx, timeout, keys...).Result()
}

// LLen 获取列表长度
func LLen(ctx context.Context, key string) (int64, error) {
	client := GetRedisClient()
	if client == nil {
		return 0, fmt.Errorf("redis client is nil")
	}
	return client.LLen(ctx, key).Result()
}

// LRange 获取列表指定范围的元素
func LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	return client.LRange(ctx, key, start, stop).Result()
}

// LRem 从列表中删除指定元素
func LRem(ctx context.Context, key string, count int64, value interface{}) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.LRem(ctx, key, count, value).Err()
}

// LIndex 获取列表中指定位置的元素
func LIndex(ctx context.Context, key string, index int64) (string, error) {
	client := GetRedisClient()
	if client == nil {
		return "", fmt.Errorf("redis client is nil")
	}
	return client.LIndex(ctx, key, index).Result()
}

// LSet 设置列表中指定位置的元素值
func LSet(ctx context.Context, key string, index int64, value interface{}) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.LSet(ctx, key, index, value).Err()
}

// LTrim 保留列表中指定范围的元素
func LTrim(ctx context.Context, key string, start, stop int64) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.LTrim(ctx, key, start, stop).Err()
}

// LInsert 在列表中指定位置插入元素
func LInsert(ctx context.Context, key string, before bool, pivot, value interface{}) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	position := "AFTER"
	if before {
		position = "BEFORE"
	}
	return client.LInsert(ctx, key, position, pivot, value).Err()
}

// =================== Hash 操作 ===================
// 常用操作:
// [x] HSet - 设置字段值
// [x] HGet - 获取字段值
// [x] HGetAll - 获取所有字段和值
// [x] HDel - 删除字段
// [x] HExists - 检查字段是否存在
// [x] HIncrBy - 字段值递增
// [x] HKeys - 获取所有字段
// [x] HVals - 获取所有值
// [x] HMGet - 批量获取字段
// [x] HScan - 迭代哈希表
// [x] HSetNX - 字段不存在时设置

// HSet 设置哈希表字段
func HSet(ctx context.Context, key, field string, value interface{}) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.HSet(ctx, key, field, value).Err()
}

// HGet 获取哈希表字段
func HGet(ctx context.Context, key, field string) (string, error) {
	client := GetRedisClient()
	if client == nil {
		return "", fmt.Errorf("redis client is nil")
	}
	return client.HGet(ctx, key, field).Result()
}

// HGetAll 获取哈希表所有字段和值
func HGetAll(ctx context.Context, key string) (map[string]string, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	return client.HGetAll(ctx, key).Result()
}

// HDel 删除哈希表字段
func HDel(ctx context.Context, key string, fields ...string) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.HDel(ctx, key, fields...).Err()
}

// HExists 检查哈希表字段是否存在
func HExists(ctx context.Context, key, field string) (bool, error) {
	client := GetRedisClient()
	if client == nil {
		return false, fmt.Errorf("redis client is nil")
	}
	return client.HExists(ctx, key, field).Result()
}

// HIncrBy 哈希表字段值递增
func HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
	client := GetRedisClient()
	if client == nil {
		return 0, fmt.Errorf("redis client is nil")
	}
	return client.HIncrBy(ctx, key, field, incr).Result()
}

// HKeys 获取哈希表所有字段
func HKeys(ctx context.Context, key string) ([]string, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	return client.HKeys(ctx, key).Result()
}

// HVals 获取哈希表所有值
func HVals(ctx context.Context, key string) ([]string, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	return client.HVals(ctx, key).Result()
}

// HMGet 批量获取哈希表字段
func HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	return client.HMGet(ctx, key, fields...).Result()
}

// HScan 迭代哈希表
func HScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, 0, fmt.Errorf("redis client is nil")
	}
	return client.HScan(ctx, key, cursor, match, count).Result()
}

// HSetNX 设置哈希表字段，仅当字段不存在时
func HSetNX(ctx context.Context, key, field string, value interface{}) (bool, error) {
	client := GetRedisClient()
	if client == nil {
		return false, fmt.Errorf("redis client is nil")
	}
	return client.HSetNX(ctx, key, field, value).Result()
}

// =================== Set 操作 ===================
// 常用操作:
// [x] SAdd - 添加成员
// [x] SMembers - 获取所有成员
// [x] SIsMember - 判断是否为成员
// [x] SRem - 删除成员
// [x] SCard - 获取成员数量
// [x] SInter - 交集
// [x] SUnion - 并集
// [x] SDiff - 差集
// [x] SPop - 随机弹出成员
// [x] SRandMember - 随机获取成员
// [x] SScan - 迭代集合

// SAdd 添加集合成员
func SAdd(ctx context.Context, key string, members ...interface{}) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.SAdd(ctx, key, members...).Err()
}

// SMembers 获取集合所有成员
func SMembers(ctx context.Context, key string) ([]string, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	return client.SMembers(ctx, key).Result()
}

// SIsMember 判断成员是否在集合中
func SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	client := GetRedisClient()
	if client == nil {
		return false, fmt.Errorf("redis client is nil")
	}
	return client.SIsMember(ctx, key, member).Result()
}

// SRem 移除集合成员
func SRem(ctx context.Context, key string, members ...interface{}) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.SRem(ctx, key, members...).Err()
}

// SCard 获取集合成员数量
func SCard(ctx context.Context, key string) (int64, error) {
	client := GetRedisClient()
	if client == nil {
		return 0, fmt.Errorf("redis client is nil")
	}
	return client.SCard(ctx, key).Result()
}

// SInter 获取多个集合的交集
func SInter(ctx context.Context, keys ...string) ([]string, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	return client.SInter(ctx, keys...).Result()
}

// SUnion 获取多个集合的并集
func SUnion(ctx context.Context, keys ...string) ([]string, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	return client.SUnion(ctx, keys...).Result()
}

// SDiff 获取多个集合的差集
func SDiff(ctx context.Context, keys ...string) ([]string, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	return client.SDiff(ctx, keys...).Result()
}

// SPop 随机移除并返回集合中的一个成员
func SPop(ctx context.Context, key string) (string, error) {
	client := GetRedisClient()
	if client == nil {
		return "", fmt.Errorf("redis client is nil")
	}
	return client.SPop(ctx, key).Result()
}

// SRandMember 随机返回集合中的一个成员
func SRandMember(ctx context.Context, key string) (string, error) {
	client := GetRedisClient()
	if client == nil {
		return "", fmt.Errorf("redis client is nil")
	}
	return client.SRandMember(ctx, key).Result()
}

// SScan 迭代集合
func SScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, 0, fmt.Errorf("redis client is nil")
	}
	return client.SScan(ctx, key, cursor, match, count).Result()
}

// =================== Sorted Set 操作 ===================
// 常用操作:
// [x] ZAdd - 添加成员
// [x] ZRange - 按位置获取成员
// [x] ZRank - 获取成员排名
// [x] ZScore - 获取成员分数
// [x] ZRevRange - 按位置倒序获取
// [x] ZIncrBy - 增加成员分数
// [x] ZCount - 指定分数范围的成员数
// [x] ZRemRangeByRank - 按排名删除成员
// [x] ZRangeByScore - 按分数范围获取成员
// [x] ZScan - 迭代有序集合

// ZAdd 添加有序集合成员
func ZAdd(ctx context.Context, key string, members ...redis.Z) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.ZAdd(ctx, key, members...).Err()
}

// ZRange 获取有序集合指定范围的成员
func ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	return client.ZRange(ctx, key, start, stop).Result()
}

// ZRank 获取有序集合成员的排名
func ZRank(ctx context.Context, key, member string) (int64, error) {
	client := GetRedisClient()
	if client == nil {
		return 0, fmt.Errorf("redis client is nil")
	}
	return client.ZRank(ctx, key, member).Result()
}

// ZScore 获取有序集合成员的分数
func ZScore(ctx context.Context, key, member string) (float64, error) {
	client := GetRedisClient()
	if client == nil {
		return 0, fmt.Errorf("redis client is nil")
	}
	return client.ZScore(ctx, key, member).Result()
}

// ZRevRange 获取有序集合指定范围的成员(倒序)
func ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	return client.ZRevRange(ctx, key, start, stop).Result()
}

// ZIncrBy 增加有序集合成员的分数
func ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error) {
	client := GetRedisClient()
	if client == nil {
		return 0, fmt.Errorf("redis client is nil")
	}
	return client.ZIncrBy(ctx, key, increment, member).Result()
}

// ZCount 获取有序集合指定分数范围的成员数量
func ZCount(ctx context.Context, key, min, max string) (int64, error) {
	client := GetRedisClient()
	if client == nil {
		return 0, fmt.Errorf("redis client is nil")
	}
	return client.ZCount(ctx, key, min, max).Result()
}

// ZRemRangeByRank 删除有序集合指定排名范围的成员
func ZRemRangeByRank(ctx context.Context, key string, start, stop int64) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.ZRemRangeByRank(ctx, key, start, stop).Err()
}

// ZRangeByScore 获取有序集合指定分数范围的成员
func ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	return client.ZRangeByScore(ctx, key, opt).Result()
}

// ZScan 迭代有序集合
func ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, 0, fmt.Errorf("redis client is nil")
	}
	return client.ZScan(ctx, key, cursor, match, count).Result()
}

// =================== 分布式锁操作 ===================
// 常用操作:
// [x] TryLock - 获取分布式锁
// [x] Unlock - 释放分布式锁
// [x] RenewLock - 续期分布式锁
// [x] GetLockTTL - 获取锁剩余时间

// TryLock 尝试获取分布式锁
func TryLock(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	client := GetRedisClient()
	if client == nil {
		return false, fmt.Errorf("redis client is nil")
	}
	return client.SetNX(ctx, key, value, expiration).Result()
}

// Unlock 释放分布式锁
func Unlock(ctx context.Context, key string) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis client is nil")
	}
	return client.Del(ctx, key).Err()
}

// RenewLock 续期分布式锁
func RenewLock(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	client := GetRedisClient()
	if client == nil {
		return false, fmt.Errorf("redis client is nil")
	}

	// 使用 EXPIRE 命令续期
	ok, err := client.Expire(ctx, key, expiration).Result()
	if err != nil {
		return false, fmt.Errorf("failed to renew lock: %v", err)
	}
	return ok, nil
}

// GetLockTTL 获取锁的剩余过期时间
func GetLockTTL(ctx context.Context, key string) (time.Duration, error) {
	client := GetRedisClient()
	if client == nil {
		return 0, fmt.Errorf("redis client is nil")
	}
	return client.TTL(ctx, key).Result()
}
