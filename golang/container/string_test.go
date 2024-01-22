package container

import (
	"fmt"
	"testing"
	"unsafe"
)

/**
来实现string相关内容呗

1、 字符集
ASCII字符集-1967，  扩展字符集到256个  英文字母、阿拉伯数字、控制字符
GB2312字符集-1980   简体中文、拉丁字母、日文
BIG5字符集-1984     增加繁体字
Unicode字符集-1994  全球化 实现了跨语言、跨平台
	               但是
                   但是定义一个字符串 hello，你好世界  与字符集进行映射可以得到二进制编码，但是解码操作分不清楚边界
解决方式：定长编码，每个字符都是占用四个字节，但是这种浪费内存

UTF8编码:变长编码  编号			模版
				【0,127】        [0???????]
			    【128，,2047】   [110?????,10??????]
			    【2048，,65535】 [110?????,10??????,10??????]

简单 for(i:=0;i<len(str);i++) 这种是遍历字节
    for(index,value:=range str) 这种遍历的是字符，在遍历过程中进行解码操作 golang默认是UTF-8编码格式

-------------------------------------------------------------------------------------------
golang string底层 -> golang没有采用类似于c语言使用\0结尾作为标识符，而是以长度len代表字节个数，限制了开头和结尾
data len
*/

// string底层数组
type stringStruct struct {
	str unsafe.Pointer
	len int
}

// 创建字符串
//
//go:nosplit 在runtime.string.go中
func gostringnocopy(str *byte) string {
	ss := stringStruct{str: unsafe.Pointer(str), len: findnull(str)}
	s := *(*string)(unsafe.Pointer(&ss))
	return s
}

// 计算字符串长度呗
func findnull(str *byte) int {
	if str == nil {
		return 0
	}
	return 1
}

func TestString(t *testing.T) {
	str := "你好世界"
	fmt.Println(len(str)) //这里边就是输出12，代表这个是12个字节
}
