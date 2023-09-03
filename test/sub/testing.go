package sub

import (
	"github.com/google/uuid"
)

//go:generate go run ../../main.go go-builder-generator/test/sub.Product
type Product struct {
	ID                uuid.UUID
	ArticleNumber     *string
	Name              string
	Description       string
	Color             string
	Pippo             *Whatever
	Size              string
	StockAvailability int
	PriceCents        int
	OnSale            bool
}

//go:generate go run ../../main.go go-builder-generator/test/sub.Whatever
type Whatever struct {}
