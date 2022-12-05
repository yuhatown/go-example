package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World Response")
}

func main() {
	// 요청한 라우팅 경로에 대한 함수 지정
	http.HandleFunc("/hello", helloHandler)

	// 웹서버 실제로 동작시키기 위한 함수. 서버가 동작할 포트 지정
	err := http.ListenAndServe(":8080", nil)

	// 예외 처리
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		fmt.Println("ListenAndServe Started! -> Port(8080)")
	}
}