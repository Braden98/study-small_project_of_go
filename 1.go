package main

import "fmt"

func main(){
   /* 这是我的第一个简单的程序 */
   fmt.Println("Hello, World!")
   var a int
   var b []int
   var c bool
   var d string
   
   a,c,e:=1,true,29
   a=3
   //f:=22
   _,e=e,a
   fmt.Printf("%v %v %v %q %v\n", a,b,c,d,e)
   const o,p=111,222
   const aa=3
   fmt.Printf("%v %v\n",o,p)
   const (
      zz="hello"
      yy="world"
      xx="."
   )
   fmt.Printf("%v %v %v",zz,yy,xx)
   bit:=12
   fmt.Printf("%v\n",bit)
   fmt.Printf("%v\n",^1==-2)
   fmt.Printf("%v\n",8^bit)

   //fmt.Printf("%v\n",!bit)
   //fmt.Printf("%v\n",~bit)
   switch a{
   case 3:
      fmt.Printf("case is 3!\n")
   case 4:
      fmt.Printf("case is 4!\n")
   }
   str:="nihao"
   switch {
   case "niha"==str:
      fmt.Printf("niha")
   case str=="nihao":
      fmt.Printf("this is nihao\n")
   }

   numbers:=[5]int{1,2,3,4,5}
   for i,x:=range numbers{
      fmt.Printf("%v %v",i,x)
   }
   num1,num2:=10,12
   max(num1,num2)
   fmt.Printf("%v %v",num1,num2)
  // fmt.Printf("%v %v",swap1(1,2))//printf函数参数无法接收
  // fmt.Println(swap2("adam","ubik"))//
   
  //数组
   //var int_arr [10] int
   var float_arr=[...]float64{1.0,2.0}
   fmt.Printf("\n\n%v",float_arr[1])

    arr_1:=[3][4]int{
      {1,2,3,4},
      {5,6,7,8},
      {9,10,11,12},//注意这里有逗号
      }

      fmt.Printf("\n\n%x\n",&arr_1[2][2])

     // var i2 int
      var ptr [3]*int
      var i uint8
      for i=0;i<3;i++{
         //ptr[i]=&a[i]
      }

      for i=0;i<3;i++{
         fmt.Printf("a[%d]=%d\n",i,*ptr[i])
      }

}
func max(num1 int,num2 int) int {
   //return num1>num2? = num1:num2
   if num1-num2>=0 {
      return num1
   }
   return num2
}
func swap1(a int,b int)(int ,int){
   a,b=b,a;
   return b,a
}

func swap2(a string,b string) (string, string){
   a,b=b,a;
   return b,a
}
