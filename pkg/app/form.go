package app

import (
	"github.com/gin-gonic/gin"
	// validator 的翻译器
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	"strings"
)

// ValidError 自定义错误类型
type ValidError struct {
	Key     string
	Message string
}

// ValidErrors 自定义数组错误类型
type ValidErrors []*ValidError

// Error 实现接口才被识别为自定义错误
func (v *ValidError) Error() string {
	return v.Message
}

// Error 实现接口才被识别为自定义错误
func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

// Errors 实现接口才被识别为自定义错误，错误可以是一个列表，包含多个错误
func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// bindAndValid 绑定校验，校验请求参数是否合法与翻译成对应语言
func bindAndValid(c *gin.Context, err error) (bool, ValidErrors) {
	var errs ValidErrors
	if err != nil {
		v := c.Value("trans")
		// 翻译模块
		trans, _ := v.(ut.Translator)
		// 参数错误验证
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs
		}
		// 返回哪些参数有问题，并翻译
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}

		return false, errs
	}

	return true, nil
}
func BindAndValidBody(c *gin.Context, v interface{}) (bool, ValidErrors) {
	err := c.ShouldBindJSON(v)
	if err != nil && err.Error() == "EOF" {
		err = c.ShouldBind(v)
	}
	return bindAndValid(c, err)
}
func BindAndValidHeader(c *gin.Context, v interface{}) (bool, ValidErrors) {
	err := c.ShouldBindHeader(v)
	return bindAndValid(c, err)

}
