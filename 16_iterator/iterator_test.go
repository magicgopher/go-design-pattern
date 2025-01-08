package iterator_test

import (
	iterator "github.com/magicgopher/go-design-pattern/16_iterator"
	"strings"
	"testing"
)

func TestBookIterator(t *testing.T) {
	// 创建图书集合
	bookShelf := iterator.NewBookShelf()

	// 添加图书
	books := []*iterator.Book{
		{ID: "1", Title: "Go Programming", Author: "John Doe", ISBN: "111-1111111"},
		{ID: "2", Title: "Python Basics", Author: "Jane Smith", ISBN: "222-2222222"},
		{ID: "3", Title: "Java Advanced", Author: "Bob Johnson", ISBN: "333-3333333"},
		{ID: "4", Title: "Go Advanced", Author: "Alice Brown", ISBN: "444-4444444"},
	}

	for _, book := range books {
		bookShelf.AddBook(book)
	}

	// 测试正向迭代器
	t.Run("Normal Iterator", func(t *testing.T) {
		it := bookShelf.GetIterator()
		count := 0
		for it.HasNext() {
			book := it.Next().(*iterator.Book)
			if book != books[count] {
				t.Errorf("期望得到书籍 %v，但得到 %v", books[count], book)
			}
			count++
		}
		if count != len(books) {
			t.Errorf("期望迭代 %d 本书，但实际迭代了 %d 本", len(books), count)
		}
	})

	// 测试反向迭代器
	t.Run("Reverse Iterator", func(t *testing.T) {
		it := bookShelf.GetReverseIterator()
		count := 0
		for it.HasNext() {
			book := it.Next().(*iterator.Book)
			if book != books[len(books)-1-count] {
				t.Errorf("期望得到书籍 %v，但得到 %v", books[len(books)-1-count], book)
			}
			count++
		}
		if count != len(books) {
			t.Errorf("期望迭代 %d 本书，但实际迭代了 %d 本", len(books), count)
		}
	})

	// 测试过滤迭代器（只返回Go相关的书籍）
	t.Run("Filter Iterator", func(t *testing.T) {
		it := bookShelf.GetFilterIterator(func(b *iterator.Book) bool {
			return strings.Contains(b.Title, "Go")
		})

		expectedGoBooks := []string{"Go Programming", "Go Advanced"}
		count := 0
		for it.HasNext() {
			book := it.Next().(*iterator.Book)
			if !strings.Contains(book.Title, "Go") {
				t.Errorf("过滤器应该只返回Go相关的书籍，但返回了 %s", book.Title)
			}
			count++
		}
		if count != len(expectedGoBooks) {
			t.Errorf("期望找到 %d 本Go相关的书，但实际找到 %d 本", len(expectedGoBooks), count)
		}
	})

	// 测试迭代器重置功能
	t.Run("Iterator Reset", func(t *testing.T) {
		it := bookShelf.GetIterator()
		// 先迭代一些元素
		it.Next()
		it.Next()
		// 重置迭代器
		it.Reset()
		// 检查第一个元素
		book := it.Next().(*iterator.Book)
		if book != books[0] {
			t.Errorf("重置后应该返回第一本书，但返回了 %v", book)
		}
	})
}
