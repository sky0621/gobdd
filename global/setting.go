package global

import (
	"sync"
)

var isLocal bool
var isLocalOnce sync.Once

func IsLocal() bool {
	return isLocal
}

func InitIsLocal(f bool) {
	// 当関数が何度呼ばれても１度きりしかフラグセットされないことを保証
	isLocalOnce.Do(func() { isLocal = f })
}
