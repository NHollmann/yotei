package main

import "os"

func main() {
	print("Hello ")
	println("World")
	println(os.Getenv("YOTEI_ADDR"))
}
