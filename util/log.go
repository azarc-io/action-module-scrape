package util

type Logger interface {
	Fatalf(msg string, args ...any)
	Infof(msg string, args ...any)
}
