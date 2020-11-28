package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := &Capper{os.Stdout}
	fmt.Fprintln(content, "Hello world")
}

//Capper type
type Capper struct {
	writer io.Writer
}

//Write content
func (content *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A')
	fmt.Printf("current diff : %d \n", diff)

	out := make([]byte, len(p))
	for i, content := range p {
		if content >= 'a' && content <= 'z' {
			fmt.Printf("current content : %d \n", content)
			content -= diff
		}
		out[i] = content
	}
	return content.writer.Write(out)
}
