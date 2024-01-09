package ctrl

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	zh_translation "github.com/go-playground/validator/v10/translations/zh"
	"github.com/go-playground/validator/v10/translations/zh_tw"
	"go.uber.org/zap"
	"net/http"
)

type TranslateErr []string

func (t *TranslateErr) Error() string {
	var (
		str = bytes.NewBufferString("")
	)
	for i, v := range *t {
		_, _ = str.WriteString(v)
		if i != len(*t)-1 {
			_, _ = str.WriteString(", ")
		}
	}
	return str.String()
}

// HandlerError 错误统一处理
func HandlerError(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if c.Errors != nil {
			resp := Response{
				Code: "FAIL",
				Data: nil,
			}
			last := c.Errors.Last()
			switch last.Err.(type) {
			case validator.ValidationErrors: // 判断错误是否是validator错误
				validationErrors := last.Err.(validator.ValidationErrors)
				// 翻译验证错误消息
				translatedErrors := make(TranslateErr, 0, len(validationErrors))
				for _, e := range validationErrors {
					translatedErrors = append(translatedErrors, e.Translate(translate("zh")))
				}
				resp.Msg = translatedErrors.Error()
			default:
				log.Error("HandlerError | "+c.Request.RequestURI, zap.Error(last.Err))
				resp.Msg = c.Errors.String()
			}
			c.JSON(200, resp)
			return
		}
	}
}

func translate(language string) ut.Translator {
	var trans ut.Translator
	var validate = binding.Validator.Engine().(*validator.Validate)
	switch language {
	case "en":
		defaultEn := en.New()
		uni := ut.New(defaultEn, defaultEn)
		trans, _ = uni.GetTranslator(language)
		_ = en_translation.RegisterDefaultTranslations(validate, trans)
	case "zh-tw":
		defaultZhTw := zh_Hant_TW.New()
		uni := ut.New(defaultZhTw, defaultZhTw)
		trans, _ = uni.GetTranslator(language)

		_ = zh_tw.RegisterDefaultTranslations(validate, trans)
	default:
		defaultZh := zh.New()
		uni := ut.New(defaultZh, defaultZh)
		trans, _ = uni.GetTranslator(language)
		_ = zh_translation.RegisterDefaultTranslations(validate, trans)
	}
	return trans
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
