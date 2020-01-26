package utils

func StringSliceToInterfaces(list []string) []interface{} {
	result := make([]interface{}, len(list))

	for index, value := range list {
		result[index] = value
	}
	return result
}

func SliceFind(list []interface{}, callback func (item interface{}, index int) bool) interface{} {
	for index, item := range list {
		if isExist := callback(item, index); isExist {
			return item
		}
	}
	return nil
}
