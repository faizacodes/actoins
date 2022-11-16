// package main

// import "fmt"

// func pupils(){

// type employee struct{
// 	name string
// 	age int
// 	average_grade float32
// }
// employee1 := employee{
// 	name : "Милана",
// 	age : 15,
// 	average_grade : 2.8,
// }
// employee2 := employee{
// 	name : "Матвей",
// 	age : 16,
// 	average_grade : 3,
// }
// fmt.Printf("%+v\n", employee1)
// fmt.Printf("%+v\n", employee2)

// }

// func main(){
// 	pupils()
// }


// type Pupil struct{
// 	name string
// 	age int
// 	average_grade float32
// 	srez []int
// }

// func NewEmployee(name string, age int, average_grade float32) employee{
// 	return employee{
// 		name: name,
// 		age: age, 
// 		average_grade: average_grade,
// 	}
// }


// type age int

// func (a age) isAdult() bool {
// 	return a >= 18
// }


// func (p Pupil) getInfo() string{
// 	return fmt.Sprintf ("Ученик: %s\nВозраст: %d\nСредняя оценка: %f\n", p.name, p.age, p.average_grade)
// }

// func (p Pupil) getAge() int{
// 	return p.age
// }

// func (p Pupil) getMaxDigit() int{
// 	max := 0
// 	for _, i := range p.srez{
// 		if max < i{
// 			max = i

// 		}


// 	}
// 	return max 
// }

// func main(){
// 	pupil1 := Pupil{"Faiza", 13, 4.3, []int{1,45, 3,87, 67, 101}}
// 	fmt.Println(pupil1)
	// pupil1.name = "Alex"
	// fmt.Println(pupil1)
	// fmt.Println(pupil1.getAge())

	// fmt.Printf("%s\n", pupil1.getInfo())
	// faiza_age := age(19)
	// fmt.Println("Am i adult?:", faiza_age.isAdult())
// 	fmt.Printf("Наибольшее число: %d\n", pupil1.getMaxDigit())
// }


// func change(s *string){
// 	*s = "moloko"
// }

// func main(){
// 	a := "hleb"
// 	fmt.Println(a)
// 	var pointer *string = &a 
// 	change(pointer)
// 	fmt.Println(a)
// }
