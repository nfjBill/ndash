package ndash

import (
	"math/rand"
)

func factorial2(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial2(num-1)
}

func Random(i int) int {
	rand.Seed(Timezone8().UnixNano())
	//生成10个0-99之间的随机数
	return rand.Intn(i)
}
