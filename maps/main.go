package main

import "fmt"

func main() {

user := map[string]string {
	"name": "Rodrigo",
	"lastName": "Silva",
}
fmt.Println(user["name"])


student := map[string]map[string]string {
	"name": {
		"firstName": "Ana",
		"lastName": "Ferreira",
	},
	"course": {
		"nameCourse": "Back-End",
		"local": "Alura",
	},
}
fmt.Println(student)
delete(student, "name")
fmt.Println(student)


schoolGrades := map[string]int {
	"Lucas": 8,
	"Thiago": 9,
	"Cris": 10,
}

for nome, schoolGrades := range schoolGrades {
	fmt.Println(nome, "tirou", schoolGrades)
	}
}