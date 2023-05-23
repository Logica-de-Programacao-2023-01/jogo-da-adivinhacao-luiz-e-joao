package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	result := rand.Intn(100) + 1
	inp := 0
	fmt.Println("Ache o numero inteiro de 1 a 100:")
	s_n := "x"

	n_Tentativas := make(map[int]int)
	tentativa := 1

	for {
		fmt.Scanln(&inp)
		n_Tentativas[tentativa]++
		if inp == result {
			fmt.Println("Você acertou!!!, o recomendado agora é apostar na mega senna")
			fmt.Printf("Você chutou %d vezes\n", n_Tentativas[tentativa])
			fmt.Println("Quer deleitar-se a jogar esse jogo novamente? (s/n)")
			fmt.Scanln(&s_n)
			if s_n == "n" {
				break
			} else {
				tentativa++
				result = rand.Intn(100) + 1
				fmt.Println("Ache o numero inteiro de 1 a 100:")
			}

		} else if inp > result {
			fmt.Println("Chutou errado, o número é menor, tente denovo:")
		} else if inp < result {
			fmt.Println("Chutou errado, o número é maior, tente denovo:")
		}
	}
	n_Tentativas_ord := make([]int, 0, len(n_Tentativas))

	for n_index := range n_Tentativas {
		n_Tentativas_ord = append(n_Tentativas_ord, n_index)

	}

	sort.Ints(n_Tentativas_ord)

	for _, num_tent := range n_Tentativas_ord {
		fmt.Printf("Você chutou %d vezes e apenas no chute %d você acertou", n_Tentativas[num_tent], num_tent)

	}

}
