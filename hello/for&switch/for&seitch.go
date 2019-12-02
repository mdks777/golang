package main

import (
	"fmt"
)

func main() {
	i:=0
	for{
        //無限ループ
        fmt.Printf("%d,",i)
        i++
        if i==15 {
            fmt.Printf("\n")
            break
        }
    }//0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,
    i=0
    for i<15{
        if i==4 || i==9{
            i++
            continue//飛ばす
        }
        fmt.Printf("%d,",i)
        i++
    }//0,1,2,3,5,6,7,8,10,11,12,13,14,
    fmt.Printf("\n")
	for i := 0; i < 15; i++ {
		if i % 7==0 && i!=0 {
			fmt.Printf("luckly!")
		} else {
			fmt.Printf("%d,", i)
		}
    } //0,1,2,3,4,5,6,luckly!8,9,10,11,12,13,luckly!
    
    fmt.Printf("\n")
	xx, yy := 5, 6
	if xx, yy := 1, 2; xx < yy {
		fmt.Printf("\n%d<%d\n", xx, yy) //1<2
	}
	fmt.Printf("%d<%d\n", xx, yy) //5<6

	fruits := [3]string{"Apple","Banana","Cherry"}
	for i,s:= range fruits {
		fmt.Printf("fruits%d=%s\n",i,s)
	}//fruits0=Apple
	//fruits1=Banana
	//fruits2=Cherry
	for i,r := range "ABC"{
		fmt.Printf("%d=%d\n",i,r)
	}
	//0=65
	//1=66
	//2=67
	for i,r:=range "あいうえお"{
		fmt.Printf("%d=%d\n",i,r)
	}//0=12354
    //3=12356
    //6=12358
	//9=12360
	//12=12362

	n:=2
	switch n {
	case 1,2:
		fmt.Printf("1or2\n")
	case 3,4:
		fmt.Printf("3or4\n")
	default://以外
		fmt.Printf("unknown!\n")
	}//1or2
	switch n {
	case 1,2:
		fmt.Printf("%d\n",n)//2
		n+=5
		fallthrough//無条件で次のcase
	case 3,4:
		fmt.Printf("%d\n",n)//7
		n+=3

	case 5,6,7:
		fmt.Printf("%d\n",n)//case3.4でfallthroughされていないため条件でも不適応
		n+=6
	}
	var a interface{} =3
	ai := a.(int)
	//af := a.(float64)->error
	fmt.Printf("%d\n",ai)//3
    var b interface{} = 3.14
	bi, isInt := b.(int)//3.14はfloat64なのでfalse
	_, IsInt := b.(int)
	bf, isFloat64 := b.(float64)//bfには数値が入る
	bs, isString := b.(string)
	fmt.Printf("%d,%v,%v,%f,%v,%s,%v\n",bi,isInt,IsInt,bf,isFloat64,bs,isString)
	//0,false,false,3.140000,true,,false
	var c interface{} =5
	switch c.(type) {
	case bool:
		fmt.Printf("bool")
	case int, uint:
		fmt.Printf("int or uint")
	case string :
		fmt.Printf("string")
	default:
		fmt.Printf("unknown")
	}//int or uint
	var d interface{} = "hello"
	fmt.Printf("\n%T\n",d)//string
	
	fmt.Printf("A")
	goto A //Aへ飛ぶ
B:	fmt.Printf("B")
	goto C//Cがなければ無限ループ
A:  fmt.Printf("C")
	goto B
C:	fmt.Printf("D\n")
	//ACBD　　関数間は行き来出来ない
LOOP:
	for {
		for{
			for{
				for{
					fmt.Printf("DONE ")
					break LOOP//LOOPのfor全てをbreak
				}
				fmt.Printf("A")
			}
			fmt.Printf("B")
		}
		fmt.Printf("C")
	}
	fmt.Printf("D\n")
	//DONE D
L:	
	for i:=1;i<=3;i++{
		for j:=1;j<=3;j++{
			if j>1{
				continue L//Lのforを一回分飛ばすー＞ｊは動かないがiは動く
			}
			fmt.Printf("%d*%d=%d\n",i,j,i*j)//continueはforが終わってから効力をだす
		}
		fmt.Printf("unsee")//一回分のforがcontinueで飛ばされるからここは表示されない
	}//1*1=1
	//2*1=2
	//3*1=3
	
	defer func() {
		defer fmt.Printf("1")
		defer fmt.Printf("2")
		defer fmt.Printf("3")
	}()//無名関数であるため()は必須
	runDefer()//done321
	/*func runDefer() {
		defer fmt.Printf("1\n")
		defer fmt.Printf("2")
		defer fmt.Printf("3")
		fmt.Printf("done")
	}*/

	/*go func() {
		for {
			fmt.Printf("sub loop")
		}
	}()//無名関数で並列処理
	for {
		fmt.Printf("main loop")
	}-> "sub loop"と"main loop"が不規則にループ*/


}

func runDefer() {
	defer fmt.Printf("1\n")
	defer fmt.Printf("2")
	defer fmt.Printf("3")
	fmt.Printf("done")
}

var I = ""
func isit() {
	I =I+"A"
}
func init() {
	I =I+"B"
}
func init() {
	I =I+"C"
}