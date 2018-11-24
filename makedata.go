package main

// func buildGraph(tx neo4j.Transaction) error {
// 	fake.Seed(11)
// 	for i := 0; i < 50; i++ {
// 		supplier := strings.Title(strings.Replace(fake.HipsterWord(), "'", "", -1) + " " + fake.CompanySuffix())
// 		supplierID := fmt.Sprintf("%s-%d",
// 			strings.Replace(strings.ToUpper(supplier), " ", "", -1)[:2],
// 			fake.Number(1000, 9999))

// 		stmt := new(strings.Builder)
// 		// stmt.WriteString(fmt.Sprintf(`CREATE (s:Supplier {name: '%s', id: '%s'})`, supplier, supplierID))
// 		tx.Run(`CREATE (s:Supplier {name: $n, id: $id})`, map[string]interface{}{
// 			"n":  supplier,
// 			"id": supplierID,
// 		})

// 		count := 1 //fake.Number(1, 2)
// 		for count > 0 {
// 			bolref := fmt.Sprintf("V%s%d", fake.HackerAdjective()[:2], fake.Number(10000, 99999))
// 			id := fake.Uint64()
// 			value := (rand.Float64() * 8) + 1
// 			ean := fake.CreditCardNumber()
// 			cdtype := fake.RandString([]string{"fixed", "percent", "percent"})
// 			start := fake.DateRange(time.Now(), time.Now().Add(time.Hour*24*14))
// 			end := start.Add(time.Hour * 24 * time.Duration(fake.Number(60, 150)))

// 			// stmt.WriteString(fmt.Sprintf(`CREATE (c%d%d:Condition {
// 			// 	id: %d, bolref: '%s', value: %.2f, type: '%s', ean: '%d', start: '%s', end: '%s'})`,
// 			// 	count, i,
// 			// 	id, strings.ToUpper(bolref),
// 			// 	value, cdtype, ean,
// 			// 	start.Format("2006-01-02"), end.Format("2006-01-02")))

// 			// stmt.WriteString(fmt.Sprintf(`CREATE (s)-[:AGREED_TO]->(c%d%d)`, count, i))

// 			tx.Run(`CREATE (c:Condition {
// 				id: $id,
// 				bolref: $bolref,
// 				value: $value,
// 				type: $type,
// 				ean: $ean,
// 				start: $start,
// 				end: $end})`,
// 				map[string]interface{}{
// 					"id":     id,
// 					"bolref": strings.ToUpper(bolref),
// 					"value":  value,
// 					"type":   cdtype,
// 					"ean":    ean,
// 					"start":  start.Format("2006-01-02"),
// 					"end":    end.Format("2006-01-02"),
// 				})

// 			tx.Run(`CREATE (s)-[:AGREED_TO]->(c)`, nil)

// 			// generate shop orders
// 			numShopOrders := 250 //fake.Number(20, 50)
// 			currency := "EUR"
// 			orderValue := fake.Float64Range(18, 945)
// 			for numShopOrders > 0 {
// 				valueFlux := fake.Float64Range(1, 12)
// 				totalValue := orderValue - valueFlux
// 				matchDate := fake.DateRange(start, end)
// 				orderID := fake.Uint64()
// 				shipOrderID := fake.Uint32()
// 				// stmt.WriteString(fmt.Sprintf(`
// 				// 	CREATE (so%d%d%d:ShopOrder {
// 				// 		orderId: %d,
// 				// 		shippingId: %d,
// 				// 		currency: '%s',
// 				// 		value: %.2f,
// 				// 		ean: %d
// 				// 	})`, numShopOrders, count, i, orderID, shipOrderID, currency, totalValue, ean))

// 				tx.Run(`
// 				CREATE (so:ShopOrder {
// 					orderId: $orderId,
// 					shipId: $shipId,
// 					currency: $currency,
// 					value: $value,
// 					ean: $ean
// 				})`, map[string]interface{}{
// 					"orderId":  orderID,
// 					"shipId":   shipOrderID,
// 					"currency": currency,
// 					"value":    totalValue,
// 					"ean":      ean,
// 				})

// 				// stmt.WriteString(fmt.Sprintf(`CREATE (c%d%d)<-[:MATCHES_CDN { date: '%s' }]-(so%d%d%d)`,
// 				// 	count, i, matchDate.Format("2006-01-02"), numShopOrders, count, i))

// 				tx.Run(`CREATE (c)<-[:MATCHES_CDN { date: $date }]-(so)`,
// 					map[string]interface{}{
// 						"date": matchDate.Format("2006-01-02"),
// 					})

// 				numShopOrders--
// 			}

// 			count--
// 		}
// 		// fmt.Println(stmt.String())
// 		res, err := tx.Run(stmt.String(), nil)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		if res.Err() != nil {
// 			log.Fatal(res.Err())
// 		} else {
// 			_, err := res.Summary()
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			// fmt.Printf("%+v\n", summary.Counters())
// 		}
// 	}
// 	fmt.Println("Done :)")
// 	return nil
// }
