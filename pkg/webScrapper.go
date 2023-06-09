package pkg

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

func GetInfo() []string {
	c := colly.NewCollector()
	var info []string
	c.OnHTML("div[class=text-center]", func(e *colly.HTMLElement) {
		info = append(info, strings.Trim(e.Text, " \n"))
	})

	err := c.Visit("https://cabinet.unimetriq.com/client/b3e7bd77b5e8b89bedcf7d5f57021a84/?nonAuth=1")
	if err != nil {
		log.Fatalln(err)
	}
	return info
}

func Reply(info []string) {
	var numberOfFreeMachines int
	for _, elem := range info {
		if elem == "Свободно" {
			numberOfFreeMachines += 1
		}
	}
	switch numberOfFreeMachines {
	case 0:
		fmt.Println("Все машинки заняты")
	case 1:
		fmt.Println("Свободно", numberOfFreeMachines, "стиральная машина из", len(info))
	case 3:
		fmt.Println("Все машинки свободны")
	default:
		fmt.Println("Свободно", numberOfFreeMachines, "стиральные машины из", len(info))
	}
}
