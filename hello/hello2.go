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

	fmt.Printf("NimCPU:=%d\n",runtime.NumCPU())
	fmt.Printf("NumGoroutine:=%d\n",runtime.NumGoroutine())
	fmt.Printf("Version:=%s\n",runtime.Version())
	//NimCPU:=8
    //NumGoroutine:=1
	//Version:=go1.12.6

	fmt.Println(I)
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

