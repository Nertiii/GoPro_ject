package main
import "fmt"
//import "math/rand"
//import "time"
func main(){
	// for i:=58; i>=23; i--{
	// 	fmt.Println(i)
	// } 

	// sum := 0 
	// for i := 1; i<=100; i++{
	// 	sum += i
	// }
	// fmt.Println(sum)

	// count := 0
	// for i:=1; i<=100; i++{//被3整除且不能被5整除的数，每5个换一次行
	// 	if i % 5 != 0 && i % 3 == 0{
	// 		count++
	// 		fmt.Print(i,"\t")
	// 		if count % 5 == 0{
	// 			fmt.Print("\n")
	// 		}
			
	// 	}
	// } 

	// count := 0
	// for i:=1; i<=25; i++{ //5星好评
	// 	count++
	// 	fmt.Print("*")
	// 	if count % 5 == 0{
	// 		fmt.Println()
	// 	}

	// }

	// for i:=1; i<=9; i++{//99乘法表
	// 	for j:=1 ; j<=i ; j++{
	// 		fmt.Printf("%dX%d=%d\t",j,i,i*j)
	// 	}
	// 	fmt.Println()
	// }

	//求素数
	// for i:=2; i<=100; i++{
	// 	flag := true
	// 	for j:=2; j<i; j++{
	// 		if i%j==0{
	// 			flag = false
	// 			break
	// 		}
	// 	}
	// 	if flag == true {
	// 		fmt.Println(i)
	// 	}
		
	// }

	//随机数
	// rand.Seed(1)
	// i := rand.Intn(10)
	// j := rand.Int()
	// fmt.Println(i)
	// fmt.Println(j)

	// //时间戳
	// ti := time.Now()
	// fmt.Println(ti)
	// tj := ti.Unix()
	// fmt.Println(tj)
	
	//数组
	// var arr1 [5] int
	// for i:=0; i<5; i++{
	// 	arr1[i]=i
	// }

	// arr2 := [5] int {1,2,3,4,5}

	// fmt.Println(arr1)
	// fmt.Println(arr2)

	// for i,j := range arr2{
	// 	fmt.Printf("数组下标是：%d,数组内容为：%d\n",i,j)
	// }

	// sum := 0
	// for _,v := range arr2{
	// 	sum+=v
	// }
	// fmt.Println(sum)

	//冒泡排序
	// arr1 := [5] int {15, 12, 9, 10, 8}
	// fmt.Printf("数组为：%d\n",arr1)
	// for i:=1 ; i<=len(arr1)-1; i++{
	// 	for j:=0; j<len(arr1)-1 ; j++{
	// 		//num:=0
	// 		if(arr1[j]>arr1[j+1]){
	// 			//num = arr1[j+1]
	// 			arr1[j], arr1[j+1] = arr1[j+1], arr1[j]
	// 		}
	// 	}
	// 	fmt.Printf("第%d轮冒泡排序结果为：%d\n",i,arr1)
	// }

	//二维数组
	// arr := [3][4] int {{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	// fmt.Println(arr)

	// for i:=0 ; i<len(arr); i++{
	// 	for j:=0 ; j<len(arr[i]); j++{
	// 		fmt.Printf("第%d-%d个数组元素为：%d\n",i,j,arr[i][j])
	// 	}
	// }

	//切片indentifier
	// arr := []int{123,456,789}
	// fmt.Println(arr)
	// arr=append(arr,147,258,369)
	// fmt.Println(arr)
	// fmt.Printf("%T\n",arr)

	// arr2 := [] int {1,2,3}
	// arr = append(arr,arr2...)
	// fmt.Println(arr)

	// arr := [10] int {1,2,3,4,5,6,7,8,9,10}
	// fmt.Printf("%T\n",arr)
	// a1 := arr[:]
	// a2 := arr[:4]
	// a3 := arr[5:]
	// a4 := arr[4:8]
	// fmt.Printf("%T\n",a1)
	// fmt.Println(arr)
	// arr[3]=100
	// fmt.Println(a1)
	// fmt.Println(a2)
	// fmt.Println(a3)
	// fmt.Println(a4)

	//map
	// var arr = map[string]int{"Java":100,"C语言":98}
	// fmt.Println(arr)
	// arr["Java"]=1
	// fmt.Println(arr)
	// var map1 = make(map[int]string)
	// map1[1]="123"
	// fmt.Println(map1)

	// fmt.Println("________________________________")
	// map1 := make(map[int]string)

	// map1[1]="内蒙古自治区"
	// map1[2]="呼伦贝尔市"
	// map1[3]="莫力达瓦达斡尔族自治旗"
	// map1[4]="尼尔基镇"
	// map1[5]="纳文西大街"
	// map1[6]="一中家属楼"

	// for k,v := range map1{
	// 	fmt.Println(k,v)
	// }
	// fmt.Println("________________________________")
	// for i:=1;i<=len(map1);i++{
	// 	fmt.Println(map1[i])
	// }
	// fmt.Println("________________________________")

	map1 := map[string]string{"name":"nerti1","sex":"male","age":"26","homeland":"inner mongolia"}
	map2 := map[string]string{"name":"nerti2","sex":"female","age":"20","homeland":"peking"}
	map3 := map[string]string{"name":"nerti3","sex":"bisexual","age":"23","homeland":"chengdu"}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	all := make([]map[string]string,0,3)
	all = append(all,map1)
	all = append(all,map2)
	all = append(all,map3)

	for i,val := range all{
		fmt.Printf("第%d个人的信息为：",i)
		fmt.Printf("name:%s\t",val["name"])
		fmt.Printf("sex:%s\t",val["sex"])
		fmt.Printf("age:%s\t",val["age"])
		fmt.Printf("homeland:%s\n",val["homeland"])
	}
	//fmt.Println(all)


}