package models

type Location struct {
	Lat  float64 `bson:"lat" json:"lat"`
	Long float64 `bson:"long" json:"long"`
}

type Item struct {
	Name   string `bson:"name" json:"name"`
	Points int    `bson:"points" json:"points"`
}

type InventoryItem struct {
	Item     Item `bson:"item" json:"item"`
	Quantity int  `bson:"quantity" json:"quantity"`
}

type Survivor struct {
	Id                   string          `bson:"_id,omitempty" json:"id"`
	Name                 string          `bson:"name" json:"name"`
	Age                  int             `bson:"age" json:"age"`
	LastLocation         Location        `bson:"last_location" json:"last_location"`
	CountInfectionReport int             `bson:"count_infection_report" json:"count_infection_report"`
	InfectedFlag         bool            `bson:"infected_flag" json:"infected_flag"`
	Gender               string          `bson:"gender" json:"gender"`
	Items                []InventoryItem `bson:"itens" json:"itens"`
}
