package main

import "fmt"

type product struct {
	Id    string
	Title string
	Price string
}

func main() {
	// 1
	var hobbies [3]string = [3]string{"first", "second", "last"}
	output(hobbies)

	// 2
	output("2 -----")
	output(hobbies[0])
	output(hobbies[1:3])

	// 3
	output("3 -----")
	mainHob := hobbies[:2]
	output(mainHob)
	// 4
	output("4 -----")
	output(mainHob[0:1])
	output(len(mainHob[0:1]), cap(mainHob[0:1]))

	// 5
	output("5 -----")
	courseGoal := []string{"Learn Go", "Apply Go"}
	output(courseGoal)
	// 6
	output("6 -----")
	courseGoal[1] = "Write RestAPI"
	output(courseGoal)
	courseGoal = append(courseGoal, "Apply Go")
	output(courseGoal)

	// Bonus
	output("Bonus -----")
	products := []product{
		{
			Id: "1", Title: "qk65", Price: "49.99",
		},
		{
			Id: "2", Title: "qk75", Price: "59.99",
		},
	}
	output(products)
	products = append(products, product{
		Id: "3", Title: "qk100", Price: "69.99",
	})
	output(products)
}

func output(array ...interface{}) {
	fmt.Println(array...)
}
