package common

const (
	ENV_PRO        = "pro"
	ENV_QA         = "qa"
	UPLOAD_TMP_DIR = "/tmp/"

	F_DATE         = "2006-01-02"          // 长日期格式
	F_DATETIME     = "2006-01-02 15:04:05" // 日期时间格式
	F_DATETIME_CN  = "01月02日15:04"         // 日期时间格式
	F_DATETIME_DOT = "2006.01.02 15:04:05" // 日期时间格式

	F_SHORTTIME = "15:04"    // 短时间格式
	F_TIMES     = "15:04:05" // 长时间格式
)

// 转换为json的空数组
var JsonEmptyArr = []struct{}{}

// 转换为json的空对象
var JsonEmptyObj = struct{}{}
