package controller

func responseTemplate(t string) (res string) {
	switch t {
	case "oksave":
		res = "record saved"
	case "okupdate":
		res = "record updated"
	case "okdelete":
		res = "record deleted"
	case "failsave":
		res = "failed to save record"
	case "failupdate":
		res = "failed to update record"
	case "faildelete":
		res = "failed to delete record"
	default:
		res = ""
	}

	return
}
