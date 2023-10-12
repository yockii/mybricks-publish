package constant

const NeedLogin = "user"

// 特殊权限
const ()

// 后台
const ResourceManage = "manage"

// 首页
const (
	ResourceHome = ResourceManage + ":home"
	// 仪表盘
	ResourceDashboard = ResourceHome + ":dashboard"
)

// 目录管理
const (
	ResourceCategory = ResourceManage + ":category"
	// 增删改查
	ResourceCategoryAdd    = ResourceCategory + ":add"
	ResourceCategoryUpdate = ResourceCategory + ":update"
	ResourceCategoryDelete = ResourceCategory + ":delete"
	ResourceCategoryList   = ResourceCategory + ":list"
	// 单个实例
	ResourceCategoryDetail = ResourceCategory + ":detail"
)

// 用户管理
const (
	ResourceUser = ResourceManage + ":user"
	// 增删改查
	ResourceUserAdd    = ResourceUser + ":add"
	ResourceUserUpdate = ResourceUser + ":update"
	ResourceUserDelete = ResourceUser + ":delete"
	ResourceUserList   = ResourceUser + ":list"
	// 单个实例
	ResourceUserDetail = ResourceUser + ":detail"
	// 重置密码
	ResourceUserResetPassword = ResourceUser + ":resetPassword"
	// 分配角色
	ResourceUserAssignRole = ResourceUser + ":assignRole"
)

// 角色管理
const (
	ResourceRole = ResourceManage + ":role"
	// 增删改查
	ResourceRoleAdd    = ResourceRole + ":add"
	ResourceRoleUpdate = ResourceRole + ":update"
	ResourceRoleDelete = ResourceRole + ":delete"
	ResourceRoleList   = ResourceRole + ":list"
	// 单个实例
	ResourceRoleDetail = ResourceRole + ":detail"
	// 分配资源
	ResourceRoleAssignResource = ResourceRole + ":assignResource"
	// 设置默认角色
	ResourceRoleSetDefaultRole = ResourceRole + ":setDefaultRole"
)

// 字典管理
const (
	ResourceDict = ResourceManage + ":dict"
	// 增删改查
	ResourceDictAdd    = ResourceDict + ":add"
	ResourceDictUpdate = ResourceDict + ":update"
	ResourceDictDelete = ResourceDict + ":delete"
	ResourceDictList   = ResourceDict + ":list"
	// 单个实例
	ResourceDictDetail = ResourceDict + ":detail"
)

// 路由管理
const (
	ResourceRoute = ResourceManage + ":route"
	// 增删改查
	ResourceRouteAdd    = ResourceRoute + ":add"
	ResourceRouteUpdate = ResourceRoute + ":update"
	ResourceRouteDelete = ResourceRoute + ":delete"
	ResourceRouteList   = ResourceRoute + ":list"
	// 单个实例
	ResourceRouteDetail = ResourceRoute + ":detail"
)

// 应用管理
const (
	ResourceApplication = ResourceManage + ":application"
	// 增删改查
	ResourceApplicationAdd    = ResourceApplication + ":add"
	ResourceApplicationUpdate = ResourceApplication + ":update"
	ResourceApplicationDelete = ResourceApplication + ":delete"
	ResourceApplicationList   = ResourceApplication + ":list"
	// 单个实例
	ResourceApplicationDetail = ResourceApplication + ":detail"
)

// oss配置管理
const (
	ResourceOssConfig = ResourceManage + ":ossConfig"
	// 增删改查
	ResourceOssConfigAdd    = ResourceOssConfig + ":add"
	ResourceOssConfigUpdate = ResourceOssConfig + ":update"
	ResourceOssConfigDelete = ResourceOssConfig + ":delete"
	ResourceOssConfigList   = ResourceOssConfig + ":list"
	// 单个实例
	ResourceOssConfigDetail = ResourceOssConfig + ":detail"
)
