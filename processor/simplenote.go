package processor

import (
	"film/config"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"math/rand"
	"time"
)

func Run(config config.Config)  {

	doc, _ := goquery.NewDocument(config.Api.Endpoint)
	selected := doc.Find(config.Api.Key)
	rand.Seed(time.Now().Unix())

	fmt.Printf("Select from %v elements\n", len(selected.Nodes))
	fmt.Printf("And random select %v\n", selected.Nodes[rand.Intn(len(selected.Nodes))].Attr[0].Val)
}