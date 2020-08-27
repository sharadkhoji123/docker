package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

var visits = 0

func main() {

	http.HandleFunc("/test", test)
	http.ListenAndServe(":8080", nil)
}

func test(w http.ResponseWriter, r *http.Request) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := rdb.Context()
	ping, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connection established")
	err = rdb.Set(ctx, "name", "Elliot", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

	fmt.Println(ping, err)
	visits++
	fmt.Fprintln(w, "No of visits: ", visits)

}
