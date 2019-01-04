package core

type Channel interface {
	Name() string
	Users() []string
	Ops() []string
	Voiced() []string
	HasVoiceOrOp(name string) bool
	HasOp(name string) bool
	HasVoice(name string) bool
}