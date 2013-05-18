/*
 * 内置引用类型的一些用法,如: map slice
 * File: RefreceType.go
 * date: 2013-05-18
 */
 package main
 
 import(
  "fmt"
 )
 
 //这几个类型都是引用类型
 //1 slice 切片
 var s1 string = "abcdef"
 var slice1 []byte = s1[2:5]//表示s1变量的第2个位置（包括），到第5个位置（不包括）
 var slice2 []byte = s1[1:4]
 var slice3 []byte = slice1[1:]//这个:后的数据都可以省略，如果没有前边的数据，比如 :5，则表示从第0个到第5个，省略后边则表示一直到结尾
 
 //然后就可以修改slice引用的变量了
 slice1[2]='g'
 //这样，相同引用的slice2 slice3的值都会变化了
 
 //2 map 需要使用make进行内存分配
 var map1 map[int]string = make(map[int]string)
 map1[1] = "a"
 map1[2] = "哈"
 
 
 func main(){
  fmt.Println("");
 }
