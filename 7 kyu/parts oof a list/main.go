package main

func PartList(arr []string) string {
	res := ""
	for i := 0; i < len(arr)-1; i++ {
		for j := range arr {
			if j == 0 {
				res += "(" + arr[j]
				continue
			}
			if i+1 == j {
				res += ","
			}
			res += " " + arr[j]

			if j == len(arr)-1 {
				res += ")"
			}
		}
	}
	return res
}
