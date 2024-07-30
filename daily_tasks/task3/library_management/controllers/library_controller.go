package main
Define a LibraryManager interface with the following methods:
AddBook(book Book)
RemoveBook(bookID int)
BorrowBook(bookID int, memberID int) error
ReturnBook(bookID int, memberID int) error
ListAvailableBooks() []Book
ListBorrowedBooks(memberID int) []Book

type librarymanager interface {
	addBook(book Book)
	remo
}
