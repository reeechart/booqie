package strings

import "strconv"

func ConvertInt32ArrayToStringArray(intArray *[]int32) string {
	str := "{"
	if intArray != nil {
		for i, v := range *intArray {
			str += strconv.Itoa(int(v))
			if i != len(*intArray)-1 {
				str += ","
			}
		}
	}
	str += "}"
	return str
}
