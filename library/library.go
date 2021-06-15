package library

import "fmt"

// huruf besar pertama bersifat public (exported)
func SayHello(name string) {
	fmt.Println("hello")
	introduce(name)
}

// huruf kecil pertama bersifat private (undexported)
func introduce(name string) {
	fmt.Println("nama saya", name)
}

// private
type student struct {
	name string // private
	age  string // private
}

type People struct {
	Name string // private
	age  int    // private
}
