package schema

import (
	"errors"
	"fmt"
	"log"
)

/* 全局接口错误码定义 99开头的为系统级错误 */

// 错误码：10***
var (
	// 通用错误码
	SUCCESS         = errors.New("00000:ok")
	FAIL            = errors.New("99999:%s")
	MONGO_ERROR     = errors.New("99100:mongo error -> %s")
	TCC_VALUE_ERROR = errors.New("99201:无效tcc值 %s")

	// 业务错误码
	E10101 = errors.New("10101:名称已经存在，不能重复新增")
	E10102 = errors.New("10102:名称不存在")
)

func Error(userDefinedErr error, args ...interface{}) error {
	if userDefinedErr == nil {
		log.Fatal("错误码定义不能为空")
	}
	e := userDefinedErr.Error()
	errMsg := fmt.Sprintf(e, args...)
	return errors.New(errMsg)
}

func Panic(userDefinedErr error, args ...interface{}) {
	if userDefinedErr == nil {
		log.Fatal("错误码定义不能为空")
	}
	e := userDefinedErr.Error()
	errMsg := fmt.Sprintf(e, args...)
	panic(errors.New(errMsg))
}
