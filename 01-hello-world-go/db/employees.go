package db

import (
	"01-hello-world-go/models"
	"github.com/brianvoe/gofakeit/v7"
	"math/rand"
)

// department -> employees
var Employees = map[string][]models.Employee{}
var departmentNames = []string{"Engineering", "Marketing", "Sales", "HR", "Finance"}

const employeeCount = 1000

func randomGender() string {
	n := rand.Intn(3)
	switch n {
	case 0:
		return "Not Want to Disclose"
	case 1:
		return "Male"
	case 2:
		return "Female"
	}
	return "Not Want to Disclose"
}

func init() {
	for i := 0; i < employeeCount; i++ {
		employee := models.Employee{
			ID:         gofakeit.UUID(),
			Name:       gofakeit.Name(),
			Age:        rand.Intn(30) + 20,
			Department: gofakeit.RandomString(departmentNames),
			Gender:     randomGender(),
		}
		Employees[employee.Department] = append(Employees[employee.Department], employee)
	}
}
