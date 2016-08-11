// funcstruct
package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books // declare a struct variable Book1
	var Book2 Books // declare a struct variable Book2

	// book 1 specification
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	// book 2 specification
	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700

	// Print Book 1 info
	printBook(Book1)

	fmt.Println()

	// Print Book 2 info
	printBook(Book2)

}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
