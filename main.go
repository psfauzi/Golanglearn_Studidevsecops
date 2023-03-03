package main

import "fmt" //paket untuk mencetak fungsi yang sudah dibuat ke consol/terminal

// membuat interface bernama bookstore dengan metode Addbook
type Bookstore interface {
	AddBook(book Book) //parameter book dan tipe data book
}

// Mendefinisikan nama  book dengan tipe data struct dan mendifiniskan beberapa jenis properti dibawah ini.
type Book struct {
	ID        int
	Title     string
	Author    string
	Quantity  int
	Available bool
}

// Mendefinisikan nama BookstoreImpl dengan tipe data struct
type BookstoreImpl struct {
	books []Book //membuat properti books dengan tipe data array []book
}

/*
metode AddBook pada tipe data BookstoreImpl digunakan untuk menambahkan buku baru ke dalam
toko buku dengan menambahkan buku tersebut ke dalam array books pada tipe data BookstoreImpl.
*/
func (b *BookstoreImpl) AddBook(book Book) error {
	b.books = append(b.books, book) //
	return nil
}

//Membuat fungsi getbook
func (b *BookstoreImpl) GetBook(id int) (*Book, error) {
	for i := range b.books {
		if b.books[i].ID == id {
			return &b.books[i], nil
		}
	}
	return nil, fmt.Errorf("Book with ID %v Not Found", id)
}

//Membuat fungsi mengahapus buku
func (b *BookstoreImpl) DeleteBook(id int) error {
	for i, book := range b.books {
		if book.ID == id {
			b.books = append(b.books[:i], b.books[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Book with ID %v Not Found", id)
}

func main() {
	// Membuat objek BookstoreImpl
	bookstore := &BookstoreImpl{}

	// Data buku yang ingin ditambahkan dalam bentuk array.
	books := [4]Book{
		{ID: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5, Available: true},
		{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee", Quantity: 3, Available: true},
		{ID: 3, Title: "1984", Author: "George Orwell", Quantity: 10, Available: false},
		{ID: 4, Title: "1984", Author: "Fauzi Bariq Mahya", Quantity: 23, Available: true},
	}

	//Membuat Prosess Menambahkan buku
	for _, book := range books {
		err := bookstore.AddBook(book)
		if err != nil {
			fmt.Printf("Ooups,,,Failed Adding Books: %v\n", err)
		}
	}

	// Dapatkan data buku berdasarkan ID
	bookID := 1
	book, err := bookstore.GetBook(bookID)
	if err != nil {
		fmt.Printf("Failed to retreive books by ID %v: %v\n", bookID, err)
	} else {
		fmt.Printf("List of books with ID %v:\n", bookID)
		fmt.Printf("ID: %v, Title: %v, Author: %v, Quantity: %v, Available: %v\n",
			book.ID, book.Title, book.Author, book.Quantity, book.Available)
	}

	//Prosess Mendelete Buku
	booksID := 3
	err = bookstore.DeleteBook(booksID)
	if err != nil {
		fmt.Printf("Failed Delete books with ID %v: %v\n", bookID, err)
	} else {
		fmt.Printf("Books with ID %v has been succesfully deleted\n", bookID)
	}

	// Cetak data buku yang sudah ditambahkan
	fmt.Printf("The list of books was succesfully added:\n")
	for _, book := range bookstore.books {
		fmt.Printf("ID: %v, Title: %v, Author: %v, Quantity: %v, Available: %v\n",
			book.ID, book.Title, book.Author, book.Quantity, book.Available)
	}

}
