package main

import "net/http"
import "fmt"
import "math"

func looping() float64 {
    x := 0.0001
    for i := 0; i <= 1000000; i++ {
        x += math.Sqrt(x)
    }
    return x
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        looping()
        fmt.Fprintf(w, "Code.education Rocks! ")
    })

    http.ListenAndServe(":8000", nil)
}
