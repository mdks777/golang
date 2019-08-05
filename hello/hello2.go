package main

import (
	"fmt" //f"fmt"とすればfmtとはかかずfだけでいい
	//"."にするとpackage無名でいける
	"runtime"
)

func main() {
	var xy int
	xy = plusAlias(1, 2)
	fmt.Printf("\n%d\n", xy)
	q, r := div(19, 7)
	fmt.Printf("商=%d剰余=%d\n", q, r)
	hello()
	fmt.Println(dosomething()) //5 0

	f := func(x, y int) int { return x + y }
	fmt.Println(f(2, 3))                                           // 5
	fmt.Printf("%T\n", func(x, y int) int { return x + y })        //func(int, int) int
	fmt.Printf("%#v\n", func(x, y int) int { return x + y })       //(func(int, int) int)(0x4813a0)
	fmt.Printf("%#v\n", func(x, y int) int { return x + y }(2, 3)) // 5
	returnFunc()()                                                 //I am a function
	ff := returnFunc()
	ff() //I am a function
	callFunction(func() {
		fmt.Printf("I am a function\n") //I am a function
	})

	fa := later()
	fmt.Printf(fa("Golang\n")) //
	fmt.Printf(fa("is\n"))     //Golamg
	fmt.Printf(fa("God\n"))    //is

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	one, two, x, y, z, w, s1, s2, s := one()
	fmt.Printf("x=%d,y=%d,%f,%f,%f,%f,%s,%s,%s", one, two, x, y, z, w, s1, s2, s)
	あいさつ(昼の挨拶)

	fmt.Println(appName())
	//fmt.Println(AppName +" " + Version) -> error
    
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
            continue
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
	default:
		fmt.Printf("unknown!\n")
	}//1or2
	switch n {
	case 1,2:
		fmt.Printf("%d\n",n)//2
		n+=5
		fallthrough//無条件で次の項目
	case 3,4:
		fmt.Printf("%d\n",n)//7
		n+=3
		
	case 5,6,7:
		fmt.Printf("%d\n",n)//fallthroughされていないため条件でも不適応
		n+=6
	}
	var a interface{} =3
	ai := a.(int)
	//af := a.(float64)->error
	fmt.Printf("%d\n",ai)
	fmt.Printf("%d\n",n)//10
    var b interface{} = 3.14
	bi, isInt := b.(int)
	_, IsInt := b.(int)
	bf, isFloat64 := b.(float64)
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
	fmt.Printf("\n%T\n",d)
	
	fmt.Printf("A")
	goto A
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
					break LOOP
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
				continue L
			}
			fmt.Printf("%d*%d=%d\n",i,j,i*j)
		}
		fmt.Printf("unsee")
	}
	
	defer func() {
		defer fmt.Printf("1")
		defer fmt.Printf("2")
		defer fmt.Printf("\n3")
	}()//一番最後（文末）に321
	runDefer()//done321
	/*go func() {
		for {
			fmt.Printf("sub loop")
		}
	}()//並列処理
	for {
		fmt.Printf("main loop")
	}*/
	//-> "sub loop"と"main loop"が不規則にループ

	fmt.Printf("NimCPU:=%d\n",runtime.NumCPU())
	fmt.Printf("NumGoroutine:=%d\n",runtime.NumGoroutine())
	fmt.Printf("Version:=%s\n",runtime.Version())
	//NimCPU:=8
    //NumGoroutine:=1
	//Version:=go1.12.6

	fmt.Println(I)

	ss := make([]int,5,10)
	for i:=0;i<5;i++ {ss[i]=i}
	fmt.Println(ss,len(ss),cap(ss))//[0 1 2 3 4] 5 10
	fmt.Println(ss[2:4])//[2 3]
	//スライス名[該当数字(無ければ[0])以上：該当数字(なければ[len(変数名)])未満]
	Ss:=make([]int,4)
	for i:=5;i<9;i++{Ss[i-5]=i}
	fmt.Println(append(ss,Ss...)) //[0 1 2 3 4 5 6 7 8]

	sS:=make([]int,0,0)
	fmt.Println(len(sS),cap(sS))// 0 0
	sS=append(sS,1)
	fmt.Println(len(sS),cap(sS))// 1 2
	sS=append(sS,[]int{2,3,4}...)
	fmt.Println(len(sS),cap(sS))// 4 4
	sS=append(sS,5)
	fmt.Println(len(sS),cap(sS))// 5 8
	sS=append(sS,6,7,8,9)
	fmt.Println(len(sS),cap(sS))// 9 16
	//capは足りなくなったら2のｎ乗になる,とも限らんらしい
	SS:=make([]int,1024,1024)
	fmt.Println(len(SS),cap(SS))//1024 1024
	SS=append(SS,0)
	fmt.Println(len(SS),cap(SS))//1025 1344
	//ちなみに本では1312だった	

	fmt.Println(copy(ss,Ss),ss) //4 [5 6 7 8 4]
	//変更した個数が値として返還され最初のスライスが変更される
	//二番目のほうが要素数が多い場合最初のスライスがすべて変更
	
	sss:= [10]int {}
	for i:=0;i<10;i++{sss[i]=i}
	fmt.Println(sss)
	ss1:=sss[2:4]
	fmt.Println(ss1,len(ss1),cap(ss1))//[2 3] 2 8
	ss2:=sss[2:4:4]
	fmt.Println(ss2,len(ss2),cap(ss2))//[2 3] 2 2
	ss3:=sss[2:4:6]
	fmt.Println(ss3,len(ss3),cap(ss3))//[2 3] 2 4
	//n:=s[a:b:c]
	//n==[ s[a]<= X <s[b] ]
	//len(n)==b-a
	//cap(n)==c-a
	//[要素and容量の最小値：要素の最大値：容量の最大値]

	sas :=[]string{"a","b","c"}
	for i, r:=range sas {
		fmt.Printf("%d,%s ",i,r)
	}//0,0 1,1 2,2 3,3 4,4 5,5 6,6 7,7 8,8 9,9
	
	for i:=0;i<len(sas);i++ {
		fmt.Printf("%d,%s ",i,sas[i])
		sas = append(sas, "d")//要素の追加
		if i == 10 {break}
	}//0,a 1,b 2,c 0,a 1,b 2,c 3,d 4,d 5,d 6,d 7,d 8,d 9,d 10,d
	//途中で追加しても結局は最後尾に追加される	

	fmt.Println(sum(1,2,3),sum(1,2,3,4,5),sum())//6 15 0
	
	a1:=[3]int{1,2,3}
	a2:=[]int{1,2,3}
	pow1(a1)//値をコピーしただけだからmainのa1は動いていない
	pow2(a2)//とりま、動くらしい
	fmt.Println(a1,a2)//[1 2 3] [1 4 9]

}


func plus(x, y int) int {
	return x + y
}

var plusAlias = plus

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func hello() {
	fmt.Printf("Hello,World\n")
}

func dosomething() (int, int) {
	var x, y int
	x = 5
	return x, y
}

func returnFunc() func() {
	return func() {
		fmt.Printf("I am a function\n")
	}
}

func callFunction(f func()) {
	f()
}

func later() func(string) string {
	//ひとつ前に与えられた文字列を保存するための変数
	var store string
	//引数に文字列をとり文字列を返す関数を返す
	return func(next string) string {
		s := store
		store = next //クロージャ内（return？）にいるためローカル変数で消えるのではなく残る
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

const (
	X  = float64(3.5) + iota //3.5+0
	Y                        //3.5+1
	Z                        //3.5+2
	W  = X + Y + Z
	S1 = "なんて日だ" //なんて日だ
	S2           //なんて日だ
	S  = S1 + "は" + S2
)
const One = 1

func one() (int, int, float64, float64, float64, float64, string, string, string) {
	const Two = 2
	return One, Two, X, Y, Z, W, S1, S2, S
}

const (
	朝の挨拶 = "おはよう"
	昼の挨拶 = "こんにちは"
	夜の挨拶 = "こんばんは"
)

func あいさつ(m string) {
	fmt.Println(m)
}

func appName() string {
	const AppName = "Go Application"
	var Version = "1.0"
	return AppName + " " + Version
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

func sum(s ...int) int{
	n:=0
	for _,v:=range s {
		n+=v
	}
	return n
}

func pow1(a1 [3]int){
	for i,v := range a1 {
		a1[i]=v*v
	}
	return
}
func pow2(a2 []int){
	for i,v := range a2 {
		a2[i]=v*v
	}
	return
}