package main

import (
  "fmt"
  "io"
  "net/http"
  "log"
)

func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
  plot(w, []complex64{complex(2,2),complex(3,2),complex(1.75,1),complex(2,1),complex(2.25,1),
  complex(2.5,1),complex(2.75,1),complex(3,1),complex(3.25,1)}, 4, 3)
}

func plot(out io.Writer, points []complex64, scale int, dotSize int) {
  scaler := 200/scale
  fmt.Fprintf(out, "<!DOCTYPE html>\n<head>\n<title>plotter</title>\n</head>\n<body>\n" +
  "<svg height='420' width='420' xmlns='http://www.w3.org/2000/svg'>\n<line x1='0' y1='210' x2='420' y2='210' style='stroke:rgb(0,0,0);stroke-width:2'/>\n<line x1='210' y1='0' x2='210' y2='420' style='stroke:rgb(0,0,0);stroke-width:2'/>\n")
  for _, point := range points {
    fmt.Println(real(point))
    fmt.Fprintf(out, "<circle cx='%d' cy='%d' r='%d' fill='red'/>\n", int(float64(210)+float64(scaler)*float64(real(point))), int(float64(210)-float64(scaler)*float64(imag(point))), dotSize)
  }
  fmt.Fprintf(out, "</svg>\n</body>\n</html>")
}
