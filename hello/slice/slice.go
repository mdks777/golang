package main

import (
	"fmt"
)

func main() {
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