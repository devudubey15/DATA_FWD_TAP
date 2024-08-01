package md5

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"hash"
)

const (
	MD5_SIZE        = 16
	MD5_STRING_SIZE = 33
	BLOCK_SIZE      = 64
)

var (
	T = [64]uint32{
		0xd76aa478, 0xe8c7b756, 0x242070db, 0xc1bdceee, 0xf57c0faf, 0x4787c62a, 0xa8304613, 0xfd469501,
		0x698098d8, 0x8b44f7af, 0xffff5bb1, 0x895cd7be, 0x6b901122, 0xfd987193, 0xa679438e, 0x49b40821,
		0xf61e2562, 0xc040b340, 0x265e5a51, 0xe9b6c7aa, 0xd62f105d, 0x02441453, 0xd8a1e681, 0xe7d3fbc8,
		0x21e1cde6, 0xc33707d6, 0xf4d50d87, 0x455a14ed, 0xa9e3e905, 0xfcefa3f8, 0x676f02d9, 0x8d2a4c8a,
		0xfffa3942, 0x8771f681, 0x6d9d6122, 0xfde5380c, 0xa4beea44, 0x4bdecfa9, 0xf6bb4b60, 0xbebfbc70,
		0x289b7ec6, 0xeaa127fa, 0xd4ef3085, 0x04881d05, 0xd9d4d039, 0xe6db99e5, 0x1fa27cf8, 0xc4ac5665,
		0xf4292244, 0x432aff97, 0xab9423a7, 0xfc93a039, 0x655b59c3, 0x8f0ccc92, 0xffeff47d, 0x85845dd1,
		0x6fa87e4f, 0xfe2ce6e0, 0xa3014314, 0x4e0811a1, 0xf7537e82, 0xbd3af235, 0x2ad7d2bb, 0xeb86d391,
	}

	S1 = [4]uint32{7, 12, 17, 22}
	S2 = [4]uint32{5, 9, 14, 20}
	S3 = [4]uint32{4, 11, 16, 23}
	S4 = [4]uint32{6, 10, 15, 21}
)

type MD5 struct {
	A, B, C, D    uint32
	messageLength [2]uint32
	storedSize    uint32
	stored        [BLOCK_SIZE * 2]byte
	finished      bool
	signature     [MD5_SIZE]byte
	str           [MD5_STRING_SIZE]byte
	hash          hash.Hash
}

func NewMD5() *MD5 {
	md5Instance := &MD5{}
	md5Instance.initialise()
	return md5Instance
}

func NewMD5WithBuffer(input []byte) *MD5 {
	md5Instance := &MD5{}
	md5Instance.initialise()
	md5Instance.process(input)
	return md5Instance
}

func (m *MD5) initialise() {
	m.hash = md5.New()
	m.storedSize = 0
	m.finished = false
	m.messageLength[0], m.messageLength[1] = 0, 0
	m.A, m.B, m.C, m.D = 0x67452301, 0xefcdab89, 0x98badcfe, 0x10325476
}

func (m *MD5) process(input []byte) {
	m.hash.Write(input)
	m.messageLength[0] += uint32(len(input))
}

func (m *MD5) finish() {
	copy(m.signature[:], m.hash.Sum(nil))
	m.finished = true
}

func (m *MD5) getResult() []byte {
	if !m.finished {
		m.finish()
	}
	return m.signature[:]
}

func SigToString(signature []byte) (string, error) {
	if len(signature) != MD5_SIZE {
		return "", errors.New("invalid signature length")
	}
	return hex.EncodeToString(signature), nil
}

func SigFromString(str string) ([]byte, error) {
	if len(str) != MD5_SIZE*2 {
		return nil, errors.New("invalid string length")
	}
	return hex.DecodeString(str)
}

// Cyclic left rotation
func cyclicLeftRotate(x uint32, n uint32) uint32 {
	return (x << n) | (x >> (32 - n))
}

func F(x, y, z uint32) uint32 {
	return (x & y) | (^x & z)
}

func G(x, y, z uint32) uint32 {
	return (x & z) | (y & ^z)
}

func H(x, y, z uint32) uint32 {
	return x ^ y ^ z
}

func I(x, y, z uint32) uint32 {
	return y ^ (x | ^z)
}

func FF(a, b, c, d, xk, s, ti uint32) uint32 {
	return b + cyclicLeftRotate(a+F(b, c, d)+xk+T[ti], s)
}

func GG(a, b, c, d, xk, s, ti uint32) uint32 {
	return b + cyclicLeftRotate(a+G(b, c, d)+xk+T[ti], s)
}

func HH(a, b, c, d, xk, s, ti uint32) uint32 {
	return b + cyclicLeftRotate(a+H(b, c, d)+xk+T[ti], s)
}

func II(a, b, c, d, xk, s, ti uint32) uint32 {
	return b + cyclicLeftRotate(a+I(b, c, d)+xk+T[ti], s)
}
