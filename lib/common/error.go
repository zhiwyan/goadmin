package common

import "errors"

//错误类
type Err struct {
	ErrorNo  int
	ErrorMsg string
}

func (this *Err) Error() string {
	return this.ErrorMsg
}

//在这里添加错误码
//100x 请求相关错误
//110x 用户状态相关错误
//120x 数据库相关报错
//130x 系统错误
//1401 请求远程接口错误
//业务其他错误码自行添加2001等等
//
//@wiki 参考 http://10.2.1.12:8090/pages/viewpage.action?pageId=9461426
var (
	ERR_SUC               = &Err{ErrorNo: 0, ErrorMsg: "OK"}
	ERR_INPUT             = &Err{ErrorNo: 1001, ErrorMsg: "缺少参数"}
	ERR_INPUT_FMT         = &Err{ErrorNo: 1002, ErrorMsg: "参数格式错误"}
	ERR_SIGN              = &Err{ErrorNo: 1003, ErrorMsg: "签名错误"}
	ERR_FLOOD_REQUEST     = &Err{ErrorNo: 1004, ErrorMsg: "重复请求"}
	ERR_NOT_FOUND_REQUEST = &Err{ErrorNo: 1005, ErrorMsg: "请求地址错误"}
	ERR_PARAM_ERR         = &Err{ErrorNo: 1006, ErrorMsg: "参数错误"}
	ERR_CAPTCHA_ERR       = &Err{ErrorNo: 1007, ErrorMsg: "验证码错误"}
	ERR_ILLEGAL_REQUEST   = &Err{ErrorNo: 1007, ErrorMsg: "请求不合法"}

	ERR_NOT_LOGIN                  = &Err{ErrorNo: 102, ErrorMsg: "用户未登录"} //session 过期 1101,兼容ios特殊处理逻辑
	ERR_LOGIN_ERROR                = &Err{ErrorNo: 1102, ErrorMsg: "用户名或密码错误"}
	ERR_USER_NOT_IN_ERROR          = &Err{ErrorNo: 1103, ErrorMsg: "用户不存在或远程连接错误"}
	ERR_USER_PASSWD_NOT_SAME_ERROR = &Err{ErrorNo: 1104, ErrorMsg: "两次输入密码不一致"}
	ERR_USER_RESET_PASSWD_ERROR    = &Err{ErrorNo: 1105, ErrorMsg: "重置密码错误"}
	ERR_NOT_REGIST_ERROR           = &Err{ErrorNo: 1107, ErrorMsg: "您还不是学霸君的用户，请注册后再使用"}
	ERR_NOT_ACCOUNT_ALREADY_ERROR  = &Err{ErrorNo: 1108, ErrorMsg: "账户已存在"}

	ERR_MYSQL   = &Err{ErrorNo: 1201, ErrorMsg: "MySQL数据库报错"}
	ERR_MONGODB = &Err{ErrorNo: 1202, ErrorMsg: "MongoDB数据库报错"}
	ERR_REDIS   = &Err{ErrorNo: 1203, ErrorMsg: "Redis数据库报错"}

	ERR_UNKNOW                = &Err{ErrorNo: 1301, ErrorMsg: "系统未知错误"}
	ERR_RESP_NOSUCCESS_NOZERO = &Err{ErrorNo: 1302, ErrorMsg: "返回值失败操作"}

	ERR_REMOTE_CURL = &Err{ErrorNo: 1401, ErrorMsg: "远程请求错误"}

	ERR_OUTPUT = &Err{ErrorNo: 1502, ErrorMsg: "输出错误"}
	ERR_CUSTOM = &Err{ErrorNo: 1603, ErrorMsg: "用户自定义错误"}

	ERR_JOSON_DECODE = &Err{ErrorNo: 1700, ErrorMsg: "json 解析错误"}
)

// 其他错误
var (
	DB_ERROR            = errors.New("db error!")        // 操作数据库错误
	DB_NO_ROWS_AFFECTED = errors.New("no rows affected") // 没有数据修改错误
)
