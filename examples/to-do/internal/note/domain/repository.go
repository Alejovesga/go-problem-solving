package domain

type Repository interface {
	Save(n *Note) error
}
