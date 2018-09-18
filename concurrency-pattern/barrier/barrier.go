package barrier

import (
	"bytes"
	"io"
	"os"
)

/*	Barrier pattern : its purpose is simple,
	put up a barrier so that nobody passes until we have
	all the results we need
	-An HTTP GET aggregator-
	For our example, we are going to write a very typical situation
	in a microservices application-an app that performs
	two HTTP GET calls and joins them in a single response
	that will be printed on the console.*/

func barrier(endpoints ...string) {}

func captureEndpoints(endpoints ...string) string {
	reader, writer, _ := os.Pipe()

	os.Stdout = writer
	out := make(chan string)

	go func() {
		var buff bytes.Buffer
		io.Copy(&buff, reader)
		out <- buff.String()
	}()

	barrier(endpoints...)

	writer.Close()
	temp := <-out
	return temp
}
