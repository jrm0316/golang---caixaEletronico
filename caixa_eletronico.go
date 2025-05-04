// Cédulas
// Leia um valor inteiro. A seguir, calcule o menor número de notas possíveis (cédulas) no qual o
// valor pode ser decomposto. As notas consideradas são de 100, 50, 20, 10, 5, 2 e 1. A seguir mostre
// o valor lido e a relação de notas necessárias.

// Entrada
// O arquivo de entrada contém um valor inteiro N (0 < N < 1000000).

// Saída
// Imprima o valor lido e, em seguida, a quantidade mínima de notas de cada tipo necessárias, conforme
// o exemplo fornecido. Não esqueça de imprimir o fim de linha após cada linha, caso contrário seu
// programa apresentará a mensagem: “Presentation Error”.

// 576

// 576
// 5 nota(s) de R$ 100,00
// 1 nota(s) de R$ 50,00
// 1 nota(s) de R$ 20,00
// 0 nota(s) de R$ 10,00
// 1 nota(s) de R$ 5,00
// 0 nota(s) de R$ 2,00
// 1 nota(s) de R$ 1,00

// 11257

// 11257
// 112 nota(s) de R$ 100,00
// 1 nota(s) de R$ 50,00
// 0 nota(s) de R$ 20,00
// 0 nota(s) de R$ 10,00
// 1 nota(s) de R$ 5,00
// 1 nota(s) de R$ 2,00
// 0 nota(s) de R$ 1,00

// 503

// 503
// 5 nota(s) de R$ 100,00
// 0 nota(s) de R$ 50,00
// 0 nota(s) de R$ 20,00
// 0 nota(s) de R$ 10,00
// 0 nota(s) de R$ 5,00
// 1 nota(s) de R$ 2,00
// 1 nota(s) de R$ 1,00

package main

import (
	"fmt"
)

func main() {
	n := 576
	ced_100 := 0
	ced_50 := 0
	ced_20 := 0
	ced_10 := 0
	ced_5 := 0
	ced_2 := 0
	ced_1 := 0

	if n > 0 {
		for n >= 100 {
			ced_100 = ced_100 + 1
			n = n - 100
		}
		for n >= 50 {
			ced_50 = ced_50 + 1
			n = n - 50
		}
		for n > 20 {
			ced_20 = ced_20 + 1
			n = n - 20
		}
		for n >= 10 {
			ced_10 = ced_10 + 1
			n = n - 10
		}
		for n >= 5 {
			ced_5 = ced_5 + 1
			n = n - 5
		}
		for n >= 2 {
			ced_2 = ced_2 + 1
			n = n - 2
		}
		for n >= 1 {
			ced_1 = ced_1 + 1
			n = n - 1
		}

		fmt.Println("cédulas de 100: ", ced_100)
		fmt.Println("cédulas de 50: ", ced_50)
		fmt.Println("cédulas de 20: ", ced_20)
		fmt.Println("cédulas de 10: ", ced_10)
		fmt.Println("cédulas de 5: ", ced_5)
		fmt.Println("cédulas de 2: ", ced_2)
		fmt.Println("cédulas de 1: ", ced_1)
		fmt.Println("n: ", n)
	}
}
