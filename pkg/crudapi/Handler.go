package crudapi

type Handler[F any] interface {
	Create() F
	ReadAll() F
	ReadOne() F
	Update() F
	Delete() F
}
