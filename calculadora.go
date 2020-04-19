package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//tipo de dato que permite tener propiedades
// y funciones, muy similiar a un como se usa en POO
type calc struct{}

//con(cal) estamos donde la posibilida de tener el metodo
//dentro de cal
func (calc) operator(entrada string, operador string) int {
	//Separa un string dado un separador
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])

	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1 - operador2)
		return operador1 - operador2
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2
	case "/":
		fmt.Println(operador1 / operador2)
		return operador1 / operador2
	default:
		fmt.Println("Operación no soportada")
		return 0
	}

}

func parsear(entrada string) int {
	// strconv.Atoi Convierte un string a enteros.
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func LeerEntrada() string {
	//Nos permite leer un input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
