package main

import (
	"calc-tut/logs"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	var userInputs []int
	var operators []string
	logs.D("welcome to bigad calculator it supports +-/* operations")
	//-- underdev -- if advanced calculator uncomment this --- underdev
	//	for  {
	//if 1 operation calculator uncomment this
	for i := 0; i < 2; i++ {
		var isNotValidOperator = true
		var userInput int

		var operator string
		logs.D("please enter number")

		for {
			_, err := fmt.Scanf("%d", &userInput)
			if err == nil {
				break
			}
			logs.W("you have entered invalid input you should enter number 0-9")
			var dump string
			fmt.Scanln(&dump)
		}
		userInputs = append(userInputs, userInput)
		/*-- underdev --
		comment if condition if you going for advanced calculator
		 --- underdev
		*/
		if i != 1 {
			logs.D("please enter operator + - / * ")

			for isNotValidOperator {
				fmt.Scan(&operator)
				if validateOperatiors(operator) {
					isNotValidOperator = false
				}
			}

			//-- underdev -- if advanced calculator uncomment this --- underdev
			// if operator == "=" {
			// 	calc(userInputs, operators)
			// 	break
			// }
			operators = append(operators, operator)
		}
		//if 1 operation calculator uncomment this
		if i == 1 {
			calc(userInputs, operators)
		}

	}

}

//calc with only 1 operation
func calc(userInputs []int, operators []string) {
	var operation int
	var debugOperation string

	var lenghtOfUserInputs = len(userInputs)
	for i := 0; i < lenghtOfUserInputs; i++ {
		if i < lenghtOfUserInputs-1 {
			debugOperation += strconv.Itoa(userInputs[i]) + operators[i] + strconv.Itoa(userInputs[i+1])
			switch operators[i] {
			case "-":
				operation += userInputs[i] - userInputs[i+1]

			case "*":
				operation += userInputs[i] * userInputs[i+1]
			case "/":
				operation += userInputs[i] / userInputs[i+1]
			default:
				operation += userInputs[i] + userInputs[i+1]
			}
		}
	}
	logs.W("debug user inputs slice")
	logs.W(userInputs, operators)
	logs.W("debug Operation as string")
	logs.W(debugOperation)
	logs.W("result")
	logs.W(operation)
}

//under dev -- advanced calc with continuous calculations -- under dev
// func calc(userInputs []int, operators []string) {
// 	var operation int
// 	var debugOperation string

// 	var lenthOfInputs = len(operators)
// 	for i := 0; i < lenthOfInputs; i++ {
// 		if i < lenthOfInputs {
// 			debugOperation += strconv.Itoa(userInputs[i]) + operators[i] + strconv.Itoa(userInputs[i+1])
// 			switch operators[i] {
// 			case "-":
// 				operation += userInputs[i] - userInputs[i+1]

// 			case "*":
// 				operation += userInputs[i] * userInputs[i+1]
// 			case "/":
// 				operation += userInputs[i] / userInputs[i+1]
// 			default:
// 				operation += userInputs[i] + userInputs[i+1]
// 			}
// 		}
// 	}
// 	logs.W("debug user inputs slice")
// 	logs.W(userInputs, operators)
// 	logs.W("debug Operation as string")
// 	logs.W(debugOperation)
// 	logs.W("result")
// 	logs.W(operation)
// }

func validateOperatiors(operators string) bool {
	var validOperatiors = regexp.MustCompile("^[-+*/=]")
	var result = validOperatiors.MatchString(operators)
	if !result {

		logs.E("Error you entered wrong input please enter operator + - / *")

	}
	return result
}
