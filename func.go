package ndash

var retryNum int8

func FuncRetries(condition bool, retries int8, noMore func(), exceed func(), success func()) {
	// condition为true 进入错误判断
	if condition {
		if retryNum < retries {
			retryNum += 1
			noMore()
		} else {
			retryNum = 0
			exceed()
		}
	} else {
		retryNum = 0
		success()
	}
}
