package main

import "fmt"

type Worker interface {
	Work()
}

type Employee struct {
	name string
	age  int
}

// реализуем метод Work() чтобы структура Employee
//удовлетворяла интерфейсу Worker

func (e Employee) Work() {
	fmt.Println("Now employee with name", e.name, "is working")
}

func Describer(w Worker) {
	fmt.Printf("Interface with type %T and value %v\n", w, w)
}

type Student struct {
	name         string
	courseNumber int
}

func (s Student) WantToEat() {
	fmt.Println("Student with name", s.name, "want to eat")
}

func (s Student) Work() {
	fmt.Println("Now student with name", s.name, "is working")
}

func main() {

	emp := Employee{"Bob", 34}
	var workerEmployee Worker = emp //присваиваем сотрудника в переменную типа Worker
	workerEmployee.Work()
	Describer(workerEmployee)

	std := Student{"Petr", 19}
	var workerStudent Worker = std
	workerStudent.Work()
	Describer(workerStudent)

	var workers []Worker = []Worker{workerEmployee, workerStudent}
	for _, worker := range workers {
		Describer(worker)
	}

}
