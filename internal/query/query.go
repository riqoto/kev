package query

import (
	"fmt"
	"kev/internal/store"
	"strings"
)

type Query struct {
	Operand OperandType
	Key     string
	Value   store.Entry
	Options []string
}

func (q Query) String() string {
	return fmt.Sprintf("Operand: %v\nKey: %v\nValue: %v\nOptions: %q\n", q.Operand, q.Key, q.Value.Value, q.Options)
}
func (uq UnparsedQuery) Parse() *Query {
	var q *Query = &Query{}

	uqs := strings.TrimSpace(uq.String())

	uqsa := strings.Split(uqs, " ")

	var oqv []string
	for _, keyword := range uqsa {
		if keyword[0] != '-' {
			oqv = append(oqv, keyword)
		} else {
			q.Options = append(q.Options, keyword)
		}
	}
	//	var ot OperandType

	q.Operand = ToOperand(oqv[0])
	q.Key = oqv[1]
	q.Value.Value = []byte(oqv[2])

	return q
}
