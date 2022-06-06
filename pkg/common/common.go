package common

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"time"
)

// Response 响应共有响应头
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

const MySqlDSN = "byte_dance:7efftEaAtzjEwfT4@tcp(106.15.107.229:3306)/byte_dance?charset=utf8mb4&parseTime=True&loc=Local"

// MD5Salt MD5加密时的盐
const MD5Salt = "UII34HJ6OIO"

// JWT
const (
	Issuer              = "xhx" // 签发人
	MySecret            = "Fy3Jfa5AD"
	TokenExpirationTime = 2 * time.Hour * time.Duration(1) // Token过期时间
)

//参数检验器
var (
	Validate = validator.New()          // 实例化验证器
	Chinese  = zh.New()                 // 获取中文翻译器
	Uni      = ut.New(Chinese, Chinese) // 设置成中文翻译器
	Trans, _ = Uni.GetTranslator("zh")  // 获取翻译字典
)

// OSSPreURL OSS前缀
const OSSPreURL = "https://byte-dance-01.oss-cn-shanghai.aliyuncs.com/test/"
