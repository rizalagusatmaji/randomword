package main

import (
	"quote/handler"

	"github.com/gin-gonic/gin"
)

//"https://animechan.vercel.app/api/random"
//bikin endpoint mengembalikan data dari animechan
// 1. cek spesifikasi https://random-indonesian-word.p.rapidapi.com/words/random?limit=5
// 2. bikin variabel baru untuk menampung respon API
// 3. bikin http client
// 4. membuat http server

func main() {
	// membuat router gin
	r := gin.Default()

	// menambah router dengan method GET
	r.GET("/randomword", handler.HandlerGetRandomWord)

	// menjalankan server di port :8080
	r.Run(":8080")
}
