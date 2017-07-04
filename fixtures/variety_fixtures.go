package fixtures

func VarietyFixtures() map[string]map[string]string {
	return map[string]map[string]string{
		"validVarietyRecord": {
			"id":           "8f071b7e-555e-4c73-b9dd-2a86da728d32",
			"commodity_id": CommodityFixtures()["validCommodityRecord"]["id"],
			"name":         "Black Sea Wheat (BS Wheat)",
			"origin":       "Black Sea",
		},
		"validVarietyRecord2": {
			"id":           "ac3958c6-30c7-4f57-b6b1-e3255d292bb8",
			"commodity_id": CommodityFixtures()["validCommodityRecord2"]["id"],
			"name":         "Brazillian Corn (Bzl Corn)",
			"origin":       "Brazil",
		},
	}
}
