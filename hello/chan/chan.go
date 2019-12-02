package main

import (
	"fmt"
	"time"
)

func main() {
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
	fmt.Printf("\n")

	ch4:=make(chan string, 3)
	ch4<-"APPLE"
	fmt.Println(len(ch4),cap(ch4))//1 3

	ch5:= make(chan int, 20)
	go receive("1st goroutine",ch5)//ランダムに
	go receive("2nd goroutine",ch5)//表示されるため
	go receive("3rd goroutine",ch5)//毎回表示が違う
	for i:=0;i<20;i++{ch5 <- i}
	close(ch5)
	time.Sleep(3* time.Second)//これないと動かない

	ch6:= make(chan int, 1)
	ch7:= make(chan int, 1)
	ch8:= make(chan int, 1)
	ch6 <- 1
	ch7 <- 2
	select {
	case <-ch6 :
		fmt.Println("from 1")
	case <-ch7 :
		fmt.Println("from 2")
	case ch8 <- 3:
		fmt.Println("to 3")
	default:
		fmt.Printf("Nothing!")//毎回表示が代わる
	}
}

func receiver(ch2 <-chan int){//ch2は送信専用のチャネル
	for {
		i:=<-ch2
		fmt.Printf("%d",i)
	}//0123456789
}

func receive(name string, ch5 <-chan int) {
	for {
		i,ok:=<-ch5
		if ok == false {
			break
		}
		fmt.Println(name,i)
	}
	fmt.Println(name + " is done")
}