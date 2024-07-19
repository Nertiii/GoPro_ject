package main
import (
	"fmt"
	//"strings"
)
func main(){
	//strings
	// mac := "hello nerti"

	// fmt.Println(strings.Contains(mac,"true"))
	// fmt.Println(strings.ContainsAny(mac,"abcde"))
	// fmt.Println(strings.Count(mac,"n"))

	// ss1 := []string{"百事","无糖","生可乐"}
	// aa := strings.Join(ss1,"-")
	// fmt.Println(aa)

	// fmt.Println(strings.Repeat(aa,5))
	// fmt.Println(strings.Replace(aa,"-","*",-1))

	// fmt.Println(strings.ToUpper(mac))
	// s2 := mac[:5]
	// fmt.Println(s2)

	//函数
	
	// aa, bb := plus(1,100)
	// fmt.Println(aa)
	// fmt.Println(bb)
	// s1 := [4]int{1,2,3,4}
	// s2 := []int{5,6,7,8}
	// fmt.Println(s1)
	// cha1(s1)
	// fmt.Printf("修改后:%d\n",s1)
	// fmt.Println(s2)
	// cha2(s2)
	// fmt.Printf("修改后:%d\n",s2)


	//递归函数
	sum:=getFbnx(12)
	fmt.Println(sum)
}
func getFbnx(n int)int{
	if n==1||n==2{
		return 1
	}
	return getFbnx(n-1) + getFbnx(n-2)
}

// func getsum(n int)int{
// 	if n>=1{
// 		return getsum(n-1) + n
// 	}
// 	return 0
// }

// func plus (a int, b int)(int){
// 	sum := 0
// 	for i := a; i<=b; i++{
// 		sum+=i
// 	}
// 	// fmt.Println(sum)
// 	return sum
// }

// func plus (a int, b int)(int, int){
// 	sum := 0
// 	arc := 0
// 	for i := a; i<=b; i++{
// 		sum+=i
// 		arc += i*2
// 	}
// 	// fmt.Println(sum)
// 	return sum,arc
// }

// func plus(arg ...int){
// 	sum := 0 
// 	for i:=arg[0]; i<=arg[1]; i++{
// 		sum+=i
// 	}
// 	fmt.Println(sum)
// }

// func cha1(cha [4]int){
// 	cha[0] = 100
// }
// func cha2(cha []int){
// 	cha[0] = 100
// }