package group

import (
	"encoding/binary"

	"github.com/jadeydi/jubjub/pkg/jubjub/extended"
	"github.com/jadeydi/jubjub/pkg/jubjub/fr"
)

// A "w-ary non-adjacent form" exponentiation context.
type Wnaf struct {
	base       []*extended.ExtendedPoint
	scalar     []int64
	windowSize uint64
}

func NewWnaf() *Wnaf {
	return &Wnaf{}
}

func (w *Wnaf) WnafForm(c []byte, window uint64) {
	wnaf := make([]int64, 0)

	bit_len := uint64(len(c)) * 8
	//u64_len := (bit_len + 1) / 64

	var c_u64 []uint64
	for i := 0; i < len(c); i += 8 {
		c_u64 = append(c_u64, binary.LittleEndian.Uint64(c[i:i+8]))
	}
	c_u64 = append(c_u64, 0)

	width := uint64(1) << window
	window_mask := window - 1

	var pos, carry uint64
	for pos < bit_len {
		u64_idx := pos / 64
		bit_idx := pos % 64

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
				wnaf = append(wnaf, int64(window_val))
			} else {
				carry = 1
				wnaf = append(wnaf, int64(window_val)-int64(width))
			}

			var r []int64
			for i := uint64(0); i < window-1; i++ {
				r = append(r, 0)
			}
			wnaf = append(wnaf, r...)
			pos += window
		}
	}

	w.scalar = wnaf
}

// base is jubjub::ExtendedPoint
func (w *Wnaf) WnafTable(base *extended.ExtendedPoint) {
	w.base = []*extended.ExtendedPoint{}

	dbl := base.Double()
	for i := 0; i < (1 << (w.windowSize - 1)); i++ {
		w.base = append(w.base, base)
		base = base.Add(dbl)
	}
}

// scalar is jubjub::Fr
func (w *Wnaf) WnafExp(wnaf []int64) *extended.ExtendedPoint {
	result := extended.Identity()
	found_one := false

	for i := len(wnaf) - 1; i >= 0; i-- {
		if found_one {
			result = result.Double()
		}

		if i != 0 {
			found_one = true

			if i > 0 {
				result = result.Add(w.base[i/2])
			} else {
				result = result.Sub(w.base[(-i)/2])
			}
		}
	}

	return result
}

func (w *Wnaf) Scalar(esk *fr.Fr) *Wnaf {
	window_size := uint64(4)
	w.WnafForm(esk.Bytes(), window_size)

	return &Wnaf{
		base:       w.base,
		scalar:     w.scalar,
		windowSize: window_size,
	}
}

func (w *Wnaf) Base(base *extended.ExtendedPoint) *extended.ExtendedPoint {
	w.WnafTable(base)
	return w.WnafExp(w.scalar)
}
