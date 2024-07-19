package main
import "fmt"
func main(){
	// //匿名函数
	// fun1()

	// func (){
	// 	fmt.Println("这是一个匿名函数")
	// }()

	// func (a,b int){
	// 	fmt.Println(a,b)
	// }(1,2)

	//回调函数
	// re1 := ope(2,4,add)
	// fmt.Println(re1)

	// re2 := ope(2,4,sub)
	// fmt.Println(re2)

	//指针数组
	// arr := [4] int {1,2,3,4}
	// fmt.Printf("数组：%d\n",arr)
	// arr1(arr)
	// fmt.Printf("最终1：%d\n",arr)
	// arr2(&arr)
	// fmt.Printf("最终2：%d\n",arr)

	//结构体
	// var p1 Person
	// p1.name = "nerti"
	// p1.age = 26
	// p1.sex = "male"
	// p1.address = "chengdu"
	// fmt.Println(p1)

	// p1 := Person{}
	// p1.name = "nerti"
	// p1.age = 26
	// p1.sex = "male"
	// p1.address = "chengdu"
	// fmt.Println(p1)

	//结构体 接口
	m1 := Mouse{"罗技"}
	fmt.Println(m1.name)
	test(m1)

	f1 := Flash{"金士顿"}
	fmt.Println(f1.name)
	test(f1)
	

}
//定义接口
type USB interface{
	start()
	end()
}
//实现类
type Mouse struct{
	name string
}

type Flash struct{
	name string
}

func (m Mouse)start(){
	fmt.Println(m.name,"鼠标开始工作啦！")
}

func (m Mouse)end(){
	fmt.Println(m.name,"鼠标结束工作咯~")
}

func (f Flash)start(){
	fmt.Println(f.name,"键盘开始工作啦！")
}

func (f Flash)end(){
	fmt.Println(f.name,"键盘结束工作咯~")
}

//测试方法
func test(usb USB){
	usb.start()
	usb.end()
}
// ------------------------------------------

// type Person struct{
// 	name string
// 	age int
// 	sex string
// 	address string
// }


// func arr1(arr1 [4]int){
// 	fmt.Printf("调用fun1的数组：%d\n",arr1)
// 	arr1[0] = 1000
// 	fmt.Printf("调用fun1后的数组：%d\n",arr1)
	
// }

// func arr2(arr2 *[4]int){
// 	fmt.Printf("调用fun2的数组：%d\n",arr2)
// 	arr2[0] = 1000
// 	fmt.Printf("调用fun2后的数组：%d\n",arr2)
	
// }

// func add(a,b int)int{
// 	return a+b
// }
// func sub(a,b int)int{
// 	return a-b
// }

// func ope(a,b int, fun func(int,int)int)int{
// 	res := fun(a,b)
// 	return res
// }

// func fun1(){
// 	fmt.Println("fun1")
// }




