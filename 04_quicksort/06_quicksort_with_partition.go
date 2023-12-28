package main

import (
	"fmt"
)

// Função principal do quicksort
func quickSort(arr []int, low int, high int) {
	// Condição de parada para recursão
	if low < high {
		// Obtém o índice do pivô após a partição
		pi := partition(arr, low, high)
		fmt.Println("pi : ", pi)

		// Recursivamente ordena os elementos antes e depois da partição
		fmt.Println("ordenando os elementos antes da partição")
		quickSort(arr, low, pi-1)
		fmt.Println("ordenando elementos depois da partição")
		quickSort(arr, pi+1, high)
	}
}

// Função para realizar a partição do array
func partition(arr []int, low int, high int) int {
	// Escolhe o pivô como o último elemento do array
	pivot := arr[high]
	i := low - 1

	// Itera sobre os elementos de low a high - 1
	for j := low; j < high; j++ {
		fmt.Println("j atual entre low e high -1:", j)
		// Se o elemento atual for menor que o pivô
		if arr[j] < pivot {
			// Incrementa o índice i e troca arr[i] com arr[j]
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Troca o pivô para a posição correta
	arr[i+1], arr[high] = arr[high], arr[i+1]

	// Retorna o índice do pivô
	return i + 1
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	low := 0
	high := len(arr) - 1

	fmt.Println("Array antes da ordenação:", arr)

	quickSort(arr, low, high)

	fmt.Println("Array depois da ordenação:", arr)
}
