// package main
// import "fmt"
// func main(){
// 	fmt.Println("hello world")
// }

// package main
// import "fmt"
// func main(){
// 	var age int
// 	age = 26
// 	fmt.Println("age = ",age)
// }

// package main
// import "fmt"
// func main(){
// 	var num1 int = 26
// 	fmt.Println("num1 = ",num1)
// 	var num2 int
// 	fmt.Println("num2 = ",num2)
// 	num3 := "男"
// 	fmt.Println("num3 = ",num3)
// 	var num4 = "abcdefg"
// 	fmt.Println("num4 = ",num4)
// }

// package main
// import "fmt"
// func main(){
// 	var num int
// 	num = 100
// 	fmt.Printf("num:%d, num的地址=%p\n",num, &num)
// }

// package main
// import "fmt"
// func main(){
// 	a := 100
// 	const baidu string = "www.baidu.com"
// 	const pi = 3.14
// 	fmt.Println(a,baidu,pi)
// }

// package main
// import "fmt"
// func main(){
// 	const (
// 		a = 100
// 		b
// 		c = "baidu"
// 		d
// 		e
// 	)
// 	fmt.Printf("%T,%d\n",a,a)
// 	fmt.Printf("%T,%d\n",b,b)
// 	fmt.Printf("%T,%d\n",c,c)
// 	fmt.Printf("%T,%d\n",d,d)
// 	fmt.Printf("%T,%d\n",e,e)
// }

// package main
// import "fmt"

// func main(){
// 	const(
// 		A = iota
// 		B = 100
// 		C
// 		D
// 		E = "NIHAO"
// 		F
// 		G = iota
// 		H
// 	)
// 	const(
// 		I = iota
// 	)
// 	fmt.Println(A)
// 	fmt.Println(B)
// 	fmt.Println(C)
// 	fmt.Println(D)
// 	fmt.Println(E)
// 	fmt.Println(F)
// 	fmt.Println(G)
// 	fmt.Println(H)
// 	fmt.Println(I)
// }

// package main
// import "fmt"

// func main(){
// 	var a1 = "Steve"
// 	a2 := 'S'
// 	a3 := '武'
// 	a4 := '昕'
// 	a5 := '昆'

// 	fmt.Printf("%s\n",a1)
// 	fmt.Printf("%T,%d\n",a2,a2)
// 	fmt.Printf("%T,%d,%c,%q\n",a3,a3,a3,a3)
// 	fmt.Printf("%T,%d,%c,%q\n",a4,a4,a4,a4)
// 	fmt.Printf("%T,%d,%c,%q\n",a5,a5,a5,a5)
// }

package main
import "fmt"

func main(){
	var a int
	var b string
	fmt.Println("请输入两个字符：")
	fmt.Scanln(&a,&b)
	fmt.Printf("a=%d,b=%s",a,b)
}