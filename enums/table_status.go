package enums

type TableStatus int

const (
	TableAllStatus TableStatus = iota
	TableOccupied
	TableVacant
)
