package container

import (
	"fmt"
	"testing"
)

/**
切片相比于 Array 引入了一个抽象层 提供了对数组部分连续片段的引用 作为数组的引用 可以在运行期间修改他的长度和范围 扩容底层进行扩容 不过切片本身无变化 所以说切片只是在数组上引入了一个抽象层

扩容：
1.容量预估
2.确定容量
3.申请内存  类似于string切片 新的容量是5 进行申请内存是5*16=80字节，需要看golang内存管理当中是否有该内存型号呗，如若没有正常型号，需要向上兼容呗
*/

// SliceHeader slice底层数据结构
type SliceHeader struct { //运行期间切片结构体
	Data uintptr //指向数组的指针 底层数据是一片连续的内存空间
	Len  int     //当前切片的长度
	Cap  int     //切片的容量也即就是底层数组的长度
}

/*
var来初始化切片 切片本身就是一个nil类型，没有赋初始值
*/
func TestInitSliceVar(t *testing.T) {
	var ints []int
	fmt.Println(ints)
}

/*
*
测试slice扩容
例如：
slice:=[]int{1,2}
slice=append(slice,3,4,5)     扩容前容量cap=2 预估扩容到cap=5

预估规则：1. oldCap*2<cap-->newCap=cap
 2. oldLen<1024 newCap=oldCap*2
 3. oldLen>1024 newCap=oldCap*0.25
*/
func TestSliceGrow(t *testing.T) {

}
