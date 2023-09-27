# Mybricks发布服务
本服务配合Mybricks系统使用，为该系统的页面制品发布和管理、浏览服务

服务编译后直接启动即可，配置文件置于同级目录下的conf/config.toml文件中。

页面可直接使用mybricks配置，因此本系统不再制作配套页面，仅就接口进行定义

## 功能说明
- [x] 用户管理、登录
- [ ] 应用管理
- [ ] 应用可用页面管理
- [ ] 页面管理
- [ ] 页面版本管理
- [ ] OSS对象存储管理
- [ ] 资源管理
- [ ] 角色权限管理

## 接口说明
除登录外，所有接口均需带入 header['Authorization'] = 'Bearer ' + token

token由登陆时服务提供

### 登录
- 请求方式：POST
- 请求地址：/api/v1/user/login
- 请求参数：
```json
{
    "username": "admin",
    "password": "sm2国密加密后的密文"
}
```
- 成功响应参数
```json
{
  "code": 0,
  "data": {
    "token": "",
    "user": {
      // 用户信息
    }
  }
}
```

### 新增用户
- 请求方式：POST
- 请求地址：/api/v1/user/add
- 请求参数：
```json
{
    "username": "admin",
    "password": "sm2国密加密后的密文",
    "nickName": "管理员",
    "email": "",
    "mobile": "",
    "status": 1 // 1-正常
}
```
- 成功响应参数
```json
{
  "code": 0,
  "data": {
    // 用户信息
  }
}
```

