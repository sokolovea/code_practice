package main

import (
	"strings"
)

// начало решения
//
// Bit masks for each code point under U+0100, for fast lookup.
const (
	pC      = 1 << iota // a control character.
	pD                  // a punctuation character.
	pN                  // a numeral.
	pS                  // a symbolic character.
	pZ                  // a spacing character.
	pLu                 // an upper-case letter.
	pLl                 // a lower-case letter.
	pp                  // a printable character according to Go's definition.
	pg      = pp | pZ   // a graphical character according to the Unicode definition.
	pLo     = pLl | pLu // a letter that is neither upper nor lower case.
	pLmask  = pLo
	pLValid = pLmask | pN | pD // '-'
)

var properties = [256]uint8{
	0x00: pC,       // '\x00'
	0x01: pC,       // '\x01'
	0x02: pC,       // '\x02'
	0x03: pC,       // '\x03'
	0x04: pC,       // '\x04'
	0x05: pC,       // '\x05'
	0x06: pC,       // '\x06'
	0x07: pC,       // '\a'
	0x08: pC,       // '\b'
	0x09: pC,       // '\t'
	0x0A: pC,       // '\n'
	0x0B: pC,       // '\v'
	0x0C: pC,       // '\f'
	0x0D: pC,       // '\r'
	0x0E: pC,       // '\x0e'
	0x0F: pC,       // '\x0f'
	0x10: pC,       // '\x10'
	0x11: pC,       // '\x11'
	0x12: pC,       // '\x12'
	0x13: pC,       // '\x13'
	0x14: pC,       // '\x14'
	0x15: pC,       // '\x15'
	0x16: pC,       // '\x16'
	0x17: pC,       // '\x17'
	0x18: pC,       // '\x18'
	0x19: pC,       // '\x19'
	0x1A: pC,       // '\x1a'
	0x1B: pC,       // '\x1b'
	0x1C: pC,       // '\x1c'
	0x1D: pC,       // '\x1d'
	0x1E: pC,       // '\x1e'
	0x1F: pC,       // '\x1f'
	0x20: pZ | pp,  // ' '
	0x21: pp,       // '!'
	0x22: pp,       // '"'
	0x23: pp,       // '#'
	0x24: pS | pp,  // '$'
	0x25: pp,       // '%'
	0x26: pp,       // '&'
	0x27: pp,       // '\''
	0x28: pp,       // '('
	0x29: pp,       // ')'
	0x2A: pp,       // '*'
	0x2B: pS | pp,  // '+'
	0x2C: pp,       // ','
	0x2D: pp | pD,  // '-'
	0x2E: pp,       // '.'
	0x2F: pp,       // '/'
	0x30: pN | pp,  // '0'
	0x31: pN | pp,  // '1'
	0x32: pN | pp,  // '2'
	0x33: pN | pp,  // '3'
	0x34: pN | pp,  // '4'
	0x35: pN | pp,  // '5'
	0x36: pN | pp,  // '6'
	0x37: pN | pp,  // '7'
	0x38: pN | pp,  // '8'
	0x39: pN | pp,  // '9'
	0x3A: pp,       // ':'
	0x3B: pp,       // ';'
	0x3C: pS | pp,  // '<'
	0x3D: pS | pp,  // '='
	0x3E: pS | pp,  // '>'
	0x3F: pp,       // '?'
	0x40: pp,       // '@'
	0x41: pLu | pp, // 'A'
	0x42: pLu | pp, // 'B'
	0x43: pLu | pp, // 'C'
	0x44: pLu | pp, // 'D'
	0x45: pLu | pp, // 'E'
	0x46: pLu | pp, // 'F'
	0x47: pLu | pp, // 'G'
	0x48: pLu | pp, // 'H'
	0x49: pLu | pp, // 'I'
	0x4A: pLu | pp, // 'J'
	0x4B: pLu | pp, // 'K'
	0x4C: pLu | pp, // 'L'
	0x4D: pLu | pp, // 'M'
	0x4E: pLu | pp, // 'N'
	0x4F: pLu | pp, // 'O'
	0x50: pLu | pp, // 'P'
	0x51: pLu | pp, // 'Q'
	0x52: pLu | pp, // 'R'
	0x53: pLu | pp, // 'S'
	0x54: pLu | pp, // 'T'
	0x55: pLu | pp, // 'U'
	0x56: pLu | pp, // 'V'
	0x57: pLu | pp, // 'W'
	0x58: pLu | pp, // 'X'
	0x59: pLu | pp, // 'Y'
	0x5A: pLu | pp, // 'Z'
	0x5B: pp,       // '['
	0x5C: pp,       // '\\'
	0x5D: pp,       // ']'
	0x5E: pS | pp,  // '^'
	0x5F: pp,       // '_'
	0x60: pS | pp,  // '`'
	0x61: pLl | pp, // 'a'
	0x62: pLl | pp, // 'b'
	0x63: pLl | pp, // 'c'
	0x64: pLl | pp, // 'd'
	0x65: pLl | pp, // 'e'
	0x66: pLl | pp, // 'f'
	0x67: pLl | pp, // 'g'
	0x68: pLl | pp, // 'h'
	0x69: pLl | pp, // 'i'
	0x6A: pLl | pp, // 'j'
	0x6B: pLl | pp, // 'k'
	0x6C: pLl | pp, // 'l'
	0x6D: pLl | pp, // 'm'
	0x6E: pLl | pp, // 'n'
	0x6F: pLl | pp, // 'o'
	0x70: pLl | pp, // 'p'
	0x71: pLl | pp, // 'q'
	0x72: pLl | pp, // 'r'
	0x73: pLl | pp, // 's'
	0x74: pLl | pp, // 't'
	0x75: pLl | pp, // 'u'
	0x76: pLl | pp, // 'v'
	0x77: pLl | pp, // 'w'
	0x78: pLl | pp, // 'x'
	0x79: pLl | pp, // 'y'
	0x7A: pLl | pp, // 'z'
	0x7B: pp,       // '{'
	0x7C: pS | pp,  // '|'
	0x7D: pp,       // '}'
	0x7E: pS | pp,  // '~'
}

// slugify возвращает "безопасный" вариант заголовка:
// только латиница, цифры и дефис
func slugify(src string) string {
	result := strings.Builder{}
	result.Grow(len(src))
	startIndex := 0
	isFirstWord := true
	for i := 0; i < len(src); i++ {
		c := src[i]
		if properties[c]&pLValid == 0 {
			if startIndex != i {
				if !isFirstWord {
					result.WriteByte('-')
				}
				isFirstWord = false
				subStr := src[startIndex:i]
				for j := range len(subStr) {
					ch2 := subStr[j] | ' '
					result.WriteByte(ch2)
				}
			}
			startIndex = i + 1
		}
	}
	if startIndex != len(src) {
		if !isFirstWord {
			result.WriteByte('-')
		}
	}
	subStr := src[startIndex:]
	for j := range len(subStr) {
		ch2 := subStr[j] | ' '
		result.WriteByte(ch2)
	}
	return result.String()
}
