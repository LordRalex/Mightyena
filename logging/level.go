package logging

type Level interface {
	Name() string
	Rank() int
}

type level struct {
	name string
	rank int
}

func (l *level) Name() string {
	return l.name
}

func (l *level) Rank() int {
	return l.rank
}

var (
	Info     = &level{name: "INFO", rank: 100}
	Warning  = &level{name: "WARN", rank: 200}
	Error    = &level{name: "ERROR", rank: 300}
	Critical = &level{name: "CRITICAL", rank: 1000}
	Debug    = &level{name: "DEBUG", rank: 0}
)
