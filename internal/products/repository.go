package products

import (
	"fmt"
	"test/pkg/store"
)

type Repository interface {
	GetAll() ([]Product, error)
	Create(id int, name, productType string, count int, price float64) (Product, error)
	LastID() (int, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(d store.Store) *repository {
	return &repository{
		db: d,
	}
}

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product
	r.db.Read(&ps)
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].ID, nil
}

func (r *repository) Create(id int, nombre, tipo string, cantidad int, precio float64) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	p := Product{id, nombre, tipo, cantidad, precio}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	p := Product{Name: name, Type: productType, Count: count, Price: price}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("producto %d no encontrado", id)
	}
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) UpdateName(id int, name string) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	var p Product
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Name = name
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("producto %d no encontrado", id)
	}
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	var ps []Product
	r.db.Read(&ps)
	deleted := false
	var index int
	for i := range ps {
		if ps[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("producto %d no encontrado", id)
	}
	ps = append(ps[:index], ps[index+1:]...)
	if err := r.db.Write(ps); err != nil {
		return err
	}
	return nil
}
