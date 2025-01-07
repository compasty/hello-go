package bitmap

import (
	"bytes"
	"fmt"
	"math/bits"
)

type Bitmap struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (b *Bitmap) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(b.words) && b.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (b *Bitmap) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(b.words) {
		b.words = append(b.words, 0)
	}
	b.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (b *Bitmap) UnionWith(t *Bitmap) {
	for i, tword := range t.words {
		if i < len(b.words) {
			b.words[i] |= tword
		} else {
			b.words = append(b.words, tword)
		}
	}
}

// return the number of elements
func (b *Bitmap) Len() int {
	len := 0
	for _, word := range b.words {
		len += bits.OnesCount64(word)
	}
	return len
}

// remove x from the set
func (b *Bitmap) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(b.words) && b.words[word]&(1<<bit) != 0 {
		b.words[word] &^= 1 << bit
	}
}

// remove all elements from the set
func (b *Bitmap) Clear() {
	for i := range b.words {
		b.words[i] = 0
	}
}

// return a copy of the set
func (b *Bitmap) Copy() *Bitmap {
	dst := make([]uint64, len(b.words))
	copy(dst, b.words)
	return &Bitmap{words: dst}
}

func (b *Bitmap) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range b.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
