package main

import "fmt"

const (
	gs = "甲乙丙丁戊己庚辛壬癸"
	zs = "子丑寅卯辰巳午未申酉戌亥"
)

func main() {
	//f1()
	//f2()
	f3()
}

func f1() {
	// 天干數組
	garr := []rune(gs)
	// 地支數組
	zarr := []rune(zs)
	// 保存合法的干支組合的Map
	// Key爲干支，Value爲天干索引與地支索引之和
	tdmap := make(map[string]int)

	// 將干支組合保存到Map中，直到出現第一個重複項就停止遍歷
	for i, j := 0, 0; i < len(garr) && j < len(zarr); i, j = (i+1)%len(garr), (j+1)%len(zarr) {
		k := string(garr[i]) + string(zarr[j])
		if _, ok := tdmap[k]; ok {
			break
		}
		tdmap[k] = i + j
	}

	// 輸出所有合法的組合
	println("\n所有合法的組合：")

	for k, v := range tdmap {
		fmt.Printf("%s:%d ", k, v)
	}

	// 輸出所有非法的組合
	println("\n所有非法的組合：")
	for i := 0; i < len(garr); i++ {
		for j := 0; j < len(zarr); j++ {
			k := string(garr[i]) + string(zarr[j])
			if _, ok := tdmap[k]; !ok {
				fmt.Printf("%s:%d ", k, i+j)
			}
		}
	}
}

func f2() {
	garr := []rune(gs)
	zarr := []rune(zs)

	// 輸出所有非法的組合
	println("\n所有非法的組合：")
	for i := 0; i < len(garr); i++ {
		for j := 0; j < len(zarr); j++ {
			if (i+j)%2 != 0 {
				fmt.Printf("%s:%d ", string(garr[i])+string(zarr[j]), i+j)
			}
		}
	}
}

func f3() {
	garr := []rune(gs)
	zarr := []rune(zs)

	// 輸出所有非法的組合
	println("\n所有非法的組合：")
	for i := 0; i < len(garr); i++ {
		if i%2 == 0 {
			for j := 1; j < len(zarr); j = j + 2 {
				fmt.Printf("%s:%d ", string(garr[i])+string(zarr[j]), i+j)
			}
		} else {
			for j := 0; j < len(zarr); j = j + 2 {
				fmt.Printf("%s:%d ", string(garr[i])+string(zarr[j]), i+j)
			}
		}
	}
}
