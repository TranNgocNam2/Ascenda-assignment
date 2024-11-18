package utils

import "strconv"

func ConvertInterfaceToFloat32(i interface{}) float32 {
	var f float64 = 0
	switch iType := i.(type) {
	case float64:
		f = iType
		break
	case string:
		f, err := strconv.ParseFloat(iType, 64)
		if err != nil {
			return float32(f)
		}
		break
	default:
		return 0
	}

	return float32(f)
}

func ChooseNonEmpty(existingValue, newValue string) string {
	if newValue != "" {
		return newValue
	}
	return existingValue
}

func ChooseNonZero(existingValue, newValue float32) float32 {
	if newValue != 0 {
		return newValue
	}
	return existingValue
}

func AppendUniqueStrSlice(existingValue, newValue []string) []string {
	unique := make(map[string]bool)
	for _, item := range existingValue {
		unique[item] = true
	}
	for _, item := range newValue {
		if !unique[item] {
			existingValue = append(existingValue, item)
			unique[item] = true
		}
	}
	return existingValue
}

func ContainsValue(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
