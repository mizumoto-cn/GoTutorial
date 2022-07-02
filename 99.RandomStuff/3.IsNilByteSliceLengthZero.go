package main

func main() {
	var data []byte
	data = nil
	println(len(data))
	// and it's 0
	// so when you do nil check for byte slice, just check length
}
