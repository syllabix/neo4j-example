package neox

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Session struct {
	neo4j.Session
}

func (s *Session) Runx(cypher string, params map[string]interface{}, configurers ...func(*neo4j.TransactionConfig)) (*Result, error) {
	res, err := s.Run(cypher, params, configurers...)
	if err != nil {
		return nil, err
	}
	return &Result{
		Result: res,
	}, nil
}
