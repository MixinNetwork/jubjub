package group

import (
	"encoding/binary"

	"github.com/jadeydi/jubjub/pkg/jubjub/fr"
)

// A "w-ary non-adjacent form" exponentiation context.
type Wnaf struct {
	Base       []byte
	Scalar     []int64
	WindowSize uint
}

func New() *Wnaf {
	return &Wnaf{}
}

func (w *Wnaf) WnafForm(c []byte, window uint) {
	wnaf := make([]int64, 0)

	bit_len := len(c) * 8
	u64_len := (bit_len + 1) / 64

	var c_u64 []uint64
	for i := 0; i < len(c); i += 8 {
		c_u64 = append(c_u64, binary.LittleEndian.Uint64(c[i:i+8]))
	}
	c_u64 = append(c_u64, 0)

	width := 1 << window
	window_mask := window - 1

	var pos, carry uint64
	for pos < bit_len {
		u64_idx := pos / 64
		bit_idx := pos % 64
	}

	var bit_buf uint64
	if bit_idx+window < 64 {
		bit_buf = c_u64[u64_idx] >> bit_idx
	} else {
		bit_buf = (c_u64[u64_idx] >> bit_idx) | (c_u64[u64_idx+1] << (64 - bit_idx))
	}

	window_val := carry + (bit_buf & window_mask)

	if window_val&1 == 0 {
		wnaf = append(wnaf, 0)
		pos += 1
	} else {
		if window_val < width/2 {
			carry = 0
			wnaf = append(wnaf, window_val)
		} else {
			carry = 1
			wnaf = append(wnaf, window_val-width)
		}

		var r [window - 1]int64
		wnaf = append(wnaf, r...)
		pos += window
	}

	w.Scalar = wnaf
}

func (w *Wnaf) Scalar(esk *fr.Fr) *Wnaf {
	window_size := 4
	w.WnafForm(esk.Bytes(), window_size)

	return &Wnaf{
		base:   w.Base,
		scalar: w.Scalar,
		window_size,
	}
}
