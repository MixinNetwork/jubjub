package fq

// q
var q = Fq{0xffffffff00000001, 0x53bda402fffe5bfe, 0x3339d80809a1d805, 0x73eda753299d7d48}

var zero = Fq{0, 0, 0, 0}

// INV = -(q^{-1} mod 2^64) mod 2^64
const INV uint64 = 0xfffffffeffffffff

const S int = 32

// R = 2^256 mod q
var R = Fq{0x00000001fffffffe, 0x5884b7fa00034802, 0x998c4fefecbc4ff5, 0x1824b159acc5056f}

// R2 = 2^512 mod q
var R2 = Fq{0xc999e990f3f29c6d, 0x2b6cedcb87925c23, 0x05d314967254398f, 0x0748d9d99f59ff11}

// R3 = R2 * 2^256 mod q = 2^768 mod q
var R3 = Fq{0xc62c1807439b73af, 0x1b3e0d188cf06990, 0x73d13c71c7b5f418, 0x6e2a5bb9c8db33e9}

// ROOTOFUNITY GENERATOR^t where t * 2^s + 1 = q with t odd.
var ROOTOFUNITY = Fq{0xb9b58d8c5f0e466a, 0x5b1b4c801819d7ec, 0x0af53ae352a31e64, 0x5bf3adda19e9b27b}

// D = -(10240/10241)
var D = Fq{0x2a522455b974f6b0, 0xfc6cc9ef0d9acab3, 0x7a08fb94c27628d1, 0x57f8f6a8fe0e262e}

// D2 = 2 * d
var D2 = Fq{0x54a448ac72e9ed5f, 0xa51befdb1b373967, 0xc0d81f217b4a799e, 0x3c0445fed27ecf14}

// `d = -(10240/10241)`
var EDWARDS_D = Fq{
	0x0106_5fd6_d634_3eb1,
	0x292d_7f6d_3757_9d26,
	0xf5fd_9207_e6bd_7fd4,
	0x2a93_18e7_4bfa_2b48,
}

// `2*d`
var EDWARDS_D2 = Fq{
	0x020c_bfad_ac68_7d62,
	0x525a_feda_6eaf_3a4c,
	0xebfb_240f_cd7a_ffa8,
	0x5526_31ce_97f4_5691,
}
