//package BusinessPortal
package main

const (
	INDEX_EMPTY = iota

	ASSET_INDEX_HOUSE
	ASSET_INDEX_FLAT
	ASSET_INDEX_SHOP

	GENERATOR_INDEX_FOOD
	GENERATOR_INDEX_WOOD
	GENERATOR_INDEX_METAL
	GENERATOR_INDEX_CEMENT
	GENERATOR_INDEX_CLOTHES
	GENERATOR_INDEX_FUEL
)

type Building struct {
	Type  int `bson:"Type" json:"Type"`
	Level int `bson:"Level" json:"Level"`
}

type BuildingInfo struct {
	Production [10]float32   `bson:"Production" json:"Production"`
	BuildTime  [10]int       `bson:"BuildTime" json:"BuildTime"`
	Cost       [10]Resources `bson:"Cost" json:"Cost"`
}

var BUILDING_INFO = []BuildingInfo{
	EMPTY_INFO,
	ASSET_INFO_HOUSE,
	ASSET_INFO_FLAT,
	ASSET_INFO_SHOP,
	GENERATOR_INFO_FOOD,
	GENERATOR_INFO_WOOD,
	GENERATOR_INFO_METAL,
	GENERATOR_INFO_CEMENT,
	GENERATOR_INFO_CLOTHES,
	GENERATOR_INFO_FUEL,
}
var EMPTY_INFO = BuildingInfo{
	Production: [10]float32{20, 30, 50, 100, 200, 300, 500, 1000, 1200, 1500},
	BuildTime:  [10]int{10, 20, 30, 120, 180, 300, 600, 900, 1200, 1500},
	Cost: [10]Resources{
		Resources{0, 0, 0, 0, 0, 0, 0},
		Resources{0, 0, 0, 0, 0, 0, 0},
		Resources{0, 0, 0, 0, 0, 0, 0},
		Resources{0, 0, 0, 0, 0, 0, 0},
		Resources{0, 0, 0, 0, 0, 0, 0},
		Resources{0, 0, 0, 0, 0, 0, 0},
		Resources{0, 0, 0, 0, 0, 0, 0},
		Resources{0, 0, 0, 0, 0, 0, 0},
		Resources{0, 0, 0, 0, 0, 0, 0},
		Resources{0, 0, 0, 0, 0, 0, 0},
	},
}
var ASSET_INFO_HOUSE = BuildingInfo{
	Production: [10]float32{20, 30, 50, 100, 200, 300, 500, 1000, 1200, 1500},
	BuildTime:  [10]int{10, 20, 30, 120, 180, 300, 600, 900, 1200, 1500},
	Cost: [10]Resources{
		Resources{100, 0, 100, 100, 100, 0, 10},
		Resources{100, 0, 100, 100, 100, 0, 10},
		Resources{100, 0, 100, 100, 100, 0, 10},
		Resources{100, 0, 100, 100, 100, 0, 10},
		Resources{100, 0, 100, 100, 100, 0, 10},
		Resources{100, 0, 100, 100, 100, 0, 10},
		Resources{100, 0, 100, 100, 100, 0, 10},
		Resources{100, 0, 100, 100, 100, 0, 10},
		Resources{100, 0, 100, 100, 100, 0, 10},
		Resources{100, 0, 100, 100, 100, 0, 10},
	},
}

var ASSET_INFO_FLAT = BuildingInfo{
	Production: [10]float32{100, 200, 300, 500, 1000, 1200, 1300, 1500, 1700, 2000},
	BuildTime:  [10]int{10, 20, 30, 300, 600, 900, 1200, 1800, 2700, 3600},
	Cost: [10]Resources{
		Resources{150, 0, 150, 150, 150, 0, 20},
		Resources{150, 0, 150, 150, 150, 0, 20},
		Resources{150, 0, 150, 150, 150, 0, 20},
		Resources{150, 0, 150, 150, 150, 0, 20},
		Resources{150, 0, 150, 150, 150, 0, 20},
		Resources{150, 0, 150, 150, 150, 0, 20},
		Resources{150, 0, 150, 150, 150, 0, 20},
		Resources{150, 0, 150, 150, 150, 0, 20},
		Resources{150, 0, 150, 150, 150, 0, 20},
		Resources{150, 0, 150, 150, 150, 0, 20},
	},
}

var ASSET_INFO_SHOP = BuildingInfo{
	Production: [10]float32{500, 700, 1000, 1200, 1300, 1500, 1700, 2000, 2200, 2500},
	BuildTime:  [10]int{10, 20, 30, 300, 600, 1200, 2700, 3600, 5400, 7200},
	Cost: [10]Resources{
		Resources{200, 0, 200, 200, 200, 0, 30},
		Resources{200, 0, 200, 200, 200, 0, 30},
		Resources{200, 0, 200, 200, 200, 0, 30},
		Resources{200, 0, 200, 200, 200, 0, 30},
		Resources{200, 0, 200, 200, 200, 0, 30},
		Resources{200, 0, 200, 200, 200, 0, 30},
		Resources{200, 0, 200, 200, 200, 0, 30},
		Resources{200, 0, 200, 200, 200, 0, 30},
		Resources{200, 0, 200, 200, 200, 0, 30},
		Resources{200, 0, 200, 200, 200, 0, 30},
	},
}

var GENERATOR_INFO_FOOD = BuildingInfo{
	Production: [10]float32{100, 200, 300, 500, 1000, 1200, 1300, 1500, 1700, 2000},
	BuildTime:  [10]int{10, 20, 30, 300, 600, 900, 1200, 1800, 2700, 3600},
	Cost: [10]Resources{
		Resources{100, 0, 200, 0, 0, 0, 30},
		Resources{100, 0, 200, 0, 0, 0, 30},
		Resources{100, 0, 200, 0, 0, 0, 30},
		Resources{100, 0, 200, 0, 0, 0, 30},
		Resources{100, 0, 200, 0, 0, 0, 30},
		Resources{100, 0, 200, 0, 0, 0, 30},
		Resources{100, 0, 200, 0, 0, 0, 30},
		Resources{100, 0, 200, 0, 0, 0, 30},
		Resources{100, 0, 200, 0, 0, 0, 30},
		Resources{100, 0, 200, 0, 0, 0, 30},
	},
}

var GENERATOR_INFO_WOOD = BuildingInfo{
	Production: [10]float32{100, 200, 300, 500, 1000, 1200, 1300, 1500, 1700, 2000},
	BuildTime:  [10]int{10, 20, 30, 300, 600, 900, 1200, 1800, 2700, 3600},
	Cost: [10]Resources{
		Resources{100, 0, 100, 200, 100, 0, 30},
		Resources{100, 0, 100, 200, 100, 0, 30},
		Resources{100, 0, 100, 200, 100, 0, 30},
		Resources{100, 0, 100, 200, 100, 0, 30},
		Resources{100, 0, 100, 200, 100, 0, 30},
		Resources{100, 0, 100, 200, 100, 0, 30},
		Resources{100, 0, 100, 200, 100, 0, 30},
		Resources{100, 0, 100, 200, 100, 0, 30},
		Resources{100, 0, 100, 200, 100, 0, 30},
		Resources{100, 0, 100, 200, 100, 0, 30},
	},
}

var GENERATOR_INFO_METAL = BuildingInfo{
	Production: [10]float32{100, 200, 300, 500, 1000, 1200, 1300, 1500, 1700, 2000},
	BuildTime:  [10]int{10, 20, 30, 300, 600, 900, 1200, 1800, 2700, 3600},
	Cost: [10]Resources{
		Resources{200, 0, 200, 100, 100, 0, 100},
		Resources{200, 0, 200, 100, 100, 0, 100},
		Resources{200, 0, 200, 100, 100, 0, 100},
		Resources{200, 0, 200, 100, 100, 0, 100},
		Resources{200, 0, 200, 100, 100, 0, 100},
		Resources{200, 0, 200, 100, 100, 0, 100},
		Resources{200, 0, 200, 100, 100, 0, 100},
		Resources{200, 0, 200, 100, 100, 0, 100},
		Resources{200, 0, 200, 100, 100, 0, 100},
		Resources{200, 0, 200, 100, 100, 0, 100},
	},
}

var GENERATOR_INFO_CEMENT = BuildingInfo{
	Production: [10]float32{100, 200, 300, 500, 1000, 1200, 1300, 1500, 1700, 2000},
	BuildTime:  [10]int{10, 20, 30, 300, 600, 900, 1200, 1800, 2700, 3600},
	Cost: [10]Resources{
		Resources{200, 0, 200, 100, 0, 0, 100},
		Resources{200, 0, 200, 100, 0, 0, 100},
		Resources{200, 0, 200, 100, 0, 0, 100},
		Resources{200, 0, 200, 100, 0, 0, 100},
		Resources{200, 0, 200, 100, 0, 0, 100},
		Resources{200, 0, 200, 100, 0, 0, 100},
		Resources{200, 0, 200, 100, 0, 0, 100},
		Resources{200, 0, 200, 100, 0, 0, 100},
		Resources{200, 0, 200, 100, 0, 0, 100},
		Resources{200, 0, 200, 100, 0, 0, 100},
	},
}

var GENERATOR_INFO_CLOTHES = BuildingInfo{
	Production: [10]float32{100, 200, 300, 500, 1000, 1200, 1300, 1500, 1700, 2000},
	BuildTime:  [10]int{10, 20, 30, 300, 600, 900, 1200, 1800, 2700, 3600},
	Cost: [10]Resources{
		Resources{200, 0, 100, 200, 0, 200, 100},
		Resources{200, 0, 100, 200, 0, 200, 100},
		Resources{200, 0, 100, 200, 0, 200, 100},
		Resources{200, 0, 100, 200, 0, 200, 100},
		Resources{200, 0, 100, 200, 0, 200, 100},
		Resources{200, 0, 100, 200, 0, 200, 100},
		Resources{200, 0, 100, 200, 0, 200, 100},
		Resources{200, 0, 100, 200, 0, 200, 100},
		Resources{200, 0, 100, 200, 0, 200, 100},
		Resources{200, 0, 100, 200, 0, 200, 100},
	},
}
var GENERATOR_INFO_FUEL = BuildingInfo{
	Production: [10]float32{100, 200, 300, 500, 1000, 1200, 1300, 1500, 1700, 2000},
	BuildTime:  [10]int{10, 20, 30, 300, 600, 900, 1200, 1800, 2700, 3600},
	Cost: [10]Resources{
		Resources{300, 0, 200, 300, 0, 300, 50},
		Resources{300, 0, 200, 300, 0, 300, 50},
		Resources{300, 0, 200, 300, 0, 300, 50},
		Resources{300, 0, 200, 300, 0, 300, 50},
		Resources{300, 0, 200, 300, 0, 300, 50},
		Resources{300, 0, 200, 300, 0, 300, 50},
		Resources{300, 0, 200, 300, 0, 300, 50},
		Resources{300, 0, 200, 300, 0, 300, 50},
		Resources{300, 0, 200, 300, 0, 300, 50},
		Resources{300, 0, 200, 300, 0, 300, 50},
	},
}
