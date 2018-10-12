package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFiler_Upload(t *testing.T) {
	f := &Filer{
		nrUrl: "http://127.0.0.1:8881",
	}

	dat, err := os.Open("/home/lee/go/src/github.com/lisuiheng/fsm/file/hello.txt")
	assert.Nil(t, err)
	r, err := f.Upload("/hello", "", dat, NO_REP)
	assert.Nil(t, err)
	fmt.Println(r)
}
