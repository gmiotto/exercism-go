package robotname

import (
	"math/rand"
	"time"
)

// Robot 1
type Robot struct {
	name string
}

// New 1
func New() *Robot {
	return (new(Robot))
}

// Name 1
func (r *Robot) Name() string {
	if r.name == "" {
		r.name = randSeq()
	}
	return r.name
}

// Reset 1
func (r *Robot) Reset() {
	r.name = ""
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var digits = []rune("1234567890")

func randSeq() string {
	b := make([]rune, 5)
	for i := range b {
		if i < 2 {
			b[i] = letters[rand.Intn(len(letters))]
		} else {
			b[i] = digits[rand.Intn(len(digits))]
		}
	}
	return string(b)
}
