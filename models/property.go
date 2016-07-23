package models

type Property struct {
	Address string  `json:"address"`
	Lng     float32 `json:"lng"`
	Lat     float32 `json:"lat"`
}

func AllProperties() []Property {
	var p Property
	properties := make([]Property, 0)

	rows, err := db.Query("SELECT address, latitude, longitude FROM properties")
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

func (p *Property) Create() error{
  _, err := db.Exec("INSERT INTO properties(address, latitude, longitude) VALUES($1, $2, $3)",
    p.Address, p.Lat, p.Lng)

  return err
}