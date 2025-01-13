package redis

import (
	"context"
	"fmt"
	"nexus-ai/utils"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

// =================== 通用键操作基准测试 ===================
// 包含以下操作的基准测试:
// - Keys   - 获取匹配模式的所有键
// - Exists - 检查键是否存在
// - Delete - 删除键
// - Expire - 设置过期时间
// - TTL    - 获取键的剩余过期时间
// - Rename - 重命名键
// - Type   - 获取键的数据类型

func BenchmarkKeyOperations(b *testing.B) {
	// 直接执行子测试，而不是使用 b.Run
	benchKeys(b)
	utils.SysInfo("Keys benchmark completed")

	benchExists(b)
	utils.SysInfo("Exists benchmark completed")

	benchDelete(b)
	utils.SysInfo("Delete benchmark completed")

	benchExpire(b)
	utils.SysInfo("Expire benchmark completed")

	benchTTL(b)
	utils.SysInfo("TTL benchmark completed")

	benchRename(b)
	utils.SysInfo("Rename benchmark completed")

	benchType(b)
	utils.SysInfo("Type benchmark completed")
}

func benchKeys(b *testing.B) {
	ctx := context.Background()

	// 准备测试数据
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("test:key:%d", i)
		if err := Set(ctx, key, "value", 0); err != nil {
			b.Fatal(err)
		}
	}

	// 重置计时器
	b.ResetTimer()

	// 执行测试
	for i := 0; i < b.N; i++ {
		if _, err := Keys(ctx, "test:key:*"); err != nil {
			b.Fatal(err)
		}
	}
}

func benchExists(b *testing.B) {
	ctx := context.Background()
	key := "test:exists"

	// 准备测试数据
	if err := Set(ctx, key, "value", 0); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Exists(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchDelete(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("test:delete:%d", i)
		if err := Set(ctx, key, "value", 0); err != nil {
			b.Fatal(err)
		}
		if err := Delete(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchExpire(b *testing.B) {
	ctx := context.Background()
	key := "test:expire"

	// 准备测试数据
	if err := Set(ctx, key, "value", 0); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := Expire(ctx, key, time.Hour); err != nil {
			b.Fatal(err)
		}
	}
}

func benchTTL(b *testing.B) {
	ctx := context.Background()
	key := "test:ttl"

	// 准备测试数据
	if err := Set(ctx, key, "value", time.Hour); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := TTL(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchRename(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("test:rename:old:%d", i)
		newKey := fmt.Sprintf("test:rename:new:%d", i)

		if err := Set(ctx, key, "value", 0); err != nil {
			b.Fatal(err)
		}
		if err := Rename(ctx, key, newKey); err != nil {
			b.Fatal(err)
		}
	}
}

func benchType(b *testing.B) {
	ctx := context.Background()
	key := "test:type"

	// 准备测试数据
	if err := Set(ctx, key, "value", 0); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Type(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

// =================== String操作基准测试 ===================
// 包含以下操作的基准测试:
// - Set/Get - 基本键值对操作
// - SetObj/GetObj - 对象序列化存储/获取
// - MGet/MSet - 批量操作
// - Incr/IncrBy - 递增操作
// - Decr/DecrBy - 递减操作
// - SetNX - 键不存在时设置
// - GetSet - 设置新值并返回旧值
// - Append - 追加字符串
// - SetEX - 设置值和过期时间

func BenchmarkStringOperations(b *testing.B) {
	benchMSet(b)
	utils.SysInfo("MSet benchmark completed")

	benchGet(b)
	utils.SysInfo("Get benchmark completed")

	benchGetObj(b)
	utils.SysInfo("GetObj benchmark completed")

	benchMGet(b)
	utils.SysInfo("MGet benchmark completed")

	benchIncr(b)
	utils.SysInfo("Incr benchmark completed")

	benchIncrBy(b)
	utils.SysInfo("IncrBy benchmark completed")

	benchDecr(b)
	utils.SysInfo("Decr benchmark completed")

	benchDecrBy(b)
	utils.SysInfo("DecrBy benchmark completed")

	benchSetNX(b)
	utils.SysInfo("SetNX benchmark completed")

	benchGetSet(b)
	utils.SysInfo("GetSet benchmark completed")

	benchAppend(b)
	utils.SysInfo("Append benchmark completed")

	benchSetEX(b)
	utils.SysInfo("SetEX benchmark completed")
}

type testStruct struct {
	Name  string
	Value int
}

func benchGet(b *testing.B) {
	ctx := context.Background()
	key := "test:get"
	if err := Set(ctx, key, "value", 0); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Get(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchGetObj(b *testing.B) {
	ctx := context.Background()
	key := "test:getobj"
	obj := testStruct{Name: "test", Value: 123}
	if err := SetObj(ctx, key, obj, 0); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result testStruct
		if err := GetObj(ctx, key, &result); err != nil {
			b.Fatal(err)
		}
	}
}

func benchMGet(b *testing.B) {
	ctx := context.Background()
	keys := []string{"test:mget1", "test:mget2"}
	if err := MSet(ctx, "test:mget1", "value1", "test:mget2", "value2"); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := MGet(ctx, keys...); err != nil {
			b.Fatal(err)
		}
	}
}

func benchMSet(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := MSet(ctx, fmt.Sprintf("test:mset1:%d", i), "value1",
			fmt.Sprintf("test:mset2:%d", i), "value2"); err != nil {
			b.Fatal(err)
		}
	}
}

func benchIncr(b *testing.B) {
	ctx := context.Background()
	key := "test:incr"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Incr(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchIncrBy(b *testing.B) {
	ctx := context.Background()
	key := "test:incrby"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := IncrBy(ctx, key, 2); err != nil {
			b.Fatal(err)
		}
	}
}

func benchDecr(b *testing.B) {
	ctx := context.Background()
	key := "test:decr"
	if err := Set(ctx, key, 10000, 0); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Decr(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchDecrBy(b *testing.B) {
	ctx := context.Background()
	key := "test:decrby"
	if err := Set(ctx, key, 10000, 0); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := DecrBy(ctx, key, 2); err != nil {
			b.Fatal(err)
		}
	}
}

func benchSetNX(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("test:setnx:%d", i)
		if _, err := SetNX(ctx, key, "value", time.Hour); err != nil {
			b.Fatal(err)
		}
	}
}

func benchGetSet(b *testing.B) {
	ctx := context.Background()
	key := "test:getset"
	if err := Set(ctx, key, "old", 0); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := GetSet(ctx, key, "new"); err != nil {
			b.Fatal(err)
		}
	}
}

func benchAppend(b *testing.B) {
	ctx := context.Background()
	key := "test:append"
	if err := Set(ctx, key, "hello", 0); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Append(ctx, key, "world"); err != nil {
			b.Fatal(err)
		}
	}
}

func benchSetEX(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("test:setex:%d", i)
		if err := SetEX(ctx, key, "value", time.Hour); err != nil {
			b.Fatal(err)
		}
	}
}

// =================== List操作基准测试 ===================
// 包含以下操作的基准测试:
// - LPush/RPush - 左/右端推入
// - LPop/RPop - 左/右端弹出
// - BLPop/BRPop - 阻塞式弹出
// - LLen - 获取长度
// - LRange - 获取指定范围元素
// - LRem - 删除指定元素
// - LIndex - 获取指定位置元素
// - LSet - 设置指定位置元素
// - LTrim - 保留指定范围的元素
// - LInsert - 在指定位置插入元素

func BenchmarkListOperations(b *testing.B) {
	benchLPush(b)
	utils.SysInfo("LPush benchmark completed")

	benchRPush(b)
	utils.SysInfo("RPush benchmark completed")

	benchLPop(b)
	utils.SysInfo("LPop benchmark completed")

	benchRPop(b)
	utils.SysInfo("RPop benchmark completed")

	benchBLPop(b)
	utils.SysInfo("BLPop benchmark completed")

	benchBRPop(b)
	utils.SysInfo("BRPop benchmark completed")

	benchLLen(b)
	utils.SysInfo("LLen benchmark completed")

	benchLRange(b)
	utils.SysInfo("LRange benchmark completed")

	benchLRem(b)
	utils.SysInfo("LRem benchmark completed")

	benchLIndex(b)
	utils.SysInfo("LIndex benchmark completed")

	benchLSet(b)
	utils.SysInfo("LSet benchmark completed")

	benchLTrim(b)
	utils.SysInfo("LTrim benchmark completed")

	benchLInsert(b)
	utils.SysInfo("LInsert benchmark completed")
}

func benchLPush(b *testing.B) {
	ctx := context.Background()
	key := "test:lpush"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := LPush(ctx, key, fmt.Sprintf("value%d", i)); err != nil {
			b.Fatal(err)
		}
	}
}

func benchRPush(b *testing.B) {
	ctx := context.Background()
	key := "test:rpush"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := RPush(ctx, key, fmt.Sprintf("value%d", i)); err != nil {
			b.Fatal(err)
		}
	}
}

func benchLPop(b *testing.B) {
	ctx := context.Background()
	key := "test:lpop"

	// 准备测试数据
	for i := 0; i < b.N; i++ {
		if err := RPush(ctx, key, fmt.Sprintf("value%d", i)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := LPop(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchRPop(b *testing.B) {
	ctx := context.Background()
	key := "test:rpop"

	// 准备测试数据
	for i := 0; i < b.N; i++ {
		if err := LPush(ctx, key, fmt.Sprintf("value%d", i)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := RPop(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchBLPop(b *testing.B) {
	ctx := context.Background()
	key := "test:blpop"

	// 准备测试数据
	for i := 0; i < b.N; i++ {
		if err := RPush(ctx, key, fmt.Sprintf("value%d", i)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := BLPop(ctx, time.Second, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchBRPop(b *testing.B) {
	ctx := context.Background()
	key := "test:brpop"

	// 准备测试数据
	for i := 0; i < b.N; i++ {
		if err := LPush(ctx, key, fmt.Sprintf("value%d", i)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := BRPop(ctx, time.Second, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchLLen(b *testing.B) {
	ctx := context.Background()
	key := "test:llen"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		if err := RPush(ctx, key, fmt.Sprintf("value%d", i)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := LLen(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchLRange(b *testing.B) {
	ctx := context.Background()
	key := "test:lrange"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		if err := RPush(ctx, key, fmt.Sprintf("value%d", i)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := LRange(ctx, key, 0, 10); err != nil {
			b.Fatal(err)
		}
	}
}

func benchLRem(b *testing.B) {
	ctx := context.Background()
	key := "test:lrem"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 先插入
		if err := RPush(ctx, key, "value"); err != nil {
			b.Fatal(err)
		}
		// 再删除
		if err := LRem(ctx, key, 1, "value"); err != nil {
			b.Fatal(err)
		}
	}
}

func benchLIndex(b *testing.B) {
	ctx := context.Background()
	key := "test:lindex"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		if err := RPush(ctx, key, fmt.Sprintf("value%d", i)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := LIndex(ctx, key, 50); err != nil {
			b.Fatal(err)
		}
	}
}

func benchLSet(b *testing.B) {
	ctx := context.Background()
	key := "test:lset"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		if err := RPush(ctx, key, fmt.Sprintf("value%d", i)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := LSet(ctx, key, 50, fmt.Sprintf("newvalue%d", i)); err != nil {
			b.Fatal(err)
		}
	}
}

func benchLTrim(b *testing.B) {
	ctx := context.Background()
	key := "test:ltrim"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 先插入100个元素
		for j := 0; j < 100; j++ {
			if err := RPush(ctx, key, fmt.Sprintf("value%d", j)); err != nil {
				b.Fatal(err)
			}
		}
		// 只保留前10个
		if err := LTrim(ctx, key, 0, 9); err != nil {
			b.Fatal(err)
		}
	}
}

func benchLInsert(b *testing.B) {
	ctx := context.Background()
	key := "test:linsert"

	// 准备测试数据
	if err := RPush(ctx, key, "pivot"); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := LInsert(ctx, key, true, "pivot", fmt.Sprintf("value%d", i)); err != nil {
			b.Fatal(err)
		}
	}
}

// =================== Hash操作基准测试 ===================
// 包含以下操作的基准测试:
// - HSet/HGet - 基本字段操作
// - HGetAll - 获取所有字段和值
// - HDel - 删除字段
// - HExists - 检查字段是否存在
// - HIncrBy - 字段值递增
// - HKeys/HVals - 获取所有字段/值
// - HMGet - 批量获取字段
// - HScan - 迭代哈希表
// - HSetNX - 字段不存在时设置

func BenchmarkHashOperations(b *testing.B) {
	benchHSet(b)
	utils.SysInfo("HSet benchmark completed")

	benchHGet(b)
	utils.SysInfo("HGet benchmark completed")

	benchHGetAll(b)
	utils.SysInfo("HGetAll benchmark completed")

	benchHDel(b)
	utils.SysInfo("HDel benchmark completed")

	benchHExists(b)
	utils.SysInfo("HExists benchmark completed")

	benchHIncrBy(b)
	utils.SysInfo("HIncrBy benchmark completed")

	benchHKeys(b)
	utils.SysInfo("HKeys benchmark completed")

	benchHVals(b)
	utils.SysInfo("HVals benchmark completed")

	benchHMGet(b)
	utils.SysInfo("HMGet benchmark completed")

	benchHScan(b)
	utils.SysInfo("HScan benchmark completed")

	benchHSetNX(b)
	utils.SysInfo("HSetNX benchmark completed")
}

func benchHSet(b *testing.B) {
	ctx := context.Background()
	key := "test:hset"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field := fmt.Sprintf("field%d", i)
		if err := HSet(ctx, key, field, fmt.Sprintf("value%d", i)); err != nil {
			b.Fatal(err)
		}
	}
}

func benchHGet(b *testing.B) {
	ctx := context.Background()
	key := "test:hget"
	field := "field"

	// 准备测试数据
	if err := HSet(ctx, key, field, "value"); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := HGet(ctx, key, field); err != nil {
			b.Fatal(err)
		}
	}
}

func benchHGetAll(b *testing.B) {
	ctx := context.Background()
	key := "test:hgetall"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		if err := HSet(ctx, key, fmt.Sprintf("field%d", i), fmt.Sprintf("value%d", i)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := HGetAll(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchHDel(b *testing.B) {
	ctx := context.Background()
	key := "test:hdel"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field := fmt.Sprintf("field%d", i)
		// 先设置
		if err := HSet(ctx, key, field, "value"); err != nil {
			b.Fatal(err)
		}
		// 再删除
		if err := HDel(ctx, key, field); err != nil {
			b.Fatal(err)
		}
	}
}

func benchHExists(b *testing.B) {
	ctx := context.Background()
	key := "test:hexists"
	field := "field"

	// 准备测试数据
	if err := HSet(ctx, key, field, "value"); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := HExists(ctx, key, field); err != nil {
			b.Fatal(err)
		}
	}
}

func benchHIncrBy(b *testing.B) {
	ctx := context.Background()
	key := "test:hincrby"
	field := "counter"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := HIncrBy(ctx, key, field, 1); err != nil {
			b.Fatal(err)
		}
	}
}

func benchHKeys(b *testing.B) {
	ctx := context.Background()
	key := "test:hkeys"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		if err := HSet(ctx, key, fmt.Sprintf("field%d", i), "value"); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := HKeys(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchHVals(b *testing.B) {
	ctx := context.Background()
	key := "test:hvals"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		if err := HSet(ctx, key, fmt.Sprintf("field%d", i), fmt.Sprintf("value%d", i)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := HVals(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchHMGet(b *testing.B) {
	ctx := context.Background()
	key := "test:hmget"
	fields := []string{"field1", "field2", "field3"}

	// 准备测试数据
	for _, field := range fields {
		if err := HSet(ctx, key, field, "value"); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := HMGet(ctx, key, fields...); err != nil {
			b.Fatal(err)
		}
	}
}

func benchHScan(b *testing.B) {
	ctx := context.Background()
	key := "test:hscan"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		if err := HSet(ctx, key, fmt.Sprintf("field%d", i), "value"); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var cursor uint64
		for {
			var err error
			_, cursor, err = HScan(ctx, key, cursor, "*", 10)
			if err != nil {
				b.Fatal(err)
			}
			if cursor == 0 {
				break
			}
		}
	}
}

func benchHSetNX(b *testing.B) {
	ctx := context.Background()
	key := "test:hsetnx"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field := fmt.Sprintf("field%d", i)
		if _, err := HSetNX(ctx, key, field, "value"); err != nil {
			b.Fatal(err)
		}
	}
}

// =================== Set操作基准测试 ===================
// 包含以下操作的基准测试:
// - SAdd - 添加成员
// - SMembers - 获取所有成员
// - SIsMember - 判断是否为成员
// - SRem - 删除成员
// - SCard - 获取成员数量
// - SInter - 交集
// - SUnion - 并集
// - SDiff - 差集
// - SPop - 随机弹出成员
// - SRandMember - 随机获取成员
// - SScan - 迭代集合

func BenchmarkSetOperations(b *testing.B) {
	benchSAdd(b)
	utils.SysInfo("SAdd benchmark completed")

	benchSMembers(b)
	utils.SysInfo("SMembers benchmark completed")

	benchSIsMember(b)
	utils.SysInfo("SIsMember benchmark completed")

	benchSRem(b)
	utils.SysInfo("SRem benchmark completed")

	benchSCard(b)
	utils.SysInfo("SCard benchmark completed")

	benchSInter(b)
	utils.SysInfo("SInter benchmark completed")

	benchSUnion(b)
	utils.SysInfo("SUnion benchmark completed")

	benchSDiff(b)
	utils.SysInfo("SDiff benchmark completed")

	benchSPop(b)
	utils.SysInfo("SPop benchmark completed")

	benchSRandMember(b)
	utils.SysInfo("SRandMember benchmark completed")

	benchSScan(b)
	utils.SysInfo("SScan benchmark completed")
}

func benchSAdd(b *testing.B) {
	ctx := context.Background()
	key := "test:sadd"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := SAdd(ctx, key, fmt.Sprintf("member%d", i)); err != nil {
			b.Fatal(err)
		}
	}
}

func benchSMembers(b *testing.B) {
	ctx := context.Background()
	key := "test:smembers"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		if err := SAdd(ctx, key, fmt.Sprintf("member%d", i)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := SMembers(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchSIsMember(b *testing.B) {
	ctx := context.Background()
	key := "test:sismember"
	member := "member"

	// 准备测试数据
	if err := SAdd(ctx, key, member); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := SIsMember(ctx, key, member); err != nil {
			b.Fatal(err)
		}
	}
}

func benchSRem(b *testing.B) {
	ctx := context.Background()
	key := "test:srem"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		member := fmt.Sprintf("member%d", i)
		// 先添加
		if err := SAdd(ctx, key, member); err != nil {
			b.Fatal(err)
		}
		// 再删除
		if err := SRem(ctx, key, member); err != nil {
			b.Fatal(err)
		}
	}
}

func benchSCard(b *testing.B) {
	ctx := context.Background()
	key := "test:scard"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		if err := SAdd(ctx, key, fmt.Sprintf("member%d", i)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := SCard(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchSInter(b *testing.B) {
	ctx := context.Background()
	key1 := "test:sinter1"
	key2 := "test:sinter2"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		if err := SAdd(ctx, key1, fmt.Sprintf("member%d", i)); err != nil {
			b.Fatal(err)
		}
		if err := SAdd(ctx, key2, fmt.Sprintf("member%d", i*2)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := SInter(ctx, key1, key2); err != nil {
			b.Fatal(err)
		}
	}
}

func benchSUnion(b *testing.B) {
	ctx := context.Background()
	key1 := "test:sunion1"
	key2 := "test:sunion2"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		if err := SAdd(ctx, key1, fmt.Sprintf("member%d", i)); err != nil {
			b.Fatal(err)
		}
		if err := SAdd(ctx, key2, fmt.Sprintf("member%d", i*2)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := SUnion(ctx, key1, key2); err != nil {
			b.Fatal(err)
		}
	}
}

func benchSDiff(b *testing.B) {
	ctx := context.Background()
	key1 := "test:sdiff1"
	key2 := "test:sdiff2"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		if err := SAdd(ctx, key1, fmt.Sprintf("member%d", i)); err != nil {
			b.Fatal(err)
		}
		if err := SAdd(ctx, key2, fmt.Sprintf("member%d", i*2)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := SDiff(ctx, key1, key2); err != nil {
			b.Fatal(err)
		}
	}
}

func benchSPop(b *testing.B) {
	ctx := context.Background()
	key := "test:spop"

	// 准备测试数据
	for i := 0; i < b.N; i++ {
		if err := SAdd(ctx, key, fmt.Sprintf("member%d", i)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := SPop(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchSRandMember(b *testing.B) {
	ctx := context.Background()
	key := "test:srandmember"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		if err := SAdd(ctx, key, fmt.Sprintf("member%d", i)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := SRandMember(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchSScan(b *testing.B) {
	ctx := context.Background()
	key := "test:sscan"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		if err := SAdd(ctx, key, fmt.Sprintf("member%d", i)); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var cursor uint64
		for {
			_, cursor, err := SScan(ctx, key, cursor, "*", 10)
			if err != nil {
				b.Fatal(err)
			}
			if cursor == 0 {
				break
			}
		}
	}
}

// =================== Sorted Set操作基准测试 ===================
// 包含以下操作的基准测试:
// - ZAdd - 添加成员
// - ZRange - 按位置获取成员
// - ZRank - 获取成员排名
// - ZScore - 获取成员分数
// - ZRevRange - 按位置倒序获取
// - ZIncrBy - 增加成员分数
// - ZCount - 指定分数范围的成员数
// - ZRemRangeByRank - 按排名删除成员
// - ZRangeByScore - 按分数范围获取成员
// - ZScan - 迭代有序集合

func BenchmarkSortedSetOperations(b *testing.B) {
	benchZAdd(b)
	utils.SysInfo("ZAdd benchmark completed")

	benchZRange(b)
	utils.SysInfo("ZRange benchmark completed")

	benchZRank(b)
	utils.SysInfo("ZRank benchmark completed")

	benchZScore(b)
	utils.SysInfo("ZScore benchmark completed")

	benchZRevRange(b)
	utils.SysInfo("ZRevRange benchmark completed")

	benchZIncrBy(b)
	utils.SysInfo("ZIncrBy benchmark completed")

	benchZCount(b)
	utils.SysInfo("ZCount benchmark completed")

	benchZRemRangeByRank(b)
	utils.SysInfo("ZRemRangeByRank benchmark completed")

	benchZRangeByScore(b)
	utils.SysInfo("ZRangeByScore benchmark completed")

	benchZScan(b)
	utils.SysInfo("ZScan benchmark completed")
}

func benchZAdd(b *testing.B) {
	ctx := context.Background()
	key := "test:zadd"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		member := redis.Z{
			Score:  float64(i),
			Member: fmt.Sprintf("member%d", i),
		}
		if err := ZAdd(ctx, key, member); err != nil {
			b.Fatal(err)
		}
	}
}

func benchZRange(b *testing.B) {
	ctx := context.Background()
	key := "test:zrange"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		member := &redis.Z{
			Score:  float64(i),
			Member: fmt.Sprintf("member%d", i),
		}
		if err := ZAdd(ctx, key, *member); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := ZRange(ctx, key, 0, 10); err != nil {
			b.Fatal(err)
		}
	}
}

func benchZRank(b *testing.B) {
	ctx := context.Background()
	key := "test:zrank"
	member := "member50"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		z := &redis.Z{
			Score:  float64(i),
			Member: fmt.Sprintf("member%d", i),
		}
		if err := ZAdd(ctx, key, *z); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := ZRank(ctx, key, member); err != nil {
			b.Fatal(err)
		}
	}
}

func benchZScore(b *testing.B) {
	ctx := context.Background()
	key := "test:zscore"
	member := "member50"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		z := &redis.Z{
			Score:  float64(i),
			Member: fmt.Sprintf("member%d", i),
		}
		if err := ZAdd(ctx, key, *z); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := ZScore(ctx, key, member); err != nil {
			b.Fatal(err)
		}
	}
}

func benchZRevRange(b *testing.B) {
	ctx := context.Background()
	key := "test:zrevrange"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		z := &redis.Z{
			Score:  float64(i),
			Member: fmt.Sprintf("member%d", i),
		}
		if err := ZAdd(ctx, key, *z); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := ZRevRange(ctx, key, 0, 10); err != nil {
			b.Fatal(err)
		}
	}
}

func benchZIncrBy(b *testing.B) {
	ctx := context.Background()
	key := "test:zincrby"
	member := "member"

	// 准备测试数据
	z := &redis.Z{
		Score:  0,
		Member: member,
	}
	if err := ZAdd(ctx, key, *z); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := ZIncrBy(ctx, key, 1.0, member); err != nil {
			b.Fatal(err)
		}
	}
}

func benchZCount(b *testing.B) {
	ctx := context.Background()
	key := "test:zcount"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		z := &redis.Z{
			Score:  float64(i),
			Member: fmt.Sprintf("member%d", i),
		}
		if err := ZAdd(ctx, key, *z); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := ZCount(ctx, key, "0", "50"); err != nil {
			b.Fatal(err)
		}
	}
}

func benchZRemRangeByRank(b *testing.B) {
	ctx := context.Background()
	key := "test:zremrangebyrank"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 准备测试数据
		for j := 0; j < 100; j++ {
			z := &redis.Z{
				Score:  float64(j),
				Member: fmt.Sprintf("member%d", j),
			}
			if err := ZAdd(ctx, key, *z); err != nil {
				b.Fatal(err)
			}
		}
		// 删除前50个成员
		if err := ZRemRangeByRank(ctx, key, 0, 49); err != nil {
			b.Fatal(err)
		}
	}
}

func benchZRangeByScore(b *testing.B) {
	ctx := context.Background()
	key := "test:zrangebyscore"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		z := &redis.Z{
			Score:  float64(i),
			Member: fmt.Sprintf("member%d", i),
		}
		if err := ZAdd(ctx, key, *z); err != nil {
			b.Fatal(err)
		}
	}

	opt := &redis.ZRangeBy{
		Min: "0",
		Max: "50",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := ZRangeByScore(ctx, key, opt); err != nil {
			b.Fatal(err)
		}
	}
}

func benchZScan(b *testing.B) {
	ctx := context.Background()
	key := "test:zscan"

	// 准备测试数据
	for i := 0; i < 100; i++ {
		z := &redis.Z{
			Score:  float64(i),
			Member: fmt.Sprintf("member%d", i),
		}
		if err := ZAdd(ctx, key, *z); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var cursor uint64
		for {
			_, cursor, err := ZScan(ctx, key, cursor, "*", 10)
			if err != nil {
				b.Fatal(err)
			}
			if cursor == 0 {
				break
			}
		}
	}
}

// =================== 分布式锁操作基准测试 ===================
// 包含以下操作的基准测试:
// - TryLock - 获取分布式锁
// - Unlock - 释放分布式锁
// - RenewLock - 续期分布式锁
// - GetLockTTL - 获取锁剩余时间

func BenchmarkLockOperations(b *testing.B) {
	benchTryLock(b)
	utils.SysInfo("TryLock benchmark completed")

	benchUnlock(b)
	utils.SysInfo("Unlock benchmark completed")

	benchRenewLock(b)
	utils.SysInfo("RenewLock benchmark completed")

	benchGetLockTTL(b)
	utils.SysInfo("GetLockTTL benchmark completed")
}

func benchTryLock(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("test:lock:%d", i)
		if _, err := TryLock(ctx, key, "value", time.Second); err != nil {
			b.Fatal(err)
		}
	}
}

func benchUnlock(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("test:lock:%d", i)
		// 先加锁
		if _, err := TryLock(ctx, key, "value", time.Second); err != nil {
			b.Fatal(err)
		}
		// 再解锁
		if err := Unlock(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

func benchRenewLock(b *testing.B) {
	ctx := context.Background()
	key := "test:lock:renew"

	// 先获取锁
	if _, err := TryLock(ctx, key, "value", time.Second); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := RenewLock(ctx, key, time.Second); err != nil {
			b.Fatal(err)
		}
	}
}

func benchGetLockTTL(b *testing.B) {
	ctx := context.Background()
	key := "test:lock:ttl"

	// 先获取锁
	if _, err := TryLock(ctx, key, "value", time.Second); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := GetLockTTL(ctx, key); err != nil {
			b.Fatal(err)
		}
	}
}

// RunBenchmarks 执行所有Redis基准测试并统计耗时
func RunBenchmarks() error {
	// 创建一个新的基准测试实例
	b := &testing.B{
		N: 1000, // 设置固定的迭代次数
	}

	utils.SysInfo("开始Redis基准测试...")

	// 定义一个函数来执行基准测试并记录耗时
	runBenchWithTimer := func(name string, benchFunc func(*testing.B)) time.Duration {
		start := time.Now()
		benchFunc(b)
		duration := time.Since(start)
		utils.SysInfo(fmt.Sprintf("%s 基准测试完成，耗时: %v", name, duration))
		return duration
	}

	// 记录所有测试结果
	benchResults := make(map[string]time.Duration)

	// 执行所有操作类型的基准测试并记录耗时
	benchResults["键操作"] = runBenchWithTimer("键操作", BenchmarkKeyOperations)
	benchResults["字符串操作"] = runBenchWithTimer("字符串操作", BenchmarkStringOperations)
	benchResults["列表操作"] = runBenchWithTimer("列表操作", BenchmarkListOperations)
	benchResults["哈希操作"] = runBenchWithTimer("哈希操作", BenchmarkHashOperations)
	benchResults["集合操作"] = runBenchWithTimer("集合操作", BenchmarkSetOperations)
	benchResults["有序集合操作"] = runBenchWithTimer("有序集合操作", BenchmarkSortedSetOperations)
	benchResults["分布式锁操作"] = runBenchWithTimer("分布式锁操作", BenchmarkLockOperations)

	// 计算总耗时
	var totalDuration time.Duration
	for _, duration := range benchResults {
		totalDuration += duration
	}

	// 输出汇总信息
	utils.SysInfo("基准测试汇总:")
	utils.SysInfo(fmt.Sprintf("总耗时: %v", totalDuration))
	utils.SysInfo("各操作耗时占比:")
	for name, duration := range benchResults {
		percentage := float64(duration) / float64(totalDuration) * 100
		utils.SysInfo(fmt.Sprintf("- %s: %v (%.2f%%)", name, duration, percentage))
	}

	return nil
}
