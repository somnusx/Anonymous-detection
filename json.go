
package main

import (
    "strings"
    "net/http"
    "encoding/json"
)
func Header(w http.ResponseWriter, r *http.Request) {
    head := r.Header
    headers, _ := json.Marshal(head)
    w.Write(headers)
}

func Ip(w http.ResponseWriter, r *http.Request) {
    a := r.RemoteAddr
    ip := strings.Split(a,":")[0]
    w.Write([]byte(ip))
}
 
func main() {
    http.HandleFunc("/ip", Ip)
    http.HandleFunc("/header", Header)
    http.ListenAndServe(":3000", nil)
}
