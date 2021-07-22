package main

import (
	"fmt"
	"github.com/scritch007/fiscrap"
	"io"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Printf("Usage %s url download_or_not\n", os.Args[0])
		os.Exit(1)
	}
	s := fiscrap.New()
	res, err := s.Parse(os.Args[1])
	if err != nil {
		panic(err)
	}
	for k, v := range res {
		if os.Args[2] == "true" {
			fmt.Printf("Downloading %s\n", v)
			resp, err := http.DefaultClient.Get(v)
			if err != nil {
				fmt.Printf("Couldn't download file %s", v)
				os.Exit(2)
			}
			f, err := os.Create(fmt.Sprintf("%s.mp3", k))
			if err != nil {
				fmt.Printf("Couldn't create file %s", k)
				os.Exit(3)
			}
			_, err = io.Copy(f, resp.Body)
			if err != nil {
				panic(err)
			}
			f.Close()
			resp.Body.Close()
		} else {
			fmt.Printf(`"%s"=>%s
`, k, v)
		}
	}

}
