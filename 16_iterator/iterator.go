package iterator

// 迭代器模式

// Iterator 迭代器接口
type Iterator interface {
	HasNext() bool
	Next() interface{}
	CurrentItem() interface{}
	Reset()
}

// Book 图书结构体
type Book struct {
	ID     string
	Title  string
	Author string
	ISBN   string
}

// BookShelf 图书集合
type BookShelf struct {
	books []*Book
}

// NewBookShelf 创建新的图书集合
func NewBookShelf() *BookShelf {
	return &BookShelf{
		books: make([]*Book, 0),
	}
}

// AddBook 添加图书
func (bs *BookShelf) AddBook(book *Book) {
	bs.books = append(bs.books, book)
}

// GetIterator 获取迭代器
func (bs *BookShelf) GetIterator() Iterator {
	return &BookIterator{
		bookShelf: bs,
		index:     0,
	}
}

// GetReverseIterator 获取反向迭代器
func (bs *BookShelf) GetReverseIterator() Iterator {
	return &ReverseBookIterator{
		bookShelf: bs,
		index:     len(bs.books) - 1,
	}
}

// GetFilterIterator 获取过滤迭代器
func (bs *BookShelf) GetFilterIterator(filter func(*Book) bool) Iterator {
	return &FilterBookIterator{
		bookShelf: bs,
		index:     0,
		filter:    filter,
	}
}

// BookIterator 图书迭代器
type BookIterator struct {
	bookShelf *BookShelf
	index     int
}

func (it *BookIterator) HasNext() bool {
	return it.index < len(it.bookShelf.books)
}

func (it *BookIterator) Next() interface{} {
	if it.HasNext() {
		book := it.bookShelf.books[it.index]
		it.index++
		return book
	}
	return nil
}

func (it *BookIterator) CurrentItem() interface{} {
	if it.index < len(it.bookShelf.books) {
		return it.bookShelf.books[it.index]
	}
	return nil
}

func (it *BookIterator) Reset() {
	it.index = 0
}

// ReverseBookIterator 反向图书迭代器
type ReverseBookIterator struct {
	bookShelf *BookShelf
	index     int
}

func (it *ReverseBookIterator) HasNext() bool {
	return it.index >= 0
}

func (it *ReverseBookIterator) Next() interface{} {
	if it.HasNext() {
		book := it.bookShelf.books[it.index]
		it.index--
		return book
	}
	return nil
}

func (it *ReverseBookIterator) CurrentItem() interface{} {
	if it.index >= 0 {
		return it.bookShelf.books[it.index]
	}
	return nil
}

func (it *ReverseBookIterator) Reset() {
	it.index = len(it.bookShelf.books) - 1
}

// FilterBookIterator 过滤图书迭代器
type FilterBookIterator struct {
	bookShelf *BookShelf
	index     int
	filter    func(*Book) bool
}

func (it *FilterBookIterator) HasNext() bool {
	for i := it.index; i < len(it.bookShelf.books); i++ {
		if it.filter(it.bookShelf.books[i]) {
			return true
		}
	}
	return false
}

func (it *FilterBookIterator) Next() interface{} {
	for i := it.index; i < len(it.bookShelf.books); i++ {
		book := it.bookShelf.books[i]
		if it.filter(book) {
			it.index = i + 1
			return book
		}
	}
	return nil
}

func (it *FilterBookIterator) CurrentItem() interface{} {
	if it.index < len(it.bookShelf.books) {
		book := it.bookShelf.books[it.index]
		if it.filter(book) {
			return book
		}
	}
	return nil
}

func (it *FilterBookIterator) Reset() {
	it.index = 0
}
