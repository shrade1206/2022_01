package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var (
		A         = make(chan string, 2)
		B         = make(chan string, 2)
		C         = make(chan string, 2)
		A_Msg     = make(chan string, 1)
		B_GetMsg  = make(chan string, 1)
		B_SendMsg = make(chan string, 1)
		C_Msg     = make(chan string)
	)
	total := 1
	go A_ch(A, B, A_Msg)
	go B_ch(B, C, B_GetMsg, B_SendMsg)
	go C_ch(C, C_Msg)
	go Message_Sender(A_Msg, B_GetMsg, B_SendMsg, C_Msg)
	for range time.Tick(time.Second * 1) {
		A <- "ball"
		// log.Printf("第%d次發球，A有%d個球，B有%d個球，C有%d個球", total, len(A), len(B), len(C))
		log.Printf("第%d次發球", total)
		total++
		log.Printf("A有%d個球，B有%d個球，C有%d個球", len(A), len(B), len(C))
	}
}

// 接球、傳球、發送訊息
func A_ch(A <-chan string, B chan<- string, A_Msg chan string) {
	for {
		ball := <-A
	SEND:
		for {
			select {
			case B <- ball:
				break SEND
			case <-time.After(time.Second):
				A_Msg <- "B快點接球"
			}
		}
	}
}

// 跟Ａ接球、傳球、發送訊息
func B_ch(B <-chan string, C chan<- string, B_Get <-chan string, B_Send chan<- string) {
	for {
		ball := <-B
		fmt.Println("B got ", ball)
	SEND:
		for {
			select {
			case C <- ball:
				break SEND
			case msg := <-B_Get:
				log.Println(msg)
				// break SEND
			case <-time.After(time.Second):
				B_Send <- "C快點接球"
			}
		}
	}
}

// 跟Ｂ接球、丟球
func C_ch(C <-chan string, C_Msg <-chan string) {
	for msg := range C_Msg {
		log.Println(msg)
		<-C
		log.Println("收到訊息，C已丟球")
		time.Sleep(time.Second * 2)
	}
}

// 傳訊息給Ｂ、Ｃ
func Message_Sender(A_Msg chan string, B_Get chan string, B_Send chan string, C_Msg chan string) {
	for {
		select {
		case msg := <-A_Msg:
			B_Get <- msg
			break
		case msg := <-B_Send:
			C_Msg <- msg
			break
		}
	}
}
