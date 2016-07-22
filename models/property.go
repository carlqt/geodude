package models

type Property struct {
	Address string  `json:"address"`
	Lng     float32 `json:"lng"`
	Lat     float32 `json:"lat"`
}

func AllProperties() []Property {
	p := Property{}
	properties := make([]Property, 0)

	rows, _ := db.Query("SELECT address, latitude, longitude FROM properties")

	for rows.Next() {
		rows.Scan(&p.Address, &p.Lat, &p.Lng)
		properties = append(properties, p)
	}

	return properties
}
