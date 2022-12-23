package write

import (
	. "EasyDI/print"
	. "EasyDI/read"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func WriteData() {
	total := 0.0
	for _, p := range Products {
		total += p.Price
	}

	dataStr := fmt.Sprintf("Time: %v, Total: $%.2f\n", time.Now().Format("Mon 15:04:05"), total)
	file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		defer file.Close()
		file.WriteString(dataStr)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func WriteJson() {
	cheapProducts := []Product{}
	for _, p := range Products {
		if p.Price < 100 {
			cheapProducts = append(cheapProducts, p)
		}
	}
	file, err := os.OpenFile("cheap.json", os.O_WRONLY|os.O_CREATE, 0666)
	if err == nil {
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
