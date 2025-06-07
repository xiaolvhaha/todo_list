package types

type TaskDomain struct {
	Id       int64
	Title    string
	Desc     string
	Property int64
	Deadline int64
	Status   int64
	Category CategoryDomain
}
