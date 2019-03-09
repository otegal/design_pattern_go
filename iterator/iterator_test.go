package iterator

import (
	"testing"
)

func TestIterator(t *testing.T) {
	// BookShelfを作成
	bookShelf := BookShelf{}

	// BookをBookShelfに追加
	asserts := []string{"すごいHaskell本", "Java言語で学ぶデザインパターン入門", "Goならわかるシステムプログラミング"}
	for _, assert := range asserts {
		bookShelf.AppendBook(&Book{assert})
	}

	iterator := bookShelf.Iterator()
	i := 0
	for iterator.HasNext() {
		book := iterator.Next()
		// fmt.Println(book.(*Book).name) // インターフェースにして失ったBookの情報を復元する方法
		if book.(*Book).name != asserts[i] {
			t.Errorf("Expect Book.name to %s, but %s", asserts[i], book.(*Book).name)
		}
		i++
	}
}
