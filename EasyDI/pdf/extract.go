package pdf

import (
	"fmt"
	"log"

	"code.sajari.com/docconv"
)

func Extract() {
	res, err := docconv.ConvertPath("./my.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
