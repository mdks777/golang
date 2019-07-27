package main



import (
    "fmt"
    "math"
)

func one() int{
    return 1
}
//nnはパッケージ変数
var nn =100

func main() {
	fmt.Println("Hello,World!")//改行付き
	fmt.Print("Hello,World!")
	println("Hello World!")//改行付き
    print("Hello World!")
    fmt.Println(1,2,3)//改行付き1 2 3
	fmt.Print(1,2,3)//1 2 3
	println(1,2,3)//1 2 3
    print(1,2,3)//123
    
    fmt.Printf("数値＝%v 文字列＝%v 文字列＝%v\n",5, "Golang", [...]int{1, 2, 3})
    //%vはさまぁまな型のデータを埋め込む
    //数値＝5　文字列＝Golang 配列＝[1 2 3]
    fmt.Printf("数値＝%#v 文字列＝%#v 文字列＝%#v\n",5, "Golang", [...]int{1, 2, 3})
    //%#vはgoのリテラル表現のデータを埋め込む
    //数値＝5　文字列＝"Golang" 配列＝[3]int{1, 2, 3}
    fmt.Printf("数値＝%T 文字列＝%T 文字列＝%T\n",5, "Golang", [...]int{1, 2, 3})
    //%Tはデータの型情報を埋め込む
    //数値＝int　文字列＝string 配列＝[3]int

    var (
        //intのx,yとstringのnameを定義
		x, y int
		name string
    )
    //xとyに1,2が代入されている
	x, y = 1, 2
	name = "tatata"
    fmt.Printf("%d, %d, %s",x, y, name)
    //1, 2, tatata

    //bool型の変数を使ってtrueを代入
    b := true
    //int型の変数を使って１を代入
    i := 1
    //float型の変数を使って3.14を代入
    f := 3.14
    //string型の変数を使って"abc"を代入
    s := "abc"
    fmt.Printf("\n%v, %d, %f, %s\n",b, i, f, s)
    //true, 1, 3.140000, abc

    n := one()
    fmt.Printf("%d\n",n)
    //1

    nn += 1
    fmt.Printf("n=%d\n",nn)
    //n=101

    aa := 256 //int型
    bb := byte(aa)//ｂｂにaaの数字をbyte型に置き換え代入
    cc := int64(aa)//ccにaaの数字をint64型に置き換え代入
    dd := uint32(aa)//ddにaaの数字をuint32型に置き換え代入
    fmt.Printf("%d, %d, %d, %d\n", aa, bb, cc, dd)
    //256, 0. 256, 256
    //byte型は256項目あるけど0から数えてるから数値の最大数は２５５
    ee := uint32(math.MaxUint32)
    fmt.Printf("%d\n", ee)//4294967295
    ff := 1.0//float64型(基本はこっち)
    gg := float32(ff)//float32型(不精密かつ大量計算にむいてる)
    fmt.Printf("%f, %f\n",ff, gg)

    r := '松'
    fmt.Printf("%v\n", r)
    //'A' == '[\]101'
    //'A' == '[\x]41'
    //'松' == '[\u]677E'
    //'群' == '[\U]000298C6

    fmt.Println("aa \a ab \b ac \f ad \n ")
    fmt.Println("ae \r af \t ag \v ah \\ ai")
    //aa  ab ac  ad
    //
    //af      ag  ah \ ai
    //効果音（ベル）あり

    ss := "GO言語の文字列"
    fmt.Printf("\n%v\n", ss)//GO言語の文字列

    aaa := [5]int{1, 2, 3, 4, 5}
    fmt.Printf("%v, %d, %d, %d\n", aaa, aaa[0], aaa[1], aaa[2])
    //[1 2 3 4 5], 1, 2, 3
    bbb := [5]int{}
    ccc := [5]int{1, 2, 3}
    fmt.Printf("%v\n%v\n",bbb, ccc)
    //[0 0 0 0 0]
    //[1 2 3 0 0]
    ia := [3]int{}
    ua := [3]uint{}
    ba := [3]bool{}
    fa := [3]float64{}
    ra := [3]rune{}
    sa :=[3]string{}
    fmt.Printf("%v,%v,%v,%v,%v,%v\n",ia, ua, ba, fa, ra, sa)
    //[0 0 0],[0 0 0],[false false false],[0 0 0],[0 0 0],[  ]
    aaa[0] = 0
    aaa[2] = 0
    //ia != ua 型が違うため。要素数が同じでないとダメ
    a1 := [3]int{1, 2, 3}
    a2 := [3]int{4, 5, 6}
    a1 = a2
    fmt.Printf("%v\n",a1)//[4 5 6]
    a1[0]=0
    a1[2]=0
    fmt.Printf("%v,%v\n",a1, a2)//[0 5 0],[4 5 6]

    var xx interface{}
    fmt.Printf("%#v\n",xx)//<nil>（具体的な数字をもっていない）
    xx =1
    xx =3.14
    xx ="松"
    xx ="松竹梅"
    xx =[...]int{1, 2, 3, 4}
    fmt.Printf("%v\n",xx)
    /*var yy interface{}
    xx =4
    yy =5
    fmt.Printf("%d",xx+yy)
    エラー（演算の対象ではない）
    */

    /*
    優先度
    * / % << >> & &^
    + - | ^
    == != < <= >=
    &&
    ||
    */
    fmt.Printf("%d,%d,%d,%d,%d\n",4+2,5-2,2*3,4/2,7%3)
    //6,3,6,2,1
    fmt.Printf("%d\n",165&155)
    /* 129
    165 == 1010 0101
    155 == 1001 1011
    129 == 1000 0001
    */
    fmt.Printf("%d\n",197|169)
    /*237
    197 == 1100 0101
    169 == 1010 1001
    237 == 1110 1101
    */
    fmt.Printf("%d\n",92^137)
    /*213
    92  == 0101 1100
    137 == 1000 1001
    213 == 1101 0101
    */
    fmt.Printf("%d\n",108&^13)
    /*96
    108 == 0110 1100 -> 0110 1100
    13  == 0000 1101 -> 1111 0010
    96               == 0110 0000
    */
    as := byte(42)
    fmt.Printf("%d,%d,%d,%d,%d\n",as,as<<1,as<<2,as<<3,as<<4)
    /*42,84,168,80,160
    42  == 0010 1010
    84  == 0101 0100
    168 == 1010 1000
    80  == 0101 0000
    160 == 1010 0000
    */
    as =byte(168)
    fmt.Printf("%d,%d,%d,%d,%d\n",as,as>>1,as>>2,as>>3,as>>4)
    /*168,84,42,21,10
    168 == 1010 1000
    84  == 0101 0100
    42  == 0010 1010
    21  == 0001 0101
    10  == 0000 1010
    */
    fmt.Printf("%v,%v,%v,%v\n",1==1,2<5,6!=4,3>=7)
    //true,true,true,false


    iia :=5
    iib :=5
    if !(iia!=4 && iib!=6){
        fmt.Printf("true or false")
    }//「両方不正解」のみfalse
    if!(iia!=4 || iib!=6){
        fmt.Printf("true and true")
    }//[両方正解]のみtrue
    if(iia!=4 && iib!=6){
        fmt.Printf("false and false")
    }//[両方不正解]のみtrue
    if(iia!=4 || iib!=6){
        fmt.Printf("false or true")
    }//[両方正解]のみfalse

    if !(iia==4 && iib==6){
        fmt.Printf("true or false")
    }//「両方正解」のみfalse
    if!(iia==4 || iib==6){
        fmt.Printf("false and false")
    }//[両方不正解]のみtrue
    if(iia==4 && iib==6){
        fmt.Printf("true and true")
    }//[両方正解]のみtrue
    if(iia==4 || iib==6){
        fmt.Printf("false or true")
    }//[両方不正解]のみfalse
    
}

