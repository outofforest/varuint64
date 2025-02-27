package varuint64

const (
	// MaxSize is the maximum possible size of varuint64.
	MaxSize = 9

	followMask = 0x80
	numberMask = followMask - 1
	shiftBits  = 7
)

const (
	max1Byte = 1<<((iota+1)*shiftBits) - 1
	max2Bytes
	max3Bytes
	max4Bytes
	max5Bytes
	max6Bytes
	max7Bytes
	max8Bytes
)

// Size computes the size of the varuint64.
func Size(v uint64) uint64 {
	// For last byte last bit is used for data, not for continuation flag.
	// That's why we may fit 64-bit number in 9 bytes, not 10.

	switch {
	case v <= max1Byte:
		return 1
	case v <= max2Bytes:
		return 2
	case v <= max3Bytes:
		return 3
	case v <= max4Bytes:
		return 4
	case v <= max5Bytes:
		return 5
	case v <= max6Bytes:
		return 6
	case v <= max7Bytes:
		return 7
	case v <= max8Bytes:
		return 8
	default:
		return 9
	}
}

// Contains returns true if slice contains complete varuint64.
func Contains(b []byte) bool {
	switch len(b) {
	case 0:
		return false
	case 1:
		return b[0]&followMask == 0
	case 2:
		return b[0]&b[1]&followMask == 0
	case 3:
		return b[0]&b[1]&b[2]&followMask == 0
	case 4:
		return b[0]&b[1]&b[2]&b[3]&followMask == 0
	case 5:
		return b[0]&b[1]&b[2]&b[3]&b[4]&followMask == 0
	case 6:
		return b[0]&b[1]&b[2]&b[3]&b[4]&b[5]&followMask == 0
	case 7:
		return b[0]&b[1]&b[2]&b[3]&b[4]&b[5]&b[6]&followMask == 0
	case 8:
		return b[0]&b[1]&b[2]&b[3]&b[4]&b[5]&b[6]&b[7]&followMask == 0
	default:
		return true
	}
}

// Put puts varuint64 in the buffer.
func Put(b []byte, v uint64) uint64 {
	// For last byte last bit is used for data, not for continuation flag.
	// That's why we may fit 64-bit number in 9 bytes, not 10.

	switch {
	case v <= max1Byte:
		b[0] = byte(v)
		return 1
	case v <= max2Bytes:
		b[0] = byte(v) | followMask
		v >>= shiftBits
		b[1] = byte(v)
		return 2
	case v <= max3Bytes:
		b[0] = byte(v) | followMask
		v >>= shiftBits
		b[1] = byte(v) | followMask
		v >>= shiftBits
		b[2] = byte(v)
		return 3
	case v <= max4Bytes:
		b[0] = byte(v) | followMask
		v >>= shiftBits
		b[1] = byte(v) | followMask
		v >>= shiftBits
		b[2] = byte(v) | followMask
		v >>= shiftBits
		b[3] = byte(v)
		return 4
	case v <= max5Bytes:
		b[0] = byte(v) | followMask
		v >>= shiftBits
		b[1] = byte(v) | followMask
		v >>= shiftBits
		b[2] = byte(v) | followMask
		v >>= shiftBits
		b[3] = byte(v) | followMask
		v >>= shiftBits
		b[4] = byte(v)
		return 5
	case v <= max6Bytes:
		b[0] = byte(v) | followMask
		v >>= shiftBits
		b[1] = byte(v) | followMask
		v >>= shiftBits
		b[2] = byte(v) | followMask
		v >>= shiftBits
		b[3] = byte(v) | followMask
		v >>= shiftBits
		b[4] = byte(v) | followMask
		v >>= shiftBits
		b[5] = byte(v)
		return 6
	case v <= max7Bytes:
		b[0] = byte(v) | followMask
		v >>= shiftBits
		b[1] = byte(v) | followMask
		v >>= shiftBits
		b[2] = byte(v) | followMask
		v >>= shiftBits
		b[3] = byte(v) | followMask
		v >>= shiftBits
		b[4] = byte(v) | followMask
		v >>= shiftBits
		b[5] = byte(v) | followMask
		v >>= shiftBits
		b[6] = byte(v)
		return 7
	case v <= max8Bytes:
		b[0] = byte(v) | followMask
		v >>= shiftBits
		b[1] = byte(v) | followMask
		v >>= shiftBits
		b[2] = byte(v) | followMask
		v >>= shiftBits
		b[3] = byte(v) | followMask
		v >>= shiftBits
		b[4] = byte(v) | followMask
		v >>= shiftBits
		b[5] = byte(v) | followMask
		v >>= shiftBits
		b[6] = byte(v) | followMask
		v >>= shiftBits
		b[7] = byte(v)
		return 8
	default:
		b[0] = byte(v) | followMask
		v >>= shiftBits
		b[1] = byte(v) | followMask
		v >>= shiftBits
		b[2] = byte(v) | followMask
		v >>= shiftBits
		b[3] = byte(v) | followMask
		v >>= shiftBits
		b[4] = byte(v) | followMask
		v >>= shiftBits
		b[5] = byte(v) | followMask
		v >>= shiftBits
		b[6] = byte(v) | followMask
		v >>= shiftBits
		b[7] = byte(v) | followMask
		v >>= shiftBits
		b[8] = byte(v)
		return 9
	}
}

// Parse parses varuint64 present in the slice.
func Parse(b []byte) (uint64, uint64) {
	v := uint64(b[0] & numberMask)
	if b[0]&followMask == 0 {
		return v, 1
	}

	v |= uint64(b[1]&numberMask) << shiftBits
	if b[1]&followMask == 0 {
		return v, 2
	}

	v |= uint64(b[2]&numberMask) << (2 * shiftBits)
	if b[2]&followMask == 0 {
		return v, 3
	}

	v |= uint64(b[3]&numberMask) << (3 * shiftBits)
	if b[3]&followMask == 0 {
		return v, 4
	}

	v |= uint64(b[4]&numberMask) << (4 * shiftBits)
	if b[4]&followMask == 0 {
		return v, 5
	}

	v |= uint64(b[5]&numberMask) << (5 * shiftBits)
	if b[5]&followMask == 0 {
		return v, 6
	}

	v |= uint64(b[6]&numberMask) << (6 * shiftBits)
	if b[6]&followMask == 0 {
		return v, 7
	}

	v |= uint64(b[7]&numberMask) << (7 * shiftBits)
	if b[7]&followMask == 0 {
		return v, 8
	}

	v |= uint64(b[8]) << (8 * shiftBits)
	return v, 9
}
