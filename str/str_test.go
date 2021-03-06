package str

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUUID(t *testing.T) {
	id1 := UUID()
	id2 := UUID()
	assert.NotEqual(t, id1, id2)
}

func TestGenerateStrings(t *testing.T) {
	fmt.Println(RandomAlphabetsLower(15))
	fmt.Println(RandomAlphabetsUpper(15))
	fmt.Println(RandomNumAlphabets(15))
	fmt.Println(RandomNumbers(15))
	fmt.Println(RandomStrWithSpecialChars(15))
	str := UUIDShort()
	assert.NotNil(t, str)
	assert.True(t, !IsBlank(str))
	assert.True(t, !IsEmpty(str))
	r := ReplaceAll(str, "a", "A")
	assert.NotNil(t, r)
}

func BenchmarkKrand(b *testing.B) {
	RandStr(30, KindAll)
	RandStr(30, KindUpper)
	RandStr(30, KindNumber)
	b.ReportAllocs()
}

// test uuid generator to make sure it won't get duplicated keys
// in not very large range
func BenchmarkGenerateUniqueId(b *testing.B) {
	m := make(map[string]bool, 1000000)
	println(UUIDShort())
	count := 0
	for i := 0; i < 100000000; i++ {
		k := UUIDShort()
		if _, ok := m[k]; ok {
			count++
		}
		m[k] = true
	}
	b.ReportAllocs()
}
