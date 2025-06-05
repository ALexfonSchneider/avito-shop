package merch

import (
	"context"
	merchdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/merch"
	"time"
)

type Deps interface {
	CreateMerch(ctx context.Context, merch *merchdomain.Merch) error
	FindMerchByName(ctx context.Context, name string) (*merchdomain.Merch, error)
}

var Catalog = []struct {
	Name  string
	Price int64
}{
	{"t-shirt", 80},
	{"cup", 20},
	{"book", 50},
	{"pen", 10},
	{"powerbank", 200},
	{"hoody", 300},
	{"umbrella", 200},
	{"socks", 10},
	{"wallet", 50},
	{"pink-hoody", 500},
}

type InitMerch struct {
	deps Deps
}

func New(deps Deps) *InitMerch {
	return &InitMerch{
		deps: deps,
	}
}

func (i *InitMerch) Init(ctx context.Context) error {
	for _, item := range Catalog {
		merchStored, err := i.deps.FindMerchByName(ctx, item.Name)
		if err != nil {
			return err
		}

		if merchStored != nil {
			continue
		}

		merch := merchdomain.NewMerch(item.Name, "", item.Price, time.Now())
		if err = merch.Validate(); err != nil {
			return err
		}

		if err = i.deps.CreateMerch(ctx, merch); err != nil {
			return err
		}
	}

	return nil
}
