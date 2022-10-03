package main

func main() {
	var v interface{}

	v = 20

	//v = v * 9

	//type assertion
	if value, ok := v.(int); ok == true {
		v = value * 9
	}
}
