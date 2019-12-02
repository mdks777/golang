package main

import (
	"fmt"
)

func main() {
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
}