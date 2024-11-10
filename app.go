package main

import (
 "fmt"
 "net/http"
 "github.com/go-redis/redis" 
 "os"
)

func main() {
 redisURL := os.Getenv("REDIS_URL")
 client := redis.NewClient(&redis.Options{
  Addr: redisURL, 
 })

 // Ping для проверки соединения
 _, err := client.Ping().Result()


 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Запуск\n")
	if err != nil {
	  fmt.Fprintf(w,  "Redis: Не удалось подключиться!\n")
	} else {
	  fmt.Fprintf(w, "Redis: успешно подлючен!\n")
	}
 })

 fmt.Println("Порт если чо 8080(http://localhost:8080/)")
 http.ListenAndServe(":8080", nil)
}
