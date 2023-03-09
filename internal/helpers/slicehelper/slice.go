package slicehelper

import "strconv"

func UniqueInt(slices []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range slices {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func UniqueInt64(slices []int64) []int64 {
	keys := make(map[int64]bool)
	list := []int64{}
	for _, entry := range slices {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func UniqueString(slices []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slices {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func IsContainsString(slices []string, str string) bool {
	for _, slice := range slices {
		if str == slice {
			return true
		}
	}

	return false
}

func IsContainsInt64(slices []int64, val int64) bool {
	for _, slice := range slices {
		if val == slice {
			return true
		}
	}

	return false
}

func IsContainsInt(slices []int, val int) bool {
	for _, slice := range slices {
		if val == slice {
			return true
		}
	}

	return false
}

func SliceStringToInterface(slices []string) (results []interface{}) {
	for _, slice := range slices {
		results = append(results, slice)
	}
	return
}

func SliceIntToString(slices []int) (results []string) {
	for _, slice := range slices {
		results = append(results, strconv.Itoa(slice))
	}
	return
}

func SliceIntToInterface(slices []int) (results []interface{}) {
	for _, slice := range slices {
		results = append(results, slice)
	}
	return
}

func SliceInt64ToString(slices []int64) (results []string) {
	for _, slice := range slices {
		results = append(results, strconv.FormatInt(slice, 10))
	}
	return
}

func SliceInt64ToInterface(slices []int64) (results []interface{}) {
	for _, slice := range slices {
		results = append(results, slice)
	}
	return
}

func ChunkSliceInterface(data []interface{}, chunkSize int) (chunks [][]interface{}) {
	for {
		if len(data) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(data) < chunkSize {
			chunkSize = len(data)
		}

		chunks = append(chunks, data[0:chunkSize])
		data = data[chunkSize:]
	}

	return chunks
}

func ChunkSliceMapInterface(data []map[string]interface{}, chunkSize int) (chunks [][]map[string]interface{}) {
	for {
		if len(data) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(data) < chunkSize {
			chunkSize = len(data)
		}

		chunks = append(chunks, data[0:chunkSize])
		data = data[chunkSize:]
	}

	return chunks
}

func ChunkSliceString(data []string, chunkSize int) (chunks [][]string) {
	for {
		if len(data) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(data) < chunkSize {
			chunkSize = len(data)
		}

		chunks = append(chunks, data[0:chunkSize])
		data = data[chunkSize:]
	}

	return chunks
}

func ChunkSliceInt64(data []int64, chunkSize int) (chunks [][]int64) {
	for {
		if len(data) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(data) < chunkSize {
			chunkSize = len(data)
		}

		chunks = append(chunks, data[0:chunkSize])
		data = data[chunkSize:]
	}

	return chunks
}

func ChunkSliceInt(data []int, chunkSize int) (chunks [][]int) {
	for {
		if len(data) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(data) < chunkSize {
			chunkSize = len(data)
		}

		chunks = append(chunks, data[0:chunkSize])
		data = data[chunkSize:]
	}

	return chunks
}

func SliceInterfaceToInt(slices []interface{}) (results []int) {
	for _, slice := range slices {
		results = append(results, slice.(int))
	}
	return
}