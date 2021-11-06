package strategies

type Strategy interface {
	Scout() error
}

type ratioPair struct {
	origin string
	target string
	ratio  float64
}
