package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Chào mừng đến với thư viện Library Mangement System!!!")
	})
	fmt.Println("Hiện server đang chạy !!!")
	log.Fatal(http.ListenAndServe(":8080",nil))
}