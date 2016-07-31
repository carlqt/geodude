package models

import (
	"github.com/carlqt/geodude/geocode"
	"github.com/shopspring/decimal"
)

type Property struct {
	ID      int     `json:"id"`
	Name    string 	`json:"name"`
	Address string  `json:"address"`
	Price		decimal.Decimal		`json:"price"`
	Description string `json:"description"`
	Type    string	`json:"type"`
	Lng     float32 `json:"lng"`
	Lat     float32 `json:"lat"`
}

func AllProperties() []Property {
	var p Property
	properties := make([]Property, 0)

	rows, err := db.Query("SELECT id, address, latitude, longitude FROM properties")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&p.ID, &p.Address, &p.Lat, &p.Lng)
		properties = append(properties, p)
	}

	return properties
}

func (p *Property) Create() error {
	err := db.QueryRow("INSERT INTO properties(address, latitude, longitude, type, description, price, name) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		p.Address, p.Lat, p.Lng, p.Type, p.Description, p.Price.String(), p.Name).Scan(&p.ID)

	return err
}

func (p *Property) GeocodeAndCreate() (*Property, error) {
	results, err := geocode.Geocode(p.Address)

	if err != nil {
		return nil, err
	}

	p.Lat = results.Geometry.Location["lat"]
	p.Lng = results.Geometry.Location["lng"]
	p.Address = results.FormattedAddress

	// err = db.QueryRow("INSERT INTO properties(address, latitude, longitude, type, description, price, name) VALUES($1, $2, $3) RETURNING id",
	// 	p.Address, p.Lat, p.Lng).Scan(&p.ID)
	err = p.Create()

	return p, err
}

func NearbyProperty(lat float32, lng float32) []Property {
	var p Property

	properties := make([]Property, 0)

	// Given the lat and lng, search the database for nearby address within a 10km radius
	rows, err := db.Query(`SELECT address, latitude, longitude FROM properties 
		WHERE earth_box(ll_to_earth($1, $2), 200) @> ll_to_earth(latitude, longitude)`, lat, lng)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&p.Address, &p.Lat, &p.Lng)
		properties = append(properties, p)
	}

	return properties
}

func (p *Property) Delete() error {
	_, err := db.Exec("DELETE FROM properties WHERE id = $1", p.ID)
	return err
}

func PropertyDelete(id int) error {
	_, err := db.Exec("DELETE FROM properties WHERE id = $1", id)
	return err
}
