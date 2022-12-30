package main

import (
	"fmt"
	"time"
)

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
	Book string
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

func PrintAllBooks(books map[int]*Book, i int) {

	fmt.Println("Name: ", books[i].Name)
	fmt.Println("Author: ", books[i].Author)
	fmt.Println("Page: ", books[i].Page)
	fmt.Println("In: ", books[i].InDate.Format("2006-01-02 15:04:05"))
	fmt.Println("Out: ", books[i].OutDate.Format("2006-01-02 15:04:05"))
	fmt.Println("In: ", books[i].IsIn)
	fmt.Println()

}

func PrintAllMembers(members map[int]*Member, i int) {

	fmt.Println("Name: ", members[i].Name)
	fmt.Println("Age: ", members[i].Age)
	fmt.Println()

}

func main() {

	in := make(map[int]*Book)
	out := make(map[int]*Book)
	books := AddBooks()

	members := AddMembers()

	fmt.Println("BOOKS:")
	fmt.Println("************")
	for i := 0; i < 5; i++ {
		PrintAllBooks(books, i)
	}

	fmt.Println("MEMBERS:")
	fmt.Println("************")
	for i := 0; i < 5; i++ {
		PrintAllMembers(members, i)
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

	members[0].Book = books[0].Name
	members[2].Book = books[1].Name
	members[4].Book = books[3].Name

	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Println("LIBRARY")
	fmt.Println()
	fmt.Println("Inside Library")
	fmt.Println("************")
	fmt.Println()
	for i := 0; i < 5; i++ {
		if books[i].IsIn {
			in[i] = books[i]
		}
		if in[i] != nil {
			PrintAllBooks(in, i)
		}
	}

	fmt.Println("Enrolments")
	fmt.Println("************")
	fmt.Println()
	for i := 0; i < 5; i++ {
		if members[i].Book != "" {
			fmt.Println("->", members[i].Name, members[i].Book)
			fmt.Println()
		}

	}

	delete(enrolment, members[0])
	books[0].IsIn = true

	fmt.Println("---------------------------------------------------------")
	fmt.Println("Inside Library")
	fmt.Println("************")
	fmt.Println()
	for i := 0; i < 5; i++ {
		if books[i].IsIn {
			in[i] = books[i]
		}
		if in[i] != nil {
			PrintAllBooks(in, i)
		}

	}
	fmt.Println()
	fmt.Println("Outside Library")
	fmt.Println("************")
	fmt.Println()
	for i := 0; i < 5; i++ {
		if !books[i].IsIn {
			out[i] = books[i]
		}
		if out[i] != nil {
			PrintAllBooks(out, i)
		}
	}

}
