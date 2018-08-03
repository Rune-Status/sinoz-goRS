package message

type Details struct {
	Members   bool
	ProcessId int
}

type Logout struct { /* EMPTY */ }

type SkillUpdate struct {
	Id         int
	Level      int
	Experience float32
}
