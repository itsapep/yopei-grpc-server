package repository

import "github.com/itsapep/yopei-grpc/server/model"

type YopeiRepository interface {
	RetrieveById(id int32) (model.Customer, error)
}

type yopeiRepository struct {
	db []model.Customer
}

// RetrieveById implements YopeiRepository
func (y *yopeiRepository) RetrieveById(id int32) (model.Customer, error) {
	for _, customer := range y.db {
		if customer.YopeiId == id {
			return customer, nil
		}
	}
	return model.Customer{}, nil
}

func NewYopeiRepository() YopeiRepository {
	repo := new(yopeiRepository)
	repo.db = []model.Customer{
		{YopeiId: 1, Balance: 2000000},
		{YopeiId: 2, Balance: 1000000},
		{YopeiId: 3, Balance: 1500000},
	}
	return repo
}
