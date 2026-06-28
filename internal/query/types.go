package query

import (
	"fmt"
	"strings"
)

// --- Operand Type ---
type OperandType int

const (
	SET OperandType = iota
	GET
	DELETE
	UNKNOWN
)

var OperandName = map[OperandType]string{
	SET:     "set",
	GET:     "get",
	DELETE:  "delete",
	UNKNOWN: "unknown",
}

func (ot OperandType) String() string {
	return OperandName[ot]
}

func ToOperand(s string) OperandType {
	switch strings.ToLower(s) {
	case "get":
		return GET
	case "set":
		return SET
	case "delete":
		return DELETE
	default:
		return UNKNOWN
	}
}

// --- Unparsed Query Type ---

type UnparsedQuery string

func (uq UnparsedQuery) String() string {
	return fmt.Sprintf("%v", string(uq))
}
