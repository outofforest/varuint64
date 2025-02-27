package varuint64

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSize(t *testing.T) {
	requireT := require.New(t)

	requireT.EqualValues(1, Size(0))
	requireT.EqualValues(1, Size(1))
	requireT.EqualValues(1, Size(max1Byte))
	requireT.EqualValues(2, Size(max1Byte+1))
	requireT.EqualValues(2, Size(max2Bytes))
	requireT.EqualValues(3, Size(max2Bytes+1))
	requireT.EqualValues(3, Size(max3Bytes))
	requireT.EqualValues(4, Size(max3Bytes+1))
	requireT.EqualValues(4, Size(max4Bytes))
	requireT.EqualValues(5, Size(max4Bytes+1))
	requireT.EqualValues(5, Size(max5Bytes))
	requireT.EqualValues(6, Size(max5Bytes+1))
	requireT.EqualValues(6, Size(max6Bytes))
	requireT.EqualValues(7, Size(max6Bytes+1))
	requireT.EqualValues(7, Size(max7Bytes))
	requireT.EqualValues(8, Size(max7Bytes+1))
	requireT.EqualValues(8, Size(max8Bytes))
	requireT.EqualValues(9, Size(max8Bytes+1))
	requireT.EqualValues(9, Size(math.MaxUint64))
}

func TestContains(t *testing.T) {
	requireT := require.New(t)

	requireT.False(Contains([]byte{}))
	requireT.True(Contains([]byte{0x00}))
	requireT.True(Contains([]byte{0x00, 0x00}))
	requireT.True(Contains([]byte{0x00, 0x80}))
	requireT.True(Contains([]byte{0x00, 0xff}))
	requireT.True(Contains([]byte{0x7F, 0x7F}))
	requireT.True(Contains([]byte{0x7F, 0x80}))
	requireT.False(Contains([]byte{0x80}))
	requireT.False(Contains([]byte{0x80, 0x80}))
	requireT.False(Contains([]byte{0x80, 0x80, 0x80}))
	requireT.False(Contains([]byte{0x80, 0x80, 0x80, 0x80}))
	requireT.False(Contains([]byte{0x80, 0x80, 0x80, 0x80}))
	requireT.False(Contains([]byte{0x80, 0x80, 0x80, 0x80, 0x80}))
	requireT.False(Contains([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80}))
	requireT.False(Contains([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80}))
	requireT.False(Contains([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80}))
	requireT.True(Contains([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80}))
	requireT.True(Contains([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x7F}))
	requireT.True(Contains([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x7F}))
	requireT.True(Contains([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x7F}))
	requireT.True(Contains([]byte{0x80, 0x80, 0x80, 0x80, 0x7F}))
	requireT.True(Contains([]byte{0x80, 0x80, 0x80, 0x7F}))
	requireT.True(Contains([]byte{0x80, 0x80, 0x7F}))
	requireT.True(Contains([]byte{0x80, 0x7F}))
	requireT.True(Contains([]byte{0x7F}))
	requireT.True(Contains([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x7F, 0x80}))
	requireT.True(Contains([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x7F, 0x80}))
	requireT.True(Contains([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x7F, 0x80}))
	requireT.True(Contains([]byte{0x80, 0x80, 0x80, 0x80, 0x7F, 0x80}))
	requireT.True(Contains([]byte{0x80, 0x80, 0x80, 0x7F, 0x80}))
	requireT.True(Contains([]byte{0x80, 0x80, 0x7F, 0x80}))
	requireT.True(Contains([]byte{0x80, 0x7F, 0x80}))
	requireT.True(Contains([]byte{0x7F, 0x80}))
}

func TestPut(t *testing.T) {
	requireT := require.New(t)
	b := make([]byte, 9)

	requireT.EqualValues(1, Put(b, 0))
	requireT.Equal([]byte{0x00}, b[:1])

	requireT.EqualValues(1, Put(b, 1))
	requireT.Equal([]byte{0x01}, b[:1])

	requireT.EqualValues(1, Put(b, max1Byte))
	requireT.Equal([]byte{numberMask}, b[:1])

	requireT.EqualValues(2, Put(b, max1Byte+1))
	requireT.Equal([]byte{followMask, 0x01}, b[:2])

	requireT.EqualValues(2, Put(b, max2Bytes))
	requireT.Equal([]byte{0xFF, numberMask}, b[:2])

	requireT.EqualValues(3, Put(b, max2Bytes+1))
	requireT.Equal([]byte{followMask, followMask, 0x01}, b[:3])

	requireT.EqualValues(3, Put(b, max3Bytes))
	requireT.Equal([]byte{0xFF, 0xFF, numberMask}, b[:3])

	requireT.EqualValues(4, Put(b, max3Bytes+1))
	requireT.Equal([]byte{followMask, followMask, followMask, 0x01}, b[:4])

	requireT.EqualValues(4, Put(b, max4Bytes))
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, numberMask}, b[:4])

	requireT.EqualValues(5, Put(b, max4Bytes+1))
	requireT.Equal([]byte{followMask, followMask, followMask, followMask, 0x01}, b[:5])

	requireT.EqualValues(5, Put(b, max5Bytes))
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, numberMask}, b[:5])

	requireT.EqualValues(6, Put(b, max5Bytes+1))
	requireT.Equal([]byte{followMask, followMask, followMask, followMask, followMask, 0x01}, b[:6])

	requireT.EqualValues(6, Put(b, max6Bytes))
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, numberMask}, b[:6])

	requireT.EqualValues(7, Put(b, max6Bytes+1))
	requireT.Equal([]byte{followMask, followMask, followMask, followMask, followMask, followMask, 0x01}, b[:7])

	requireT.EqualValues(7, Put(b, max7Bytes))
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, numberMask}, b[:7])

	requireT.EqualValues(8, Put(b, max7Bytes+1))
	requireT.Equal([]byte{followMask, followMask, followMask, followMask, followMask, followMask, followMask, 0x01}, b[:8])

	requireT.EqualValues(8, Put(b, max8Bytes))
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, numberMask}, b[:8])

	requireT.EqualValues(9, Put(b, max8Bytes+1))
	requireT.Equal([]byte{followMask, followMask, followMask, followMask, followMask, followMask, followMask, followMask,
		0x01}, b[:9])

	requireT.EqualValues(9, Put(b, math.MaxUint64))
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, b[:9])
}

func TestParse(t *testing.T) {
	requireT := require.New(t)

	v, n := Parse([]byte{0x00})
	requireT.EqualValues(1, n)
	requireT.EqualValues(0, v)

	v, n = Parse([]byte{0x00, 0x80})
	requireT.EqualValues(1, n)
	requireT.EqualValues(0, v)

	v, n = Parse([]byte{0x01})
	requireT.EqualValues(1, n)
	requireT.EqualValues(1, v)

	v, n = Parse([]byte{0x01, 0x80})
	requireT.EqualValues(1, n)
	requireT.EqualValues(1, v)

	v, n = Parse([]byte{numberMask})
	requireT.EqualValues(1, n)
	requireT.EqualValues(max1Byte, v)

	v, n = Parse([]byte{followMask, 0x01})
	requireT.EqualValues(2, n)
	requireT.EqualValues(max1Byte+1, v)

	v, n = Parse([]byte{followMask, 0x01, 0x80})
	requireT.EqualValues(2, n)
	requireT.EqualValues(max1Byte+1, v)

	v, n = Parse([]byte{0xFF, numberMask})
	requireT.EqualValues(2, n)
	requireT.EqualValues(max2Bytes, v)

	v, n = Parse([]byte{followMask, followMask, 0x01})
	requireT.EqualValues(3, n)
	requireT.EqualValues(max2Bytes+1, v)

	v, n = Parse([]byte{0xFF, 0xFF, numberMask})
	requireT.EqualValues(3, n)
	requireT.EqualValues(max3Bytes, v)

	v, n = Parse([]byte{0xFF, 0xFF, numberMask, 0x80})
	requireT.EqualValues(3, n)
	requireT.EqualValues(max3Bytes, v)

	v, n = Parse([]byte{followMask, followMask, followMask, 0x01})
	requireT.EqualValues(4, n)
	requireT.EqualValues(max3Bytes+1, v)

	v, n = Parse([]byte{followMask, followMask, followMask, 0x01, 0x80})
	requireT.EqualValues(4, n)
	requireT.EqualValues(max3Bytes+1, v)

	v, n = Parse([]byte{0xFF, 0xFF, 0xFF, numberMask})
	requireT.EqualValues(4, n)
	requireT.EqualValues(max4Bytes, v)

	v, n = Parse([]byte{0xFF, 0xFF, 0xFF, numberMask, 0x80})
	requireT.EqualValues(4, n)
	requireT.EqualValues(max4Bytes, v)

	v, n = Parse([]byte{followMask, followMask, followMask, followMask, 0x01})
	requireT.EqualValues(5, n)
	requireT.EqualValues(max4Bytes+1, v)

	v, n = Parse([]byte{followMask, followMask, followMask, followMask, 0x01, 0x80})
	requireT.EqualValues(5, n)
	requireT.EqualValues(max4Bytes+1, v)

	v, n = Parse([]byte{0xFF, 0xFF, 0xFF, 0xFF, numberMask})
	requireT.EqualValues(5, n)
	requireT.EqualValues(max5Bytes, v)

	v, n = Parse([]byte{0xFF, 0xFF, 0xFF, 0xFF, numberMask, 0x80})
	requireT.EqualValues(5, n)
	requireT.EqualValues(max5Bytes, v)

	v, n = Parse([]byte{followMask, followMask, followMask, followMask, followMask, 0x01})
	requireT.EqualValues(6, n)
	requireT.EqualValues(max5Bytes+1, v)

	v, n = Parse([]byte{followMask, followMask, followMask, followMask, followMask, 0x01, 0x80})
	requireT.EqualValues(6, n)
	requireT.EqualValues(max5Bytes+1, v)

	v, n = Parse([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, numberMask})
	requireT.EqualValues(6, n)
	requireT.EqualValues(max6Bytes, v)

	v, n = Parse([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, numberMask, 0x80})
	requireT.EqualValues(6, n)
	requireT.EqualValues(max6Bytes, v)

	v, n = Parse([]byte{followMask, followMask, followMask, followMask, followMask, followMask, 0x01})
	requireT.EqualValues(7, n)
	requireT.EqualValues(max6Bytes+1, v)

	v, n = Parse([]byte{followMask, followMask, followMask, followMask, followMask, followMask, 0x01, 0x80})
	requireT.EqualValues(7, n)
	requireT.EqualValues(max6Bytes+1, v)

	v, n = Parse([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, numberMask})
	requireT.EqualValues(7, n)
	requireT.EqualValues(max7Bytes, v)

	v, n = Parse([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, numberMask, 0x80})
	requireT.EqualValues(7, n)
	requireT.EqualValues(max7Bytes, v)

	v, n = Parse([]byte{followMask, followMask, followMask, followMask, followMask, followMask, followMask, 0x01})
	requireT.EqualValues(8, n)
	requireT.EqualValues(max7Bytes+1, v)

	v, n = Parse([]byte{followMask, followMask, followMask, followMask, followMask, followMask, followMask, 0x01, 0x80})
	requireT.EqualValues(8, n)
	requireT.EqualValues(max7Bytes+1, v)

	v, n = Parse([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, numberMask})
	requireT.EqualValues(8, n)
	requireT.EqualValues(max8Bytes, v)

	v, n = Parse([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, numberMask, 0x80})
	requireT.EqualValues(8, n)
	requireT.EqualValues(max8Bytes, v)

	v, n = Parse([]byte{followMask, followMask, followMask, followMask, followMask, followMask, followMask, followMask,
		0x01})
	requireT.EqualValues(9, n)
	requireT.EqualValues(max8Bytes+1, v)

	v, n = Parse([]byte{followMask, followMask, followMask, followMask, followMask, followMask, followMask, followMask,
		0x01, 0x80})
	requireT.EqualValues(9, n)
	requireT.EqualValues(max8Bytes+1, v)

	v, n = Parse([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})
	requireT.EqualValues(9, n)
	requireT.EqualValues(uint64(math.MaxUint64), v)

	v, n = Parse([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80})
	requireT.EqualValues(9, n)
	requireT.EqualValues(uint64(math.MaxUint64), v)
}
