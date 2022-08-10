package service

import (
	"context"
	"encoding/json"

	"github.com/itsapep/yopei-grpc/server/repository"
)

// type YopeiService interface {
//  CheckBalance(ctx context.Context, in *CheckBalanceMessage) (*ResultMessage, error)
//  DoPayment(ctx context.Context, in *PaymentMessage) (*ResultMessage, error)
// }

type YopeiService struct {
	repo repository.YopeiRepository
	UnimplementedYopeiPaymentServer
}

// CheckBalance implements YopeiService
func (l *YopeiService) CheckBalance(ctx context.Context, in *CheckBalanceMessage) (*ResultMessage, error) {
	yopeiId := in.YopeiId
	customer, err := l.repo.RetrieveById(yopeiId)
	if err != nil {
		return nil, err
	}

	c, err := json.Marshal(customer)
	if err != nil {
		return nil, err
	}
	resultMessage := &ResultMessage{
		Result: string(c),
		Error:  nil,
	}
	return resultMessage, nil
}

// DoPayment implements YopeiService
func (l *YopeiService) DoPayment(ctx context.Context, in *PaymentMessage) (*ResultMessage, error) {
	yopeiId := in.YopeiId
	amount := in.Amount
	customer, err := l.repo.RetrieveById(yopeiId)
	if err != nil {
		return nil, err
	}

	if customer.Balance < amount {
		return &ResultMessage{
			Result: "FAILED",
			Error: &Error{
				Code:    "X07",
				Message: "Insufficient Balance",
			},
		}, nil
	}

	resultMessage := &ResultMessage{
		Result: "SUCCESS",
		Error:  nil,
	}
	return resultMessage, nil
}

func NewYopeiService(repo repository.YopeiRepository) *YopeiService {
	service := new(YopeiService)
	service.repo = repo
	return service
}
