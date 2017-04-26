
package main

import (
    "net/http"
    "encoding/json"
)
func SayHello(w http.ResponseWriter, r *http.Request) {
	map1 = r.Header
    str, _ := json.Marshal(map1)
    w.Write(str)
}
 
func main() {
    http.HandleFunc("/", SayHello)
    http.ListenAndServe(":3000", nil)
}
