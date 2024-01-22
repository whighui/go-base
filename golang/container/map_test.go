package container

import "unsafe"

/**
golang map考点

1. 扩容 渐进式扩容，避免一次扩容带来的瞬时抖动
   count/(2^B)>6.5 ------> 翻倍扩容hmap.B++
   loadFactory没超标，但是noverflow较多，这里边需要对B进行细分 这时候进行等量扩容

2.  mapaccess 查找
    makemap   创建map
    mapassign 赋值
	hashGrow  扩容 evacuate 迁移
    mapdelete 删除
	mapiterinit 遍历选择初始桶 重点是fastrand选择随机数






考点：哪些类型可以作为map key
基本类型：int、float、bool、string、接口类型、指针、数组可以作为map key
然而，以下类型不能作为 map 的键：slice、map、function、包含上述类型的结构体
这是因为 slice、map 和 function 等类型的值不是固定的（它们在内存中的表示可能会改变），因此不能用于比较。例如，两个包含相同元素的 slice 在使用 == 运算符进行比较时会产生编译错误，因此 slice 不能作为 map 的键。
*/

const bucketCnt = 8

type hmap struct {
	count      int            //键值对数目
	flags      uint8          //
	B          uint8          //buckets=2^B
	noverflow  uint16         //记录溢出桶使用的数量
	hash0      uint32         // hash seed
	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer //旧桶
	nevacuate  uintptr        //即将迁移的旧桶编号
	extra      *mapextra      //记录溢出桶信息 在这里哈呢
}

type mapextra struct {
	overflow     *[]*bmap //切片溢出桶数量 如果B>4 在创建的时候会创建2^B-4个溢出桶  overflow代表已使用溢出桶数切片
	oldoverflow  *[]*bmap //扩容之前旧map使用的溢出桶数量
	nextOverflow *bmap    //下个空闲溢出桶
}

// A bucket for a Go map. 不是key/value key/value每个节点进行存放，避免每个节点进行内存对齐，只需要最后进行一次内存对齐即可
type bmap struct {
	tophash  [bucketCnt]uint8
	keys     [8]uint16 //mkaetype 编译期间进行确定
	values   [8]uint16
	overflow uintptr
}
