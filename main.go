package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/syllabix/neo4j-example/data"
	"github.com/syllabix/neo4j-example/data/graphdb"
	"github.com/syllabix/neo4j-example/neox"
)

var bolref string
var seed bool
var clear bool
var dbtype string
var supplierID string
var nx bool

var generator data.Generator

type calc struct {
	SupplierName string  `neo:"s.name"`
	CdnValue     float64 `neo:"value"`
	CdnType      string  `neo:"type"`
	OrderPrice   float64 `neo:"price"`
}

func init() {
	flag.StringVar(&bolref, "bolref", "", "the bolref of the condition you would like to calculate")
	flag.StringVar(&supplierID, "sid", "", "the supplier id of which you would like to calculate to current total value across all conditions")
	flag.BoolVar(&seed, "seed", false, "generate mock data to run benchmarks with")
	flag.BoolVar(&clear, "clear", false, "delete all nodes and relationships in the database")
	flag.StringVar(&dbtype, "dbtype", "neo4j", "the database type to work with")
	flag.BoolVar(&nx, "neox", false, "run with neox")
	flag.Parse()
}

func main() {

	driver, err := neox.NewDriver("bolt://localhost:7687",
		neo4j.BasicAuth("neo4j", "admin123", ""))

	if err != nil {
		log.Fatal(err)
		return
	}

	defer driver.Close()

	generator := graphdb.New(driver)

	if clear && seed {
		log.Fatal("Conflicting commands provided, -clear and -seed")
	}

	if seed {
		err := generator.Generate()
		if err != nil {
			log.Printf("[ERROR] - %v", err)
		}
		return
	}

	if clear {
		err := generator.Reset()
		if err != nil {
			log.Printf("[ERROR] - %v", err)
		}
		return
	}

	session, err := driver.Sessionx(neo4j.AccessModeWrite)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()

	clc := new(calc)

	if bolref != "" {
		start := time.Now()

		fmt.Println("Calculating total value of condition for BolRef", bolref)
		result, err := session.Runx(`
		MATCH (c:Condition)<-[:MATCHED_TO]-(order)
		WHERE c.bolref = $bolref
		RETURN c.value, c.type, toFloat(order.price) as price`,
			map[string]interface{}{
				"bolref": bolref,
			})

		if err != nil {
			log.Fatal(err)
		}

		var total float64
		for result.Next() {
			r := result.Record()

			cVal := r.GetByIndex(0).(float64)
			cType := r.GetByIndex(1).(string)
			oVal := r.GetByIndex(2).(float64)

			result.ToStruct(clc)

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

		duration := time.Since(start)

		fmt.Printf("\n================================\n")
		fmt.Printf("Condition %s value is €%.2f\n", bolref, total)
		fmt.Printf("================================\n\n\n")

		fmt.Printf("Calculation Duration: %v", duration)
	}

	if supplierID != "" {
		start := time.Now()
		fmt.Println()
		fmt.Println("Calculating total value of conditions for supplier id", supplierID)

		result, err := session.Runx(`
		MATCH (s:Supplier)-[:AGREED_TO]->(c:Condition)<-[:MATCHED_TO]-(order)
		WHERE s.id = $supplierId
		RETURN s.name, c.value as value, c.type as type, toFloat(order.price) as price`,
			map[string]interface{}{
				"supplierId": supplierID,
			})

		if err != nil {
			log.Fatal(err)
		}

		// var sname string
		var total float64
		for result.Next() {

			err = result.ToStruct(clc)
			if err != nil {
				log.Println(err)
			}

			switch clc.CdnType {
			case "percent":
				t := clc.OrderPrice * clc.CdnValue
				total += t
			case "fixed":
				total += clc.CdnValue
			}
		}

		duration := time.Since(start)

		if err != result.Err() {
			log.Println(err)
			return
		}

		fmt.Printf("\n=========================================================\n")
		fmt.Printf("%s has conditions with a total value of\n€%.2f\n", clc.SupplierName, total)
		fmt.Printf("=========================================================\n")

		fmt.Printf("Calculation Duration: %v\n\n", duration)
	}

}
