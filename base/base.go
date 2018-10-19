package base

import (
	"math"
	"strings"

	"shortme/conf"
)

// Int2String converts an unsigned 64bit integer to a string.
func Int2String(seq uint64) (shortURL string) {
	charSeq := []rune{}
	if seq != 0 {
		for seq != 0 {
			mod := seq % conf.Conf.Common.BaseStringLength
			div := seq / conf.Conf.Common.BaseStringLength
			charSeq = append(charSeq, rune(conf.Conf.Common.BaseString[mod]))
			seq = div
		}
	} else {
		charSeq = append(charSeq, rune(conf.Conf.Common.BaseString[seq]))
	}

	tmpShortURL := string(charSeq)
	shortURL = reverse(tmpShortURL)
	return
}

// String2Int converts a short URL string to an unsigned 64bit integer.
func String2Int(shortURL string) (seq uint64) {
	shortURL = reverse(shortURL)
	for index, char := range shortURL {
		base := uint64(math.Pow(float64(conf.Conf.Common.BaseStringLength), float64(index)))
		seq += uint64(strings.Index(conf.Conf.Common.BaseString, string(char))) * base
	}
	return
}
