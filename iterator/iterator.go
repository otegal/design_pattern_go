package iterator

import "fmt"

// Book は本を表す構造体
type Book struct {
	name string
}

// BookShelf は本棚を表す構造体。複数の本(Book)を所有可能
type BookShelf struct {
	Books []*Book
}

// AppendBook は本棚に本を格納するBookShelfのメソッド
func (bookShelf *BookShelf) AppendBook(book *Book) {
	bookShelf.Books = append(bookShelf.Books, book)
}

// Execute はプログラム実行用のメソッド
func Execute() {
	// Bookを作成
	firstBook := Book{name: "first book"}
	secondBook := Book{name: "second book"}

	fmt.Println(firstBook, secondBook) // デバッグ用

	// BookShelfを作成
	bookShelf := BookShelf{}
	bookShelf.AppendBook(&firstBook)
	bookShelf.AppendBook(&secondBook)

	// デバッグ用
	for _, book := range bookShelf.Books {
		fmt.Println(book)
	}

}
