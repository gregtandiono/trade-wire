package adaptors

import "errors"

// ResponseTemplate generates an interface response for JSON responses
func ResponseTemplate(resType string) (res map[string]string) {
	switch resType {
	case "insert:fail":
		res = map[string]string{
			"error": "failed to insert record",
		}
	case "insert:success":
		res = map[string]string{
			"message": "record succesfully created",
		}
	case "fetch:fail":
		res = map[string]string{
			"error": "failed to fetch record(s)",
		}
	case "update:fail":
		res = map[string]string{
			"error": "failed to update record",
		}
	case "update:success":
		res = map[string]string{
			"message": "record successfully updated",
		}
	case "delete:fail":
		res = map[string]string{
			"error": "failed to delete record",
		}
	case "delete:success":
		res = map[string]string{
			"message": "record successfully deleted",
		}
	default:
		// no op
		err := errors.New("no resType found")
		panic(err)
	}

	return
}
