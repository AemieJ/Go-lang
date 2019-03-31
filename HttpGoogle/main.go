package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logwriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	/*bs := make([]byte, 99999) //creates a byte slice with 99999 elements
	resp.Body.Read(bs)        //reads the entire body
	fmt.Println(string(bs))*/
	lw := logwriter{}
	io.Copy(lw, resp.Body)
	/*EXPLANATION OF THIS CODE :
	so basically we know what reader interface is so opp to it
	is writer interface which takes in the byte slice and gives
	the output like outgoing of http request and terminal and all.

	So we have some values in our program already which like basically provides
	the input to the writer interface .

	type Writer interface {
		Write(p []byte) ( n int , err error)
	}
	it is similar to read but the byte slice in write is truly used as a source
	of input .

	NOW FOR THE io.copy :
	func Copy(w Writer , r Reader) 	(written int64,err error)
	*/
}

/*In the body field from docs it is from io.ReadCloser which is
an interface of Reader and Closer .So the body can be anything which
satisfied the interface . Now the reader and  closer are interfaces in
io.ReadCloser
io.Reader Interface : typr Reader interface {
	Read(p []byte) (n int,err error)
}
n is the number of bytes that were just read in the slice
err will just tell like something went wrong

And for ReaderCloser interface we need both the interfaces
Reader and Closer too.

io.Closer : type Closer interface {
	close() error
}
*/
func (logwriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	//int will be number of bytes written from bs
	return len(bs), nil
}
