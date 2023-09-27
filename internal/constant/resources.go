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

// 字典管理
const (
	ResourceNotice = ResourceManage + ":notice"
	// 增删改查
	ResourceNoticeAdd    = ResourceNotice + ":add"
	ResourceNoticeUpdate = ResourceNotice + ":update"
	ResourceNoticeDelete = ResourceNotice + ":delete"
	ResourceNoticeList   = ResourceNotice + ":list"
	// 单个实例
	ResourceNoticeDetail = ResourceNotice + ":detail"
)

// 横幅管理
const (
	ResourceBanner = ResourceManage + ":banner"
	// 增删改查
	ResourceBannerAdd    = ResourceBanner + ":add"
	ResourceBannerUpdate = ResourceBanner + ":update"
	ResourceBannerDelete = ResourceBanner + ":delete"
	ResourceBannerList   = ResourceBanner + ":list"
	// 单个实例
	ResourceBannerDetail = ResourceBanner + ":detail"
)

// 团队管理
const (
	ResourceTeam = ResourceManage + ":team"
	// 增删改查
	ResourceTeamAdd    = ResourceTeam + ":add"
	ResourceTeamUpdate = ResourceTeam + ":update"
	ResourceTeamDelete = ResourceTeam + ":delete"
	ResourceTeamList   = ResourceTeam + ":list"
	// 单个实例
	ResourceTeamDetail = ResourceTeam + ":detail"
)

// 评价管理
const (
	ResourceOpinion = ResourceManage + ":opinion"
	// 增删改查
	ResourceOpinionAdd    = ResourceOpinion + ":add"
	ResourceOpinionUpdate = ResourceOpinion + ":update"
	ResourceOpinionDelete = ResourceOpinion + ":delete"
	ResourceOpinionList   = ResourceOpinion + ":list"
	// 单个实例
	ResourceOpinionDetail = ResourceOpinion + ":detail"
)

// 订单管理
const (
	ResourceOrder = ResourceManage + ":order"
	// 增删改查
	ResourceOrderAdd    = ResourceOrder + ":add"
	ResourceOrderUpdate = ResourceOrder + ":update"
	ResourceOrderDelete = ResourceOrder + ":delete"
	ResourceOrderList   = ResourceOrder + ":list"
	// 单个实例
	ResourceOrderDetail = ResourceOrder + ":detail"
)
