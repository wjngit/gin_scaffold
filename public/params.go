package public

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

func DefaultGetValidParams(c *gin.Context, params interface{}) error {
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	//获取验证器
	valid, err0 := GetValidator(c)
	if err0 != nil {
		return err0
	}
	//获取翻译器
	trans, err := GetTranslation(c)
	if err != nil {
		return err
	}
	err = valid.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var sliceErrs []string
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}

func GetValidator(c *gin.Context) (*validator.Validate, error) {
	val, ok0 := c.Get(ValidatorKey)
	if !ok0 {
		return nil, errors.New("未设置验证器")
	}
	validate, ok := val.(*validator.Validate)
	if !ok {
		return nil, errors.New("获取验证器失败")
	}
	return validate, nil
}

func GetTranslation(c *gin.Context) (ut.Translator, error) {
	trans, ok0 := c.Get(TranslatorKey)
	if !ok0 {
		return nil, errors.New("未设置翻译器")
	}
	translator, ok := trans.(ut.Translator)
	if !ok {
		return nil, errors.New("获取翻译器失败")
	}
	return translator, nil
}
