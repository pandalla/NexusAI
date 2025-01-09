# API 文档

## 用户认证模块 (Auth)

### 模块概述

用户认证模块提供用户注册、登录、第三方认证等功能，支持多种认证方式和安全策略。基于users表设计，实现完整的用户身份验证和授权管理。

### 功能特点

- 支持邮箱/手机号注册和登录
- 支持多种第三方平台OAuth认证
- 完整的用户信息管理功能
- 安全的密码管理机制
- 基于Token的身份验证
- 用户状态和权限管理

### 接口列表

- [注册新用户](#注册新用户)
- [用户登录](#用户登录)
- [第三方认证](#第三方认证)
- [获取用户信息](#获取用户信息)
- [更新用户信息](#更新用户信息)
- [修改密码](#修改密码)
- [重置密码](#重置密码)
- [登出](#登出)
- [发送手机验证码](#发送手机验证码)
- [发送邮箱验证码](#发送邮箱验证码)
- [验证码校验](#验证码校验)

### 接口详情

#### 注册新用户

```http
POST /api/v1/auth/register
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| username | string | 否 | 用户名(2-50字符) |
| email | string | 是* | 邮箱地址 |
| phone | string | 是* | 手机号码 |
| password | string | 是 | 密码(8-32字符) |
| user_group | string | 否 | 用户组(默认: normal) |

注：email和phone至少提供一个

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "user_id": "12345678",
        "username": "example_user",
        "email": "user@example.com",
        "user_group": "normal",
        "created_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 用户登录

```http
POST /api/v1/auth/login
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| account | string | 是 | 账号(邮箱/手机号) |
| password | string | 是 | 密码 |
| device_info | object | 否 | 设备信息 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "user_id": "12345678",
        "access_token": "eyJhbGciOiJIUzI1NiIs...",
        "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
        "expires_in": 3600
    }
}
```

#### 第三方认证

```http
POST /api/v1/auth/oauth/{platform}
```

支持的platform: github, wechat, google

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| code | string | 是 | OAuth授权码 |
| state | string | 是 | 状态码(防CSRF) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "user_id": "12345678",
        "access_token": "eyJhbGciOiJIUzI1NiIs...",
        "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
        "expires_in": 3600,
        "oauth_info": {
            "platform": "github",
            "oauth_id": "gh_123456"
        }
    }
}
```

#### 获取用户信息

```http
GET /api/v1/auth/user
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "user_id": "12345678",
        "username": "example_user",
        "email": "user@example.com",
        "phone": "1234567890",
        "user_group": "normal",
        "oauth_info": {
            "github": {
                "oauth_id": "gh_123456",
                "username": "gh_user"
            }
        },
        "user_quota": {
            "total": 1000.00,
            "used": 100.00,
            "remaining": 900.00
        },
        "user_options": {
            "language": "zh-CN",
            "timezone": "Asia/Shanghai",
            "notification": {
                "email": true,
                "sms": false
            }
        },
        "last_login_time": "2024-03-20T12:00:00Z",
        "last_login_ip": "192.168.1.1",
        "status": 1,
        "created_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 更新用户信息

```http
PUT /api/v1/auth/user
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| username | string | 否 | 用户名(2-50字符) |
| email | string | 否 | 邮箱地址 |
| phone | string | 否 | 手机号码 |
| user_options | object | 否 | 用户配置选项 |
| notification_settings | object | 否 | 通知设置 |

请求示例：

```json
{
    "username": "new_username",
    "user_options": {
        "language": "en-US",
        "timezone": "UTC",
        "notification": {
            "email": true,
            "sms": true
        }
    }
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "user_id": "12345678",
        "username": "new_username",
        "user_options": {
            "language": "en-US",
            "timezone": "UTC",
            "notification": {
                "email": true,
                "sms": true
            }
        },
        "updated_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 修改密码

```http
POST /api/v1/auth/password/change
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| old_password | string | 是 | 原密码 |
| new_password | string | 是 | 新密码(8-32字符) |
| verify_code | string | 否 | 验证码(开启二次验证时必填) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "user_id": "12345678",
        "updated_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 重置密码

```http
POST /api/v1/auth/password/reset
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| account | string | 是 | 账号(邮箱/手机号) |
| verify_code | string | 是 | 验证码 |
| new_password | string | 是 | 新密码(8-32字符) |
| device_info | object | 否 | 设备信息 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "user_id": "12345678",
        "updated_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 登出

```http
POST /api/v1/auth/logout
```

请求头：

```http
Authorization: Bearer {access_token}
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| all_devices | boolean | 否 | 是否登出所有设备(默认: false) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "logout_time": "2024-03-20T12:00:00Z"
    }
}
```

#### 发送手机验证码

```http
POST /api/v1/auth/verify-code/sms
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| phone | string | 是 | 手机号码 |
| scene | string | 是 | 验证场景(register/login/reset/bind) |
| template | string | 否 | 短信模板ID |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "request_id": "sms_123456",
        "expire_time": "2024-03-20T12:05:00Z"
    }
}
```

#### 发送邮箱验证码

```http
POST /api/v1/auth/verify-code/email
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| email | string | 是 | 邮箱地址 |
| scene | string | 是 | 验证场景(register/login/reset/bind) |
| template | string | 否 | 邮件模板ID |
| language | string | 否 | 邮件语言(zh-CN/en-US) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "request_id": "email_123456",
        "expire_time": "2024-03-20T12:05:00Z"
    }
}
```

#### 验证码校验

```http
POST /api/v1/auth/verify-code/check
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| account | string | 是 | 账号(手机号/邮箱) |
| code | string | 是 | 验证码 |
| scene | string | 是 | 验证场景(register/login/reset/bind) |
| request_id | string | 是 | 发送验证码时返回的请求ID |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "is_valid": true,
        "verified_at": "2024-03-20T12:01:00Z"
    }
}
```

错误码补充：

| 错误码 | 描述 | 解决方案 |
|--------|------|----------|
| 422004 | 验证码发送过于频繁 | 请等待60秒后重试 |
| 422005 | 验证码已过期 | 请重新获取验证码 |
| 422006 | 验证码错误 | 请检查验证码是否正确 |
| 422007 | 请求ID无效 | 请使用最新发送的验证码 |

注意事项补充：

11. 验证码有效期为5分钟
12. 同一手机号/邮箱1分钟内只能发送1次验证码
13. 同一IP每小时最多发送10次验证码
14. 验证码连续错误5次将被锁定15分钟
15. 建议在敏感操作时使用验证码二次验证

### 错误码

| 错误码 | 描述 | 解决方案 |
|--------|------|----------|
| 401001 | 未授权访问 | 请检查access_token是否有效 |
| 401002 | access_token已过期 | 请使用refresh_token刷新令牌 |
| 401003 | 无效的refresh_token | 请重新登录 |
| 403001 | 账号已被禁用 | 请联系管理员 |
| 404001 | 用户不存在 | 请检查账号是否正确 |
| 409001 | 邮箱已存在 | 请更换邮箱或找回密码 |
| 409002 | 手机号已存在 | 请更换手机号或找回密码 |
| 422001 | 密码格式错误 | 密码必须包含8-32个字符 |
| 422002 | 验证码错误 | 请检查验证码是否正确 |
| 422003 | 原密码错误 | 请检查原密码是否正确 |
| 429001 | 请求过于频繁 | 请稍后重试 |

### 注意事项

1. 所有密码传输都必须经过加密处理
2. access_token默认有效期为1小时
3. refresh_token默认有效期为30天
4. 第三方认证需要先在平台配置对应的OAuth信息
5. 用户状态为禁用时，所有接口都将返回403001错误
6. 重要操作(如修改密码)建议增加二次验证
7. 建议使用HTTPS协议进行通信
8. 密码必须满足复杂度要求：
   - 长度8-32个字符
   - 必须包含大小写字母和数字
   - 建议包含特殊字符
9. 验证码有效期为5分钟
10. 连续5次密码错误将触发账号保护机制

## 令牌管理模块 (Token)

### 模块概述

令牌管理模块提供API访问令牌的创建、管理和监控功能。基于tokens表设计，实现细粒度的访问控制和用量管理。支持多种令牌类型和权限级别，确保API调用的安全性和可控性。

### 功能特点

- 支持多种令牌类型（普通/高级/专线）
- 灵活的权限和配额管理
- 实时的用量统计和监控
- 支持令牌自动过期和续期
- 可配置的访问限制和安全策略
- 详细的使用日志和审计记录

### 接口列表

- [创建令牌](#创建令牌)
- [获取令牌列表](#获取令牌列表)
- [获取令牌详情](#获取令牌详情)
- [更新令牌配置](#更新令牌配置)
- [更新令牌状态](#更新令牌状态)
- [删除令牌](#删除令牌)
- [获取令牌用量](#获取令牌用量)
- [获取令牌日志](#获取令牌日志)

### 接口详情

#### 创建令牌

```http
POST /api/v1/tokens
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| token_name | string | 是 | 令牌名称(2-100字符) |
| token_group | string | 是 | 令牌组(normal/premium/exclusive) |
| token_models | array | 是 | 支持的模型ID列表 |
| token_options | object | 否 | 令牌配置选项 |
| expire_time | string | 否 | 过期时间(ISO8601格式) |

请求示例：

```json
{
    "token_name": "My API Token",
    "token_group": "premium",
    "token_models": [1, 2, 3],
    "token_options": {
        "rate_limit": {
            "requests_per_minute": 60,
            "tokens_per_minute": 40000
        },
        "ip_whitelist": ["192.168.1.1", "10.0.0.1"],
        "allowed_origins": ["https://example.com"]
    },
    "expire_time": "2024-12-31T23:59:59Z"
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "token_id": "87654321",
        "token_key": "sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        "token_name": "My API Token",
        "token_group": "premium",
        "token_models": [1, 2, 3],
        "token_options": {
            "rate_limit": {
                "requests_per_minute": 60,
                "tokens_per_minute": 40000
            },
            "ip_whitelist": ["192.168.1.1", "10.0.0.1"],
            "allowed_origins": ["https://example.com"]
        },
        "expire_time": "2024-12-31T23:59:59Z",
        "created_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 获取令牌列表

```http
GET /api/v1/tokens
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| token_group | string | 否 | 按令牌组筛选 |
| status | integer | 否 | 按状态筛选(1:正常 0:禁用) |
| keyword | string | 否 | 按名称搜索 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 100,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "token_id": "87654321",
                "token_name": "My API Token",
                "token_group": "premium",
                "token_models": [1, 2, 3],
                "status": 1,
                "created_at": "2024-03-20T12:00:00Z",
                "expire_time": "2024-12-31T23:59:59Z",
                "last_used_at": "2024-03-20T12:30:00Z",
                "usage_stats": {
                    "requests_today": 1000,
                    "tokens_today": 50000
                }
            }
            // ... 更多令牌
        ]
    }
}
```

#### 获取令牌详情

```http
GET /api/v1/tokens/{token_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "token_id": "87654321",
        "token_name": "My API Token",
        "token_group": "premium",
        "token_models": [1, 2, 3],
        "token_options": {
            "rate_limit": {
                "requests_per_minute": 60,
                "tokens_per_minute": 40000
            },
            "ip_whitelist": ["192.168.1.1", "10.0.0.1"],
            "allowed_origins": ["https://example.com"]
        },
        "status": 1,
        "created_at": "2024-03-20T12:00:00Z",
        "expire_time": "2024-12-31T23:59:59Z",
        "last_used_at": "2024-03-20T12:30:00Z",
        "usage_stats": {
            "requests_today": 1000,
            "tokens_today": 50000,
            "requests_total": 10000,
            "tokens_total": 500000
        }
    }
}
```

#### 更新令牌配置

```http
PUT /api/v1/tokens/{token_id}
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| token_name | string | 否 | 令牌名称 |
| token_models | array | 否 | 支持的模型ID列表 |
| token_options | object | 否 | 令牌配置选项 |
| expire_time | string | 否 | 过期时间 |

请求示例：

```json
{
    "token_name": "Updated Token Name",
    "token_options": {
        "rate_limit": {
            "requests_per_minute": 100,
            "tokens_per_minute": 50000
        },
        "ip_whitelist": ["192.168.1.1", "10.0.0.1", "10.0.0.2"]
    },
    "expire_time": "2025-12-31T23:59:59Z"
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "token_id": "87654321",
        "token_name": "Updated Token Name",
        "updated_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 更新令牌状态

```http
PATCH /api/v1/tokens/{token_id}/status
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| status | integer | 是 | 状态(1:启用 0:禁用) |
| reason | string | 否 | 状态变更原因 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "token_id": "87654321",
        "status": 1,
        "updated_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 删除令牌

```http
DELETE /api/v1/tokens/{token_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| confirm | boolean | 是 | 确认删除(必须为true) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "token_id": "87654321",
        "deleted_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 获取令牌用量

```http
GET /api/v1/tokens/{token_id}/usage
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| start_time | string | 否 | 开始时间(ISO8601格式) |
| end_time | string | 否 | 结束时间(ISO8601格式) |
| group_by | string | 否 | 分组方式(hour/day/month) |
| model_id | integer | 否 | 按模型ID筛选 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "summary": {
            "total_requests": 10000,
            "total_tokens": 500000,
            "total_amount": 50.00,
            "average_latency": 500
        },
        "models": [
            {
                "model_id": 1,
                "model_name": "gpt-4",
                "requests": 5000,
                "tokens": 300000,
                "amount": 30.00
            }
        ],
        "timeline": [
            {
                "timestamp": "2024-03-20T00:00:00Z",
                "requests": 100,
                "tokens": 5000,
                "amount": 0.50,
                "latency": 450
            }
        ]
    }
}
```

#### 获取令牌日志

```http
GET /api/v1/tokens/{token_id}/logs
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| start_time | string | 否 | 开始时间 |
| end_time | string | 否 | 结束时间 |
| status | integer | 否 | 状态筛选 |
| model_id | integer | 否 | 模型筛选 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 1000,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "log_id": "123456789",
                "request_id": "req_xxxxx",
                "model_id": 1,
                "model_name": "gpt-4",
                "request_time": "2024-03-20T12:00:00Z",
                "response_time": "2024-03-20T12:00:01Z",
                "latency": 1000,
                "status": 1,
                "tokens": 500,
                "amount": 0.05,
                "client_ip": "192.168.1.1",
                "error_message": null
            }
        ]
    }
}
```

### 错误码

| 错误码 | 描述 | 解决方案 |
|--------|------|----------|
| 401001 | 未授权访问 | 请检查access_token是否有效 |
| 403001 | 令牌已被禁用 | 请检查令牌状态 |
| 403002 | 超出使用限制 | 请检查令牌配额或升级令牌组 |
| 403003 | IP未授权 | 请检查IP白名单配置 |
| 404001 | 令牌不存在 | 请检查令牌ID是否正确 |
| 409001 | 令牌名称重复 | 请更换令牌名称 |
| 422001 | 参数格式错误 | 请检查请求参数格式 |
| 429001 | 请求频率超限 | 请降低请求频率 |

### 注意事项

1. 令牌创建后，token_key 只显示一次，请妥善保管
2. 令牌默认有效期为永久，可以通过expire_time设置过期时间
3. 令牌组的权限和限制：
   - normal: 基础API访问权限
   - premium: 高级功能和更高配额
   - exclusive: 专属资源和定制化配置
4. 令牌配置选项包括：
   - 速率限制(rate_limit)
   - IP白名单(ip_whitelist)
   - 来源域名限制(allowed_origins)
   - 并发请求数(max_concurrent)
5. 用量统计支持按小时/天/月分组查看
6. 日志保留时间为30天
7. 建议定期检查令牌使用情况，及时发现异常
8. 重要操作会记录操作日志，便于审计
9. 支持批量管理令牌（未列出相关接口）
10. 令牌删除后不可恢复，请谨慎操作

## 模型服务模块 (Model)

### 模块概述

模型服务模块提供AI模型的管理和监控功能。基于models表设计，实现模型配置管理、状态监控和用量统计，为中继服务提供模型元信息支持。

### 功能特点

- 支持多种模型类型配置
- 灵活的模型参数管理
- 实时的模型状态监控
- 详细的用量统计分析
- 完整的操作审计日志
- 智能的调度策略配置

### 接口列表

#### 模型管理接口

- [获取模型列表](#获取模型列表)
- [获取模型详情](#获取模型详情)
- [创建模型配置](#创建模型配置)
- [更新模型配置](#更新模型配置)
- [删除模型配置](#删除模型配置)

#### 模型监控接口

- [获取模型状态](#获取模型状态)
- [获取模型用量](#获取模型用量)
- [获取模型日志](#获取模型日志)

### 接口详情

#### 获取模型列表

```http
GET /api/v1/models
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| model_type | string | 否 | 模型类型(text/image/audio) |
| provider | string | 否 | 服务商筛选 |
| status | integer | 否 | 状态筛选(1:正常 0:维护) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 100,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "model_id": 1,
                "model_name": "gpt-4",
                "model_type": "text",
                "provider": "openai",
                "description": "GPT-4是一个大型多模态模型",
                "model_config": {
                    "max_tokens": 8192,
                    "supports_stream": true
                },
                "status": 1,
                "pricing": {
                    "input": 0.01,
                    "output": 0.03,
                    "unit": "1K tokens"
                },
                "stats": {
                    "total_calls": 1000000,
                    "success_rate": 99.9,
                    "avg_latency": 500
                },
                "created_at": "2024-03-20T12:00:00Z"
            }
        ]
    }
}
```

#### 获取模型详情

```http
GET /api/v1/models/{model_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "model_id": 1,
        "model_name": "gpt-4",
        "model_type": "text",
        "provider": "openai",
        "description": "GPT-4是一个大型多模态模型",
        "model_config": {
            "max_tokens": 8192,
            "supports_stream": true,
            "temperature_range": [0, 2],
            "top_p_range": [0, 1],
            "stop_sequence": ["\n", "。", ".", "!"],
            "supports_functions": true
        },
        "routing_config": {
            "strategy": "round_robin",
            "timeout": 30,
            "retry_count": 3,
            "fallback_models": [2, 3]
        },
        "status": 1,
        "pricing": {
            "input": 0.01,
            "output": 0.03,
            "unit": "1K tokens",
            "currency": "USD",
            "min_charge": 0.01
        },
        "stats": {
            "total_calls": 1000000,
            "success_rate": 99.9,
            "avg_latency": 500,
            "p99_latency": 1000,
            "daily_calls": 10000,
            "daily_tokens": 500000
        },
        "created_at": "2024-03-20T12:00:00Z",
        "updated_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 创建模型配置

```http
POST /api/v1/models
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| model_name | string | 是 | 模型名称 |
| model_type | string | 是 | 模型类型(text/image/audio) |
| provider | string | 是 | 服务提供商 |
| description | string | 否 | 模型描述 |
| model_config | object | 是 | 模型配置 |
| routing_config | object | 否 | 路由配置 |
| pricing | object | 是 | 价格配置 |

请求示例：

```json
{
    "model_name": "gpt-4-turbo",
    "model_type": "text",
    "provider": "openai",
    "description": "GPT-4 Turbo版本，支持更长上下文",
    "model_config": {
        "max_tokens": 16384,
        "supports_stream": true,
        "temperature_range": [0, 2],
        "top_p_range": [0, 1],
        "supports_functions": true,
        "context_window": 128000
    },
    "routing_config": {
        "strategy": "weighted_random",
        "timeout": 60,
        "retry_count": 2,
        "fallback_models": [1, 2]
    },
    "pricing": {
        "input": 0.01,
        "output": 0.03,
        "unit": "1K tokens",
        "currency": "USD",
        "min_charge": 0.01
    }
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "model_id": 4,
        "model_name": "gpt-4-turbo",
        "created_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 更新模型配置

```http
PUT /api/v1/models/{model_id}
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| description | string | 否 | 模型描述 |
| model_config | object | 否 | 模型配置 |
| routing_config | object | 否 | 路由配置 |
| pricing | object | 否 | 价格配置 |
| status | integer | 否 | 模型状态 |

请求示例：

```json
{
    "model_config": {
        "max_tokens": 32768,
        "temperature_range": [0, 1.5]
    },
    "routing_config": {
        "timeout": 30,
        "retry_count": 3
    },
    "status": 1
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "model_id": 4,
        "updated_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 删除模型配置

```http
DELETE /api/v1/models/{model_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| confirm | boolean | 是 | 确认删除(必须为true) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "model_id": 4,
        "deleted_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 获取模型状态

```http
GET /api/v1/models/{model_id}/status
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "model_id": 1,
        "status": 1,
        "health_check": {
            "last_check_time": "2024-03-20T12:00:00Z",
            "is_available": true,
            "latency": 200,
            "error_rate": 0.1
        },
        "current_load": {
            "qps": 10.5,
            "concurrent_requests": 5,
            "queue_length": 0
        },
        "provider_status": {
            "is_available": true,
            "quota_remaining": 1000000,
            "rate_limit_remaining": 50
        }
    }
}
```

#### 获取模型用量

```http
GET /api/v1/models/{model_id}/usage
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| start_time | string | 否 | 开始时间(ISO8601格式) |
| end_time | string | 否 | 结束时间(ISO8601格式) |
| group_by | string | 否 | 分组方式(hour/day/month) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "summary": {
            "total_requests": 100000,
            "total_tokens": 5000000,
            "total_amount": 500.00,
            "success_rate": 99.9,
            "avg_latency": 500
        },
        "timeline": [
            {
                "timestamp": "2024-03-20T00:00:00Z",
                "requests": 1000,
                "tokens": 50000,
                "amount": 5.00,
                "success_rate": 99.9,
                "avg_latency": 480,
                "error_count": 1
            }
        ],
        "error_stats": {
            "timeout": 5,
            "rate_limit": 3,
            "invalid_request": 2
        }
    }
}
```

#### 获取模型日志

```http
GET /api/v1/models/{model_id}/logs
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| start_time | string | 否 | 开始时间 |
| end_time | string | 否 | 结束时间 |
| log_type | string | 否 | 日志类型(operation/error/all) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 1000,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "log_id": "123456789",
                "log_type": "operation",
                "operation": "update_config",
                "operator": "admin",
                "details": {
                    "changed_fields": ["model_config", "status"],
                    "old_values": {
                        "status": 0
                    },
                    "new_values": {
                        "status": 1
                    }
                },
                "created_at": "2024-03-20T12:00:00Z"
            }
        ]
    }
}
```

### 错误码

| 错误码 | 描述 | 解决方案 |
|--------|------|----------|
| 401001 | 未授权访问 | 请检查access_token是否有效 |
| 403001 | 权限不足 | 请检查用户权限 |
| 404001 | 模型不存在 | 请检查模型ID是否正确 |
| 409001 | 模型名称重复 | 请更换模型名称 |
| 422001 | 参数格式错误 | 请检查请求参数格式 |
| 500001 | 服务商接口异常 | 请稍后重试或联系技术支持 |

### 注意事项

1. 模型配置更新后需要一定时间生效
2. 建议定期检查模型状态和性能指标
3. 重要配置变更会记录在操作日志中
4. 模型用量统计有5分钟左右的延迟
5. 日志默认保留30天
6. 删除模型前请确保没有令牌正在使用
7. 路由配置变更可能影响正在处理的请求
8. 建议为核心模型配置备用模型
9. 性能指标每分钟更新一次
10. 模型定价变更不影响已创建的令牌

## 渠道管理模块 (Channel)

### 模块概述

渠道管理模块提供AI服务提供商的配置和管理功能。基于channels表设计，实现多渠道接入、负载均衡和故障转移，确保服务的稳定性和可用性。

### 功能特点

- 支持多个服务提供商接入
- 灵活的密钥和配置管理
- 实时的可用性监控
- 智能的负载均衡策略
- 自动的故障检测和处理
- 详细的用量统计和分析

### 接口列表

#### 渠道管理接口

- [获取渠道列表](#获取渠道列表)
- [获取渠道详情](#获取渠道详情)
- [创建渠道配置](#创建渠道配置)
- [更新渠道配置](#更新渠道配置)
- [删除渠道配置](#删除渠道配置)

#### 渠道监控接口

- [获取渠道状态](#获取渠道状态)
- [获取渠道用量](#获取渠道用量)
- [获取渠道日志](#获取渠道日志)

### 接口详情

#### 获取渠道列表

```http
GET /api/v1/channels
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| provider | string | 否 | 服务商筛选 |
| status | integer | 否 | 状态筛选(1:正常 0:禁用) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 100,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "channel_id": 1,
                "channel_name": "OpenAI-Channel-1",
                "provider": "openai",
                "description": "OpenAI生产环境主通道",
                "channel_config": {
                    "base_url": "https://api.openai.com/v1",
                    "timeout": 30,
                    "max_retries": 3
                },
                "status": 1,
                "stats": {
                    "total_requests": 1000000,
                    "success_rate": 99.9,
                    "avg_latency": 500
                },
                "created_at": "2024-03-20T12:00:00Z"
            }
        ]
    }
}
```

#### 获取渠道详情

```http
GET /api/v1/channels/{channel_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "channel_id": 1,
        "channel_name": "OpenAI-Channel-1",
        "provider": "openai",
        "description": "OpenAI生产环境主通道",
        "channel_config": {
            "base_url": "https://api.openai.com/v1",
            "api_version": "2024-03",
            "timeout": 30,
            "max_retries": 3,
            "proxy": "http://proxy.example.com",
            "retry_codes": [429, 500, 502, 503, 504]
        },
        "auth_config": {
            "auth_type": "api_key",
            "api_key": "sk-xxx...xxx",
            "org_id": "org-xxx"
        },
        "routing_config": {
            "weight": 100,
            "models": ["gpt-4", "gpt-3.5-turbo"],
            "backup_channels": [2, 3]
        },
        "status": 1,
        "stats": {
            "total_requests": 1000000,
            "success_rate": 99.9,
            "avg_latency": 500,
            "error_rate": 0.1,
            "quota_used": 450000,
            "quota_remaining": 550000
        },
        "created_at": "2024-03-20T12:00:00Z",
        "updated_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 创建渠道配置

```http
POST /api/v1/channels
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| channel_name | string | 是 | 渠道名称 |
| provider | string | 是 | 服务提供商 |
| description | string | 否 | 渠道描述 |
| channel_config | object | 是 | 渠道配置 |
| auth_config | object | 是 | 认证配置 |
| routing_config | object | 否 | 路由配置 |

请求示例：

```json
{
    "channel_name": "OpenAI-Channel-2",
    "provider": "openai",
    "description": "OpenAI备用通道",
    "channel_config": {
        "base_url": "https://api.openai.com/v1",
        "api_version": "2024-03",
        "timeout": 30,
        "max_retries": 3,
        "proxy": "http://proxy.example.com",
        "retry_codes": [429, 500, 502, 503, 504]
    },
    "auth_config": {
        "auth_type": "api_key",
        "api_key": "sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        "org_id": "org-xxxxxxxxxxxxxxxx"
    },
    "routing_config": {
        "weight": 50,
        "models": ["gpt-4", "gpt-3.5-turbo"],
        "backup_channels": [1, 3]
    }
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "channel_id": 2,
        "channel_name": "OpenAI-Channel-2",
        "created_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 更新渠道配置

```http
PUT /api/v1/channels/{channel_id}
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| channel_name | string | 否 | 渠道名称 |
| description | string | 否 | 渠道描述 |
| channel_config | object | 否 | 渠道配置 |
| auth_config | object | 否 | 认证配置 |
| routing_config | object | 否 | 路由配置 |
| status | integer | 否 | 渠道状态 |

请求示例：

```json
{
    "channel_config": {
        "timeout": 60,
        "max_retries": 5
    },
    "routing_config": {
        "weight": 80,
        "backup_channels": [1, 4]
    },
    "status": 1
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "channel_id": 2,
        "updated_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 删除渠道配置

```http
DELETE /api/v1/channels/{channel_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| confirm | boolean | 是 | 确认删除(必须为true) |
| force | boolean | 否 | 强制删除(默认: false) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "channel_id": 2,
        "deleted_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 获取渠道状态

```http
GET /api/v1/channels/{channel_id}/status
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "channel_id": 1,
        "status": 1,
        "health_check": {
            "last_check_time": "2024-03-20T12:00:00Z",
            "is_available": true,
            "latency": 200,
            "error_rate": 0.1
        },
        "current_load": {
            "qps": 10.5,
            "concurrent_requests": 5,
            "queue_length": 0
        },
        "provider_status": {
            "is_available": true,
            "quota_remaining": 1000000,
            "rate_limit_remaining": 50
        }
    }
}
```

#### 获取渠道用量

```http
GET /api/v1/channels/{channel_id}/usage
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| start_time | string | 否 | 开始时间(ISO8601格式) |
| end_time | string | 否 | 结束时间(ISO8601格式) |
| group_by | string | 否 | 分组方式(hour/day/month) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "summary": {
            "total_requests": 100000,
            "total_tokens": 5000000,
            "total_amount": 500.00,
            "success_rate": 99.9,
            "avg_latency": 500
        },
        "timeline": [
            {
                "timestamp": "2024-03-20T00:00:00Z",
                "requests": 1000,
                "tokens": 50000,
                "amount": 5.00,
                "success_rate": 99.9,
                "avg_latency": 480,
                "error_count": 1
            }
        ],
        "error_stats": {
            "timeout": 5,
            "rate_limit": 3,
            "invalid_request": 2
        }
    }
}
```

#### 获取渠道日志

```http
GET /api/v1/channels/{channel_id}/logs
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| start_time | string | 否 | 开始时间 |
| end_time | string | 否 | 结束时间 |
| log_type | string | 否 | 日志类型(operation/error/all) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 1000,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "log_id": "123456789",
                "log_type": "operation",
                "operation": "update_config",
                "operator": "admin",
                "details": {
                    "changed_fields": ["channel_config", "status"],
                    "old_values": {
                        "status": 0
                    },
                    "new_values": {
                        "status": 1
                    }
                },
                "created_at": "2024-03-20T12:00:00Z"
            }
        ]
    }
}
```

### 错误码

| 错误码 | 描述 | 解决方案 |
|--------|------|----------|
| 401001 | 未授权访问 | 请检查access_token是否有效 |
| 403001 | 权限不足 | 请检查用户权限 |
| 404001 | 渠道不存在 | 请检查渠道ID是否正确 |
| 409001 | 渠道名称重复 | 请更换渠道名称 |
| 422001 | 参数格式错误 | 请检查请求参数格式 |
| 500001 | 服务商接口异常 | 请稍后重试或联系技术支持 |

### 注意事项

1. 渠道配置更新后需要一定时间生效
2. 建议定期检查渠道状态和性能指标
3. 重要配置变更会记录在操作日志中
4. 渠道用量统计有5分钟左右的延迟
5. 日志默认保留30天
6. 删除渠道前请确保没有令牌正在使用
7. 路由配置变更可能影响正在处理的请求
8. 建议为核心渠道配置备用渠道
9. 性能指标每分钟更新一次
10. 渠道定价变更不影响已创建的令牌

## 计费模块 (Billing)

### 模块概述

计费模块提供用户额度的管理、充值和消费功能。基于billings表设计，实现完整的计费流程，包括预付费充值、后付费结算、用量统计和账单管理。

### 功能特点

- 支持预付费和后付费模式
- 灵活的计费策略配置
- 多级计费方案管理
- 实时的用量计算和统计
- 完整的账单和流水记录
- 支持多种支付方式接入

### 接口列表

#### 账户管理接口

- [获取账户余额](#获取账户余额)
- [获取计费方案](#获取计费方案)
- [更新计费方案](#更新计费方案)
- [设置自动充值](#设置自动充值)

#### 充值消费接口

- [创建充值订单](#创建充值订单)
- [查询订单状态](#查询订单状态)
- [获取消费记录](#获取消费记录)
- [获取充值记录](#获取充值记录)

#### 账单管理接口

- [获取账单列表](#获取账单列表)
- [获取账单详情](#获取账单详情)
- [导出账单数据](#导出账单数据)

### 接口详情

#### 获取账户余额

```http
GET /api/v1/billing/balance
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "user_id": "12345678",
        "balance": 1000.00,
        "frozen_amount": 50.00,
        "available_amount": 950.00,
        "credit_limit": 5000.00,
        "credit_used": 0.00,
        "auto_recharge": {
            "enabled": true,
            "threshold": 100.00,
            "amount": 500.00
        },
        "billing_mode": "prepaid",
        "currency": "USD",
        "updated_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 获取计费方案

```http
GET /api/v1/billing/plans
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "user_id": "12345678",
        "billing_mode": "prepaid",
        "current_plan": {
            "plan_id": "premium",
            "plan_name": "高级版",
            "models": [
                {
                    "model_id": 1,
                    "model_name": "gpt-4",
                    "pricing": {
                        "input": 0.01,
                        "output": 0.03,
                        "unit": "1K tokens"
                    },
                    "discount": 0.9
                }
            ],
            "features": {
                "max_tokens": 8192,
                "priority_support": true,
                "concurrent_requests": 10
            }
        },
        "available_plans": [
            {
                "plan_id": "basic",
                "plan_name": "基础版",
                "pricing": [
                    {
                        "model_id": 1,
                        "pricing": {
                            "input": 0.015,
                            "output": 0.045,
                            "unit": "1K tokens"
                        }
                    }
                ]
            }
        ]
    }
}
```

#### 更新计费方案

```http
PUT /api/v1/billing/plans
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| plan_id | string | 是 | 计费方案ID |
| billing_mode | string | 否 | 计费模式(prepaid/postpaid) |
| effective_time | string | 否 | 生效时间 |

请求示例：

```json
{
    "plan_id": "premium",
    "billing_mode": "prepaid",
    "effective_time": "2024-04-01T00:00:00Z"
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "user_id": "12345678",
        "plan_id": "premium",
        "effective_time": "2024-04-01T00:00:00Z",
        "updated_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 设置自动充值

```http
PUT /api/v1/billing/auto-recharge
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| enabled | boolean | 是 | 是否启用 |
| threshold | number | 否 | 触发阈值 |
| amount | number | 否 | 充值金额 |
| payment_method | string | 否 | 支付方式 |

请求示例：

```json
{
    "enabled": true,
    "threshold": 100.00,
    "amount": 500.00,
    "payment_method": "credit_card"
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "user_id": "12345678",
        "auto_recharge": {
            "enabled": true,
            "threshold": 100.00,
            "amount": 500.00,
            "payment_method": "credit_card"
        },
        "updated_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 创建充值订单

```http
POST /api/v1/billing/recharge
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| amount | number | 是 | 充值金额 |
| currency | string | 否 | 货币类型(默认: USD) |
| payment_method | string | 是 | 支付方式(alipay/wechat/card) |
| description | string | 否 | 充值说明 |

请求示例：

```json
{
    "amount": 1000.00,
    "currency": "USD",
    "payment_method": "alipay",
    "description": "账户充值"
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "order_id": "REG20240320123456",
        "amount": 1000.00,
        "currency": "USD",
        "payment_url": "https://payment.example.com/pay/xxx",
        "qr_code": "data:image/png;base64,xxx...",
        "expire_time": "2024-03-20T12:30:00Z",
        "created_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 查询订单状态

```http
GET /api/v1/billing/orders/{order_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "order_id": "REG20240320123456",
        "user_id": "12345678",
        "amount": 1000.00,
        "currency": "USD",
        "payment_method": "alipay",
        "status": "paid",
        "paid_time": "2024-03-20T12:05:00Z",
        "description": "账户充值",
        "created_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 获取消费记录

```http
GET /api/v1/billing/consumption
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| start_time | string | 否 | 开始时间 |
| end_time | string | 否 | 结束时间 |
| model_id | integer | 否 | 模型筛选 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 1000,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "record_id": "CON20240320123456",
                "user_id": "12345678",
                "model_id": 1,
                "model_name": "gpt-4",
                "request_id": "req_xxxxx",
                "tokens": 500,
                "amount": 0.05,
                "currency": "USD",
                "created_at": "2024-03-20T12:00:00Z"
            }
        ],
        "summary": {
            "total_tokens": 50000,
            "total_amount": 5.00
        }
    }
}
```

#### 获取充值记录

```http
GET /api/v1/billing/recharge
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| start_time | string | 否 | 开始时间 |
| end_time | string | 否 | 结束时间 |
| status | string | 否 | 状态筛选 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 100,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "order_id": "REG20240320123456",
                "user_id": "12345678",
                "amount": 1000.00,
                "currency": "USD",
                "payment_method": "alipay",
                "status": "paid",
                "description": "账户充值",
                "created_at": "2024-03-20T12:00:00Z",
                "paid_time": "2024-03-20T12:05:00Z"
            }
        ],
        "summary": {
            "total_amount": 10000.00,
            "success_count": 98,
            "failed_count": 2
        }
    }
}
```

#### 获取账单列表

```http
GET /api/v1/billing/bills
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| year | integer | 否 | 年份 |
| month | integer | 否 | 月份 |
| status | string | 否 | 状态筛选 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 12,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "bill_id": "BILL202403",
                "user_id": "12345678",
                "year": 2024,
                "month": 3,
                "total_amount": 1500.00,
                "paid_amount": 1500.00,
                "status": "paid",
                "period_start": "2024-03-01T00:00:00Z",
                "period_end": "2024-03-31T23:59:59Z",
                "due_date": "2024-04-10T00:00:00Z",
                "paid_time": "2024-04-05T12:00:00Z"
            }
        ]
    }
}
```

#### 获取账单详情

```http
GET /api/v1/billing/bills/{bill_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "bill_id": "BILL202403",
        "user_id": "12345678",
        "year": 2024,
        "month": 3,
        "bill_summary": {
            "total_amount": 1500.00,
            "paid_amount": 1500.00,
            "total_requests": 10000,
            "total_tokens": 500000
        },
        "model_details": [
            {
                "model_id": 1,
                "model_name": "gpt-4",
                "requests": 5000,
                "tokens": 300000,
                "amount": 900.00
            },
            {
                "model_id": 2,
                "model_name": "gpt-3.5-turbo",
                "requests": 5000,
                "tokens": 200000,
                "amount": 600.00
            }
        ],
        "payment_records": [
            {
                "payment_id": "PAY20240405123456",
                "amount": 1500.00,
                "payment_method": "alipay",
                "paid_time": "2024-04-05T12:00:00Z"
            }
        ],
        "status": "paid",
        "period_start": "2024-03-01T00:00:00Z",
        "period_end": "2024-03-31T23:59:59Z",
        "due_date": "2024-04-10T00:00:00Z",
        "paid_time": "2024-04-05T12:00:00Z",
        "created_at": "2024-04-01T00:00:00Z"
    }
}
```

#### 导出账单数据

```http
GET /api/v1/billing/bills/{bill_id}/export
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| format | string | 否 | 导出格式(csv/xlsx，默认: xlsx) |
| type | string | 否 | 导出类型(summary/detail，默认: detail) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "download_url": "https://example.com/download/bill_202403.xlsx",
        "expire_time": "2024-03-20T13:00:00Z"
    }
}
```

### 错误码

| 错误码 | 描述 | 解决方案 |
|--------|------|----------|
| 401001 | 未授权访问 | 请检查access_token是否有效 |
| 403001 | 余额不足 | 请充值后重试 |
| 403002 | 超出信用额度 | 请结清账单后重试 |
| 404001 | 订单不存在 | 请检查订单ID是否正确 |
| 404002 | 账单不存在 | 请检查账单ID是否正确 |
| 409001 | 重复支付 | 订单已支付，请勿重复支付 |
| 422001 | 金额无效 | 请检查充值金额是否正确 |
| 422002 | 支付方式无效 | 请选择正确的支付方式 |

### 注意事项

1. 账单按自然月生成，次月1号生成上月账单
2. 账单支付截止日期为次月10号
3. 超过截止日期未支付将产生滞纳金
4. 支持多种支付方式和货币类型
5. 账单数据导出链接有效期为1小时
6. 消费记录实时更新，账单金额可能有轻微延迟
7. 建议开启自动充值，避免余额不足
8. 如有账单争议，请在15天内提出
9. 支持对账单进行分期付款（需单独申请）
10. 可以设置账单金额预警通知

## 消息管理模块 (Message)

### 模块概述

消息管理模块提供AI对话消息的存储、查询和分析功能。基于message_saves表设计，实现消息的持久化存储、标签管理和内容分析，支持对话历史的回溯和数据分析。

### 功能特点

- 支持多种消息类型存储
- 灵活的消息检索功能
- 对话上下文管理
- 消息标签和分类
- 对话数据分析
- 支持批量操作

### 接口列表

#### 消息管理接口

- [保存消息](#保存消息)
- [获取消息列表](#获取消息列表)
- [获取消息详情](#获取消息详情)
- [更新消息标签](#更新消息标签)
- [删除消息](#删除消息)

#### 对话管理接口

- [获取对话列表](#获取对话列表)
- [获取对话详情](#获取对话详情)
- [导出对话记录](#导出对话记录)

### 接口详情

#### 保存消息

```http
POST /api/v1/messages/save
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| conversation_id | string | 是 | 对话ID |
| model_id | integer | 是 | 模型ID |
| role | string | 是 | 角色(user/assistant/system) |
| content | string | 是 | 消息内容 |
| tokens | integer | 是 | token数量 |
| message_type | string | 否 | 消息类型(text/image/audio) |
| message_tags | array | 否 | 消息标签 |
| parent_id | string | 否 | 父消息ID |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "message_id": "msg_123456",
        "conversation_id": "conv_123456",
        "created_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 获取消息列表

```http
GET /api/v1/messages/list
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| conversation_id | string | 否 | 按对话筛选 |
| model_id | integer | 否 | 按模型筛选 |
| role | string | 否 | 按角色筛选 |
| message_type | string | 否 | 按类型筛选 |
| message_tags | array | 否 | 按标签筛选 |
| start_time | string | 否 | 开始时间 |
| end_time | string | 否 | 结束时间 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 100,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "message_id": "msg_123456",
                "conversation_id": "conv_123456",
                "model_id": 1,
                "role": "user",
                "content": "你好",
                "tokens": 2,
                "message_type": "text",
                "message_tags": ["greeting"],
                "created_at": "2024-03-20T12:00:00Z"
            }
        ]
    }
}
```

#### 获取消息详情

```http
GET /api/v1/messages/{message_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "message_id": "msg_123456",
        "conversation_id": "conv_123456",
        "model_id": 1,
        "role": "user",
        "content": "你好",
        "tokens": 2,
        "message_type": "text",
        "message_tags": ["greeting"],
        "created_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 更新消息标签

```http
PUT /api/v1/messages/{message_id}
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| message_tags | array | 是 | 消息标签 |

请求示例：

```json
{
    "message_tags": ["greeting", "urgent"]
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "message_id": "msg_123456",
        "conversation_id": "conv_123456",
        "updated_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 删除消息

```http
DELETE /api/v1/messages/{message_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| confirm | boolean | 是 | 确认删除(必须为true) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "message_id": "msg_123456",
        "deleted_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 获取对话列表

```http
GET /api/v1/conversations
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| conversation_id | string | 否 | 按对话筛选 |
| model_id | integer | 否 | 按模型筛选 |
| start_time | string | 否 | 开始时间 |
| end_time | string | 否 | 结束时间 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 100,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "conversation_id": "conv_123456",
                "model_id": 1,
                "start_time": "2024-03-20T12:00:00Z",
                "end_time": "2024-03-20T12:30:00Z"
            }
        ]
    }
}
```

#### 获取对话详情

```http
GET /api/v1/conversations/{conversation_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "conversation_id": "conv_123456",
        "model_id": 1,
        "messages": [
            {
                "message_id": "msg_123456",
                "role": "user",
                "content": "你好",
                "tokens": 2,
                "message_type": "text",
                "message_tags": ["greeting"],
                "created_at": "2024-03-20T12:00:00Z"
            },
            {
                "message_id": "msg_123457",
                "role": "assistant",
                "content": "你好，有什么可以帮你的吗？",
                "tokens": 10,
                "message_type": "text",
                "message_tags": ["response"],
                "created_at": "2024-03-20T12:01:00Z"
            }
        ]
    }
}
```

#### 导出对话记录

```http
GET /api/v1/conversations/{conversation_id}/export
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| format | string | 否 | 导出格式(csv/xlsx，默认: xlsx) |
| type | string | 否 | 导出类型(summary/detail，默认: detail) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "download_url": "https://example.com/download/conversation_20240320.xlsx",
        "expire_time": "2024-03-20T13:00:00Z"
    }
}
```

### 错误码

| 错误码 | 描述 | 解决方案 |
|--------|------|----------|
| 401001 | 未授权访问 | 请检查access_token是否有效 |
| 403001 | 权限不足 | 请检查用户权限 |
| 404001 | 消息不存在 | 请检查消息ID是否正确 |
| 409001 | 消息重复 | 请检查消息是否重复 |
| 422001 | 参数格式错误 | 请检查请求参数格式 |
| 500001 | 服务器异常 | 请稍后重试或联系技术支持 |

### 注意事项

1. 消息保存后，消息内容将无法修改
2. 消息删除后不可恢复，请谨慎操作
3. 消息标签和分类有助于快速检索和分类消息
4. 对话历史记录有助于回溯和分析对话内容
5. 消息详情页面提供消息的详细信息和操作日志
6. 消息列表页面提供消息的列表和搜索功能
7. 消息详情页面提供消息的详细信息和操作日志
8. 消息删除后不可恢复，请谨慎操作
9. 消息列表页面提供消息的列表和搜索功能
10. 消息详情页面提供消息的详细信息和操作日志

## 系统监控模块 (Monitor)

### 模块概述

系统监控模块提供全面的系统运行状态监控和性能分析功能。基于metrics表设计，实现系统指标的采集、分析和告警，确保系统的稳定运行和问题的及时发现。

### 功能特点

- 实时性能指标监控
- 多维度数据统计
- 自定义告警规则
- 系统健康检查
- 资源使用分析
- 性能瓶颈诊断

### 接口列表

#### 指标监控接口

- [获取系统概况](#获取系统概况)
- [获取性能指标](#获取性能指标)
- [获取资源用量](#获取资源用量)
- [获取服务状态](#获取服务状态)

#### 告警管理接口

- [获取告警规则](#获取告警规则)
- [设置告警规则](#设置告警规则)
- [获取告警历史](#获取告警历史)
- [处理告警事件](#处理告警事件)

### 接口详情

#### 获取系统概况

```http
GET /api/v1/monitor/overview
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "system_status": "healthy",
        "total_requests": 1000000,
        "success_rate": 99.9,
        "avg_latency": 200,
        "active_users": 5000,
        "concurrent_requests": 100,
        "cpu_usage": 45.5,
        "memory_usage": 65.8,
        "disk_usage": 55.2,
        "network_io": {
            "in_bytes": 1024000,
            "out_bytes": 2048000
        },
        "updated_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 获取性能指标

```http
GET /api/v1/monitor/metrics
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| metrics | array | 否 | 指标类型列表 |
| start_time | string | 否 | 开始时间 |
| end_time | string | 否 | 结束时间 |
| interval | string | 否 | 统计间隔(1m/5m/1h) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "metrics": {
            "qps": [
                {
                    "timestamp": "2024-03-20T12:00:00Z",
                    "value": 100.5
                }
            ],
            "latency": [
                {
                    "timestamp": "2024-03-20T12:00:00Z",
                    "avg": 200,
                    "p95": 500,
                    "p99": 800
                }
            ],
            "error_rate": [
                {
                    "timestamp": "2024-03-20T12:00:00Z",
                    "value": 0.1
                }
            ]
        },
        "period": {
            "start_time": "2024-03-20T11:00:00Z",
            "end_time": "2024-03-20T12:00:00Z",
            "interval": "1m"
        }
    }
}
```

#### 获取资源用量

```http
GET /api/v1/monitor/resources
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| resources | array | 否 | 资源类型列表 |
| start_time | string | 否 | 开始时间 |
| end_time | string | 否 | 结束时间 |
| interval | string | 否 | 统计间隔(1m/5m/1h) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "resources": {
            "cpu": [
                {
                    "timestamp": "2024-03-20T12:00:00Z",
                    "usage": 45.5,
                    "load": 2.5
                }
            ],
            "memory": [
                {
                    "timestamp": "2024-03-20T12:00:00Z",
                    "total": 16384,
                    "used": 10240,
                    "free": 6144
                }
            ],
            "disk": [
                {
                    "timestamp": "2024-03-20T12:00:00Z",
                    "total": 1024000,
                    "used": 512000,
                    "free": 512000
                }
            ]
        },
        "period": {
            "start_time": "2024-03-20T11:00:00Z",
            "end_time": "2024-03-20T12:00:00Z",
            "interval": "1m"
        }
    }
}
```

#### 获取服务状态

```http
GET /api/v1/monitor/services
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "services": [
            {
                "service_name": "api-gateway",
                "status": "running",
                "uptime": 864000,
                "last_check": "2024-03-20T12:00:00Z",
                "health_check": {
                    "status": "healthy",
                    "latency": 50
                }
            },
            {
                "service_name": "auth-service",
                "status": "running",
                "uptime": 864000,
                "last_check": "2024-03-20T12:00:00Z",
                "health_check": {
                    "status": "healthy",
                    "latency": 30
                }
            }
        ],
        "updated_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 获取告警规则

```http
GET /api/v1/monitor/alerts
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| alert_type | string | 否 | 告警类型筛选 |
| status | string | 否 | 状态筛选 |
| keyword | string | 否 | 关键词搜索 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 100,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "alert_id": "alert_123456",
                "alert_name": "CPU使用率告警",
                "alert_type": "system",
                "status": "active",
                "threshold": 80,
                "created_at": "2024-03-20T12:00:00Z"
            }
        ]
    }
}
```

#### 设置告警规则

```http
POST /api/v1/monitor/alerts
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| alert_name | string | 是 | 告警名称 |
| alert_type | string | 是 | 告警类型(system/resource/service) |
| threshold | number | 是 | 告警阈值 |
| duration | string | 是 | 告警持续时间 |
| notification_config | object | 否 | 通知配置 |

请求示例：

```json
{
    "alert_name": "CPU使用率告警",
    "alert_type": "system",
    "threshold": 80,
    "duration": "5m",
    "notification_config": {
        "email": true,
        "sms": false,
        "wechat": true
    }
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "alert_id": "alert_123456",
        "alert_name": "CPU使用率告警",
        "alert_type": "system",
        "status": "active",
        "threshold": 80,
        "duration": "5m",
        "notification_config": {
            "email": true,
            "sms": false,
            "wechat": true
        },
        "created_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 获取告警历史

```http
GET /api/v1/monitor/alerts
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| alert_type | string | 否 | 告警类型筛选 |
| status | string | 否 | 状态筛选 |
| keyword | string | 否 | 关键词搜索 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 100,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "alert_id": "alert_123456",
                "alert_name": "CPU使用率告警",
                "alert_type": "system",
                "status": "resolved",
                "threshold": 80,
                "created_at": "2024-03-20T12:00:00Z"
            }
        ]
    }
}
```

#### 处理告警事件

```http
POST /api/v1/monitor/alerts/{alert_id}/resolve
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| resolution_note | string | 是 | 处理说明 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "alert_id": "alert_123456",
        "status": "resolved",
        "resolution_note": "处理说明",
        "resolved_at": "2024-03-20T12:00:00Z"
    }
}
```

### 错误码

| 错误码 | 描述 | 解决方案 |
|--------|------|----------|
| 401001 | 未授权访问 | 请检查access_token是否有效 |
| 403001 | 权限不足 | 请检查用户权限 |
| 404001 | 告警规则不存在 | 请检查告警规则ID是否正确 |
| 409001 | 告警名称重复 | 请更换告警名称 |
| 422001 | 参数格式错误 | 请检查请求参数格式 |
| 500001 | 告警处理异常 | 请稍后重试或联系技术支持 |

### 注意事项

1. 告警规则设置后需要一定时间生效
2. 建议定期检查告警规则和告警历史
3. 重要告警事件会记录在操作日志中
4. 告警通知配置需要提前设置
5. 告警处理后需要及时更新告警状态
6. 告警通知配置变更可能影响正在处理的请求
7. 建议为核心告警配置备用告警规则
8. 告警处理后需要及时更新告警状态
9. 告警通知配置变更可能影响正在处理的请求
10. 告警处理后需要及时更新告警状态

## 告警通知模块 (Notification)

### 模块概述

告警通知模块提供告警事件的通知配置和管理功能。基于notification_configs表设计，实现告警事件的通知配置、发送和记录。支持多种通知方式和告警策略。

### 功能特点

- 支持多种通知方式（邮件/短信/微信）
- 灵活的告警策略配置
- 实时的告警事件记录
- 详细的告警历史查询
- 支持告警事件的批量操作

### 接口列表

#### 告警通知接口

- [获取通知配置](#获取通知配置)
- [设置通知配置](#设置通知配置)
- [发送通知](#发送通知)
- [获取通知记录](#获取通知记录)

### 接口详情

#### 获取通知配置

```http
GET /api/v1/notification/configs
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| alert_type | string | 否 | 告警类型筛选 |
| status | string | 否 | 状态筛选 |
| keyword | string | 否 | 关键词搜索 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 100,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "config_id": "config_123456",
                "alert_type": "system",
                "status": "active",
                "notification_config": {
                    "email": true,
                    "sms": false,
                    "wechat": true
                },
                "created_at": "2024-03-20T12:00:00Z"
            }
        ]
    }
}
```

#### 设置通知配置

```http
POST /api/v1/notification/configs
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| alert_type | string | 是 | 告警类型(system/resource/service) |
| notification_config | object | 是 | 通知配置 |

请求示例：

```json
{
    "alert_type": "system",
    "notification_config": {
        "email": true,
        "sms": false,
        "wechat": true
    }
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "config_id": "config_123456",
        "alert_type": "system",
        "status": "active",
        "notification_config": {
            "email": true,
            "sms": false,
            "wechat": true
        },
        "created_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 发送通知

```http
POST /api/v1/notification/send
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| alert_id | string | 是 | 告警ID |
| notification_config | object | 是 | 通知配置 |

请求示例：

```json
{
    "alert_id": "alert_123456",
    "notification_config": {
        "email": true,
        "sms": false,
        "wechat": true
    }
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "notification_id": "noti_123456",
        "status": "sent",
        "sent_time": "2024-03-20T12:00:00Z"
    }
}
```

#### 获取通知记录

```http
GET /api/v1/notification/records
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| alert_type | string | 否 | 告警类型筛选 |
| status | string | 否 | 状态筛选 |
| keyword | string | 否 | 关键词搜索 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 100,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "record_id": "noti_123456",
                "alert_id": "alert_123456",
                "status": "sent",
                "sent_time": "2024-03-20T12:00:00Z"
            }
        ]
    }
}
```

### 错误码

| 错误码 | 描述 | 解决方案 |
|--------|------|----------|
| 401001 | 未授权访问 | 请检查access_token是否有效 |
| 403001 | 权限不足 | 请检查用户权限 |
| 404001 | 通知配置不存在 | 请检查通知配置ID是否正确 |
| 409001 | 告警类型重复 | 请更换告警类型 |
| 422001 | 参数格式错误 | 请检查请求参数格式 |
| 500001 | 通知发送异常 | 请稍后重试或联系技术支持 |

### 注意事项

1. 通知配置设置后需要一定时间生效
2. 建议定期检查通知配置和通知记录
3. 重要通知配置变更会记录在操作日志中
4. 通知发送后需要及时更新通知状态
5. 通知记录保留时间为30天
6. 删除通知配置前请确保没有告警事件正在发送
7. 通知配置变更可能影响正在处理的请求
8. 建议为核心通知配置备用通知配置
9. 通知发送后需要及时更新通知状态
10. 通知记录保留时间为30天

## 请求日志模块 (RequestLog)

### 模块概述

请求日志模块提供API请求的完整记录和分析功能。基于request_logs表设计，实现请求的追踪、统计和诊断，支持问题排查和性能优化。

### 功能特点

- 完整的请求记录
- 灵活的日志检索
- 请求链路追踪
- 性能数据统计
- 错误日志分析
- 支持批量导出

### 接口列表

#### 日志查询接口

- [获取请求日志](#获取请求日志)
- [获取日志详情](#获取日志详情)
- [导出日志数据](#导出日志数据)

#### 统计分析接口

- [获取请求统计](#获取请求统计)
- [获取错误分析](#获取错误分析)
- [获取性能分析](#获取性能分析)

### 接口详情

#### 获取请求日志

```http
GET /api/v1/request-logs
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| start_time | string | 否 | 开始时间 |
| end_time | string | 否 | 结束时间 |
| status | string | 否 | 状态筛选 |
| model_id | integer | 否 | 模型筛选 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 1000,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "log_id": "123456789",
                "request_id": "req_xxxxx",
                "model_id": 1,
                "model_name": "gpt-4",
                "request_time": "2024-03-20T12:00:00Z",
                "response_time": "2024-03-20T12:00:01Z",
                "latency": 1000,
                "status": 1,
                "tokens": 500,
                "amount": 0.05,
                "client_ip": "192.168.1.1",
                "error_message": null
            }
        ]
    }
}
```

#### 获取日志详情

```http
GET /api/v1/request-logs/{log_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "log_id": "123456789",
        "request_id": "req_xxxxx",
        "model_id": 1,
        "model_name": "gpt-4",
        "request_time": "2024-03-20T12:00:00Z",
        "response_time": "2024-03-20T12:00:01Z",
        "latency": 1000,
        "status": 1,
        "tokens": 500,
        "amount": 0.05,
        "client_ip": "192.168.1.1",
        "error_message": null
    }
}
```

#### 导出日志数据

```http
GET /api/v1/request-logs/{log_id}/export
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| format | string | 否 | 导出格式(csv/xlsx，默认: xlsx) |
| type | string | 否 | 导出类型(summary/detail，默认: detail) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "download_url": "https://example.com/download/request_log_202403.xlsx",
        "expire_time": "2024-03-20T13:00:00Z"
    }
}
```

#### 获取请求统计

```http
GET /api/v1/request-logs/statistics
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| start_time | string | 否 | 开始时间 |
| end_time | string | 否 | 结束时间 |
| status | string | 否 | 状态筛选 |
| model_id | integer | 否 | 模型筛选 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total_requests": 10000,
        "success_count": 9500,
        "failed_count": 500,
        "average_latency": 500,
        "error_rate": 5
    }
}
```

#### 获取错误分析

```http
GET /api/v1/request-logs/errors
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| start_time | string | 否 | 开始时间 |
| end_time | string | 否 | 结束时间 |
| status | string | 否 | 状态筛选 |
| model_id | integer | 否 | 模型筛选 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total_errors": 500,
        "error_types": {
            "timeout": 100,
            "rate_limit": 200,
            "invalid_request": 200
        },
        "error_messages": {
            "500": 100,
            "429": 150,
            "400": 100,
            "503": 50
        }
    }
}
```

#### 获取性能分析

```http
GET /api/v1/request-logs/performance
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| start_time | string | 否 | 开始时间 |
| end_time | string | 否 | 结束时间 |
| status | string | 否 | 状态筛选 |
| model_id | integer | 否 | 模型筛选 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total_requests": 10000,
        "average_latency": 500,
        "p95_latency": 700,
        "p99_latency": 900
    }
}
```

### 错误码

| 错误码 | 描述 | 解决方案 |
|--------|------|----------|
| 401001 | 未授权访问 | 请检查access_token是否有效 |
| 403001 | 权限不足 | 请检查用户权限 |
| 404001 | 日志记录不存在 | 请检查日志ID是否正确 |
| 409001 | 日志类型重复 | 请更换日志类型 |
| 422001 | 参数格式错误 | 请检查请求参数格式 |
| 500001 | 日志分析异常 | 请稍后重试或联系技术支持 |

### 注意事项

1. 日志记录保留时间为30天
2. 日志数据导出链接有效期为1小时
3. 日志查询和统计功能需要一定时间处理
4. 日志分析结果可能会有延迟
5. 建议定期检查日志记录和统计数据
6. 日志查询和统计功能需要一定时间处理
7. 日志分析结果可能会有延迟
8. 建议为核心日志配置备用日志记录
9. 性能分析结果每分钟更新一次
10. 日志查询和统计功能需要一定时间处理

## 中继服务模块 (Relay)

### 模块概述

中继服务模块是系统的核心接入层，提供兼容OpenAI格式的API接口。支持文本、图像、音频、视频等多模态内容的处理，以及实时流式传输能力。

### 功能特点

- 完全兼容OpenAI接口格式
- 支持多种内容类型处理
- 实时流式数据传输
- 多模态内容融合
- 智能负载均衡
- 自动故障转移

### 接口列表

#### 文本处理接口

- [文本补全](#文本补全)
- [聊天对话](#聊天对话)
- [文本嵌入](#文本嵌入)

#### 图像处理接口

- [图像生成](#图像生成)
- [图像编辑](#图像编辑)
- [图像变体](#图像变体)

#### 音频处理接口

- [语音转文本](#语音转文本)
- [文本转语音](#文本转语音)
- [实时语音对话](#实时语音对话)

#### 视频处理接口

- [视频生成](#视频生成)
- [视频编辑](#视频编辑)
- [实时视频对话](#实时视频对话)

#### 多模态接口

- [多模态理解](#多模态理解)
- [多模态生成](#多模态生成)
- [多模态对话](#多模态对话)

### 接口详情

#### 文本补全

```http
POST /api/v1/completions
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| model | string | 是 | 模型ID |
| prompt | string | 是 | 提示文本 |
| max_tokens | integer | 否 | 最大生成长度(默认: 2048) |
| temperature | number | 否 | 采样温度(0-2，默认: 1) |
| top_p | number | 否 | 核采样(0-1，默认: 1) |
| n | integer | 否 | 生成数量(默认: 1) |
| stream | boolean | 否 | 是否流式输出(默认: false) |
| stop | array/string | 否 | 停止标记 |
| presence_penalty | number | 否 | 重复惩罚(默认: 0) |
| frequency_penalty | number | 否 | 频率惩罚(默认: 0) |

请求示例：

```json
{
    "model": "gpt-3.5-turbo",
    "prompt": "写一首关于春天的诗",
    "max_tokens": 100,
    "temperature": 0.7
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "id": "cmpl-123456",
        "object": "text_completion",
        "created": 1679289600,
        "model": "gpt-3.5-turbo",
        "choices": [
            {
                "text": "春风拂面百花香，\n蝴蝶翩翩舞斜阳。\n绿柳轻摇迎客至，\n小溪潺潺唱悠扬。",
                "index": 0,
                "finish_reason": "stop"
            }
        ],
        "usage": {
            "prompt_tokens": 9,
            "completion_tokens": 36,
            "total_tokens": 45
        }
    }
}
```

#### 聊天对话

```http
POST /api/v1/chat/completions
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| model | string | 是 | 模型ID |
| messages | array | 是 | 对话消息列表 |
| max_tokens | integer | 否 | 最大生成长度(默认: 2048) |
| temperature | number | 否 | 采样温度(0-2，默认: 1) |
| top_p | number | 否 | 核采样(0-1，默认: 1) |
| n | integer | 否 | 生成数量(默认: 1) |
| stream | boolean | 否 | 是否流式输出(默认: false) |
| stop | array/string | 否 | 停止标记 |
| presence_penalty | number | 否 | 重复惩罚(默认: 0) |
| frequency_penalty | number | 否 | 频率惩罚(默认: 0) |

请求示例：

```json
{
    "model": "gpt-4",
    "messages": [
        {
            "role": "system",
            "content": "你是一个专业的AI助手。"
        },
        {
        "role": "user",
            "content": "介绍一下你自己。"
        }
    ],
    "temperature": 0.7
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "id": "chatcmpl-123456",
        "object": "chat.completion",
        "created": 1679289600,
        "model": "gpt-4",
        "choices": [
            {
                "index": 0,
                "message": {
                    "role": "assistant",
                    "content": "你好！我是一个AI助手，基于GPT-4模型训练..."
                },
                "finish_reason": "stop"
            }
        ],
        "usage": {
            "prompt_tokens": 23,
            "completion_tokens": 42,
            "total_tokens": 65
        }
    }
}
```

#### 文本嵌入

```http
POST /api/v1/embeddings
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| model | string | 是 | 模型ID |
| input | string/array | 是 | 输入文本 |
| encoding_format | string | 否 | 编码格式(float/base64) |

请求示例：

```json
{
    "model": "text-embedding-ada-002",
    "input": "OpenAI API文档",
    "encoding_format": "float"
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "object": "list",
        "data": [
            {
                "object": "embedding",
                "embedding": [0.0023064255, -0.009327292, ...],
                "index": 0
            }
        ],
        "model": "text-embedding-ada-002",
        "usage": {
            "prompt_tokens": 5,
            "total_tokens": 5
        }
    }
}
```

#### 图像生成

```http
POST /api/v1/images/generations
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| model | string | 是 | 模型ID |
| prompt | string | 是 | 图像描述 |
| n | integer | 否 | 生成数量(默认: 1) |
| size | string | 否 | 图像尺寸(1024x1024/512x512/256x256) |
| quality | string | 否 | 图像质量(standard/hd) |
| style | string | 否 | 图像风格 |
| response_format | string | 否 | 返回格式(url/b64_json) |

请求示例：

```json
{
    "model": "dall-e-3",
    "prompt": "一只可爱的小猫在阳光下玩耍",
    "n": 1,
    "size": "1024x1024",
    "quality": "hd",
    "response_format": "url"
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "created": 1679289600,
        "data": [
            {
                "url": "https://example.com/images/xxx.png",
                "revised_prompt": "一只橘色的小猫在温暖的阳光下玩耍，背景是绿色的草地"
            }
        ]
    }
}
```

#### 图像编辑

```http
POST /api/v1/images/edits
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: multipart/form-data
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| model | string | 是 | 模型ID |
| image | file | 是 | 原始图像(PNG格式) |
| mask | file | 否 | 蒙版图像 |
| prompt | string | 是 | 编辑描述 |
| n | integer | 否 | 生成数量(默认: 1) |
| size | string | 否 | 图像尺寸 |
| response_format | string | 否 | 返回格式(url/b64_json) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "created": 1679289600,
        "data": [
            {
                "url": "https://example.com/images/xxx.png",
                "revised_prompt": "将图像背景改为海滩场景"
            }
        ]
    }
}
```

#### 图像变体

```http
POST /api/v1/images/variations
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: multipart/form-data
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| model | string | 是 | 模型ID |
| image | file | 是 | 原始图像(PNG格式) |
| n | integer | 否 | 生成数量(默认: 1) |
| size | string | 否 | 图像尺寸 |
| response_format | string | 否 | 返回格式(url/b64_json) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "created": 1679289600,
        "data": [
            {
                "url": "https://example.com/images/xxx.png"
            }
        ]
    }
}
```

#### 语音转文本

```http
POST /api/v1/audio/transcriptions
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: multipart/form-data
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| model | string | 是 | 模型ID |
| file | file | 是 | 音频文件 |
| language | string | 否 | 语言代码(默认: auto) |
| prompt | string | 否 | 转录提示 |
| response_format | string | 否 | 返回格式(json/text/srt/vtt) |
| temperature | number | 否 | 采样温度(0-1) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "text": "这是一段音频转文字的示例。",
        "language": "zh",
        "duration": 10.5,
        "segments": [
            {
                "id": 0,
                "start": 0.0,
                "end": 2.5,
                "text": "这是一段"
            },
            {
                "id": 1,
                "start": 2.5,
                "end": 5.2,
                "text": "音频转文字"
            },
            {
                "id": 2,
                "start": 5.2,
                "end": 10.5,
                "text": "的示例。"
            }
        ]
    }
}
```

#### 文本转语音

```http
POST /api/v1/audio/speech
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| model | string | 是 | 模型ID |
| input | string | 是 | 输入文本 |
| voice | string | 是 | 语音类型(alloy/echo/fable/onyx/nova/shimmer) |
| response_format | string | 否 | 音频格式(mp3/opus/aac/flac) |
| speed | number | 否 | 语速(0.25-4.0) |

请求示例：

```json
{
    "model": "tts-1",
    "input": "这是一段文字转语音的示例。",
    "voice": "alloy",
    "response_format": "mp3",
    "speed": 1.0
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "audio_url": "https://example.com/audio/xxx.mp3",
        "duration": 5.5
    }
}
```

#### 实时语音对话

```http
WebSocket /api/v1/audio/chat
```

连接参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| token | string | 是 | 访问令牌 |
| model | string | 是 | 模型ID |
| mode | string | 是 | 交互模式(turn/continuous) |
| input_format | string | 是 | 输入格式(wav/mp3/opus) |
| output_format | string | 是 | 输出格式(wav/mp3/opus) |
| language | string | 否 | 语言代码(默认: auto) |

请求消息：

```json
{
    "type": "audio",
    "data": "base64_encoded_audio_data",
    "sequence": 1,
    "final": true
}
```

响应消息：

```json
{
    "type": "transcript",
    "data": "你好，请问有什么可以帮你？",
    "sequence": 1
}
```

#### 视频生成

```http
POST /api/v1/videos/generations
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| model | string | 是 | 模型ID |
| prompt | string | 是 | 视频描述 |
| duration | number | 否 | 视频时长(秒) |
| resolution | string | 否 | 视频分辨率(1080p/720p/480p) |
| fps | integer | 否 | 帧率(默认: 30) |
| style | string | 否 | 视频风格 |
| audio | boolean | 否 | 是否生成音频 |

请求示例：

```json
{
    "model": "text-to-video-1",
    "prompt": "一只猫在草地上奔跑",
    "duration": 5,
    "resolution": "1080p",
    "fps": 30,
    "audio": true
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "task_id": "task_123456",
        "status": "processing",
        "progress": 0,
        "estimated_time": 300,
        "created_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 视频编辑

```http
POST /api/v1/videos/edits
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: multipart/form-data
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| model | string | 是 | 模型ID |
| video | file | 是 | 原始视频 |
| prompt | string | 是 | 编辑描述 |
| mask_video | file | 否 | 蒙版视频 |
| start_time | number | 否 | 起始时间(秒) |
| end_time | number | 否 | 结束时间(秒) |
| output_format | string | 否 | 输出格式(mp4/mov/webm) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "task_id": "task_123456",
        "status": "processing",
        "progress": 0,
        "estimated_time": 600,
        "created_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 实时视频对话

```http
WebSocket /api/v1/videos/chat
```

连接参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| token | string | 是 | 访问令牌 |
| model | string | 是 | 模型ID |
| video_format | string | 是 | 视频格式(h264/vp8/vp9) |
| resolution | string | 否 | 视频分辨率 |
| bitrate | integer | 否 | 视频比特率 |
| audio_enabled | boolean | 否 | 是否包含音频 |

请求消息：

```json
{
    "type": "video_frame",
    "data": "base64_encoded_video_frame",
    "timestamp": 1679289600000,
    "sequence": 1
}
```

响应消息：

```json
{
    "type": "analysis",
    "data": {
        "objects": ["cat", "grass"],
        "actions": ["running"],
        "emotions": ["happy"]
    },
    "timestamp": 1679289600000,
    "sequence": 1
}
```

#### 多模态理解

```http
POST /api/v1/multimodal/understanding
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: multipart/form-data
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| model | string | 是 | 模型ID |
| messages | array | 是 | 消息列表 |
| temperature | number | 否 | 采样温度(0-2，默认: 1) |
| max_tokens | integer | 否 | 最大生成长度 |

请求示例：

```json
{
    "model": "gpt-4-vision",
    "messages": [
        {
            "role": "user",
            "content": [
                {
                    "type": "text",
                    "text": "这张图片里有什么？"
                },
                {
                    "type": "image",
                    "image_url": "https://example.com/image.jpg"
                }
            ]
        }
    ],
    "temperature": 0.7
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "id": "mm_123456",
        "object": "multimodal.understanding",
        "created": 1679289600,
        "model": "gpt-4-vision",
        "choices": [
            {
                "message": {
                    "role": "assistant",
                    "content": "图片中显示一只橘色的猫咪正在草地上奔跑，阳光明媚，背景是绿色的草地和蓝色的天空。"
                },
                "finish_reason": "stop"
            }
        ],
        "usage": {
            "prompt_tokens": 50,
            "completion_tokens": 45,
            "total_tokens": 95
        }
    }
}
```

#### 多模态生成

```http
POST /api/v1/multimodal/generation
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: multipart/form-data
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| model | string | 是 | 模型ID |
| prompt | string | 是 | 生成提示 |
| modalities | array | 是 | 生成类型列表(text/image/audio/video) |
| settings | object | 否 | 各模态的具体参数 |

请求示例：

```json
{
    "model": "multimodal-1",
    "prompt": "生成一个关于猫咪的短视频，配上可爱的背景音乐",
    "modalities": ["video", "audio"],
    "settings": {
        "video": {
            "duration": 10,
            "resolution": "1080p"
        },
        "audio": {
            "style": "cheerful",
            "format": "mp3"
        }
    }
}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "task_id": "mm_task_123456",
        "status": "processing",
        "outputs": {
            "video": {
                "url": "https://example.com/video.mp4",
                "duration": 10,
                "resolution": "1080p"
            },
            "audio": {
                "url": "https://example.com/audio.mp3",
                "duration": 10,
                "format": "mp3"
            }
        },
        "created_at": "2024-03-20T12:00:00Z"
    }
}
```

## 模型训练模块 (Training)

### 模块概述

模型训练模块提供模型的微调和定制化训练功能。基于trainings表设计，支持用户使用自己的数据对基础模型进行训练，实现个性化的AI模型。

### 功能特点

- 支持多种训练方式
- 灵活的数据集管理
- 训练过程监控
- 模型性能评估
- 训练结果验证
- 模型版本管理

### 接口列表

#### 数据集管理接口

- [上传训练数据](#上传训练数据)
- [获取数据集列表](#获取数据集列表)
- [获取数据集详情](#获取数据集详情)
- [删除训练数据](#删除训练数据)

#### 训练任务接口

- [创建训练任务](#创建训练任务)
- [获取训练状态](#获取训练状态)
- [取消训练任务](#取消训练任务)
- [获取训练日志](#获取训练日志)

#### 模型管理接口

- [获取训练模型](#获取训练模型)
- [部署训练模型](#部署训练模型)
- [删除训练模型](#删除训练模型)

### 接口详情

#### 上传训练数据

```http
POST /api/v1/training/datasets/upload
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: multipart/form-data
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| file | file | 是 | 训练数据文件(JSONL格式) |
| name | string | 是 | 数据集名称 |
| description | string | 否 | 数据集描述 |
| type | string | 是 | 数据类型(text/image/audio/video) |
| format | string | 否 | 数据格式说明 |
| tags | array | 否 | 数据集标签 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "dataset_id": "ds_123456",
        "name": "对话数据集2024",
        "file_size": 1024000,
        "total_samples": 1000,
        "status": "processing",
        "created_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 获取数据集列表

```http
GET /api/v1/training/datasets
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| page | integer | 否 | 页码(默认: 1) |
| page_size | integer | 否 | 每页数量(默认: 20) |
| type | string | 否 | 数据类型筛选 |
| status | string | 否 | 状态筛选 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 100,
        "page": 1,
        "page_size": 20,
        "items": [
            {
                "dataset_id": "ds_123456",
                "name": "对话数据集2024",
                "description": "用于训练客服对话模型",
                "type": "text",
                "file_size": 1024000,
                "total_samples": 1000,
                "status": "ready",
                "created_at": "2024-03-20T12:00:00Z"
            }
        ]
    }
}
```

#### 获取数据集详情

```http
GET /api/v1/training/datasets/{dataset_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "dataset_id": "ds_123456",
        "name": "对话数据集2024",
        "description": "用于训练客服对话模型",
        "type": "text",
        "format": "jsonl",
        "file_size": 1024000,
        "total_samples": 1000,
        "valid_samples": 980,
        "invalid_samples": 20,
        "tags": ["customer_service", "chat"],
        "status": "ready",
        "validation_results": {
            "format_check": "passed",
            "content_check": "passed",
            "error_samples": []
        },
        "created_at": "2024-03-20T12:00:00Z",
        "updated_at": "2024-03-20T12:10:00Z"
    }
}
```

#### 删除训练数据

```http
DELETE /api/v1/training/datasets/{dataset_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "dataset_id": "ds_123456",
        "deleted_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 创建训练任务

```http
POST /api/v1/training/jobs
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| base_model | string | 是 | 基础模型ID |
| dataset_id | string | 是 | 数据集ID |
| training_type | string | 是 | 训练类型(fine-tune/custom) |
| hyperparameters | object | 否 | 训练超参数 |
| validation_split | number | 否 | 验证集比例(0-1) |
| max_epochs | integer | 否 | 最大训练轮数 |
| early_stopping | boolean | 否 | 是否启用早停 |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "job_id": "job_123456",
        "status": "queued",
        "estimated_time": 3600,
        "created_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 获取训练状态

```http
GET /api/v1/training/jobs/{job_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "job_id": "job_123456",
        "base_model": "gpt-3.5-turbo",
        "dataset_id": "ds_123456",
        "status": "training",
        "progress": 45.5,
        "current_epoch": 2,
        "total_epochs": 3,
        "metrics": {
            "loss": 0.123,
            "accuracy": 0.956,
            "validation_loss": 0.145
        },
        "estimated_remaining_time": 1800,
        "created_at": "2024-03-20T12:00:00Z",
        "updated_at": "2024-03-20T12:30:00Z"
    }
}
```

#### 取消训练任务

```http
POST /api/v1/training/jobs/{job_id}/cancel
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "job_id": "job_123456",
        "status": "cancelled",
        "cancelled_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 获取训练日志

```http
GET /api/v1/training/jobs/{job_id}/logs
```

请求头：

```http
Authorization: Bearer {access_token}
```

查询参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| start_time | string | 否 | 开始时间 |
| end_time | string | 否 | 结束时间 |
| level | string | 否 | 日志级别(info/warning/error) |
| limit | integer | 否 | 返回条数(默认: 100) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 1000,
        "items": [
                {
                    "timestamp": "2024-03-20T12:00:00Z",
                "level": "info",
                "message": "开始训练第1轮",
                "details": {
                    "epoch": 1,
                    "batch_size": 4,
                    "learning_rate": 1e-5
                }
            }
        ]
    }
}
```

#### 获取训练模型

```http
GET /api/v1/training/models/{model_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "model_id": "model_123456",
        "model_name": "gpt-3.5-turbo",
        "status": "ready",
        "created_at": "2024-03-20T12:00:00Z",
        "updated_at": "2024-03-20T12:30:00Z"
    }
}
```

#### 部署训练模型

```http
POST /api/v1/training/models/{model_id}/deploy
```

请求头：

```http
Authorization: Bearer {access_token}
Content-Type: application/json
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| channel_id | string | 是 | 渠道ID |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "model_id": "model_123456",
        "status": "deployed",
        "deployed_at": "2024-03-20T12:00:00Z"
    }
}
```

#### 删除训练模型

```http
DELETE /api/v1/training/models/{model_id}
```

请求头：

```http
Authorization: Bearer {access_token}
```

请求参数：

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|:----:|------|
| confirm | boolean | 是 | 确认删除(必须为true) |

响应结果：

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "model_id": "model_123456",
        "deleted_at": "2024-03-20T12:00:00Z"
    }
}
```
