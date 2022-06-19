package entity

type Sort int

const (
	DateDESC Sort = iota
	DateASC
	SumDESC
	SumASC
)

var Sorts = map[string]Sort{
	"DATE_DESC": DateDESC,
	"DATE_ASC":  DateASC,
	"SUM_DESC":  SumDESC,
	"SUM_ASC":   SumASC,
}
