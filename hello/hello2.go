package main

import (
	"fmt" //f"fmt"とすればfmtとはかかずfだけでいい
	//"."にするとpackage無名でいける
	"runtime"
)

func main() {
	var xy int
	xy = plusAlias(1, 2)
	fmt.Printf("\n%d\n", xy)//3
	q, r := div(19, 7)
	fmt.Printf("商=%d剰余=%d\n", q, r)//商=2剰余=5
	hello()//Hello,World
	fmt.Println(dosomething()) //5 0
	/*func plus(x, y int) int {
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
    }*/
	

	f := func(x, y int) int { return x + y }
	fmt.Println(f(2, 3))                                           // 5
	fmt.Printf("%T\n", func(x, y int) int { return x + y })        //func(int, int) int
	fmt.Printf("%#v\n", func(x, y int) int { return x + y })       //(func(int, int) int)(0x4813a0)
	fmt.Printf("%#v\n", func(x, y int) int { return x + y }(2, 3)) // 5
	returnFunc()()//I am a function
	ff := returnFunc()
	ff() //I am a function
	callFunction(func() {
		fmt.Printf("I am a function\n") //I am a function
	})
	/*func returnFunc() func() {
		return func() {
			fmt.Printf("I am a function\n")
		}		//return値として関数を送ってmainで実行
	}
	func callFunction(f func()) {
		f()　　　//関数を引数としてcallFunction内のfに移してfを内部で実行
	}*/

	fa := later()
	fmt.Printf(fa("Golang\n")) //
	fmt.Printf(fa("is\n"))     //Golamg
	fmt.Printf(fa("God\n"))    //is
	/*func later() func(string) string {
		var store string	//ひとつ前に与えられた文字列を保存するための変数
		return func(next string) string {	//引数に文字列をとり文字列を返す関数を返す
			s := store
			store = next //クロージャ内（return？）にいるためローカル変数で消えるのではなく残る
			return s	//最初ｓには何も入っていないので一個後になる
		}
	}*/
	ints := integers()
	fmt.Println(ints())//1
	fmt.Println(ints())//2
	fmt.Println(ints())//3
	/*func integers() func() int {
		i := 0
		return func() int {
			i++
			return i
		}//クロージャ内（return？）にいるためローカル変数で消えるのではなく残る
	}*/
	
	one, two, x, y, z, w, s1, s2, s := one()
	fmt.Printf("x=%d,y=%d,%f,%f,%f,%f,%s,%s,%s", one, two, x, y, z, w, s1, s2, s)
	あいさつ(昼の挨拶)
	/*x=1,y=2,3.500000,4.500000,5.500000,13.500000,なんて日だ,なんて日だ,なんて日だは 
	なんて日だこんにちは*/
	/*const (
		X  = float64(3.5) + iota //3.5+0
		Y                        //3.5+1　　何も書かなければ継続
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
		fmt.Println(m)//昼の挨拶という変数の中の”こんにちは”
	}*/

	fmt.Println(appName())//fmt.Println(AppName +" " + Version) -> error
	//Go Application 1.0
	/*func appName() string {
		const AppName = "Go Application"
		var Version = "1.0"
		return AppName + " " + Version
	}*/
    
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


//ここから！



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
	fmt.Println(sss)//[0 1 2 3 4 5 6 7 8 9]
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
	}//0,a 1,b 2,c
	fmt.Printf("\n")

	for i:=0;i<len(sas);i++ {
		fmt.Printf("%d,%s ",i,sas[i])
		sas = append(sas, "d")//要素の追加
		if i == 10 {break}
	}//0,a 1,b 2,c 3,d 4,d 5,d 6,d 7,d 8,d 9,d 10,d
	//途中で追加しても結局は最後尾に追加される	
	fmt.Printf("\n")
	
	asa:=[]int{1,2,3,4,5}
	fmt.Println(sum(1,2,3),sum(asa...),sum(),sum(asa[2:4]...))//6 15 0 7
	/*func sum(s ...int) int{
		n:=0
		for _,v:=range s {
			n+=v
		}
		return n
	}*/
	a1:=[3]int{1,2,3}
	a2:=[]int{1,2,3}
	pow1(a1)//値をコピーしただけだからmainのa1は動いていない
	pow2(a2)//とりま、動くらしい
	fmt.Println(a1,a2)//[1 2 3] [1 4 9]
	/*func pow1(a1 [3]int){
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
	}スライスは関数間を行き来できる*/

	ab:=[3]int{0,1,2}
	aba:=ab[:]//ab(配列)とaba(スライス)は共有している
	fmt.Println(len(aba))//3
	fmt.Println(cap(aba))//3
	ab[0]=9
	fmt.Println(ab,aba)//[9 1 2] [9 1 2]ー＞abを変えたら両方
	
	ab=[3]int{0,1,2}//リセット
	aba=append(aba,3)//要素の追加により別のメモリ上の配列にさし代わる
	fmt.Println(len(aba))//4
	fmt.Println(cap(aba))//8
	ab[0]=9
	fmt.Println(ab,aba)//[9 1 2] [0 1 2 3]ー＞abを変えてもabaに影響がない

	m:=make(map[int]string)//var m map[int]string
	m[1]="US"
	m[81]="JAPAN"
	m[83]="CHAINA"
	///m:=map[int]string{1:"US", 81:"JAPAN", 83:"CHAINA"}
	fmt.Println(m)//map[1:US 81:JAPAN 83:CHAINA]
	m1:=make(map[string]string)
	m1["Yamada"]="Taro"
	m1["Sato"]="Taichi"
	m1["Yamada"]="Jiro"
	fmt.Println(m1)//map[Sato:Taichi Yamada:Jiro]
		m2:=make(map[float64]int)
	m2[1.000000000000000000000001]=1
	m2[1.000000000000000000000002]=2
	m2[1.000000000000000000000003]=3
	fmt.Println(m2)//map[1:3]
	ma:=map[int][]int{//キーがintで要素が[]int
		1: {1},//[]int{1}と同じ
		2: {1,2},
		3: {1,2,3},
	}
	fmt.Println(ma)//map[1:[1] 2:[1 2] 3:[1 2 3]]
		ma1:=map[int]map[float64]string{
		1: {3.14: "円周率"},
		2: {1.414: "ルート２"},
	}
	fmt.Println(ma1)//map[1:map[3.14:円周率] 2:map[1.414:ルート２]]
		ma2:=map[int]string{1:"A",2:"B",3:"C"}
	fmt.Println(ma2[2],ma2[6])//B ""(空白)
	mm,ok:=ma2[1]
	fmt.Println(mm,ok)//A true
	mm,ok=ma2[7]
	fmt.Println(mm,ok)//（空白） false 
	_,ok=ma2[3]
	fmt.Println(ok)//true
	if _,ok =ma2[1];ok{
		fmt.Printf("ma2[1]の存在確認")
	}//作動
	ma2[4]="D"
	ma2[6]="F"
	ma2[5]="E"
	for k,v:=range ma2{
		fmt.Println(k,v)
	}//キーの順番で表示されるかどうか「不定」毎回変わる
	delete(ma2,2)
	fmt.Println(ma2)//map[1:A 3:C 4:D 5:E 6:F]
	//m:=make(map[int]string,100)->ヒント的な奴。スライスの容量の簡易版的な？

		ch := make(chan int, 10)//バッファサイズ10のチャネル、容量みたいなもん（超えたらエラー）
	//チャネルには「先入れ先出し」「データの順序が保証される」をもつ
	ch <- 5//チャネルに５を送信
	ch1 := <- ch//チャネルから値を受信
	fmt.Println(ch1)//5
		ch2:=make(chan int)
	go receiver(ch2)
	for i:=0;i<10;i++{
		ch2<-i
	}
		ch3:=make(chan string, 3)
	ch3<-"APPLE"
	fmt.Println(len(ch3),cap(ch3))//1 3
		ch4:= make(chan int, 5)
	go receive("1st goroutine",ch4)
	go receive("2nd goroutine",ch4)
	go receive("3rd goroutine",ch4)
	for i:=0;i<20;i++{
		ch <- i
	}
	close(ch4)
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

func receiver(ch2 <-chan int){//chは送信専用のチャネル
	for {
		i:=<-ch2
		fmt.Println(i)
	}//0123456789
}

func receive(name string, ch4 <-chan int) {
	for {
		i,ok:=<-ch4
		if ok == false {
			break
		}
		fmt.Println(i)
	}
	fmt.Println(name + " is done")
}