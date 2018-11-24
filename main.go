package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var bolref string
var genData bool

func init() {
	flag.StringVar(&bolref, "bolref", "", "the bolref of the condition you would like to calculate")
	flag.BoolVar(&genData, "makedata", false, "generate mock data to run benchmarks with")
	flag.Parse()
}

func main() {

	driver, err := neo4j.NewDriver("bolt://localhost:7687",
		neo4j.BasicAuth("neo4j", "admin123", ""))
	if err != nil {
		log.Println(err)
		return
	}
	defer driver.Close()

	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		log.Println(err)
		return
	}
	defer session.Close()

	if genData {
		session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return nil, buildGraph(tx)
		})
		return
	}

	result, err := session.Run(`
		MATCH (c:Condition)<-[:MATCHED_TO]-(shoporder)
		WHERE c.bolref = $bolref
		RETURN c.value, c.type, shoporder.value as ordervalue`,
		map[string]interface{}{
			"bolref": bolref,
		})

	if err != nil {
		log.Fatal(err)
	}

	var total float64
	for result.Next() {
		r := result.Record()
		fmt.Println(r.Values())
		fmt.Println(r.Keys())
		cVal := r.GetByIndex(0).(float64)
		cType := r.GetByIndex(1).(string)
		oVal := r.GetByIndex(2).(float64)

		switch cType {
		case "percent":
			t := oVal * float64(cVal)
			total += t
		case "fixed":
			total += cVal
		}
	}

	if err != result.Err() {
		log.Println(err)
		return
	}

	fmt.Printf(" ================================\n")
	fmt.Printf(" Condition %s value is €%.2f\n", bolref, total)
	fmt.Printf(" ================================\n")
}
