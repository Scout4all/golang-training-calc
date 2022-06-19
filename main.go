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
	logs.W("when you need to get result please add = when app ask you for operator")

	//-- underdev -- if advanced calculator uncomment this --- underdev
	for {
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
		logs.D("please enter operator + - / * ")

		for isNotValidOperator {
			fmt.Scan(&operator)
			if validateOperatiors(operator) {
				isNotValidOperator = false
			}
		}

		//-- underdev -- if advanced calculator uncomment this --- underdev
		if operator == "=" {
			calc(userInputs, operators)
			break
		}
		operators = append(operators, operator)

	}

}

//under dev -- advanced calc with continuous calculations -- under dev
func calc(userInputs []int, operators []string) {
	var operation int
	var debugOperation string

	var lenthOfInputs = len(userInputs)
	for i := 0; i < lenthOfInputs; i++ {

		if i == 0 {

			operation = userInputs[i]
			debugOperation = strconv.Itoa(userInputs[i])

		} else {

			debugOperation += operators[i-1] + strconv.Itoa(userInputs[i])

			switch operators[i-1] {
			case "-":
				operation -= userInputs[i]

			case "*":
				operation *= userInputs[i]
			case "/":
				operation /= userInputs[i]
			default:
				operation += userInputs[i]
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
func validateOperatiors(operators string) bool {
	var validOperatiors = regexp.MustCompile("^[-+*/=]")
	var result = validOperatiors.MatchString(operators)
	if !result {

		logs.E("Error you entered wrong input please enter operator + - / *")

	}
	return result
}
