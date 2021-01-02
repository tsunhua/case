package main

func main(){
	f1()
}

func f1() {
	jz := 1024
	for {
		jz = jz + 60
		if jz > 2020 {
			println(jz-60, jz)
			break
		}
	}
}