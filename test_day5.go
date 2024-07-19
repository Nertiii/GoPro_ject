package main
import "fmt"
func main(){
	// m := Mouse{"luoji"}
	// Test(m)

	// var t1 Trian = Trian{3,4,5}
	// var t2 Ant = Ant{4}
	// fmt.Println(t1.Zhou())
	// fmt.Println(t2.Zhou())
	go print()
	for i:=1; i<=100; i++{
		fmt.Print("主ggggot打印：%d\n",i)
	}

	
}

func print(){
	for i:=1; i<=100; i++{
		fmt.Print("子got打印：%d\n",i)
	}
	
}

//定义一个接口
// type Shape interface{
// 	Zhou() int
// 	Mian() int
// }
// type Trian struct{
// 	a int
// 	b int
// 	c int
// }
// func (t Trian)Zhou() int{
// 	return t.a + t.b +t.c
// }
// func (t Trian)Mian() int{
// 	return t.a * t.b *t.c
// }

// type Ant struct{
// 	bian int
// }

// func (a Ant)Zhou() int{
// 	return a.bian * 4
// }

// func (a Ant)Mian() int{
// 	return a.bian*a.bian
// }



// type USB interface{
// 	start()
// 	end()
// }

// type Mouse struct{
// 	name string
// }

// func Test(u USB){
// 	u.start()
// 	u.end()
// }

// func (m Mouse)start(){
// 	fmt.Println(m.name,"aaa")
// }

// func (m Mouse)end(){
// 	fmt.Println(m.name,"bbb")
// }

