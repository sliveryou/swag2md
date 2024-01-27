# 示例接口文档

## 接口概览（总计15个）

### example/department

| **路径**                          | **功能**     | **请求方式** | **是否需要鉴权** |
|-----------------------------------|--------------|--------------|------------------|
| [/api/department](#查询部门分页)  | 查询部门分页 | GET          | false            |
| [/api/department](#创建部门)      | 创建部门     | POST         | false            |
| [/api/department/{id}](#查询部门) | 查询部门     | GET          | false            |
| [/api/department/{id}](#更新部门) | 更新部门     | PUT          | false            |
| [/api/department/{id}](#删除部门) | 删除部门     | DELETE       | false            |

### example/job

| **路径**                   | **功能**     | **请求方式** | **是否需要鉴权** |
|----------------------------|--------------|--------------|------------------|
| [/api/job](#查询岗位分页)  | 查询岗位分页 | GET          | false            |
| [/api/job](#创建岗位)      | 创建岗位     | POST         | false            |
| [/api/job/{id}](#查询岗位) | 查询岗位     | GET          | false            |
| [/api/job/{id}](#更新岗位) | 更新岗位     | PUT          | false            |
| [/api/job/{id}](#删除岗位) | 删除岗位     | DELETE       | false            |

### example/personnel

| **路径**                         | **功能**     | **请求方式** | **是否需要鉴权** |
|----------------------------------|--------------|--------------|------------------|
| [/api/personnel](#查询人员分页)  | 查询人员分页 | GET          | false            |
| [/api/personnel](#创建人员)      | 创建人员     | POST         | false            |
| [/api/personnel/{id}](#查询人员) | 查询人员     | GET          | false            |
| [/api/personnel/{id}](#更新人员) | 更新人员     | PUT          | false            |
| [/api/personnel/{id}](#删除人员) | 删除人员     | DELETE       | false            |

## 接口详情

### example/department

### 查询部门分页

[返回概览](#exampledepartment)

---

GET /api/department

请求参数：

| **来源** | **参数**    | **描述** | **类型** | **约束** | **说明** |
|----------|-------------|----------|----------|----------|----------|
| query    | name        | 部门名称 | string   | 非必填   |          |
| query    | description | 部门描述 | string   | 非必填   |          |
| query    | page        | 页数     | integer  | 必填     |          |
| query    | page_size   | 每条页数 | integer  | 必填     |          |

请求示例：

```
Query:
/api/department?description=description&name=name&page=1&page_size=1
```

---

Content-Type: application/json

响应参数：

| **参数**           | **描述** | **类型**     | **说明** |
|--------------------|----------|--------------|----------|
| count              | 总数     | integer      |          |
| page_count         | 页数     | integer      |          |
| results            | 结果     | object array |          |
| &emsp; description | 部门描述 | string       |          |
| &emsp; id          | 部门ID   | integer      |          |
| &emsp; name        | 部门名称 | string       |          |

响应示例：

```json
{
  "trace_id": "a1b2c3d4e5f6g7h8",
  "code": 0,
  "msg": "ok",
  "data": {
    "count": 1,
    "page_count": 1,
    "results": [
      {
        "description": "description",
        "id": 1,
        "name": "name"
      }
    ]
  }
}
```

### 创建部门

[返回概览](#exampledepartment)

---

POST /api/department  
Content-Type: application/json

请求参数：

| **来源** | **参数**    | **描述** | **类型** | **约束** | **说明** |
|----------|-------------|----------|----------|----------|----------|
| body     | description | 部门描述 | string   | 非必填   |          |
| body     | name        | 部门名称 | string   | 必填     |          |

请求示例：

```json
{
  "description": "description",
  "name": "name"
}
```

---

Content-Type: application/json

响应参数：

| **参数**    | **描述** | **类型** | **说明** |
|-------------|----------|----------|----------|
| description | 部门描述 | string   |          |
| id          | 部门ID   | integer  |          |
| name        | 部门名称 | string   |          |

响应示例：

```json
{
  "trace_id": "a1b2c3d4e5f6g7h8",
  "code": 0,
  "msg": "ok",
  "data": {
    "description": "description",
    "id": 1,
    "name": "name"
  }
}
```

### 查询部门

[返回概览](#exampledepartment)

---

GET /api/department/{id}

请求参数：

| **来源** | **参数** | **描述** | **类型** | **约束** | **说明** |
|----------|----------|----------|----------|----------|----------|
| path     | id       | 部门ID   | integer  | 必填     |          |

---

Content-Type: application/json

响应参数：

| **参数**    | **描述** | **类型** | **说明** |
|-------------|----------|----------|----------|
| description | 部门描述 | string   |          |
| id          | 部门ID   | integer  |          |
| name        | 部门名称 | string   |          |

响应示例：

```json
{
  "trace_id": "a1b2c3d4e5f6g7h8",
  "code": 0,
  "msg": "ok",
  "data": {
    "description": "description",
    "id": 1,
    "name": "name"
  }
}
```

### 更新部门

[返回概览](#exampledepartment)

---

PUT /api/department/{id}  
Content-Type: application/json

请求参数：

| **来源** | **参数**    | **描述** | **类型** | **约束** | **说明** |
|----------|-------------|----------|----------|----------|----------|
| path     | id          | 部门ID   | integer  | 必填     |          |
| body     | description | 部门描述 | string   | 非必填   |          |
| body     | name        | 部门名称 | string   | 必填     |          |

请求示例：

```json
{
  "description": "description",
  "name": "name"
}
```

---

Content-Type: application/json

响应参数：

| **参数**    | **描述** | **类型** | **说明** |
|-------------|----------|----------|----------|
| description | 部门描述 | string   |          |
| id          | 部门ID   | integer  |          |
| name        | 部门名称 | string   |          |

响应示例：

```json
{
  "trace_id": "a1b2c3d4e5f6g7h8",
  "code": 0,
  "msg": "ok",
  "data": {
    "description": "description",
    "id": 1,
    "name": "name"
  }
}
```

### 删除部门

[返回概览](#exampledepartment)

---

DELETE /api/department/{id}  
Content-Type: application/json

请求参数：

| **来源** | **参数** | **描述** | **类型** | **约束** | **说明** |
|----------|----------|----------|----------|----------|----------|
| path     | id       | 部门ID   | integer  | 必填     |          |

---

Content-Type: application/json

响应参数：

| **参数** | **描述** | **类型** | **说明** |
|----------|----------|----------|----------|
| id       | 部门ID   | integer  |          |

响应示例：

```json
{
  "trace_id": "a1b2c3d4e5f6g7h8",
  "code": 0,
  "msg": "ok",
  "data": {
    "id": 1
  }
}
```

### example/job

### 查询岗位分页

[返回概览](#examplejob)

---

GET /api/job

请求参数：

| **来源** | **参数**      | **描述** | **类型** | **约束** | **说明** |
|----------|---------------|----------|----------|----------|----------|
| query    | name          | 岗位名称 | string   | 非必填   |          |
| query    | department_id | 部门ID   | integer  | 非必填   |          |
| query    | page          | 页数     | integer  | 必填     |          |
| query    | page_size     | 每条页数 | integer  | 必填     |          |

请求示例：

```
Query:
/api/job?department_id=1&name=name&page=1&page_size=1
```

---

Content-Type: application/json

响应参数：

| **参数**             | **描述** | **类型**     | **说明** |
|----------------------|----------|--------------|----------|
| count                | 总数     | integer      |          |
| page_count           | 页数     | integer      |          |
| results              | 结果     | object array |          |
| &emsp; department_id | 部门ID   | integer      |          |
| &emsp; id            | 岗位ID   | integer      |          |
| &emsp; name          | 岗位名称 | string       |          |

响应示例：

```json
{
  "trace_id": "a1b2c3d4e5f6g7h8",
  "code": 0,
  "msg": "ok",
  "data": {
    "count": 1,
    "page_count": 1,
    "results": [
      {
        "department_id": 1,
        "id": 1,
        "name": "name"
      }
    ]
  }
}
```

### 创建岗位

[返回概览](#examplejob)

---

POST /api/job  
Content-Type: application/json

请求参数：

| **来源** | **参数**      | **描述** | **类型** | **约束** | **说明** |
|----------|---------------|----------|----------|----------|----------|
| body     | department_id | 部门ID   | integer  | 必填     |          |
| body     | name          | 岗位名称 | string   | 必填     |          |

请求示例：

```json
{
  "department_id": 1,
  "name": "name"
}
```

---

Content-Type: application/json

响应参数：

| **参数**      | **描述** | **类型** | **说明** |
|---------------|----------|----------|----------|
| department_id | 部门ID   | integer  |          |
| id            | 岗位ID   | integer  |          |
| name          | 岗位名称 | string   |          |

响应示例：

```json
{
  "trace_id": "a1b2c3d4e5f6g7h8",
  "code": 0,
  "msg": "ok",
  "data": {
    "department_id": 1,
    "id": 1,
    "name": "name"
  }
}
```

### 查询岗位

[返回概览](#examplejob)

---

GET /api/job/{id}

请求参数：

| **来源** | **参数** | **描述** | **类型** | **约束** | **说明** |
|----------|----------|----------|----------|----------|----------|
| path     | id       | 岗位ID   | integer  | 必填     |          |

---

Content-Type: application/json

响应参数：

| **参数**      | **描述** | **类型** | **说明** |
|---------------|----------|----------|----------|
| department_id | 部门ID   | integer  |          |
| id            | 岗位ID   | integer  |          |
| name          | 岗位名称 | string   |          |

响应示例：

```json
{
  "trace_id": "a1b2c3d4e5f6g7h8",
  "code": 0,
  "msg": "ok",
  "data": {
    "department_id": 1,
    "id": 1,
    "name": "name"
  }
}
```

### 更新岗位

[返回概览](#examplejob)

---

PUT /api/job/{id}  
Content-Type: application/json

请求参数：

| **来源** | **参数**      | **描述** | **类型** | **约束** | **说明** |
|----------|---------------|----------|----------|----------|----------|
| path     | id            | 岗位ID   | integer  | 必填     |          |
| body     | department_id | 部门ID   | integer  | 必填     |          |
| body     | name          | 岗位名称 | string   | 必填     |          |

请求示例：

```json
{
  "department_id": 1,
  "name": "name"
}
```

---

Content-Type: application/json

响应参数：

| **参数**      | **描述** | **类型** | **说明** |
|---------------|----------|----------|----------|
| department_id | 部门ID   | integer  |          |
| id            | 岗位ID   | integer  |          |
| name          | 岗位名称 | string   |          |

响应示例：

```json
{
  "trace_id": "a1b2c3d4e5f6g7h8",
  "code": 0,
  "msg": "ok",
  "data": {
    "department_id": 1,
    "id": 1,
    "name": "name"
  }
}
```

### 删除岗位

[返回概览](#examplejob)

---

DELETE /api/job/{id}  
Content-Type: application/json

请求参数：

| **来源** | **参数** | **描述** | **类型** | **约束** | **说明** |
|----------|----------|----------|----------|----------|----------|
| path     | id       | 岗位ID   | integer  | 必填     |          |

---

Content-Type: application/json

响应参数：

| **参数** | **描述** | **类型** | **说明** |
|----------|----------|----------|----------|
| id       | 岗位ID   | integer  |          |

响应示例：

```json
{
  "trace_id": "a1b2c3d4e5f6g7h8",
  "code": 0,
  "msg": "ok",
  "data": {
    "id": 1
  }
}
```

### example/personnel

### 查询人员分页

[返回概览](#examplepersonnel)

---

GET /api/personnel

请求参数：

| **来源** | **参数**      | **描述** | **类型** | **约束** | **说明**          |
|----------|---------------|----------|----------|----------|-------------------|
| query    | name          | 人员名称 | string   | 非必填   |                   |
| query    | phone         | 联系电话 | string   | 非必填   |                   |
| query    | status        | 在岗情况 | integer  | 非必填   | 0-离岗 1-在岗     |
| query    | review_status | 审查情况 | integer  | 非必填   | 0-未审查 1-已审查 |
| query    | description   | 人员备注 | string   | 非必填   |                   |
| query    | department_id | 部门ID   | integer  | 非必填   |                   |
| query    | job_id        | 岗位ID   | integer  | 非必填   |                   |
| query    | page          | 页数     | integer  | 必填     |                   |
| query    | page_size     | 每条页数 | integer  | 必填     |                   |

请求示例：

```
Query:
/api/personnel?department_id=1&description=description&job_id=1&name=name&page=1&page_size=1&phone=phone&review_status=1&status=1
```

---

Content-Type: application/json

响应参数：

| **参数**             | **描述** | **类型**     | **说明**          |
|----------------------|----------|--------------|-------------------|
| count                | 总数     | integer      |                   |
| page_count           | 页数     | integer      |                   |
| results              | 结果     | object array |                   |
| &emsp; department_id | 部门ID   | integer      |                   |
| &emsp; description   | 人员备注 | string       |                   |
| &emsp; id            | 人员ID   | integer      |                   |
| &emsp; job_id        | 岗位ID   | integer      |                   |
| &emsp; name          | 人员名称 | string       |                   |
| &emsp; phone         | 联系电话 | string       |                   |
| &emsp; review_status | 审查情况 | integer      | 0-未审查 1-已审查 |
| &emsp; status        | 在岗情况 | integer      | 0-离岗 1-在岗     |

响应示例：

```json
{
  "trace_id": "a1b2c3d4e5f6g7h8",
  "code": 0,
  "msg": "ok",
  "data": {
    "count": 1,
    "page_count": 1,
    "results": [
      {
        "department_id": 1,
        "description": "description",
        "id": 1,
        "job_id": 1,
        "name": "name",
        "phone": "phone",
        "review_status": 1,
        "status": 1
      }
    ]
  }
}
```

### 创建人员

[返回概览](#examplepersonnel)

---

POST /api/personnel  
Content-Type: application/json

请求参数：

| **来源** | **参数**      | **描述** | **类型** | **约束** | **说明**          |
|----------|---------------|----------|----------|----------|-------------------|
| body     | department_id | 部门ID   | integer  | 必填     |                   |
| body     | description   | 人员备注 | string   | 非必填   |                   |
| body     | job_id        | 岗位ID   | integer  | 必填     |                   |
| body     | name          | 人员名称 | string   | 必填     |                   |
| body     | phone         | 联系电话 | string   | 非必填   |                   |
| body     | review_status | 审查情况 | integer  | 必填     | 0-未审查 1-已审查 |
| body     | status        | 在岗情况 | integer  | 必填     | 0-离岗 1-在岗     |

请求示例：

```json
{
  "department_id": 1,
  "description": "description",
  "job_id": 1,
  "name": "name",
  "phone": "phone",
  "review_status": 1,
  "status": 1
}
```

---

Content-Type: application/json

响应参数：

| **参数**      | **描述** | **类型** | **说明**          |
|---------------|----------|----------|-------------------|
| department_id | 部门ID   | integer  |                   |
| description   | 人员备注 | string   |                   |
| id            | 人员ID   | integer  |                   |
| job_id        | 岗位ID   | integer  |                   |
| name          | 人员名称 | string   |                   |
| phone         | 联系电话 | string   |                   |
| review_status | 审查情况 | integer  | 0-未审查 1-已审查 |
| status        | 在岗情况 | integer  | 0-离岗 1-在岗     |

响应示例：

```json
{
  "trace_id": "a1b2c3d4e5f6g7h8",
  "code": 0,
  "msg": "ok",
  "data": {
    "department_id": 1,
    "description": "description",
    "id": 1,
    "job_id": 1,
    "name": "name",
    "phone": "phone",
    "review_status": 1,
    "status": 1
  }
}
```

### 查询人员

[返回概览](#examplepersonnel)

---

GET /api/personnel/{id}

请求参数：

| **来源** | **参数** | **描述** | **类型** | **约束** | **说明** |
|----------|----------|----------|----------|----------|----------|
| path     | id       | 人员ID   | integer  | 必填     |          |

---

Content-Type: application/json

响应参数：

| **参数**      | **描述** | **类型** | **说明**          |
|---------------|----------|----------|-------------------|
| department_id | 部门ID   | integer  |                   |
| description   | 人员备注 | string   |                   |
| id            | 人员ID   | integer  |                   |
| job_id        | 岗位ID   | integer  |                   |
| name          | 人员名称 | string   |                   |
| phone         | 联系电话 | string   |                   |
| review_status | 审查情况 | integer  | 0-未审查 1-已审查 |
| status        | 在岗情况 | integer  | 0-离岗 1-在岗     |

响应示例：

```json
{
  "trace_id": "a1b2c3d4e5f6g7h8",
  "code": 0,
  "msg": "ok",
  "data": {
    "department_id": 1,
    "description": "description",
    "id": 1,
    "job_id": 1,
    "name": "name",
    "phone": "phone",
    "review_status": 1,
    "status": 1
  }
}
```

### 更新人员

[返回概览](#examplepersonnel)

---

PUT /api/personnel/{id}  
Content-Type: application/json

请求参数：

| **来源** | **参数**      | **描述** | **类型** | **约束** | **说明**          |
|----------|---------------|----------|----------|----------|-------------------|
| path     | id            | 人员ID   | integer  | 必填     |                   |
| body     | department_id | 部门ID   | integer  | 必填     |                   |
| body     | description   | 人员备注 | string   | 非必填   |                   |
| body     | job_id        | 岗位ID   | integer  | 必填     |                   |
| body     | name          | 人员名称 | string   | 必填     |                   |
| body     | phone         | 联系电话 | string   | 非必填   |                   |
| body     | review_status | 审查情况 | integer  | 必填     | 0-未审查 1-已审查 |
| body     | status        | 在岗情况 | integer  | 必填     | 0-离岗 1-在岗     |

请求示例：

```json
{
  "department_id": 1,
  "description": "description",
  "job_id": 1,
  "name": "name",
  "phone": "phone",
  "review_status": 1,
  "status": 1
}
```

---

Content-Type: application/json

响应参数：

| **参数**      | **描述** | **类型** | **说明**          |
|---------------|----------|----------|-------------------|
| department_id | 部门ID   | integer  |                   |
| description   | 人员备注 | string   |                   |
| id            | 人员ID   | integer  |                   |
| job_id        | 岗位ID   | integer  |                   |
| name          | 人员名称 | string   |                   |
| phone         | 联系电话 | string   |                   |
| review_status | 审查情况 | integer  | 0-未审查 1-已审查 |
| status        | 在岗情况 | integer  | 0-离岗 1-在岗     |

响应示例：

```json
{
  "trace_id": "a1b2c3d4e5f6g7h8",
  "code": 0,
  "msg": "ok",
  "data": {
    "department_id": 1,
    "description": "description",
    "id": 1,
    "job_id": 1,
    "name": "name",
    "phone": "phone",
    "review_status": 1,
    "status": 1
  }
}
```

### 删除人员

[返回概览](#examplepersonnel)

---

DELETE /api/personnel/{id}  
Content-Type: application/json

请求参数：

| **来源** | **参数** | **描述** | **类型** | **约束** | **说明** |
|----------|----------|----------|----------|----------|----------|
| path     | id       | 人员ID   | integer  | 必填     |          |

---

Content-Type: application/json

响应参数：

| **参数** | **描述** | **类型** | **说明** |
|----------|----------|----------|----------|
| id       | 人员ID   | integer  |          |

响应示例：

```json
{
  "trace_id": "a1b2c3d4e5f6g7h8",
  "code": 0,
  "msg": "ok",
  "data": {
    "id": 1
  }
}
```
