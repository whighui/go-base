package golang

/**
内存对齐

1. 首先理解概念 地址总线、数据总线、内存模型
地址总线：cpu寻址范围与地址总线有关，与数据总线无关，地址总线的位数决定了cpu寻址空间
数据总线：cpu位宽=数据总线的宽度=cpu内部通用寄存器的宽度=机器字长   位宽通常与cpu的字长一致  数据总线是双向通道
内存模型：单双通道->cpu并行处理内存->rank->chip->bank->row col  每个rank由8个chip组成 每个chip有8个bank，每个bank就是数据存储的实体

2.MemoryAlignment 以当前结构体为例子
先确定每个成员对齐值，分别是
						A 1
						B 2  ---->max ----->16byte  所以16是对齐值
						C 8
						D 16
进行内存对齐，最后还需要整个结构体的对齐大小是对齐值的整数倍
*/

type MemoryAlignment struct {
	A int8
	B int16
	C int64
	D string
}
