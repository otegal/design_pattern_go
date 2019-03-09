package iterator

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Aggregate interface {
	Iterator() Iterator
}

type BookShelfIterator struct {
	BookShelf *BookShelf
	Index     int
}

func (bookShelfIterator *BookShelfIterator) HasNext() bool {
	if bookShelfIterator.Index >= bookShelfIterator.BookShelf.GetLength() {
		return false
	}
	return true
}

func (bookShelfIterator *BookShelfIterator) Next() interface{} {
	book := bookShelfIterator.BookShelf.GetBookAt(bookShelfIterator.Index)
	bookShelfIterator.Index++
	return book
}

type BookShelf struct {
	Books []*Book
	Last  int
}

func (bookShelf *BookShelf) GetBookAt(Last int) *Book {
	return bookShelf.Books[Last]
}

func (bookShelf *BookShelf) AppendBook(book *Book) {
	bookShelf.Books = append(bookShelf.Books, book)
}

func (bookShelf *BookShelf) GetLength() int {
	return len(bookShelf.Books)
}

func (bookShelf *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{BookShelf: bookShelf}
}

type Book struct {
	name string
}
