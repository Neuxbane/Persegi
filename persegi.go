package main
import "fmt"

func luaspersegi(sisi float64) float64 {
	return sisi*sisi
}

func main() {
	var a float64
	for true {
		fmt.Printf("Masukan sisinya:")
		fmt.Scan(&a)
		fmt.Println(luaspersegi(a))
	}
}