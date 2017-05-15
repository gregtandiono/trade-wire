package fixtures

// CommodityFixtures return a map of a map of valid and invalid commodity
// records for test purposes
func CommodityFixtures() map[string]map[string]string {
	return map[string]map[string]string{
		"validCommodityRecord": {
			"id":   "75a5cdfe-ca69-4680-a903-af89eaaa4804",
			"name": "wheat",
		},
		"validCommodityRecord2": {
			"id":   "55ebd6ea-c4ff-4d53-923e-0768d1fe86a6",
			"name": "corn",
		},
	}
}
