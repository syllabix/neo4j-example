package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func buildGraph(tx neo4j.Transaction) error {
	gofakeit.Seed(11)
	for i := 0; i < 50; i++ {
		supplier := strings.Title(gofakeit.HipsterWord()) + " " + gofakeit.CompanySuffix()
		supplierID := fmt.Sprintf("%s-%d",
			strings.Replace(strings.ToUpper(supplier), " ", "", -1)[:2],
			gofakeit.Number(1000, 9999))

		stmt := new(strings.Builder)
		stmt.WriteString(fmt.Sprintf(`CREATE (s:Supplier {name: '%s', id: '%s'})`, supplier, supplierID))

		count := gofakeit.Number(1, 7)
		for count > 0 {
			bolref := fmt.Sprintf("V%s%d", gofakeit.HackerAdjective()[:2], gofakeit.Number(10000, 99999))
			id := gofakeit.Uint64()
			value := (rand.Float64() * 8) + 1
			ean := gofakeit.CreditCardNumber()
			cdtype := gofakeit.RandString([]string{"fixed", "percent", "percent"})
			start := gofakeit.DateRange(time.Now(), time.Now().Add(time.Hour*24*14))
			end := start.Add(time.Hour * 24 * time.Duration(gofakeit.Number(60, 150)))

			stmt.WriteString(fmt.Sprintf(`CREATE (c%d%d:Condition {id: %d, bolref: '%s', value: %.2f, type: '%s', ean: '%d', start: '%s', end: '%s'})`,
				count, i,
				id, strings.ToUpper(bolref),
				value, cdtype, ean,
				start.Format("2006-01-02"), end.Format("2006-01-02")))

			stmt.WriteString(fmt.Sprintf(`CREATE (s)-[:AGREED_TO]->(c%d%d)`, count, i))

			count--
		}
		// fmt.Println(stmt.String())
		res, err := tx.Run(stmt.String(), nil)
		if err != nil {
			log.Fatal(err)
		}

		if res.Err() != nil {
			log.Fatal(res.Err())
		} else {
			_, err := res.Summary()
			if err != nil {
				log.Fatal(err)
			}
			// fmt.Printf("%+v\n", summary.Counters())
		}
	}
	fmt.Println("Done :)")
	return nil
}
