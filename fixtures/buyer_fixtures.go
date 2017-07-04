package fixtures

// BuyerFixtures returns a set of buyer records
func BuyerFixtures() map[string]map[string]string {
	return map[string]map[string]string{
		"validBuyerRecord": {
			"id":           "f40e4dd4-f441-428b-8ff3-f893cb176819",
			"name":         "Japfa Comfeed Indonesia",
			"address":      `Wisma Millenia Lt. 7 \n Jl. MT. Haryono Kav. 16 \n Jakarta 12810, Indonesia \n T.021-2854 5680 (Hunting)`,
			"company_type": "buyer",
			// "pic": `[
			// 	{
			// 		"name": "Wahjudi Gunawan",
			// 		"position": "Head of Procurement",
			// 		"office_number": "+6221566712",
			// 		"cell_number": "+6281276488",
			// 		"notes": "Good guy, easy to talk to, prx must be very sharp"
			// 	},
			// 	{
			// 		"name": "Dewi Tjahyono",
			// 		"position": "Head of Operations",
			// 		"office_number": "+6221566712",
			// 		"cell_number": "+6281276322",
			// 		"notes": "Hard to deal with with what she wants"
			// 	}
			// ]`,
		},
		"invalidBuyerRecord": {
			"name": "",
		},
	}
}
