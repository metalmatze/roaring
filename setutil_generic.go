// +build !arm64 gccgo appengine

package roaring


func union2by2(set1 []uint16, set2 []uint16, buffer []uint16) int {
	pos := uint(0)
	k1 := uint(0)
	k2 := uint(0)
	len1 := uint(len(set1))
	len2 := uint(len(set2))
	if 0 == len(set2) {
		buffer = buffer[:len(set1)]
		copy(buffer, set1[:])
		return len(set1)
	}
	if 0 == len(set1) {
		buffer = buffer[:len(set2)]
		copy(buffer, set2[:])
		return len(set2)
	}
	s1 := set1[k1]
	s2 := set2[k2]
	buffer = buffer[:len1 + len2]
	for pos < uint(len(buffer)) {
		if pos >= uint(len(buffer)) {
			panic("You provided a buffer with insufficient capacity")
		}
		if s1 < s2 {
			buffer[pos] = s1
			pos++
			k1++
			if k1 >= len1 {
				copy(buffer[pos:], set2[k2:])
				pos += len2 - k2
				return int(pos)
			}
			s1 = set1[k1]
		} else if s1 == s2 {
			buffer[pos] = s1
			pos++
			k1++
			k2++
			if k1 >= len1 {
				copy(buffer[pos:], set2[k2:])
				pos += len2 - k2
				break
			}
			if k2 >= len2 {
				copy(buffer[pos:], set1[k1:])
				pos += len1 - k1
				break
			}
			s1 = set1[k1]
			s2 = set2[k2]
		} else { // if (set1[k1]>set2[k2])
			buffer[pos] = s2
			pos++
			k2++
			if k2 >= len2 {
				copy(buffer[pos:], set1[k1:])
				pos += len1 - k1
				break
			}
			s2 = set2[k2]
		}
	}
	return int(pos)
}