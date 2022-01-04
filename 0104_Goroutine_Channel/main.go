package main

import (
	"log"
	"strconv"
	"time"
)

type player struct {
	name string
	cn   chan string
}

var a, b, c player
var (
	A   = make(chan string, 2)
	B   = make(chan string, 2)
	C   = make(chan string, 2)
	msg = make(chan string)
)

func main() {
	// wg := new(sync.WaitGroup)
	// a := player{name: "A", cn: A}
	// b := player{name: "B", cn: B}
	// c := player{name: "C", cn: C}
	i := 1
	num := strconv.Itoa(i)
	go GetBall(num)
	log.Printf("第%d次發球\n", i)
	for range time.Tick(time.Second) {
		i++
		select {
		case <-A:
			log.Println("ok")
			log.Println(A)
			b := <-A
			B <- b
			log.Println("?")
			// A傳球給B
			// PassBall(A, B)
		case <-B:
			// B傳球給C
			PassBall(A, B)

		default: // Channel 阻塞的話執行此區域
			log.Println("boom")
		}
		log.Printf("A目前有%d個球", len(A))
		log.Printf("B目前有%d個球", len(B))
		log.Printf("C目前有%d個球", len(C))

		// totle++
	}
}

func PassBall(Front chan string, Later chan string) {
	ball := <-Front
	Later <- ball
	log.Printf("ok")
}

func GetBall(num string) {
	A <- "ball" + num
}

// 發球給A
// func GetBall() {
// 	// 球不足2，就接球
// 	if len(a.ball) > 2 {
// 		// chan接到的球丟進struct，接球才完成
// 		ball := <-a_chan
// 		a.ball = append(a.ball, ball)
// 		log.Printf("%s接球", a.name)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// // 傳球
// func PassBall(Front player, FrontMsg chan string, Later player, LaterMsg chan string, msg chan string) {
// 	// 球不足2，就接球
// 	if len(Front.ball) < 2 {
// 		// chan接到的球丟進struct，接球才完成
// 		ball := <-FrontMsg
// 		Front.ball = append(Front.ball, ball)
// 		log.Printf("%s接球", Front.name)
// 	}
// 	// 前面有球，後面也可以接球
// 	if len(Front.ball) > 0 && len(Later.ball) < 2 {
// 		Front.ball = Front.ball[:len(Front.ball)-1]
// 		LaterMsg <- "ball"
// 		ball := <-LaterMsg
// 		Later.ball = append(Later.ball, ball)
// 		log.Printf("%s接球", Later.name)
// 	}
// 	// 如果前後兩個都滿了，通知後面快點接球或丟球
// 	if len(Front.ball) == 2 && len(Later.ball) == 2 {
// 		log.Printf("%s的球已經滿了，通知下一位接球或最後一位丟球", Front.name)
// 		msg <- Later.name + "請快點接球"
// 	}
// }

// // 丟球
// func ThrowBall(msg chan string) {
// 	Msg := <-msg
// 	if Msg == "C快點接球" && len(c.ball) == 2 {
// 		c.ball = c.ball[:len(c.ball)-1]
// 	}
// 	time.Sleep(time.Second * 2)
// }
