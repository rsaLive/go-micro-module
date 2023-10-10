package msg

const (
	SERVER_CONFIG         = "config.yaml"
	SERVER_DUBBOGO_CONFIG = "dubbogo.yaml"
	MODE_ENV              = "MODE_ENV"
)

const (
	Success = "操作成功"
	Failed  = "操作失败"
)

const (
	Http = 200
)

const (
	ErrorLogin           = "账号或密码错误"
	ErrorJSONParse       = "json解析失败"
	ErrorJSONMarshal     = "json序列化失败"
	ErrorForUpdate       = "锁定错误"
	ErrorInsert          = "插入异常"
	ErrorDelete          = "删除异常"
	ErrorUpdate          = "更新异常"
	ErrorSelect          = "查询异常"
	ErrorEmptyParam      = "值为空"
	ErrorCopierStruct    = "拷贝结构体错误"
	ErrorNoAction        = "无需操作"
	ErrorInvalidParam    = "参数不合法"
	ErrorNoData          = "没有数据"
	ErrorDatetime        = "时间格式错误"
	ErrorAddFoodCode     = "取餐码生成失败"
	ErrorGetToken        = "获取token错误"
	ErrorGetOpenID       = "获取openid错误"
	ErrorGetWXPhone      = "获取微信手机号错误"
	FailedGetWXPhone     = "获取微信手机号失败"
	WXPhoneUsed          = "手机号已绑定其他微信"
	ErrorWXLoginCode     = "登录码为空"
	ErrorDestroy         = "注销错误"
	AlreadyDestroy       = "已注销"
	ErrorWxPayClient     = "微信支付client错误"
	ErrorWxPayOrder      = "微信下单错误"
	ErrorWxPayDescr256   = "微信回调解析错误"
	ErrorGetUser         = "获取用户信息错误"
	UserBalanceNotEnough = "余额不足"
	UserTicketNotEnough  = "余额不足"
	ModifyUserWealthFail = "用户余额操作失败"
	ErrorPrint           = "打印错误"
)
