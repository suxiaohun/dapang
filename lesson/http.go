package lesson

import (
	"dapang/rsa_tools"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Http() {
	http.Handle("/", timeMiddleware(hello))
	http.Handle("/key_verify", http.HandlerFunc(IsMatch))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello\n"))
}

func timeMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		next.ServeHTTP(wr, r)
		timeCost := time.Since(timeStart)
		fmt.Println(timeCost)
	})
}

type Param struct {
	Privk string `json:"privk"`
	Pubk  string `json:"pubk"`
}

func IsMatch(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := make(map[string]interface{}, 0)
	res["code"] = 1000

	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	privk := ""
	pubk := ""

	switch req.Method {
	case "GET":
		privk = req.FormValue("privk")
		pubk = req.FormValue("pubk")
	case "POST":
		decoder := json.NewDecoder(req.Body)
		var p Param
		err := decoder.Decode(&p)
		if err != nil {
			panic(err)
		}

		privk = p.Privk
		pubk = p.Pubk
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

	if privk == "" {
		res["code"] = 1200
		res["message"] = "privk is blank"
		w.Write(toJson(res))
		return
	}

	if pubk == "" {
		res["code"] = 1200
		res["message"] = "pubk is blank"
		w.Write(toJson(res))
		return
	}

	match, err := rsa_tools.VerifyKeyPair(privk, pubk)

	if match {
		res["code"] = 1000
		res["message"] = "ok"
	} else {
		res["code"] = 2000
		res["message"] = err.Error()
	}
	w.Write(toJson(res))
}

func toJson(data map[string]interface{}) []byte {
	jData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("json marshal error: %v\n", err)
	}
	return jData
}
