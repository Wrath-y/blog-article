package errcode

// 微信端错误码
var (
	WechatNetworkBusy         = &ErrCode{40000, "网络繁忙，请稍后重试", ""}
	WechatAuthTokenFailed     = &ErrCode{40001, "获取token失败，请重试", ""}
	WechatCacheTokenFailed    = &ErrCode{40002, "获取token失败，请重试", ""}
	WechatTokenExpired        = &ErrCode{40003, "token已过期，请重试", ""}
	WechatUserInfoFailed      = &ErrCode{40004, "获取用户信息失败，请重试", ""}
	WechatCacheUserInfoFailed = &ErrCode{40005, "获取用户信息失败，请重试", ""}
	WechatMobileFailed        = &ErrCode{40006, "获取用户手机号失败，请重试", ""}
	WechatInvalidParam        = &ErrCode{40100, "无效的参数", ""}
	WechatInvalidSign         = &ErrCode{40101, "无效的签名", ""}
	WechatBodyTooLarge        = &ErrCode{40102, "请求消息体过大", ""}
)
