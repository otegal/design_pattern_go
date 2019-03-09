package iterator

import "fmt"

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type BookShelfIterator struct {
	BookShelf *BookShelf
	Last      int
}

func (bookShelfIterator *BookShelfIterator) HasNext() bool {
	if bookShelfIterator.Last >= len(bookShelfIterator.BookShelf.Books) {
		return false
	}
	return true
}

func (bookShelfIterator *BookShelfIterator) Next() interface{} {
	book := bookShelfIterator.BookShelf.Books[bookShelfIterator.Last]
	bookShelfIterator.Last++
	return book
}

type BookShelf struct {
	Books []*Book
}

func (bookShelf *BookShelf) AppendBook(book *Book) {
	bookShelf.Books = append(bookShelf.Books, book)
}

func (bookShelf *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{BookShelf: bookShelf}
}

type Book struct {
	name string
}

func Execute() {
	// BookShelfを作成
	bookShelf := BookShelf{}

	// BookをBookShelfに追加
	bookShelf.AppendBook(&Book{name: "すごいHaskell本"})
	bookShelf.AppendBook(&Book{name: "Java言語で学ぶデザインパターン入門"})
	bookShelf.AppendBook(&Book{name: "リーダブルコード"})
	bookShelf.AppendBook(&Book{name: "Goならわかるシステムプログラミング"})

	iterator := bookShelf.Iterator()
	for iterator.HasNext() {
		book := iterator.Next()
		fmt.Println(book)
	}
}
