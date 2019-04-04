package main
  
import (
        "net/http"
)

func main() {
        h := func(w http.ResponseWriter, r *http.Request) {
                return
        }

        http.HandleFunc("/demo", h)

        http.ListenAndServe(":19009", nil)
}
