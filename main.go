package main

import (
	"fmt"
	"time"
)

// --Summary:
//
//	Create a program to manage lending of library books.
//
// --Requirements:
// * The library must have books and members, and must include:
//   - Which books have been checked out +
//   - What time the books were checked out +
//   - What time the books were returned +
//
// * Perform the following:
//   - Add at least 4 books and at least 3 members to the library +
//   - Check out a book +
//   - Check in a book +
//   - Print out initial library information, and after each change
//
// * There must only ever be one copy of the library in memory at any time
//
// --Notes:
// * Use the `time` package from the standard library for check in/out times
// * Liberal use of type aliases, structs, and maps will help organize this project

type Book struct {
	Name, Author string
	Page         int
	OutDate      time.Time
	InDate       time.Time
	IsIn         bool
}

var Books map[int]*Book

type Member struct {
	Name string
	Age  int
}

var Members map[int]*Member

func AddBooks() map[int]*Book {
	Books = make(map[int]*Book)

	Books[0] = &Book{
		Name:   "It",
		Author: "Stephen King",
		Page:   1000,
		InDate: time.Date(2019, time.April, 14, 3, 45, 0, 5, time.UTC),
		IsIn:   true,
	}

	Books[1] = &Book{
		Name:   "The Gunslinger",
		Author: "Stephen King",
		Page:   300,
		InDate: time.Date(2019, time.April, 14, 3, 45, 0, 5, time.UTC),
		IsIn:   true,
	}

	Books[2] = &Book{
		Name:   "Insomnia",
		Author: "Stephen King",
		Page:   783,
		InDate: time.Date(2019, time.April, 14, 3, 45, 0, 5, time.UTC),
		IsIn:   true,
	}

	Books[3] = &Book{
		Name:   "Green Mile",
		Author: "Stephen King",
		Page:   500,
		InDate: time.Date(2019, time.April, 14, 3, 45, 0, 5, time.UTC),
		IsIn:   true,
	}

	Books[4] = &Book{
		Name:   "Salem's Lot",
		Author: "Stephen King",
		Page:   902,
		InDate: time.Date(2019, time.April, 14, 3, 45, 0, 5, time.UTC),
		IsIn:   true,
	}

	return Books
}

func AddMembers() map[int]*Member {
	Members = make(map[int]*Member)

	Members[0] = &Member{
		Name: "Serap Baysal",
		Age:  23,
	}
	Members[1] = &Member{
		Name: "Sumeyye Baysal",
		Age:  23,
	}
	Members[2] = &Member{
		Name: "Ayse Hanedan",
		Age:  45,
	}
	Members[3] = &Member{
		Name: "Murat Basar",
		Age:  37,
	}
	Members[4] = &Member{
		Name: "Ceyhun Arici",
		Age:  28,
	}
	Members[5] = &Member{
		Name: "Melek Alkan",
		Age:  39,
	}

	return Members
}

func Enrolments(m Member, b Book) map[Member]Book {
	enrolment := make(map[Member]Book)

	if !b.IsIn {
		fmt.Println("INSIDE ISIN")
		b.IsIn = true
		b.InDate = time.Now()
		delete(enrolment, m)
		fmt.Println("You can't get this book")
		fmt.Println(enrolment)
		return enrolment
	}
	if b.IsIn {
		fmt.Println("INSIDE ISOUT")
		b.OutDate = time.Now()
		b.IsIn = false
		enrolment[m] = b

		return enrolment
	}
	return enrolment

}

func main() {
	books := AddBooks()

	members := AddMembers()

	fmt.Println("BOOKS:")
	for i := 0; i < len(books); i++ {
		fmt.Println("Name: ", books[i].Name)
		fmt.Println("Author: ", books[i].Author)
		fmt.Println("Page: ", books[i].Page)
		fmt.Println("In: ", books[i].InDate)
		fmt.Println("Out: ", books[i].OutDate)
		fmt.Println("In: ", books[i].IsIn)
		fmt.Println()
	}

	fmt.Println("MEMBERS:")
	for i := 0; i < len(members); i++ {
		fmt.Println("Name: ", members[i].Name)
		fmt.Println("Age: ", members[i].Age)
		fmt.Println()
	}

	enrolment := make(map[*Member]*Book)
	enrolment[members[0]] = books[0]
	enrolment[members[2]] = books[1]
	enrolment[members[4]] = books[3]

	books[0].IsIn = false
	books[0].InDate = time.Now()
	books[1].IsIn = false
	books[1].InDate = time.Now()
	books[3].IsIn = false
	books[3].InDate = time.Now()

	for i := 0; i < 5; i++ {
		fmt.Println(enrolment[members[i]])
		fmt.Println()
		fmt.Println(books[i])
		fmt.Println(members[i])
	}

	fmt.Println("-------------------------------------------------------------------------------------------")
	delete(enrolment, members[0])
	books[0].IsIn = true
	fmt.Println(enrolment[members[0]])
	fmt.Println(books[0])

	for i := 0; i < 5; i++ {
		fmt.Println(enrolment[members[i]])
		fmt.Println(books[i])
		fmt.Println(members[i])
	}

}
