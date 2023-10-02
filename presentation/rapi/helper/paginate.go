package helper

func GetOrder(order string) (string, string) {
	var orderRes string
	if order != "asc" && order != "desc" {
		orderRes = "desc"
	} else {
		orderRes = order
	}

	var operation string
	if orderRes == "asc" {
		operation = ">"
	} else {
		operation = "<"
	}

	return orderRes, operation
}
