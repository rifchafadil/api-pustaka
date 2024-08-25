package book

type BookService interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookRequest BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository BookRepository
}

func NewService(repository BookRepository) *service  {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	
	// return s.repository.FindAll()
}

func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err

	// return s.repository.FindAll()
}
func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book := Book {
		Title: bookRequest.Title,
		Price: int(price),
		Description: bookRequest.Description,
		Rating: int(rating),
		Discount: int(discount),
	}

	newBook, err := s.repository.Create(book)

	return newBook, err

	// return s.repository.FindAll()
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	book, _ := s.repository.FindById(ID)

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()
	
	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = bookRequest.Description
	book.Rating = int(rating)
	book.Discount = int(discount)
	
	newBook, err := s.repository.Update(book)

	return newBook, err

	// return s.repository.FindAll()
}

func (s *service) Delete(ID int) (Book, error) {
	book, _ := s.repository.FindById(ID)

	newBook, err := s.repository.Delete(book)

	return newBook, err

	// return s.repository.FindAll()
}


