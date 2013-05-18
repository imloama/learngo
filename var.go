/**
 * 学习变量使用章节
 * 定义的变量必须使用,否则在编译时,不能通过编译
 * File:var.go
 * time:2013-05-17
 */
package main

import(
  "fmt"
)

//定义常量,普通定义方式
const var1 int = 0

//定义多个常量，这种方式不能用在方法内
const(
  var2 int = 1
  var3 string = "定义一个变量"
)

//特殊变量 iota，初始值为0  每次赋值增加1  当遇到新的const时，再次初始化为0
const(
  var4 = iota//=0
  var5 = iota//=1
  var6       //=2  值自动累加  
)
const viota = iota //=0表示

//定义普通变量
var var7 int = 1
var var8,var9 int = 3,4

var ai int
ai = 1

//按组定义变量
var(
  var10 int
  var11 string
)

//最简化的定义方式
var12 :=2.0//自动添加变量类型   float
var13,var14 := 5,6//自动添加变量类型为int
//虚数
var c complex128 = 5+5i

var a, _ = 3, 4 //特殊变量 下拉框  _ 表示接收参数，直接丢弃

//定义byte类型
var byte1 []byte = 'abc'
var s1 string = "abc"
//string加
var s2 string = "abc"+//+只能放在这一行,不能放到下一行,这是要求.
  "def"
//定义多行文本
var s3 string = `abc
  def`
//byte与string的相互转换
var byte2 []byte= []byte(s1)
var s3 string = string(byte2)

//不能直接修改string引用的内容,如下边的用法是错误的
s3[0] = 'd'

//主函数
function main(){
  fmt.Println("var4=%d",var4);
}
