package errcode

// 后台应用错误码
var (
	AdminNetworkBusy           = &ErrCode{39000, "网络繁忙，请稍后重试", ""}
	AdminUserNotExist          = &ErrCode{39002, "用户名或密码错误", ""}
	AdminResourceForbidden     = &ErrCode{39003, "您没有权限操作此资源", ""}
	AdminAPIAuthInUse          = &ErrCode{39004, "接口已被分配给角色，请撤销相关权限后再删除", ""}
	AdminMenuCannotConvert     = &ErrCode{39005, "父菜单与子菜单不能相互转换", ""}
	AdminMenuNameRepeat        = &ErrCode{39006, "菜单名重复，请调整后重新提交", ""}
	AdminCannotModifySuperUser = &ErrCode{39007, "不能添加或修改超管用户", ""}
	AdminUserAlreadyExist      = &ErrCode{39008, "该用户名已存在", ""}
	AdminCaptchaInvalid        = &ErrCode{39009, "验证码错误或已失效", ""}
	AdminRoleUsed              = &ErrCode{39010, "角色已被占用,无法修改", ""}
	AdminMenuHasChild          = &ErrCode{39011, "菜单下还有启用的子菜单", ""}
	AdminInvalidParam          = &ErrCode{39100, "无效的参数，请检查后重试", ""}
	AdminNoRecord              = &ErrCode{39101, "记录不存在", ""}
	AdminRecordExist           = &ErrCode{39102, "记录已存在，请勿重复提交", ""}
	AdminRoleNoPermission      = &ErrCode{39103, "暂无权限，请联系管理员分配权限后重新登录", ""}
	AdminUserForbidden         = &ErrCode{39104, "用户已被禁用，请联系管理员", ""}
	AdminInvalidPassword       = &ErrCode{39105, "密码必须包含数字大小写字母且8-16位", ""}
	AdminInvalidSign           = &ErrCode{39106, "无效的签名", ""}
	AdminUploadFailed          = &ErrCode{39107, "上传失败", ""}
	AdminLoginExpired          = &ErrCode{40003, "登录已过期，请重新登录", ""}
)
