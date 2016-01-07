//package BusinessPortal
package main

const (
	RESOURCE_INDEX_MONEY   = "Money"
	RESOURCE_INDEX_FOOD    = "Food"
	RESOURCE_INDEX_WOOD    = "Wood"
	RESOURCE_INDEX_METAL   = "Metal"
	RESOURCE_INDEX_CEMENT  = "Cement"
	RESOURCE_INDEX_CLOTHES = "Clothes"
	RESOURCE_INDEX_FUEL    = "Fuel"
)

const (
	PRODUCTION_INDEX_MONEY = iota
	PRODUCTION_INDEX_FOOD
	PRODUCTION_INDEX_WOOD
	PRODUCTION_INDEX_METAL
	PRODUCTION_INDEX_CEMENT
	PRODUCTION_INDEX_CLOTHES
	PRODUCTION_INDEX_FUEL
)

type Resources struct {
	Money   float32 `bson:"Money" json:"Money"`
	Food    float32 `bson:"Food" json:"Food"`
	Wood    float32 `bson:"Wood" json:"Wood"`
	Metal   float32 `bson:"Metal" json:"Metal"`
	Cement  float32 `bson:"Cement" json:"Cement"`
	Clothes float32 `bson:"Clothes" json:"Clothes"`
	Fuel    float32 `bson:"Fuel" json:"Fuel"`
}

func (resources *Resources) Init() {
	resources.Money = 0
	resources.Food = 0
	resources.Wood = 0
	resources.Metal = 0
	resources.Cement = 0
	resources.Clothes = 0
	resources.Fuel = 0
}

func (resources *Resources) AssignResource(value int, resource string) {
	if resource == RESOURCE_INDEX_MONEY {
		resources.Money = float32(value)
	}
	if resource == RESOURCE_INDEX_FOOD {
		resources.Food = float32(value)
	}
	if resource == RESOURCE_INDEX_WOOD {
		resources.Wood = float32(value)
	}
	if resource == RESOURCE_INDEX_METAL {
		resources.Metal = float32(value)
	}
	if resource == RESOURCE_INDEX_CEMENT {
		resources.Cement = float32(value)
	}
	if resource == RESOURCE_INDEX_CLOTHES {
		resources.Clothes = float32(value)
	}
	if resource == RESOURCE_INDEX_FUEL {
		resources.Fuel = float32(value)
	}
}

var BUILDING_PRODUCTION_FACTOR = []Resources{
	Resources{1, 0, 0, 0, 0, 0, 0},       //Banker
	Resources{0.5, 0.8, 0, 0, 0, 0.8, 0}, //Dealer
	Resources{0.5, 0.5, 0, 0, 0, 0, 0.5}, //Transporter
	Resources{0.8, 0, 0, 0, 0, 0, 0},     //Engineer
	Resources{0, 1, 1, 1, 1, 1, 1},       //Entreprenuer
}

var DEFAULT_RESOURCES_BANKER = Resources{100000, 100, 100, 100, 100, 100, 0}
var DEFAULT_RESOURCES_DEALER = Resources{100, 20000, 20000, 20000, 20000, 100, 0}
var DEFAULT_RESOURCES_TRANSPORT = Resources{100, 0, 100, 100, 100, 100, 100000}
var DEFAULT_RESOURCES_ENGINEER = Resources{100, 100, 100, 100, 100, 100, 100}
var DEFAULT_RESOURCES_ENTREPRENEUR = Resources{10000, 100, 100, 100, 100, 100, 100}

var DEFAULT_PRODUCTION_BANKER = Resources{4000, 0, 0, 0, 0, 0, 0}
var DEFAULT_PRODUCTION_DEALER = Resources{0, 400, 900, 900, 900, 900, 0}
var DEFAULT_PRODUCTION_TRANSPORT = Resources{0, 0, 0, 0, 0, 0, 4000}
var DEFAULT_PRODUCTION_ENGINEER = Resources{3100, 300, 0, 0, 0, 300, 300}
var DEFAULT_PRODUCTION_ENTREPRENEUR = Resources{0, 1000, 1000, 1000, 1000, 300, 300}
