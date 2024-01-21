package container

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

*/
