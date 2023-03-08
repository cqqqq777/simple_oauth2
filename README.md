# simple_oauth2
实现一个简单的oauth2

## api示例
### 1 authorization_code 
#### 1-1 获取授权code
##### 请求方式
```
GET /authorize
```
##### 参数说明
| 参数            | 类型     | 说明                                                       |
|---------------|--------|----------------------------------------------------------|
| client_id     | string | 在oauth2 server注册的client_id,见配置文件 oauth2.client.id        |
| response_type | string | 固定值：code                                                 |
| scope         | string | 权限范围,如:str1,str2,str3,str为配置文件中 oauth2.client.scope.id的值 |
| state         | string | 表示客户端的当前状态,可以指定任意值,认证服务器会原封不动地返回这个值                      |
| redirect_uri  | string | 回调uri,会在后面添加query参数?code=xxx&state=xxx,发放的code就在其中       |

##### 请求示例
```
# 浏览器请求
http://localhost:8080/authorize?client_id=test_client_1&response_type=code&scope=all&state=xyz&redirect_uri=http://localhost:9093/cb

# 302跳转,返回code
http://localhost:9093/cb?code=282423&state=xyz
```

#### 1-2 使用code交换token
##### 请求方式
```
POST /token 
```
##### 请求头
- client_id
- client_secret

##### 请求参数
| 参数           | 类型     | 说明                    |
|--------------|--------|-----------------------|
| grand_type   | string | 固定值authorization_code |
| code         | string | 携带上次请求返回的code         |
| username     | string | 用户账号                  |
| password     | string | 用户密码                  |
| redirect_uri | string | 1-1中的redirect_url     |

##### 返回示例
```
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MCwiZXhwIjoxNjc4MjYyMjk3LCJuYmYiOjE2NzgyNjA4NTd9.Ka1XzujQyxgP4WWTWFWtV9B3NP73QxIoSmKRO7iv5nU",
    "expires_in": 1440,
    "id_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOaWNrbmFtZSI6IiIsImV4cCI6MTY3ODI2MjI5NywibmJmIjoxNjc4MjYwODU3fQ.zcJlWYIoLPeuOppvHhRNtpme5ucOnLGx7LZtmsmKqVM",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODA4NTI4NTd9.H_dxuRwzGhJo70ezT6fIGFWX-3PKsJzIa3DMMS9MDPg",
    "scope": "all",
    "token_type": "Bearer"
}
```

### 2 刷新token
待完善

### 3 退出登录
待完善

#### SSO
待完善