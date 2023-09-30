package limit

import (
	"sync"
	"time"
)

// 百度一面面试官被闻到过单机限流令牌桶如何实现
type Bucket struct {
	maxLimit int32
	token    int32
	psm      string
	lock     *sync.Mutex //这里边只声明 但是不需要进行初始化就能够使用的原因是 mutex->state seam两个Int类型字段 0值就是未上锁的状态

}

func (b *Bucket) allow(psm string) bool {
	b.lock.Lock()
	defer b.lock.Unlock() //这么做就是好像有点简单呗
	if b.token > 0 {
		b.token--
		return true
	}
	return false
}

func (b *Bucket) timer() {
	tick := time.Tick(1 * time.Second)

	for {
		select {
		case <-tick:
			//重新将令牌桶初始化为最大值
			b.lock.Lock()
			defer b.lock.Unlock()
			b.token = b.maxLimit
		default:
		}
	}
}
