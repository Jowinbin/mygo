package main

import (
        "net/http"
            "fmt"

)

func handler(w http.ResponseWriter, r *http.Request)  {
        fmt.Fprintf(w, "hello, https")

}

func main() {
//        http.HandleFunc("/",handler)
                http.Handle("/", http.FileServer(http.Dir(".")))
            err:=http.ListenAndServeTLS(":1443","server.crt","server.pem", nil)
        if err!=nil{
            fmt.Println("error \n",err)
        }
//     http.HandleFunc("/", handler)
// 	http.ListenAndServe(":12345", nil)
}
