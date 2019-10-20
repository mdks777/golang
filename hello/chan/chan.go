package main

import "fmt"

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