/**
 * 变量使用章节
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
  var5        //=1
  var6        //=2
)

//定义普通变量
var var7 int = 1
var var8,var9 int = 3,4

var(
  var10 int
  var11 string
)

//最简化的定义方式
var12 :=2.0
var13,var14 := 5,6
//虚数
var c complex128 = 5+5i

var a, _ = 3, 4 //特殊变量 下拉框  _ 表示接收参数，直接丢弃


function main(){
  fmt.Println("var4=%d",var4);
}
