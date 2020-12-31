## case

已知干支紀年法中有 10 天干，12 地支，有 60 種合法組合，求所有非法的組合，並探討其中規律。

## solution

先說結論：

1. 所有非法的組合爲：
甲丑 甲卯 甲巳 甲未 甲酉 甲亥 乙子 乙寅 乙辰 乙午 乙申 乙戌 丙丑 丙卯 丙巳 丙未 丙酉 丙亥 丁子 丁寅 丁辰 丁午 丁申 丁戌 戊丑 戊卯 戊巳 戊未 戊酉 戊亥 己子 己寅 己辰 己午 辛寅 辛辰 辛午 辛申 辛戌 壬丑 壬卯 壬巳 壬未 壬酉 壬亥 癸子 癸寅 癸辰 癸午 癸申 癸戌
2. 其中的規律是：使用 1～10 給天干編號，1～12給地支編號，分別從天干和地支的編號中任取一個，如果其和爲偶數則是合法的組合，否則就是非法的組合。例如「甲丑」中「甲」的編號是 1，「丑」的編號是「2」，1+2=3 爲奇數，是非法的組合。

試著使用編程來解決此問題。

首先，將天干和地址的字符串定義如下：

```go
var gs = "甲乙丙丁戊己庚辛壬癸"
var zs = "子丑寅卯辰巳午未申酉戌亥"
```

第一種思路是：先求出所有合法的組合，然後窮盡所有組合，如組合不在合法組合中即是非法組合，輸出即可，編程到函數 f1 如下：

```go
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
```

輸出結果爲：

```go
所有合法的組合：
庚午:12 己卯:8 甲申:8 己酉:14 丙寅:4 丁卯:6 己巳:10 癸酉:18 戊寅:6 戊戌:14 癸巳:14 癸卯:12 甲辰:4 乙卯:4 辛巳:12 丙戌:12 壬辰:12 乙未:8 辛亥:18 戊午:10 丁巳:8 己未:12 辛 壬申:16 甲午:6 庚子:6 辛丑:8 庚戌:16 丙辰:6 戊辰:8 癸未:16 甲子:0 丙申:10 丁未:10 壬子:8 庚辰:10 甲戌:10 丁丑:4 壬午:14 丙子:2 庚寅:8 壬寅:10 戊申:12 甲寅:2 乙丑:2 辛未:14 丁亥:14 戊子:4 己亥:16 丙午:8 庚申:14 癸亥:20 丁酉:12 
所有非法的組合：
甲丑:1 甲卯:3 甲巳:5 甲未:7 甲酉:9 甲亥:11 乙子:1 乙寅:3 乙辰:5 乙午:7 乙申:9 乙戌:11 丙丑:3 丙卯:5 丙巳:7 丙未:9 丙酉:11 丙亥:13 丁子:3 丁寅:5 丁辰:7 丁午:9 丁申:11 丁 己辰:9 己午:11 己申:13 己戌:15 庚丑:7 庚卯:9 庚巳:11 庚未:13 庚酉:15 庚亥:17 辛子:7 辛寅:9 辛辰:11 辛午:13 辛申:15 辛戌:17 壬丑:9 壬卯:11 壬巳:13 壬未:15 壬酉:17 壬亥:19 癸子:9 癸寅:11 癸辰:13 癸午:15 癸申:17 癸戌:19 
```

嗯，符合預期，合法的組合和非法的組合各佔 60 個。

觀察發現，所有合法的組合的Value值都是偶數，所有非法的組合的Value值都是奇數。

因此有第二種思路求出非法組合：

```go
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
```

輸出結果：

```go
所有非法的組合：
甲丑:1 甲卯:3 甲巳:5 甲未:7 甲酉:9 甲亥:11 乙子:1 乙寅:3 乙辰:5 乙午:7 乙申:9 乙戌:11 丙丑:3 丙卯:5 丙巳:7 丙未:9 丙酉:11 丙亥:13 丁子:3 丁寅:5 丁辰:7 丁午:9 丁申:11 丁 己辰:9 己午:11 己申:13 己戌:15 庚丑:7 庚卯:9 庚巳:11 庚未:13 庚酉:15 庚亥:17 辛子:7 辛寅:9 辛辰:11 辛午:13 辛申:15 辛戌:17 壬丑:9 壬卯:11 壬巳:13 壬未:15 壬酉:17 壬亥:19 癸子:9 癸寅:11 癸辰:13 癸午:15 癸申:17 癸戌:19
```

很棒，跟思路一的非法組合一致。

學過小學數學的都知道，奇數+偶數=奇數，奇數+奇數=偶數，偶數+偶數=偶數。於是有了第三種思路，只要保證 i 和 j 的奇偶性不同即可保證 i 和 j 的組合爲非法組合。代碼如下：

```go
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
```

輸出結果同思路二的一致。