package parsing

import (
	"fmt"
	"strings"
)

func BalanceEquation(reagents []string, products []string) (string, error) {
	var reagents_str string
	var products_str string
	for _, str := range reagents {
		reagents_str += str
	}
	for _, str := range products {
		products_str += str
	}
	reagents_parsed_general := ParseFormula(reagents_str)
	products_parsed_general := ParseFormula(products_str)	
	for key := range reagents_parsed_general {
		if products_parsed_general[key] == 0 {
			return "", fmt.Errorf("incorrect input")
		}
	}
	mat := make([][]float64, len(reagents_parsed_general))
	i := 0
	for key, value := range reagents_parsed_general {
		for _, val := range reagents {
			if strings.Contains(val, key) {
				mat[i] = append(mat[i], float64(value))
			} else {
				mat[i] = append(mat[i], 0.0)
			}
		}
		i++
	}
	i = 0
	for key, value := range products_parsed_general {
		for _, val := range products {
			if strings.Contains(val, key) {
				mat[i] = append(mat[i], float64(-value))
			} else {
				mat[i] = append(mat[i], 0.0)
			}
		}
		i++
	}
	// answer := gaussianElimination(mat, make([]float64, len(mat)))
	// fmt.Println(answer)
	// var str string
	// i = 0
	// for _, value := range reagents {
		// if i != 0 {
			// str += " + "
		// }
		// str += fmt.Sprintf("%f", answer[i]) + value
		// i++
	// }
	// str += " = "
	// i = 0
	// for _, value := range products {
		// if i != 0 {
			// str += " + "
		// }
		// str += fmt.Sprintf("%f", answer[i]) + value
		// i++
	// }
	return "", nil
}