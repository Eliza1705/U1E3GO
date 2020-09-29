package main
import (
	"bufio"
	"fmt"
	"log"
	"os"
)
func main() {
	var a [3]string
	a[0] = "Elizabeth"
	a[1] = "resendiz"
	a[2] = "amaro"
	var b [3]string

	file, err := os.Open("frase.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var cont = 0
	fmt.Println("******Frase desordenada************")
	for scanner.Scan() {

		fmt.Println(scanner.Text())
		b[cont] = scanner.Text()
		cont++
		//fmt.Println(scanner.Bytes())

	}
	var cadenaCorrecta [3]string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if a[i] == b[j] {
				cadenaCorrecta[i] = a[i]
			}
		}
	}
	fmt.Println("******Frase ordenada*************")
	for k := 0; k < 3; k++ {
		fmt.Println(cadenaCorrecta[k])
	}
}
