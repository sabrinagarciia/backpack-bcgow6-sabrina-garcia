package products

type Service interface {
	GetAll() ([]Product, error)
	Save(id int, name string, color string, price float64, stock int, code string, published bool, date string) (Product, error)
	Update(id int, name, color string, price float64, stock int, code string, published bool, date string) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Save(id int, name string, color string, price float64, stock int, code string, published bool, date string) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++

	producto, err := s.repository.Save(id, name, color, price, stock, code, published, date)
	if err != nil {
		return Product{}, err
	}

	return producto, nil
}

func (s *service) Update(id int, name string, color string, price float64, stock int, code string, published bool, date string) (Product, error) {

	return s.repository.Update(id, name, color, price, stock, code, published, date)
}