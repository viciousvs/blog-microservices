package post

type Repository interface {
	Create(post Post) (string, error)
	GetAll() ([]*Post, error)
	GetById(uuid string) (Post, error)
	Update(post Post) (string, error)
	Delete(uuid string) (string, error)
}
