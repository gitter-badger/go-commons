package props

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

const fileName = "test.props"

func TestLoadFile(t *testing.T) {
	p, err := LoadFile(fileName)
	assert.Nil(t, err)
	assert.NotEmpty(t, p)
	assert.Equal(t, "127.0.0.1", p["ip"])
	_, err = LoadFile("non-exist")
	assert.NotNil(t, err)
}

const propText = `ip=127.0.0.1
username=winjeg
result=success
# this is a comment`

func TestFromString(t *testing.T) {
	p := FromString(propText)
	assert.Equal(t, "winjeg", p["username"])
	q := FromString("")
	assert.Empty(t, q)
	x := FromString("tom")
	assert.Empty(t, x)
	assert.NotNil(t, p.String())
}
