package main

import "fmt"

func main() {
	p := Project{name: "Upload images"}
	t1 := Task{name: "Mockup", responsable: "UI Designer", price: 1000}
	t2 := Task{name: "Markup", responsable: "Web Designer"}
	t3 := Task{name: "Javascript", responsable: "Frontend Developer", price: 1000}
	t4 := Task{name: "API Backend", responsable: "Backend Developer"}
	t5 := Task{name: "Database", responsable: "DBA", price: 1000}

	t1.Add(&SubTask{name: "HTML", price: 300})
	t1.Add(&SubTask{name: "CSS", price: 700})

	t4.Add(&SubTask{name: "Authentication", price: 300})
	t4.Add(&SubTask{name: "DB connection", price: 900})

	p.Add(&t1)
	p.Add(&t2)
	p.Add(&t3)
	p.Add(&t4)
	p.Add(&t5)

	fmt.Println(&p)
}
