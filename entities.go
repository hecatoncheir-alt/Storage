package storage

import (
	"time"
)

// City is a structure of prices in database
type City struct {
	ID       string `json:"uid"`
	Name     string `json:"cityName,omitempty"`
	IsActive bool   `json:"cityIsActive"`
}

// Price is a structure of ont price in database
type Price struct {
	ID        string    `json:"uid"`
	Value     float64   `json:"priceValue,omitempty"`
	DateTime  time.Time `json:"priceDateTime,omitempty"`
	IsActive  bool      `json:"priceIsActive"`
	Cities    []City    `json:"belongs_to_city,omitempty"`
	Products  []Product `json:"belongs_to_product,omitempty"`
	Companies []Company `json:"belongs_to_company,omitempty"`
}

// Product is a structure of ont product in database
type Product struct {
	ID               string     `json:"uid,omitempty"`
	Name             string     `json:"productName,omitempty"`
	IRI              string     `json:"productIri,omitempty"`
	PreviewImageLink string     `json:"previewImageLink,omitempty"`
	IsActive         bool       `json:"productIsActive"`
	Categories       []Category `json:"belongs_to_category,omitempty"`
	Companies        []Company  `json:"belongs_to_company,omitempty"`
	Prices           []Price    `json:"has_price,omitempty"`
}

// ProductsByNameForPage is a structure of products find in database for one page
type ProductsByNameForPage struct {
	Products []Product

	CurrentPage,
	TotalProductsForOnePage,
	TotalProductsFound int

	SearchedName,
	Language string
}

// Category is a structure of ont Category in database
type Category struct {
	ID        string    `json:"uid,omitempty"`
	Name      string    `json:"categoryName,omitempty"`
	IsActive  bool      `json:"categoryIsActive"`
	Companies []Company `json:"belongs_to_company,omitempty"`
	Products  []Product `json:"has_product,omitempty"`
}

// Company is a structure of ont Company in database
/* Для того что бы продукт принадлежащий компании отображался в категории принадлежащей
компании нужно иметь корректные данные belongs_to_company и belongs_to_category на гранях продукта */
type Company struct {
	ID         string     `json:"uid,omitempty"`
	IRI        string     `json:"companyIri,omitempty"`
	Name       string     `json:"companyName,omitempty"`
	Categories []Category `json:"has_category,omitempty"`
	Products   []Product  `json:"has_product,omitempty"`
	IsActive   bool       `json:"companyIsActive"`
}
