package account

import "context"

type Service interface {
	CreateUser(cxt context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}
