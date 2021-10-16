package main

import (
	//"unsafe"
	"fmt"
	//"strconv"
)

// func main() {
// 	// var a int = 10
// 	// var b int = 20
// 	// var r = a + b
// 	// println("a + b = ",r)
// 	var n = 100
// 	fmt.Printf("n的数据类型是:%T,n的字节大小为:%d", n, unsafe.Sizeof(n))
// func main(){
// 	var adr string="hello changcheng!"
// 	var adr2 string="I Love you!"
// 	var add string = adr+adr2
//     fmt.Println(add)
// }
// func main(){
// 	var v float64 = 1
// 	fmt.Printf("v=%f",v)
// }
// func main(){
// 	var n int = 10
// 	var x float32 = 20.324
// 	var y bool = true
// 	var str1 string
// 	var str2 string
// 	var str3 string
// 	var str string
// 	str1 = fmt.Sprintf("%d",n)
// 	str2 = fmt.Sprintf("%f",x)
// 	str3 = fmt.Sprintf("%v",y)
// 	fmt.Printf("%q,%sq,%q",str1,str2,str3)
// 	str = strconv.FormatInt(int64(n),10)
// 	fmt.Printf("%q",str)
// 	str = strconv.FormatFloat(float64(x),'f',6,32)
// 	fmt.Printf("%q",str)
// 	str = strconv.Itoa(n)
// 	fmt.Printf("%q",str)
// }

// func main(){
// 	a := 10
// 	b := 20
// 	c := 30
// 	var max int
// 	if a > b{
//     max = a
// 	}else if a < b{
// 		max = b
// 	}
// 	if c > max{
// 		max = c
// 	}
// 	fmt.Println("max=",max)
// }
//   * 
//  ***
// *****


//  ***
//   *

  ///打印空心金字塔
func main(){

	var n int = 20
	for i:=1;i<=n;i++{
		for k:=1;k<=n-i;k++{
			fmt.Print(" ")
		}
		for j:=1;j<=2*i-1 ;j++{
        if j == 1|| j == 2*i-1 {
			fmt.Print("*")
		}else{
			fmt.Print(" ")
			}	
		
		}
		fmt.Print("\n")
	}
	for i:=1;i<=n-1;i++{
		for q:= 1;q<=i ;q++{
		fmt.Print(" ")
    }
	for k:=1;k<=2*n-2*i-1;k++{
		if k==1 || k==2*n-2*i-1 || i==n-1 {
		fmt.Print("*")
		}else{
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
}
