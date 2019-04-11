# Go
程序一般由关键字、常量、变量、运算符、类型和函数组成
`{`不可以单独一行
![-w849](media/15549824795525/15549825302508.jpg)

##赋值
正确示例
```go
var a int 
_,b,c:=3,5,7.7
a=4
```
不能声明在同一代码块未使用的局部变量，但全局变量可以声明&不使用。声明变量但不赋值可以用`var b int`，声明并赋值可以`a:=5.4`隐式判断类型,多个变量声明赋值可以`c,d,_:=5,6,7`，交换两个变量的值可以`a,b=b,a``_`是指不赋值（保持不变）而非赋默认值。
## 常量
类型只能是bool，数字（整型，浮点型和复数），字符串型，用const定义
`const o,p:=111,222`是错的，:必须去除，否则报错：`syntax error: unexpected := at end of statement`。并且不能将已经声明或者赋值的变量再次进行const声明赋值。
const也不能像变量一样只声明不赋值，否则报错`missing value in const declaration`
常量也可以用作枚举，
```go
 const (
      zz="hello"
      yy="world"
      xx="."
   )
```
## 标准输出
   fmt.Printf("%v %v %v %q %v\n", a,b,c,d,str)
其中P大写相当于java中public,小写声明相当于protected
##运算符
同java，注意`^`作为单目运算符是取反，多目运算符是xor
```go
   bit:=12
   fmt.Printf("%v ",bit)
   fmt.Printf("%v ",^bit)
   fmt.Printf("%v ",8^bit)
```
结果是`12 -13 4`-13是因为
##基本流
### 条件
条件语句同java，只是多了`select`，且布尔表达式不需()
```go
switch a{
   case 3:
      fmt.Printf("case is 3!\n")
   case 4:
      fmt.Printf("case is 4!\n")
   }
   str:="nihao"
   switch {
   case "niha":
      fmt.Printf("niha")
   case str=="nihao":
      fmt.Printf("this is nihao\n")
   }
```
switch 可以case后写布尔表达式，也可以在switch时声明，然后case一个值，自动提供==，字符串、数值都是这样。
###循环
for{} for xx;xx;xx {} for cond {} 相比java只是少了括号，break和continue也一样。
还多了一个range循环，可以对切片、map、数组、字符串迭代，执行过程。
```go
numbers:=[5]int{1,2,3,4,5}
   for i,x:=range numbers{
      fmt.Printf("%v %v",i,x)
   }
```
i,x顺序对应其index和value，若少一个，则先对应index.

##函数
函数声明告诉了编译器函数的名称，返回类型，和参数。定义如下
```go
func function_name( [parameter list] ) [return_types] {
   函数体
}
```
```go
func max(num1 int,num2 int) int {
   //return num1>num2? = num1:num2 不符合一种事情只有一种方法的golang思想
   if num1-num2>=0 {
      return num1
   }
   return num2
}
```
经典的swap问题
```go
func swap(x *int, y *int) {
   var temp int
   temp = *x    /* 保持 x 地址上的值 */
   *x = *y      /* 将 y 值赋给 x */
   *y = temp    /* 将 temp 值赋给 y */
}
```
这样是可以swap的，否则不行。
## 作用域
和c一样，哈哈哈
## 数组
一样不可以赋空值，基本操作`var int_arr [10] int` `var float64_arr=[10] float64{1.2,2.4}`，也可用:=代替var
### 多维数组
```go
var n_arr [][][] int
var arr_1 =[4][3]int{
{1,2,3,4},
{5,6,7,8},
{9,10,11,12},//注意这里有逗号
}
```
###数组传参
`void myFunction(param [10]int){}`,也可以不设置长度
