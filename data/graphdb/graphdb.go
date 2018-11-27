package graphdb

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	fake "github.com/brianvoe/gofakeit"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func init() {
	// set deterministic "fakes" (each run will produce the same output)
	fake.Seed(11)
}

type condition struct {
	id         uint64
	bolref     string
	value      float64
	ean        uint32
	valuetype  string
	start      time.Time
	end        time.Time
	supplierID string
}

type shoporder struct {
	orderID   uint64
	shipID    uint32
	matchDate time.Time
	price     float64
	currency  string
	ean       uint32
	cdnID     uint64
}

const (
	sFileName  = "supplier.csv"
	cFileName  = "condition.csv"
	soFileName = "shoporder.csv"
)

// New returns an instance of a GraphGenerator
func New(driver neo4j.Driver) *GraphGenerator {
	return &GraphGenerator{
		driver: driver,
	}
}

// GraphGenerator generates a usable data set
type GraphGenerator struct {
	driver neo4j.Driver
}

// Generate will create a data set of supplier, conditions, sales orders and financial groups
func (g *GraphGenerator) Generate() error {

	err := makeTestData()
	if err != nil {
		return err
	}

	cmd := fmt.Sprintf(`
		LOAD CSV WITH HEADERS FROM "file:///%s" AS csvline
		CREATE (s:Supplier { id: csvline.Id, name: csvline.Name })`,
		sFileName)

	err = g.run(cmd)
	if err != nil {
		return err
	}

	err = g.run("CREATE INDEX ON :Supplier(id)")
	if err != nil {
		return err
	}

	cmd = fmt.Sprintf(`
		LOAD CSV WITH HEADERS FROM "file:///%s" AS csvline
		MERGE (s:Supplier { id: csvline.SupplierId })
		CREATE (c:Condition {
			id: csvline.Id,
			bolref: csvline.BolRef,
			value: toFloat(csvline.Value),
			ean: toInteger(csvline.EAN),
			type: csvline.Type,
			startdate: csvline.StartDate,
			enddate: csvline.EndDate })
		CREATE (s)-[:AGREED_TO]->(c)`,
		cFileName)

	err = g.run(cmd)
	if err != nil {
		return err
	}

	err = g.run("CREATE INDEX ON :Condition(id)")
	if err != nil {
		return err
	}

	err = g.run("CREATE INDEX ON :Supplier(bolref)")
	if err != nil {
		return err
	}

	cmd = fmt.Sprintf(`
		USING PERIODIC COMMIT 500
		LOAD CSV WITH HEADERS FROM "file:///%s" AS csvline
		MERGE (c:Condition { id: csvline.ConditionId })
		CREATE (so:ShopOrder {
			id: csvline.Id,
			shipid: csvline.ShipId,
			currency: csvline.Currency,
			price: toFloat(csvline.Price),
			ean: toInteger(csvline.EAN) })
		CREATE (so)-[:MATCHED_TO { date: csvline.MatchDate }]->(c)`,
		soFileName)

	return g.run(cmd)
}

func (g *GraphGenerator) run(cmd string) error {

	session, err := g.driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return err
	}
	defer session.Close()

	res, err := session.Run(cmd, nil)

	if err != nil {
		return err
	}

	summary, err := res.Summary()
	if err != nil {
		return err
	}

	fmt.Printf("\nSummary:\n%+v\n\n", summary.Counters())
	return nil
}

// Reset resets any previously generated data
func (g *GraphGenerator) Reset() error {
	for _, filename := range [3]string{sFileName, cFileName, soFileName} {
		err := os.Remove("/Users/crushonly/neo4j/import/" + filename)
		if err != nil {
			log.Println(err)
		}
	}

	session, err := g.driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return err
	}

	res, err := session.Run(`MATCH (n) DETACH DELETE n`, nil)

	if err != nil {
		log.Fatal(err)
	}

	summary, err := res.Summary()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", summary.Counters())
	fmt.Println("the database is now empty")

	return nil
}

func makeTestData() error {
	sfile, cfile, sofile, close, err := makeFiles()
	if err != nil {
		return err
	}
	defer close()

	sWriter := bufio.NewWriter(sfile)
	cWriter := bufio.NewWriter(cfile)
	soWriter := bufio.NewWriter(sofile)

	sWriter.WriteString("Id,Name\n")
	cWriter.WriteString("Id,SupplierId,BolRef,Value,EAN,Type,Start Date,EndDate\n")
	soWriter.WriteString("Id,ConditionId,ShipId,Currency,Price,EAN,MatchDate\n")

	for i := 0; i < 2; i++ {
		sID, supplier := createSupplier()
		_, err = sWriter.WriteString(fmt.Sprintf("%s,%s\n", sID, supplier))
		if err != nil {
			return err
		}

		count := fake.Number(1, 8)
		for count > 0 {
			cdn := createCondition(sID)
			_, err = cWriter.WriteString(fmt.Sprintf(
				"%d,%s,%s,%.2f,%d,%s,%s,%s\n",
				cdn.id, cdn.supplierID, cdn.bolref, cdn.value, cdn.ean, cdn.valuetype,
				cdn.start.Format("2006-01-02"), cdn.end.Format("2006-01-02")))

			if err != nil {
				return err
			}

			price := fake.Price(8, 120)
			numOrders := fake.Number(40000, 50000)
			for numOrders > 0 {
				order := createShopOrder(price, cdn)
				_, err = soWriter.WriteString(fmt.Sprintf(
					"%d,%d,%d,%s,%.2f,%d,%s\n",
					order.orderID, order.cdnID,
					order.shipID, order.currency,
					order.price, order.ean, order.matchDate.Format("2006-01-02"),
				))
				numOrders--
			}

			count--
		}
	}

	cWriter.Flush()
	sWriter.Flush()
	soWriter.Flush()
	return nil
}

func makeFiles() (sfile, cfile, shfile *os.File, closer func(), err error) {
	sfile, err = os.Create("/Users/crushonly/neo4j/import/" + sFileName)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	cfile, err = os.Create("/Users/crushonly/neo4j/import/" + cFileName)
	if err != nil {
		sfile.Close()
		return nil, nil, nil, nil, err
	}

	shfile, err = os.Create("/Users/crushonly/neo4j/import/" + soFileName)
	if err != nil {
		sfile.Close()
		cfile.Close()
		return nil, nil, nil, nil, err
	}

	closer = func() {
		sfile.Close()
		cfile.Close()
		shfile.Close()
	}

	return
}

func createSupplier() (id, name string) {
	name = fake.RandString([]string{
		strings.Title(fake.HipsterWord()) + " " + strings.Replace(fake.CompanySuffix(), "and Sons", "And Co.", -1),
		strings.Title(fake.HackerNoun()) + "." + fake.DomainSuffix(),
		strings.Title(fake.BuzzWord()) + " " + strings.Replace(fake.CompanySuffix(), "and Sons", "And Co.", -1),
	})
	//name = strings.Title(strings.Replace(sname, "'", "", -1))
	id = fmt.Sprintf("%s-%d",
		strings.Replace(strings.ToUpper(name), " ", "", -1)[:2],
		fake.Number(10000, 99999))
	return
}

func createCondition(supplierID string) condition {
	start := fake.DateRange(time.Now(), time.Now().Add(time.Hour*24*14))
	end := start.Add(time.Hour * 24 * time.Duration(fake.Number(60, 150)))

	return condition{
		bolref:     fmt.Sprintf("V%s%d", strings.ToUpper(fake.HipsterWord()[:2]), fake.Number(10000, 99999)),
		id:         fake.Uint64(),
		value:      (rand.Float64() * 8) + 1,
		ean:        fake.Uint32(),
		valuetype:  fake.RandString([]string{"fixed", "percent", "percent"}),
		start:      start,
		end:        end,
		supplierID: supplierID,
	}
}

func createShopOrder(price float64, cdn condition) shoporder {
	valueFlux := fake.Float64Range(1, 7)
	totalValue := price - valueFlux
	return shoporder{
		cdnID:     cdn.id,
		orderID:   fake.Uint64(),
		shipID:    fake.Uint32(),
		matchDate: fake.DateRange(cdn.start, cdn.end),
		currency:  "EUR",
		ean:       cdn.ean,
		price:     totalValue,
	}
}
