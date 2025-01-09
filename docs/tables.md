# 数据库表设计文档

## 表属性总览

| 表名 | 软删除 | 时间戳 | JSON配置 | 状态字段 | 索引优化 | 追踪ID |
|-----|:----:|:----:|:-------:|:------:|:------:|:-----:|
| users | ✓ | ✓ | ✓ | ✓ | ✓ | - |
| tokens | ✓ | ✓ | ✓ | ✓ | ✓ | - |
| models | ✓ | ✓ | ✓ | ✓ | ✓ | - |
| channels | ✓ | ✓ | ✓ | ✓ | ✓ | - |
| usages | - | ✓ | ✓ | - | ✓ | ✓ |
| relay_logs | - | ✓ | ✓ | - | ✓ | ✓ |
| sys_logs | - | ✓ | ✓ | - | ✓ | ✓ |
| message_saves | - | ✓ | ✓ | - | ✓ | ✓ |
| billings | - | ✓ | ✓ | ✓ | ✓ | - |
| pays | - | ✓ | - | ✓ | ✓ | ✓ |
| workers | ✓ | ✓ | ✓ | ✓ | ✓ | - |
| worker_nodes | ✓ | ✓ | ✓ | ✓ | ✓ | - |
| quotas | - | ✓ | ✓ | - | ✓ | - |
| request_logs | - | ✓ | ✓ | - | ✓ | ✓ |
| model_trainings | ✓ | ✓ | ✓ | ✓ | ✓ | - |
| model_versions | ✓ | ✓ | ✓ | ✓ | ✓ | - |
| model_deployments | ✓ | ✓ | ✓ | ✓ | ✓ | - |

## 详细表结构

### 1. users（用户表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| user_id | bigint unsigned | ✓ | PK | 用户ID |
| username | varchar(50) | - | - | 用户名 |
| password | varchar(255) | - | - | 加密密码 |
| email | varchar(100) | - | UK | 邮箱 |
| phone | varchar(20) | - | UK | 手机号 |
| oauth_info | json | - | - | 第三方授权信息(github/wechat/google等) |
| user_tokens | json | - | - | 用户token列表 |
| user_group | varchar(50) | ✓ | ✓ | 用户组 |
| user_quota | json | - | - | 用户配额信息 |
| user_options | json | - | - | 用户其他配置 |
| last_login_time | timestamp | - | - | 最后登录时间 |
| last_login_ip | varchar(50) | - | - | 最后登录IP |
| status | tinyint(1) | ✓ | - | 状态 1:正常 0:禁用 |
| created_at | timestamp | ✓ | - | 创建时间 |
| updated_at | timestamp | ✓ | - | 更新时间 |

### 2. tokens（令牌表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| token_id | bigint unsigned | ✓ | PK | 令牌ID |
| user_id | bigint unsigned | ✓ | ✓ | 所属用户ID |
| token_name | varchar(100) | ✓ | - | 令牌名称 |
| token_group | json | - | - | 令牌组配置(普通/高级/尊享/专线/特权等) |
| token_models | json | - | - | 支持的模型列表(关联models表的model_id) |
| token_options | json | - | - | 令牌配置选项(频率/额度/白名单等) |
| token_key | varchar(255) | ✓ | UK | 令牌密钥 |
| expire_time | timestamp | - | - | 过期时间 |
| status | tinyint(1) | ✓ | - | 状态 1:正常 0:禁用 |
| created_at | timestamp | ✓ | - | 创建时间 |
| updated_at | timestamp | ✓ | - | 更新时间 |

### 3. models（模型表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| model_id | bigint unsigned | ✓ | PK | 模型ID |
| model_name | varchar(100) | ✓ | UK | 模型名称 |
| model_alias | json | - | - | 模型别名映射配置 |
| model_price | decimal(10,6) | ✓ | - | 模型单价 |
| price_type | varchar(20) | ✓ | - | 计费类型(usage:按量,times:按次) |
| model_description | text | - | - | 模型描述 |
| model_config | json | - | - | 模型配置(频率限制/并发数等) |
| model_type | varchar(50) | ✓ | ✓ | 模型类型 |
| provider | varchar(50) | ✓ | - | 服务提供商 |
| status | tinyint(1) | ✓ | - | 状态 1:正常 0:禁用 |
| created_at | timestamp | ✓ | - | 创建时间 |
| updated_at | timestamp | ✓ | - | 更新时间 |

### 4. channels（渠道表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| channel_id | bigint unsigned | ✓ | PK | 渠道ID |
| channel_name | varchar(100) | ✓ | - | 渠道名称 |
| channel_level | varchar(50) | ✓ | ✓ | 渠道等级(普通/高级/尊享等) |
| channel_description | text | - | - | 渠道描述 |
| channel_models | json | - | - | 渠道支持的模型配置 |
| price_factor | decimal(10,2) | ✓ | - | 价格权重(倍率) |
| upstream_config | json | ✓ | - | 上游服务配置 |
| auth_config | json | ✓ | - | 认证配置 |
| retry_config | json | - | - | 重试配置 |
| rate_limit | json | - | - | 速率限制配置 |
| model_mapping | json | - | - | 模型重定向映射 |
| test_models | json | - | - | 测试模型配置 |
| status | tinyint(1) | ✓ | - | 状态 1:正常 0:禁用 |
| created_at | timestamp | ✓ | - | 创建时间 |
| updated_at | timestamp | ✓ | - | 更新时间 |

### 5. usages（用量表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| usage_id | bigint unsigned | ✓ | PK | 用量ID |
| token_id | bigint unsigned | ✓ | ✓ | 令牌ID |
| user_id | bigint unsigned | ✓ | ✓ | 用户ID |
| channel_id | bigint unsigned | ✓ | ✓ | 渠道ID |
| model_id | bigint unsigned | ✓ | ✓ | 模型ID |
| tokens_count | int | - | - | Token数量(按量收费时使用) |
| times_count | int | - | - | 请求次数(按次收费时使用) |
| unit_price | decimal(10,6) | ✓ | - | 单价 |
| price_factor | decimal(10,2) | ✓ | - | 价格倍率 |
| total_amount | decimal(10,6) | ✓ | - | 总金额 |
| usage_type | varchar(20) | ✓ | ✓ | 计费类型(usage:按量,times:按次) |
| request_id | varchar(64) | - | - | 请求ID |
| created_at | timestamp | ✓ | ✓ | 创建时间 |

### 6. relay_logs（中继日志表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| relay_id | bigint unsigned | ✓ | PK | 中继ID |
| user_id | bigint unsigned | ✓ | ✓ | 用户ID |
| channel_id | bigint unsigned | ✓ | ✓ | 渠道ID |
| model_id | bigint unsigned | ✓ | ✓ | 模型ID |
| upstream_url | varchar(255) | ✓ | - | 上游服务地址 |
| relay_status | tinyint | ✓ | ✓ | 转发状态(1:成功 2:失败 3:超时 4:限流 5:上游异常) |
| upstream_status | int | - | - | 上游返回状态码 |
| error_type | varchar(50) | - | ✓ | 错误类型(rate_limit/server_error/timeout/network等) |
| error_message | text | - | - | 错误详细信息 |
| retry_count | int | - | - | 重试次数 |
| retry_result | json | - | - | 重试结果记录 |
| request_headers | json | - | - | 请求头信息 |
| request_body | json | - | - | 请求体信息 |
| response_headers | json | - | - | 响应头信息 |
| response_body | json | - | - | 响应体信息(可能需要脱敏) |
| upstream_latency | int | - | - | 上游服务延迟(ms) |
| total_latency | int | - | - | 总延迟时间(ms) |
| request_tokens | int | - | - | 请求token数量 |
| response_tokens | int | - | - | 响应token数量 |
| quota_consumed | decimal(10,6) | - | - | 消耗的配额数量 |
| upstream_request_id | varchar(64) | - | - | 上游请求ID |
| request_id | varchar(64) | ✓ | ✓ | 关联的请求ID |
| created_at | timestamp | ✓ | ✓ | 创建时间 |

### 7.1 gateway_logs（网关日志表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| log_id | bigint unsigned | ✓ | PK | 日志ID |
| node_id | varchar(50) | ✓ | ✓ | 网关节点ID |
| log_level | varchar(20) | ✓ | ✓ | 日志级别(info/warn/error) |
| event_type | varchar(50) | ✓ | ✓ | 事件类型(route/forward/error等) |
| request_id | varchar(64) | - | ✓ | 关联请求ID |
| target_service | varchar(100) | - | - | 目标服务 |
| response_code | int | - | - | 响应状态码 |
| error_type | varchar(50) | - | - | 错误类型 |
| error_message | text | - | - | 错误信息 |
| log_details | json | - | - | 详细日志信息 |
| created_at | timestamp | ✓ | ✓ | 创建时间 |

### 7.2 master_logs（主服务日志表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| log_id | bigint unsigned | ✓ | PK | 日志ID |
| node_id | varchar(50) | ✓ | ✓ | 主服务节点ID |
| log_level | varchar(20) | ✓ | ✓ | 日志级别 |
| event_type | varchar(50) | ✓ | ✓ | 事件类型(node/redis/mysql/worker等) |
| operation | varchar(50) | ✓ | - | 操作类型 |
| request_id | varchar(64) | - | ✓ | 关联请求ID |
| target_type | varchar(50) | - | - | 目标类型(redis/mysql/worker) |
| target_node | varchar(100) | - | - | 目标节点 |
| error_type | varchar(50) | - | - | 错误类型 |
| error_message | text | - | - | 错误信息 |
| log_details | json | - | - | 详细日志信息 |
| created_at | timestamp | ✓ | ✓ | 创建时间 |

### 7.3 worker_logs（工作节点日志表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| log_id | bigint unsigned | ✓ | PK | 日志ID |
| worker_id | bigint unsigned | ✓ | ✓ | 工作节点ID |
| node_id | varchar(50) | ✓ | ✓ | 节点实例ID |
| log_level | varchar(20) | ✓ | ✓ | 日志级别 |
| event_type | varchar(50) | ✓ | ✓ | 事件类型(task/redis/mysql等) |
| request_id | varchar(64) | - | ✓ | 关联请求ID |
| resource_type | varchar(50) | - | - | 资源类型(cpu/memory/disk等) |
| resource_usage | json | - | - | 资源使用情况 |
| error_type | varchar(50) | - | - | 错误类型 |
| error_message | text | - | - | 错误信息 |
| log_details | json | - | - | 详细日志信息 |
| created_at | timestamp | ✓ | ✓ | 创建时间 |

### 7.4 master_mysql_logs（主服务MySQL日志表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| log_id | bigint unsigned | ✓ | PK | 日志ID |
| node_id | varchar(50) | ✓ | ✓ | 节点ID |
| log_level | varchar(20) | ✓ | ✓ | 日志级别 |
| event_type | varchar(50) | ✓ | ✓ | 事件类型(connect/query/error等) |
| operation | varchar(50) | - | - | 操作类型 |
| request_id | varchar(64) | - | ✓ | 关联请求ID |
| execution_time | int | - | - | 执行时间(ms) |
| error_type | varchar(50) | - | - | 错误类型 |
| error_message | text | - | - | 错误信息 |
| log_details | json | - | - | 详细日志信息 |
| created_at | timestamp | ✓ | ✓ | 创建时间 |

### 7.5 worker_mysql_logs（工作节点MySQL日志表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| log_id | bigint unsigned | ✓ | PK | 日志ID |
| worker_id | bigint unsigned | ✓ | ✓ | 工作节点ID |
| node_id | varchar(50) | ✓ | ✓ | 节点ID |
| log_level | varchar(20) | ✓ | ✓ | 日志级别 |
| event_type | varchar(50) | ✓ | ✓ | 事件类型(connect/query/error等) |
| operation | varchar(50) | - | - | 操作类型 |
| request_id | varchar(64) | - | ✓ | 关联请求ID |
| execution_time | int | - | - | 执行时间(ms) |
| error_type | varchar(50) | - | - | 错误类型 |
| error_message | text | - | - | 错误信息 |
| log_details | json | - | - | 详细日志信息 |
| created_at | timestamp | ✓ | ✓ | 创建时间 |

### 7.6 redis_logs（Redis日志表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| log_id | bigint unsigned | ✓ | PK | 日志ID |
| node_type | varchar(20) | ✓ | ✓ | 节点类型(master/worker) |
| node_id | varchar(50) | ✓ | ✓ | 节点ID |
| log_level | varchar(20) | ✓ | ✓ | 日志级别 |
| event_type | varchar(50) | ✓ | ✓ | 事件类型(set/get/del/persist等) |
| operation | varchar(50) | ✓ | - | 操作类型 |
| key_pattern | varchar(255) | - | - | 操作的key模式 |
| request_id | varchar(64) | - | ✓ | 关联请求ID |
| execution_time | int | - | - | 执行时间(ms) |
| error_type | varchar(50) | - | - | 错误类型 |
| error_message | text | - | - | 错误信息 |
| log_details | json | - | - | 详细日志信息 |
| created_at | timestamp | ✓ | ✓ | 创建时间 |

### 7.7 redis_persist_logs（Redis持久化日志表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| log_id | bigint unsigned | ✓ | PK | 日志ID |
| node_type | varchar(20) | ✓ | ✓ | 节点类型(master/worker) |
| node_id | varchar(50) | ✓ | ✓ | 节点ID |
| persist_type | varchar(50) | ✓ | ✓ | 持久化类型(rdb/aof/混合) |
| event_type | varchar(50) | ✓ | ✓ | 事件类型(start/complete/error) |
| target_table | varchar(100) | - | - | 目标MySQL表 |
| data_size | bigint | - | - | 数据大小(bytes) |
| affected_rows | int | - | - | 影响行数 |
| start_time | timestamp | ✓ | - | 开始时间 |
| end_time | timestamp | - | - | 结束时间 |
| duration | int | - | - | 持续时间(ms) |
| error_type | varchar(50) | - | - | 错误类型 |
| error_message | text | - | - | 错误信息 |
| log_details | json | - | - | 详细日志信息 |
| created_at | timestamp | ✓ | ✓ | 创建时间 |

### 8. message_saves（消息存储表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| message_id | bigint unsigned | ✓ | PK | 消息ID |
| user_id | bigint unsigned | ✓ | ✓ | 用户ID |
| token_id | bigint unsigned | ✓ | ✓ | 令牌ID |
| model_id | bigint unsigned | ✓ | ✓ | 模型ID |
| channel_id | bigint unsigned | ✓ | ✓ | 渠道ID |
| conversation_id | varchar(64) | - | ✓ | 会话ID |
| parent_message_id | varchar(64) | - | ✓ | 父消息ID |
| message_type | varchar(20) | ✓ | ✓ | 消息类型(user/assistant/system) |
| message_role | varchar(20) | ✓ | - | 消息角色 |
| business_type | varchar(50) | - | ✓ | 业务类型(客服/营销/技术支持等) |
| industry_type | varchar(50) | - | ✓ | 行业类型(金融/教育/医疗等) |
| scene_type | varchar(50) | - | ✓ | 场景类型(咨询/问答/创作等) |
| message_title | varchar(255) | - | - | 消息标题 |
| message_content | text | ✓ | - | 消息内容 |
| message_tokens | int | - | - | 消息token数 |
| prompt_tokens | int | - | - | 提示词token数 |
| completion_tokens | int | - | - | 补全token数 |
| message_status | tinyint | ✓ | ✓ | 消息状态(1:正常 2:待审 3:违规) |
| quality_score | decimal(3,2) | - | - | 质量评分(0-5分) |
| relevance_score | decimal(3,2) | - | - | 相关性评分(0-5分) |
| user_feedback | tinyint | - | - | 用户反馈(1:点赞 2:点踩) |
| user_comment | varchar(500) | - | - | 用户评论 |
| usage_info | json | - | - | 用量信息 |
| message_config | json | - | - | 消息配置(温度/top_p等) |
| prompt_template | json | - | - | 提示词模板 |
| keywords | json | - | - | 关键词提取 |
| sentiment | varchar(20) | - | - | 情感倾向(正面/负面/中性) |
| intent_tags | json | - | - | 意图标签 |
| entity_tags | json | - | - | 实体标签 |
| classification | json | - | - | 分类信息 |
| message_extra | json | - | - | 消息额外信息 |
| message_tags | json | - | - | 消息标签 |
| is_sensitive | tinyint(1) | ✓ | - | 是否敏感内容 |
| is_public | tinyint(1) | ✓ | - | 是否公开 |
| share_code | varchar(32) | - | UK | 分享码 |
| error_type | varchar(50) | - | - | 错误类型 |
| error_message | text | - | - | 错误信息 |
| retry_count | int | - | - | 重试次数 |
| latency | int | - | - | 响应延迟(ms) |
| request_id | varchar(64) | - | ✓ | 请求ID |
| created_at | timestamp | ✓ | ✓ | 创建时间 |
| updated_at | timestamp | ✓ | - | 更新时间 |

### 9. billings（账单表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| billing_id | bigint unsigned | ✓ | PK | 账单ID |
| user_id | bigint unsigned | ✓ | ✓ | 用户ID |
| billing_no | varchar(64) | ✓ | UK | 账单编号 |
| billing_type | varchar(20) | ✓ | ✓ | 账单类型(消费/退款) |
| billing_cycle | varchar(20) | ✓ | - | 账单周期(日/周/月) |
| billing_money | decimal(10,2) | ✓ | - | 账单金额 |
| billing_currency | varchar(10) | ✓ | - | 账单币种 |
| billing_details | json | ✓ | - | 账单明细(各模型用量统计) |
| quota_details | json | ✓ | - | 配额使用明细 |
| billing_status | tinyint | ✓ | ✓ | 账单状态(1:未出账 2:已出账 3:已支付 4:已逾期) |
| start_time | timestamp | - | - | 账单开始时间 |
| end_time | timestamp | - | - | 账单结束时间 |
| due_time | timestamp | ✓ | - | 账单到期时间 |
| pay_time | timestamp | - | - | 支付时间 |
| remark | varchar(255) | - | - | 账单备注 |
| created_at | timestamp | ✓ | ✓ | 创建时间 |
| updated_at | timestamp | ✓ | - | 更新时间 |

### 10. pays（支付表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| pay_id | bigint unsigned | ✓ | PK | 支付ID |
| user_id | bigint unsigned | ✓ | ✓ | 用户ID |
| order_no | varchar(64) | ✓ | UK | 支付订单号 |
| pay_platform | varchar(50) | ✓ | ✓ | 支付平台(stripe/paddle/alipay/wechat等) |
| pay_method | varchar(50) | ✓ | - | 支付方式(信用卡/银行转账/支付宝/微信等) |
| pay_type | varchar(20) | ✓ | ✓ | 支付类型(个人/企业) |
| pay_scene | varchar(20) | ✓ | - | 支付场景(充值/订阅等) |
| pay_currency | varchar(10) | ✓ | - | 支付币种 |
| pay_amount | decimal(10,2) | ✓ | - | 支付金额 |
| pay_status | tinyint | ✓ | ✓ | 支付状态(1:待支付 2:支付中 3:支付成功 4:支付失败 5:已退款) |
| pay_title | varchar(255) | ✓ | - | 支付标题 |
| pay_desc | varchar(1000) | - | - | 支付描述 |
| payer_name | varchar(100) | - | - | 付款人姓名/企业名称 |
| payer_email | varchar(100) | - | - | 付款人邮箱 |
| payer_phone | varchar(20) | - | - | 付款人电话 |
| company_info | json | - | - | 企业付款信息(营业执照/税号等) |
| billing_info | json | - | - | 账单信息(发票信息等) |
| notify_url | varchar(255) | - | - | 支付回调通知地址 |
| transaction_id | varchar(64) | - | - | 第三方交易ID |
| callback_data | json | - | - | 支付回调数据 |
| platform_config | json | - | - | 支付平台特定配置 |
| refund_info | json | - | - | 退款信息 |
| expire_time | timestamp | ✓ | - | 支付过期时间 |
| pay_time | timestamp | - | - | 支付成功时间 |
| created_at | timestamp | ✓ | ✓ | 创建时间 |
| updated_at | timestamp | ✓ | - | 更新时间 |

### 11.1 relay_workers（中继工作节点整体表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| cluster_id | bigint unsigned | ✓ | PK | 集群ID |
| cluster_name | varchar(100) | ✓ | - | 集群名称 |
| total_workers | int | ✓ | - | 工作节点总数 |
| active_workers | int | ✓ | - | 活跃节点数 |
| total_instances | int | ✓ | - | 实例总数 |
| active_instances | int | ✓ | - | 活跃实例数 |
| cluster_status | tinyint | ✓ | ✓ | 集群状态(1:正常 2:部分可用 3:异常) |
| cluster_config | json | - | - | 集群配置信息 |
| resource_usage | json | - | - | 资源使用统计 |
| performance_stats | json | - | - | 性能统计信息 |
| alert_config | json | - | - | 告警配置 |
| maintenance_window | json | - | - | 维护窗口配置 |
| status | tinyint(1) | ✓ | - | 状态 1:正常 0:禁用 |
| created_at | timestamp | ✓ | - | 创建时间 |
| updated_at | timestamp | ✓ | - | 更新时间 |

### 11.2 workers（工作节点表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| worker_id | bigint unsigned | ✓ | PK | 工作节点ID |
| cluster_id | bigint unsigned | ✓ | ✓ | 所属集群ID |
| worker_name | varchar(100) | ✓ | - | 节点名称 |
| worker_type | varchar(50) | ✓ | ✓ | 节点类型(text/image/video/audio) |
| worker_group | varchar(50) | ✓ | ✓ | 节点组 |
| worker_role | varchar(20) | ✓ | ✓ | 节点角色(master/slave) |
| worker_priority | int | ✓ | - | 节点优先级 |
| worker_options | json | - | - | 节点配置选项 |
| worker_status | tinyint | ✓ | ✓ | 节点状态(在线/离线/维护) |
| load_balance | int | ✓ | - | 负载权重 |
| max_instances | int | ✓ | - | 最大实例数 |
| min_instances | int | ✓ | - | 最小实例数 |
| current_instances | int | ✓ | - | 当前实例数 |
| resource_limits | json | - | - | 资源限制配置 |
| scaling_rules | json | - | - | 扩缩容规则 |
| status | tinyint(1) | ✓ | - | 状态 1:正常 0:禁用 |
| created_at | timestamp | ✓ | - | 创建时间 |
| updated_at | timestamp | ✓ | - | 更新时间 |

### 11.3 worker_nodes（工作节点实例表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| node_id | bigint unsigned | ✓ | PK | 节点实例ID |
| worker_id | bigint unsigned | ✓ | ✓ | 所属工作节点ID |
| cluster_id | bigint unsigned | ✓ | ✓ | 所属集群ID |
| node_ip | varchar(50) | ✓ | - | 节点IP地址 |
| node_port | int | ✓ | - | 节点端口 |
| node_region | varchar(50) | - | ✓ | 节点区域 |
| node_zone | varchar(50) | - | ✓ | 节点可用区 |
| node_options | json | - | - | 节点实例配置 |
| node_status | tinyint | ✓ | ✓ | 实例状态 |
| cpu_usage | decimal(5,2) | - | - | CPU使用率 |
| memory_usage | decimal(5,2) | - | - | 内存使用率 |
| network_stats | json | - | - | 网络统计信息 |
| performance_stats | json | - | - | 性能统计信息 |
| health_check | timestamp | ✓ | - | 最后健康检查时间 |
| last_error | text | - | - | 最后错误信息 |
| startup_time | timestamp | ✓ | - | 启动时间 |
| status | tinyint(1) | ✓ | - | 状态 1:正常 0:禁用 |
| created_at | timestamp | ✓ | - | 创建时间 |
| updated_at | timestamp | ✓ | - | 更新时间 |

### 13. quotas（配额表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| quota_id | bigint unsigned | ✓ | PK | 配额ID |
| user_id | bigint unsigned | ✓ | ✓ | 用户ID |
| pay_id | bigint unsigned | ✓ | ✓ | 关联的支付ID |
| quota_amount | decimal(10,2) | ✓ | - | 配额金额 |
| remaining_amount | decimal(10,2) | ✓ | - | 剩余金额 |
| frozen_amount | decimal(10,2) | ✓ | - | 冻结金额(处理中的请求) |
| quota_type | varchar(20) | ✓ | ✓ | 配额类型(充值/赠送/奖励等) |
| quota_level | varchar(20) | ✓ | ✓ | 配额等级(普通/高级/尊享等) |
| valid_period | int | ✓ | - | 有效期(天) |
| start_time | timestamp | ✓ | - | 生效时间 |
| expire_time | timestamp | ✓ | - | 过期时间 |
| quota_config | json | - | - | 配额特殊配置 |
| status | tinyint(1) | ✓ | ✓ | 状态(1:正常 2:冻结 3:过期) |
| created_at | timestamp | ✓ | - | 创建时间 |
| updated_at | timestamp | ✓ | - | 更新时间 |

### 14. request_logs（请求日志表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| request_id | varchar(64) | ✓ | PK | 请求ID |
| user_id | bigint unsigned | ✓ | ✓ | 用户ID |
| token_id | bigint unsigned | ✓ | ✓ | 令牌ID |
| channel_id | bigint unsigned | ✓ | ✓ | 渠道ID |
| model_id | bigint unsigned | ✓ | ✓ | 模型ID |
| worker_id | bigint unsigned | - | ✓ | 处理的工作节点ID |
| request_type | varchar(50) | ✓ | ✓ | 请求类型 |
| request_path | varchar(255) | ✓ | - | 请求路径 |
| request_method | varchar(10) | ✓ | - | 请求方法(GET/POST等) |
| request_headers | json | - | - | 请求头信息 |
| request_params | json | - | - | 请求参数 |
| request_tokens | int | - | - | 请求内容token数 |
| request_status | tinyint | ✓ | ✓ | 请求状态(1:成功 2:失败 3:超时) |
| request_time | timestamp | ✓ | - | 请求开始时间 |
| process_time | timestamp | ✓ | - | 处理开始时间 |
| response_time | timestamp | ✓ | - | 响应结束时间 |
| upstream_time | int | - | - | 上游处理耗时(毫秒) |
| total_time | int | ✓ | - | 总耗时(毫秒) |
| response_code | int | ✓ | ✓ | 响应状态码 |
| response_headers | json | - | - | 响应头信息 |
| error_message | text | - | - | 错误信息 |
| client_ip | varchar(50) | ✓ | - | 客户端IP |
| created_at | timestamp | ✓ | ✓ | 创建时间 |

### 15.1 model_trainings（模型训练表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| training_id | bigint unsigned | ✓ | PK | 训练ID |
| user_id | bigint unsigned | ✓ | ✓ | 用户ID |
| base_model_id | bigint unsigned | ✓ | ✓ | 基础模型ID |
| training_name | varchar(100) | ✓ | - | 训练任务名称 |
| training_desc | text | - | - | 训练描述 |
| webhook_url | varchar(255) | ✓ | - | 数据接收Webhook |
| training_type | varchar(50) | ✓ | ✓ | 训练类型(增量/全量) |
| training_config | json | ✓ | - | 训练配置(学习率/批次等) |
| data_config | json | ✓ | - | 数据配置(格式/预处理等) |
| data_status | json | - | - | 数据统计(数量/质量等) |
| resource_config | json | ✓ | - | 资源配置(GPU/内存等) |
| training_status | tinyint | ✓ | ✓ | 训练状态(1:待训练 2:训练中 3:已完成 4:失败) |
| start_time | timestamp | - | - | 开始时间 |
| end_time | timestamp | - | - | 结束时间 |
| metrics | json | - | - | 训练指标 |
| error_info | json | - | - | 错误信息 |
| created_at | timestamp | ✓ | ✓ | 创建时间 |
| updated_at | timestamp | ✓ | - | 更新时间 |

### 15.2 model_versions（模型版本表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| version_id | bigint unsigned | ✓ | PK | 版本ID |
| training_id | bigint unsigned | ✓ | ✓ | 训练ID |
| user_id | bigint unsigned | ✓ | ✓ | 用户ID |
| version_name | varchar(100) | ✓ | - | 版本名称 |
| version_desc | text | - | - | 版本描述 |
| model_path | varchar(255) | ✓ | - | 模型存储路径 |
| model_size | bigint | - | - | 模型大小(bytes) |
| model_format | varchar(50) | ✓ | - | 模型格式 |
| model_hash | varchar(64) | ✓ | - | 模型哈希值 |
| performance_metrics | json | - | - | 性能指标 |
| version_status | tinyint | ✓ | ✓ | 版本状态(1:待部署 2:已部署 3:已下线) |
| created_at | timestamp | ✓ | ✓ | 创建时间 |
| updated_at | timestamp | ✓ | - | 更新时间 |

### 15.3 model_deployments（模型部署表）

| 字段名 | 类型 | 必填 | 索引 | 描述 |
|-------|------|:----:|:----:|------|
| deployment_id | bigint unsigned | ✓ | PK | 部署ID |
| version_id | bigint unsigned | ✓ | ✓ | 版本ID |
| user_id | bigint unsigned | ✓ | ✓ | 用户ID |
| deployment_name | varchar(100) | ✓ | - | 部署名称 |
| deployment_desc | text | - | - | 部署描述 |
| cluster_id | bigint unsigned | ✓ | ✓ | 部署集群ID |
| instance_count | int | ✓ | - | 实例数量 |
| resource_config | json | ✓ | - | 资源配置 |
| scaling_config | json | - | - | 扩缩容配置 |
| route_config | json | ✓ | - | 路由配置 |
| endpoint_url | varchar(255) | - | - | 访问端点 |
| deployment_status | tinyint | ✓ | ✓ | 部署状态(1:部署中 2:运行中 3:已停止) |
| health_check | json | - | - | 健康检查配置 |
| monitoring_config | json | - | - | 监控配置 |
| created_at | timestamp | ✓ | ✓ | 创建时间 |
| updated_at | timestamp | ✓ | - | 更新时间 |

## 字段说明

### 通用字段

- created_at: 创建时间，所有表都包含
- updated_at: 更新时间，所有表都包含（日志表除外）
- status: 状态字段，用于软删除和状态管理
- request_id: 请求追踪ID，用于日志关联

### 特殊说明

1. 所有表使用 utf8mb4 字符集
2. 所有金额相关字段使用 decimal 类型
3. JSON字段用于存储灵活配置
4. 主键均使用 bigint unsigned 类型
5. 时间戳默认使用 CURRENT_TIMESTAMP
