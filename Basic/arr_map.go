package main

import "fmt"

//切片转换为Map
func sliceToMap(s_key, s_value []string) map[string]string {
	mapObj := map[string]string{}
	for s_key_index := range s_key {
		mapObj[s_key[s_key_index]] = s_value[s_key_index]
	}
	return mapObj
}

func main() {
	key := []string{"a", "b", "c", "d", "e"}
	value := []string{"1", "2", "3", "4", "5"}
	map := slice_To_Map(s_key, s_value)

	for key, value := range map {
		fmt.Printf("Key: %s,Value: %s\n", key, value)
	}
}