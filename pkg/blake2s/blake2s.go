// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package blake2s implements the BLAKE2s hash algorithm as
// defined in RFC 7693.
package blake2s

import (
	"encoding/binary"
	"errors"
	"hash"
)

const (
	// The blocksize of BLAKE2s in bytes.
	BlockSize = 64
	// The hash size of BLAKE2s-256 in bytes.
	Size = 32
)

var errKeySize = errors.New("blake2s: invalid key size")

var iv = [8]uint32{
	0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a,
	0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19,
}

// Sum256 returns the BLAKE2s-256 checksum of the data.
func Sum256(data []byte) [Size]byte {
	var sum [Size]byte
	checkSum(&sum, Size, data)
	return sum
}

// New256 returns a new hash.Hash computing the BLAKE2s-256 checksum. A non-nil
// key turns the hash into a MAC. The key must between zero and 32 bytes long.
func New256(key []byte) (hash.Hash, error) { return newDigest(Size, key, nil) }
func New256WithPersonalization(key, personalization []byte) (hash.Hash, error) {
	return newDigest(Size, key, personalization)
}

func newDigest(hashSize int, key []byte, personalization []byte) (*digest, error) {
	if len(key) > Size {
		return nil, errKeySize
	}
	if personalization != nil && len(personalization) < 8 {
		for i := len(personalization); i < 8; i++ {
			personalization = append(personalization, byte(0))
		}
	}
	d := &digest{
		size:            hashSize,
		keyLen:          len(key),
		personalization: personalization,
	}
	copy(d.key[:], key)
	d.Reset()
	return d, nil
}

func checkSum(sum *[Size]byte, hashSize int, data []byte) {
	var (
		h [8]uint32
		c [2]uint32
	)

	h = iv
	h[0] ^= uint32(hashSize) | (1 << 16) | (1 << 24)

	if length := len(data); length > BlockSize {
		n := length &^ (BlockSize - 1)
		if length == n {
			n -= BlockSize
		}
		hashBlocks(&h, &c, 0, data[:n])
		data = data[n:]
	}

	var block [BlockSize]byte
	offset := copy(block[:], data)
	remaining := uint32(BlockSize - offset)

	if c[0] < remaining {
		c[1]--
	}
	c[0] -= remaining

	hashBlocks(&h, &c, 0xFFFFFFFF, block[:])

	for i, v := range h {
		binary.LittleEndian.PutUint32(sum[4*i:], v)
	}
}

type digest struct {
	h      [8]uint32
	c      [2]uint32
	size   int
	block  [BlockSize]byte
	offset int

	key    [BlockSize]byte
	keyLen int

	personalization []byte
}

func (d *digest) BlockSize() int { return BlockSize }

func (d *digest) Size() int { return d.size }

func (d *digest) Reset() {
	d.h = iv
	d.h[0] ^= uint32(d.size) | (uint32(d.keyLen) << 8) | (1 << 16) | (1 << 24)
	if d.personalization != nil {
		for i := uint(0); i < 8; i++ {
			b := uint32(d.personalization[i]) << (8 * uint(i%4))
			d.h[6+i/4] ^= b
		}
	}
	d.offset, d.c[0], d.c[1] = 0, 0, 0
	if d.keyLen > 0 {
		d.block = d.key
		d.offset = BlockSize
	}
}

func (d *digest) Write(p []byte) (n int, err error) {
	n = len(p)

	if d.offset > 0 {
		remaining := BlockSize - d.offset
		if n <= remaining {
			d.offset += copy(d.block[d.offset:], p)
			return
		}
		copy(d.block[d.offset:], p[:remaining])
		hashBlocks(&d.h, &d.c, 0, d.block[:])
		d.offset = 0
		p = p[remaining:]
	}

	if length := len(p); length > BlockSize {
		nn := length &^ (BlockSize - 1)
		if length == nn {
			nn -= BlockSize
		}
		hashBlocks(&d.h, &d.c, 0, p[:nn])
		p = p[nn:]
	}

	d.offset += copy(d.block[:], p)
	return
}

func (d *digest) Sum(b []byte) []byte {
	var block [BlockSize]byte
	h := d.h
	c := d.c

	copy(block[:], d.block[:d.offset])
	remaining := uint32(BlockSize - d.offset)
	if c[0] < remaining {
		c[1]--
	}
	c[0] -= remaining

	hashBlocks(&h, &c, 0xFFFFFFFF, block[:])

	var sum [Size]byte
	for i, v := range h {
		binary.LittleEndian.PutUint32(sum[4*i:], v)
	}

	return append(b, sum[:d.size]...)
}
