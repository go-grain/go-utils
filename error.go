package go_utils

import (
	"github.com/go-pay/gopay/pkg/xlog"
)

func GoRecoverPanic(handler func(err interface{})) {
	if err := recover(); err != nil {
		handler(err)
	}
}

func Go(always bool, str string, f func()) {
	defer GoRecoverPanic(func(err interface{}) {
		if always {
			xlog.Info(str, "协程崩溃啦... 正在重新启动: ", err)
			f()
		} else {
			xlog.Info(str, " 程序崩溃...", err)
		}
	})
	f()
}
