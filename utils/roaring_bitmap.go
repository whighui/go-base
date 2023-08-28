package utils

//redis_scene bitmap 缺点 为什么转换用roaring bitmap 开源

//https://cloud.tencent.com/developer/article/1136054
//hashmap 数组+链表 ---> 数组+红黑树
// Roaring BitMap 其实是一个意思, 数组 + array container, 后续 array container 变为 bitmap container一句话就能讲原理说清楚,
