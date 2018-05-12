package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader)Read(b []byte) (int, error){
	var err error
	i := 0
	for ; i < len(b); i++ {
		b[i] = 'A'
	}

	return i, err
}


func main() {
	reader.Validate(MyReader{})
}
