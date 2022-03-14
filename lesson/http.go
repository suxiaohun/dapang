package lesson

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Http() {
	http.Handle("/", timeMiddleware(index1))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index1(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello"))
	//fmt.Fprint(w, "http 接口测试")

}
func index2(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	w.Write([]byte("hello"))
	//fmt.Fprint(w, "http 接口测试")
	timeCost := time.Since(timeStart)
	fmt.Println(timeCost)
}

func timeMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		next.ServeHTTP(wr, r)
		timeCost := time.Since(timeStart)
		fmt.Println(timeCost)
	})
}
