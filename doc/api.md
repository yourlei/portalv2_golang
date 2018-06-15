# portalv2

## 文档大纲

根据系统的模块划分，该文档主要描述的内容有以下部分：

1 [用户管理](#1用户管理)

- 用户登录, 修改密码
- 用户增, 删, 改, 查
- 用户状态变更（审核, 启用, 禁用, 注销）

2 [验证码](#2验证码)

- 生成随机验证码
- 服务端校验提交的验证码

3 [角色管理](#3角色管理)

- 角色组增, 删, 改, 查
- 角色组下用户迁移

4 [权限管理](#4权限管理)

- 角色组权限增, 减

5 [资源管理](#5资源管理)

- 资源增, 删, 改, 查

6 [日志管理](#6日志管理)

- 日志增, 查
- 日志导出

7 [站点配置](#7站点配置)

- 自定义系统标题, logo


---

# 0 接口的相关规范

## 0.1 响应数据

- 返回的数据格式，每个接口固定返回*code*, *error*信息，当http请求资源成功时 *code*的返回的值为 *0*， 对应的非0则表示请求失败，具体的错误信息可查看*error*中的*msg*。 固定返回的字段如下：


``` json

code: 0,
error: {
  msg: '这里是错误信息，成功时为空'
}
```

- 错误代码说明

| code | 说明 |
|:----:|:----:|
|1 |请求参数错误|
|404 |资源不存在|
|500|服务出错|
|1000xx | 用户管理相关的错误码  |
|2000xx | 角色管理相关的错误码  |
|3000xx | 资源管理相关的错误码  |
|4000xx | 权限相关的错误码  |
|5000xx | 数据集管理相关的错误码  |

## 0.2 状态码

- 200: GET请求成功, 及DELETE或PATCH同步请求完成，或者PUT同步更新一个已存在的资源
- 201: POST 同步请求完成，或者PUT同步创建一个新的资源


- 401 Unauthorized: 用户未认证，请求失败
- 403 Forbidden: 用户无权限访问该资源，请求失败
- 404 Not found: 资源不存在

- 500 Internal Server Error: 服务器错误，确认状态并报告问题

## 0.3 接口权限说明

- token: 表示调用该接口须携带登录后服务端发放的token
- token + admin: 表示调用该接口须携带token, 同时用户所在角色组为管理员

---

# 1用户管理

#### 1.1 登录

- method: post
- url: domain/api/v1/admin/signin | domain/api/v1/users/signin

- 参数 

``` json
{
	"email": "admin@ibbd.net",
	"password": "scut2017",
	"uuid": "U2FsdGVkX19oujDSO4IPg7s9P6TZibiIiDG56WJ01Vw=", # 验证码标识
	"code": "TtfHg" # 前端输入的验证码
}
```

- 请求参数

| 参数 |类型  |说明  |范围及格式|
|:----:|:----:|:----:|:--------:|
|email |string|邮箱  |          |
|password | string   | 密码|    |
|uuid|string|验证码标识||
|code|string|验证码||

返回结果:

``` json
{
  code: 0,  # 成功是为0
  error: {
    msg: ''    # 错误信息 成功为空
  },
  data: {
    id: '',
    role_id: '',
    token: '',
    name: '',
    mobile: '',
    email: '',
    status: ''
  }
}
```

- *code*值有:

| code | 说明  |
|:----:|:----:|
|10001 |账户不存在|
|10002 | 密码不正确   |
|10003 | 账户未审核|
|10004| 审核不通过 |
|10005|账户已禁用|
|10006|账户已注销|
|10007|验证码不正确或已失效|
|10008|验证码失效|
|100012|没有登录权限,请使用管理员账户登录|          |

#### 1.1 注册(新增用户)

- method: post
- url: domain/api/v1/signup(doamin/api/v1/admin/users/new)

- 参数

``` json
{
  "name": "Janni",
  "mobile": "18888888885",
  "email": "fanni@gmail.com",
  "password": "a123456",
  "roleId": 1
}
```

- 返回值:

``` js
{
  code: 0, 
  error: {
    msg: ''    # 错误信息
  },
  data: {
    name: '',
    email: '',
    mobile: '',
    status: 1, 
  }
}
```

- **code说明**

| code         |  msg         |   说明   |
|:------------:|:------------:|:--------:|
| 10009        | 用户名已注册 |          |  
| 100010       | 邮箱/手机号已注册   |          | 


~~#### 1.2 注销用户~~(该接口功能已移至**1.3 禁用,启用,注销用户**)

- 请求方式: delete
- URL: domain/api/v1/users/:id

- 参数

``` json
{
  "remark": "注销leimi"
}
```

- 返回值

``` json

{
    "code": 0,
    "error": {
        "msg": ""
    }
}
```

#### 1.3 禁用,启用,注销用户

- 请求方式: patch
- URL: domain/users/status/:id
- 权限: token + admin

- 参数

``` json
{
	"status": 1,
	"remark": "启用leimi"
}
```

- 参数说明

| 参数 |类型  |说明  |范围及格式|
|:----:|:----:|:----:|:--------:|
|status |number|用户状态值  | 1 启用,-1 禁用, 0 注销|
|remark | string| 描述|    |

- 返回值

``` json

{
    "code": 0,
    "error": {
        "msg": ""
    }
}
```

#### 1.4 审核用户

- 请求方式: patch
- URL: domain/users/check/:id
- 权限: token + admin

- 参数

``` json
{
  "check_status": 1,
  "check_remark": "审核leimi通过"
}
```

- 参数说明

| 参数 |类型  |说明  |范围及格式|
|:----:|:----:|:----:|:--------:|
|check_status |number|审核状态值  |1 审核通过,-1 审核不通过, 0 未审核|
|check_remark | string| 描述|    |

- 返回值

``` json
{
    "code": 0,
    "error": {
        "msg": ""
    }
}
```

#### 1.5 编辑用户信息

- 请求方式: patch
- URL: domain/users/edit/:id
- 权限: token + admin

- 参数

``` json
{
  "name": "",
  "password": "",
  "mobile": ""
}
```

- 参数说明

| name         |  说明         |
|:------------:|:------------:|
| name        | 用户名 |
| password        | 密码 |
| mobile        | 手机号 |

- 返回值

``` json
{
    "code": 0,
    "error": {
      "msg": ""
    },
    "data": {
      "name": "不是",
      "email": "park@qq.com",
      "mobile": "18988888884",
      "status": 1,
      "check_status": 0
    }
}
```

#### 1.6 修改用户密码

- 请求方式: post
- URL: domain/api/v1/admin/passwd/change
- 权限: token 

- 请求参数:

``` js
{
  'passwd': '',      // 原密码
  'new_passwd': ''  // 新密码
}
```

- 返回结果

``` js
{

  code: 0,
  error: {
    msg: ''
  }
}
```


#### 1.7 获取用户列表

- 请求方式: get
- URL: domain/api/v1/users?query={body}
- 权限: token + admin

- body参数:

``` json
{
  "offset": 0, # 页码
  "limit": 10, # 页长
  "where": {
    "status": "0",
    "check_status": "0",
    "email": "admin@ibbd.net",
  }
}
```

- **where参数说明**

| 参数 |类型  |说明  |范围及格式|
|:----:|:----:|:----:|:--------:|
|email |string|邮箱  |          |
|status | string   | 用户状态|1:启用, -1: 禁用, 0:注销|
|check_status|string|账户审核状态||


- 返回结果

``` json
{
    "code": 0,
    "error": {
        "msg": null
    },
    "data": [
        {
            "id": "36",
            "name": "Janni",
            "role": 2,
            "mobile": "18888888887",
            "email": "Janni@gmail.com",
            "status": 1,
            "check_status": 0
        }
    ],
    "total": 1
}
```

---

# 2验证码

#### 2.0 图片验证码

- method: get
- url: domain/api/v1/image/base64

- 返回值

``` json
{
  "code": 0,
  "error": {
    "msg": ""
  },
  "data": {
    "img": "data:image/bmp;base64,...", # 图片的base64编码
    "uuid": "fc09ef70-058c-11e8-a28e-1fc716bf09a5" # 验证码id标识
  }
}
```

**接口更新**

- 考虑到安全性, 原来的客户端验证修改为服务端验证方式,该接口返回的**uuid**作为验证码标识,提交时带上用户输入的验证码及uuid,具体看**0.用户登录**
- 验证码有效时间为60s,过期后需重新获取
- 验证码不区分大小写

---

# 3角色管理

#### 3.1 角色列表

- 请求方式: get
- URL: domain/api/v1/roles?query={body}
- 权限: token + admin

- body参数

``` json
{
  "offset": 0,
  "limit": 8,
  "where": {
    "name": "",
    "created_at": {
      "$gt": "2006-01-02T15:04:05Z",
      "$lt": "2018-01-28T23:59:59Z"
    },
    "updated_at": {
      "$gt": "2006-01-02T15:04:05Z",
      "$lt": "2018-01-28T23:59:59Z"
    }
	}
}
```

- 返回结果

``` json
{
    "code": 0,
    "errors": {
        "msg": ""
    },
    "total": 3,
    "data": [
        {
            "id": 1,
            "name": "超级管理员组",
            "remark": "超级管理员",
            "status": 1,
            "created_at": "2017-10-11T17:11:54.000Z",
            "updated_at": "2017-10-11T17:11:54.000Z",
            "deleted_at": "2000-01-01T00:00:00.000Z"
        }
    ]
}
```

#### 3.2  创建角色

- 请求方式: post
- URL: domain/api/v1/roles
- 权限: token + admin

- 参数

``` json
{
	"name": "开发组",
	"remark": "开发小组"
}
```

- 错误代码

| code         |  msg         |   说明   |
|:------------:|:------------:|:--------:|
| 20001        | 角色组已存在 | 409冲突     |


- 返回值

``` json
{
    "code": 0,
    "error": {
        "msg": ""
    }
}
```

#### 3.3  删除角色

- 请求方式: delete
- URL: domain/api/v1/roles/:id
- 权限: token + admin

- 返回值

``` json

{
    "code": 0,
    "error": {
        "msg": ""
    }
}
```

#### 3.4  编辑角色

- 请求方式: patch
- URL: domain/api/v1/roles/:id
- 权限: token + admin

- 参数

``` json
{
	"name": "测试组",
	"remark": "测试小组"
}
```

- 返回值

``` json

{
    "code": 0,
    "error": {
        "msg": ""
    }
}
```

### 3.4  获取角色组下用户

- 请求方式: get
- URL: domain/api/v1/roles/users/:id

- 返回值

``` json
{
    "code": 0,
    "error": {
        "msg": ""
    },
    "total": 1,
    "data": [
        {
            "id": 4,
            "name": "雷米"
        }
    ]
}
```

### 3.5 转移角色组下用户

- 请求方式: patch
- URL: domain/api/v1/roles/users

- 参数

``` json
{
  "roleId": 3,
  "userId": [5,7]
}
```

- 返回值

``` json
{
    "code": 0,
    "error": {
        "msg": ""
    }
}
```

---

# 4权限管理

####  4.1 资源列表

- method: get
- url: domain/api/v1/resources

- 返回结果

``` json
{
    "code": 0,
    "error": {
        "msg": ""
    },
    "data": {
        "menus": [
            [
                {
                  "id": 4, # 父菜单
                  "name": "国际"
                },
                {
                  "id": 6,
                  "name": "新闻搜索"
                }
            ],
            [
                {
                  "id": 7, # 父菜单
                  "name": "人工智能"
                },
                {
                  "id": 8,
                  "name": "机器学习"
                }
            ]
        ],,
        "interfaces": {
            "users": [
              "登录",
              "注册"
            ]
        }
    }
}
```

- 结果说明:

**menus以父菜单进行分组, 第一个元素为父菜单,其余元素则为该父菜单下的子菜单**

#### 4.2 获取角色组下资源

- method: get
- url: domain/api/v1/roles/resource/:id

**id: 角色id**

``` json
{
    "code": 0,
    "error": {
        "msg": ""
    },
    "data": {
        "menus": [
            {
                "id": 4,
                "name": "国际"
            },
            {
                "id": 6,
                "name": "新闻搜索"
            }
        ],
        "interfaces": {}
    }
}
```

#### 4.3 修改权限

- method: post
- url: domain/api/v1/roles/role2resource

``` json
{
	"id": 1,
	"menus": {
		"associate": [4, 5], # 分配权限的菜单id
		"dissociate": [1, 2] # 移除权限的菜单id
	},
	"interfaces": {
		"associate": ["登录"], # 分配权限的接口名称
		"dissociate": ["删除用户"] # 移除权限的接口名称
	}
}
```

#### 4.4 获取角色组下资源(用户端)

- method: get
- url: domain/api/v1/users/resource/:roleId

- 返回结果

``` json
{
    "code": 0,
    "error": {
      "msg": ""
    },
    "data": {
        "menus": [
          {
                "id": 1,
                "name": "管理数据",
                "action": 0,
                "priority": 8,
                "parent": 0,
                "schema": "{\"params\":\"\",\"form\":[],\"iframe\":[]}",
                "item": "data",
                "child": [
                    {
                        "id": 2,
                        "name": "订单数据",
                        "action": 1,
                        "priority": 8,
                        "parent": 1,
                        "schema": "{\"params\":\"http://baidu.com\",\"form\":[],\"iframe\":[]}",
                        "item": "order"
                    },
                    {
                        "id": 3,
                        "name": "客流数据",
                        "action": 1,
                        "priority": 8,
                        "parent": 1,
                        "schema": "{\"params\":\"http://baidu.com\",\"form\":[],\"iframe\":[]}",
                        "item": "customer"
                    }
                ]
            },
            {
                "id": 4,
                "name": "搜索",
                "action": 1,
                "priority": 3,
                "parent": 0,
                "schema": "{\"params\":\"http://baidu.com\",\"form\":[],\"iframe\":[]}",
                "item": "find"
            }
        ],
        "interfaces": [
          {
            "created_at": "2018-01-26 14:25:08",
            "updated_at": "2018-01-26 14:25:08",
            "id": 2,
            "name": "登录",
            "url": "api/v1/users/signin"
          }
        ]
    }
}
```

---

# 5资源管理

### 5.1 菜单资源

#### 5.1.1 创建菜单

- method: post
- url: domain/api/v1/resource/menus

- 参数

``` json
{
  "name": "数仓dfgfd",
	"item": "data",
	"parent": 1,
	"action": 1,
  "priority": 3,
  "schema": {
    "params": "http://baidu.com",
    "formComponet": [
      {
        "name": "time", # 表单域名称
        "label": "时段", # 表单域标签
        "input": "select", # date | select
        "values": ["深圳", "广州", "香港"], # input 为date时, values 为空[]
        "size": "large",
        "rules": [] # 对应前端的表单域验证规则
      }
    ],
    "iframePage":[
      {
        "url":{},
        "width": "",
        "height": ""
      }
    ]
  }
  "remark": ""
}
```

| 字段 | 类型 | 说明 |是否必填 |
|:----:|:---:|:---:|:--- |
| name | string | 菜单名称 | Y |
| item | string | 菜单项 (一级菜单必填)| N |
| parent | nubmer | 父级菜单ID, 值为-1表示该菜单为一级菜单 | Y |
| action | number | 动作类型(1: 显示子菜单, 2: 打开iframe页面, 3: route, 4: 打开search页面) | Y |
| schema | string | 菜单对应页面的配置 | Y |
| priority | number | 权重 | Y |

- 返回值

``` json
{
  code: 0,
  error: {
    msg: ""
  }
}
```

- 错误码说明

| code |说明 |
|:----:|:---:|
| 30001 | 资源冲突, 菜单名已占用, **注意：一级菜单及相同父菜单下name, item 须唯一** |

#### 5.1.2 菜单列表

- method: get
- url: domain/api/v1/resource/menus?query={body}

- 参数

``` json
{
  "offset":: 0,
  "limit": 10,
  "where": {
    "name": "",
    "action": 0,
    "created_at": {
      $gt: '2017-02-02',
      $lt: '2017-02-10'
    },
    "updated_at": {
      $gt: '',
      $lt: ''
    }
  }
}
```

- 参数说明

| 字段 | 类型 | 说明 |是否必填 |
|:----:|:---:|:---:|:--- |
| offset | number | 页码 | N |
| limit | number | 页宽 | N |
| name | string | 菜单名 | N |
| action | number | 动作类型 | N |
| created_at | object | 创建时间 | N |

- 返回值

``` json
{
  "code": 0,
  "error": {
    "msg": ""
  },
  "data": [
    {
            "created_at": "2018-02-26 13:54:42",
            "updated_at": "2018-02-26 13:54:42",
            "id": 4,
            "name": "国际",
            "item": "FinanceNew",
            "params": "",
            "parent": "",
            "action": 0,
            "priority": 7,
            "group": "游客组,研发"
        },
  ]
}
```

#### 5.1.3 编辑菜单

- method: put
- url: domain/api/v1/resource/menus/:id

- 参数

``` json
{
  "name": "",
  "level": 0,
  "params": "",
  "remark": "",
  "priority": 9
}
```

- 参数说明

| 字段 | 类型 | 说明 |是否必填 |
|:----:|:---:|:---:|:--- |
| name | string |  | N |
| level | number |  | Y |
| params | string |  | N |
| remark | string |  | N |

- 返回值

``` json
{
  "code": 0,
  "error": {
    "msg": ""
  }
}
```

#### 5.1.4 删除菜单

- method: delete
- url: domain/api/v1/resource/menus/:id

- 返回值

``` json
{
  "code": 0,
  "error": {
    "msg": ""
  }
}
```

#### 5.1.5 获取父级菜单

- method: get
- url: domain/resources/menu/parent

- 返回值

``` json
{
  "code": 0,
  "error": {
    "msg": ""
  },
   "data": [
        {
            "id": 1,
            "name": "管理数据"
        },
        {
            "id": 5,
            "name": "搜索-A"
        }
    ]
}
```

### 5.2 接口资源

#### 5.2.1 创建

- method: post
- url: domain/api/v1/resource/interfaces

- 参数

``` json
{
  "name": "注册f",
	"url": "/api/v1/users/signup",
	"method": "post",
	"module": "user",
  "remark": ""
}
```

| 字段 | 类型 | 说明 |是否必填 |
|:----:|:---:|:---:|:--- |
| name | string | 接口名称 | Y |
| url | string | 接口地址 | Y |
| method | string | 方法 | Y |
| module | string | 模块 | Y |
| remark | string | 描述 | N |

- 返回值

``` json
{
  code: 0,
  error: {
    msg: ""
  }
}
```

#### 5.2.2 接口列表

- method: get
- url: domain/api/v1/resource/interfaces?query={body}

- 参数

``` json
{
  "offset":: 0,
  "limit": 10,
  "where": {
    "name": "",
    "module": "",
    "created_at": {
      $gt: '2017-02-02',
      $lt: '2017-02-10'
    },
    "updated_at": {
      $gt: '',
      $lt: ''
    }
  },
  "fields": []
}
```

- 参数说明

| 字段 | 类型 | 说明 |是否必填 |
|:----:|:---:|:---:|:--- |
| offset | number | 页码 | N |
| limit | number | 页宽 | N |
| name | string | 接口名 | N |
| module | string | 模块 | N |
| created_at | object | 创建时间 | N |

- 返回值

``` json
{
  "code": 0,
  "error": {
    "msg": ""
  },
  "data": []
}
```

#### 5.2.3 编辑接口

- method: put
- url: domain/api/v1/resource/interfaces/:id

- 参数

``` json
{
  "name": "",
  "url": "",
  "remark": ""
}
```

- 参数说明

| 字段 | 类型 | 说明 |是否必填 |
|:----:|:---:|:---:|:--- |
| name | string |  | N |
| url | string |  | N |
| remark | string |  | N |

- 返回值

``` json
{
  "code": 0,
  "error": {
    "msg": ""
  }
}
```

#### 5.2.4 删除接口

- method: delete
- url: domain/api/v1/resource/interfaces/:id

- 返回值

``` json
{
  "code": 0,
  "error": {
    "msg": ""
  }
}
```

---

# 6日志管理

#### 6.1 日志列表

- 请求方式: get
- URL: domain/api/v1/logs?query=body

- body参数

``` json
{
  "offset": 0,
  "limit": 8,
  "where": {
    "account": "",   # 用户名
    "account_id": 1,　# 用户id
    "action": "",
    "ip": ,
    "created_at": {
      $gt: '2017-02-02',
      $lt: '2017-02-10'
    }
   }
}
```

- 返回值

``` json
{
    "code": 0,
    "error": {
        "msg": ""
    },
    "total": 821,
    "data": [
        {
            "id": 1,
            "account_id": 1,
            "account": "sysu@portal.com",
            "action": "更新密码",
            "method": "POST",
            "url": "/api/v1/user/passwd",
            "ip": "127.0.0.1",
            "created_at": "2017-10-11T22:56:27.000Z",
            "updated_at": "2017-10-11T22:56:27.000Z"
        }]
```

### 6.2 日志导出

- method: get
- url: domain/api/v1/logs/download?query={body}


- body 

``` json

{
  "where":{
    "account": "",   # 用户名
    "action": "",
    "ip": ,
    "created_at": {
      $gt: '2017-02-02',
      $lt: '2017-02-10'
    }
  }
}
```

---

# 7站点配置

#### 7.1 编辑配置

- method: put
- url: domain/api/v1/setting/:id

**id 值为1**

- 参数

``` json
{
  "user_title": "用户端标题",
  "admin_title": "管理员端标题"
}
```

- 返回结果

``` json
{
  "code": 0,
  "error": {
    "msg": ""
  },
  "data": {
    "id": 1,
    "user_title": "用户端标题",
    "admin_title": "管理员端标题"
  }
}
```

#### 7.2 获取配置信息

- method: get
- url: domain/api/v1/setting/:id

**id 值为1,仅保存一条配置信息**

- 返回结果

``` json
{
  "code": 0,
  "error": {
    "msg": ""
  },
  "data": {
    "id": 1,
    "user_title": "用户端标题",
    "admin_title": "管理员端标题"
  }
}
```

#### 7.3 上传logo

- method: post
- url: domain/api/v1/upload/logo/:name

**name: user(用户端logo/favicon) | admin(管理员端logo/favicon)**

- 参数

**form**

| 字段 | 类型 |
|:-----:|:----:|
| input | file |

**文件格式限定为.png**

**图片地址**:

* logo: domain/public/logo/{user|admin}/logo.png
* favicon: domain/public/logo/{user|admin}/favicon.ico

---

### 8　es查询(用户端)

- method: get
- url: domain/api/v1/es/search?query={body}

- body查询体

``` json
{
  "offset":0,
  "limit": 10,
  "where": {
    "index":"qualitrip", 　 # 数据集
    "key":"天津",  　　　　　 # 搜索参数
    "name": "菜单名"
  },
  "order": ["sum", "desc"] # ["排序字段", "排序规则(asc(升序), desc(降序)"]
}
```

- 返回结果

``` json
{
    "code": 0,
    "error": {
        "msg": ""
    },
    "total": 2351,
    "columns": [
        {
            "name": "sail_time", # 列名
            "sorting": true  # 是否可排序
        },
        {
            "name": "city",
            "sorting": false
        },
     ],
    "data": [
        {
            "ticket_id": "113916317954",
            "order_status": "3",
            "trace_code": "PY",
            "sail_time": "2000-01-01T00:00:00.000+0800",
            "class_of_service": "121",
            "departure": "HKA",
            "destination": "PY-HKA",
            "customer_profile": "1",
            "price": "300",
            "currency": "RMB"
        }
    ]
}
```

**结果说明: 前端根据columns展示表头**
