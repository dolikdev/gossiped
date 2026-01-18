package utils

import (
	"strings"
	"unicode/utf8"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// CodePage895 is the Kamenický encoding.
var CodePage895 encoding.Encoding = cp895Encoding{}

type cp895Encoding struct{}

func (cp895Encoding) NewDecoder() *encoding.Decoder {
	return &encoding.Decoder{Transformer: &cp895Decoder{}}
}

func (cp895Encoding) NewEncoder() *encoding.Encoder {
	return &encoding.Encoder{Transformer: &cp895Encoder{}}
}

func (cp895Encoding) String() string {
	return "CP895"
}

type cp895Decoder struct{ transform.NopResetter }

func (d *cp895Decoder) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for nSrc < len(src) {
		r := decode895[src[nSrc]]
		size := utf8.RuneLen(r)
		if nDst+size > len(dst) {
			err = transform.ErrShortDst
			break
		}
		nDst += utf8.EncodeRune(dst[nDst:], r)
		nSrc++
	}
	return nDst, nSrc, err
}

type cp895Encoder struct{ transform.NopResetter }

func (e *cp895Encoder) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for nSrc < len(src) {
		r, size := utf8.DecodeRune(src[nSrc:])
		if r == utf8.RuneError && size == 1 {
			if !atEOF && !utf8.FullRune(src[nSrc:]) {
				err = transform.ErrShortSrc
				break
			}
		}

		b, ok := encode895[r]
		if !ok {
			if nDst >= len(dst) {
				err = transform.ErrShortDst
				break
			}
			dst[nDst] = 0x1a // replacement
			nDst++
			nSrc += size
			continue
		}

		if nDst >= len(dst) {
			err = transform.ErrShortDst
			break
		}
		dst[nDst] = b
		nDst++
		nSrc += size
	}
	return nDst, nSrc, err
}

var decode895 = [256]rune{
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
	0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
	0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, 0x2E, 0x2F,
	0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, 0x3E, 0x3F,
	0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4A, 0x4B, 0x4C, 0x4D, 0x4E, 0x4F,
	0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5A, 0x5B, 0x5C, 0x5D, 0x5E, 0x5F,
	0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A, 0x6B, 0x6C, 0x6D, 0x6E, 0x6F,
	0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7A, 0x7B, 0x7C, 0x7D, 0x7E, 0x7F,
	0x010C, 0x011B, 0x0161, 0x00FD, 0x00E1, 0x00ED, 0x00E9, 0x0159, 0x017E, 0x00FA, 0x016F, 0x0165, 0x010F, 0x0148, 0x0160, 0x011A,
	0x00C1, 0x00CD, 0x00C9, 0x0158, 0x017D, 0x00DA, 0x016E, 0x0164, 0x010E, 0x0147, 0x00D6, 0x00DC, 0x00A3, 0x00A5, 0x20A7, 0x0192,
	0x00AA, 0x02C7, 0x02D8, 0x02D9, 0x02DD, 0x02DB, 0x00A6, 0x00A7, 0x00A4, 0x00A9, 0x00AA, 0x00AB, 0x00AC, 0x00AD, 0x00AE, 0x00AF,
	0x2591, 0x2592, 0x2593, 0x2502, 0x2524, 0x2561, 0x2562, 0x2556, 0x2555, 0x2563, 0x2551, 0x2557, 0x255D, 0x255C, 0x255B, 0x2510,
	0x2514, 0x2534, 0x252C, 0x251C, 0x2500, 0x253C, 0x255E, 0x255F, 0x255A, 0x2554, 0x2569, 0x2566, 0x2560, 0x2550, 0x256C, 0x2567,
	0x2568, 0x2564, 0x2565, 0x2559, 0x2558, 0x2552, 0x2553, 0x256B, 0x256A, 0x2518, 0x250C, 0x2588, 0x2584, 0x258C, 0x2590, 0x2580,
	0x03B1, 0x00DF, 0x0393, 0x03C0, 0x03A3, 0x03C3, 0x00B5, 0x03C4, 0x03A6, 0x0398, 0x03A9, 0x03B4, 0x221E, 0x03C6, 0x03B5, 0x2229,
	0x2261, 0x00B1, 0x2264, 0x2265, 0x2320, 0x2321, 0x00F7, 0x2248, 0x00B0, 0x2219, 0x00B7, 0x221A, 0x207F, 0x00B2, 0x25A0, 0x00A0,
}

var encode895 = map[rune]byte{}

func init() {
	for i, r := range decode895 {
		if i == 0 && r == 0 {
			encode895[0] = 0
			continue
		}
		if r != 0 {
			encode895[r] = byte(i)
		}
	}
}

var (
	cDecoder = map[string]*encoding.Decoder{
		"CP866":   charmap.CodePage866.NewDecoder(),
		"+7_FIDO": charmap.CodePage866.NewDecoder(),
		"+7":      charmap.CodePage866.NewDecoder(),
		"IBM866":  charmap.CodePage866.NewDecoder(),
		"CP850":   charmap.CodePage850.NewDecoder(),
		"CP852":   charmap.CodePage852.NewDecoder(),
		"CP848":   charmap.CodePage866.NewDecoder(),
		"CP1250":  charmap.Windows1250.NewDecoder(),
		"CP1251":  charmap.Windows1251.NewDecoder(),
		"CP1252":  charmap.Windows1252.NewDecoder(),
		"CP10000": charmap.Macintosh.NewDecoder(),
		"CP437":   charmap.CodePage437.NewDecoder(),
		"CP895":   CodePage895.NewDecoder(),
		"IBMPC":   charmap.CodePage437.NewDecoder(),
		"LATIN-1": charmap.ISO8859_1.NewDecoder(),
		"LATIN-2": charmap.ISO8859_2.NewDecoder(),
		"LATIN-5": charmap.ISO8859_5.NewDecoder(),
		"LATIN-9": charmap.ISO8859_9.NewDecoder(),
		"KOI8-R":  charmap.KOI8R.NewDecoder(),
	}
	cEncoder = map[string]*encoding.Encoder{
		"CP866":   charmap.CodePage866.NewEncoder(),
		"+7_FIDO": charmap.CodePage866.NewEncoder(),
		"+7":      charmap.CodePage866.NewEncoder(),
		"IBM866":  charmap.CodePage866.NewEncoder(),
		"CP850":   charmap.CodePage850.NewEncoder(),
		"CP852":   charmap.CodePage852.NewEncoder(),
		"CP848":   charmap.CodePage866.NewEncoder(),
		"CP1250":  charmap.Windows1250.NewEncoder(),
		"CP1251":  charmap.Windows1251.NewEncoder(),
		"CP1252":  charmap.Windows1252.NewEncoder(),
		"CP10000": charmap.Macintosh.NewEncoder(),
		"CP437":   charmap.CodePage437.NewEncoder(),
		"CP895":   CodePage895.NewEncoder(),
		"IBMPC":   charmap.CodePage437.NewEncoder(),
		"LATIN-1": charmap.ISO8859_1.NewEncoder(),
		"LATIN-2": charmap.ISO8859_2.NewEncoder(),
		"LATIN-5": charmap.ISO8859_5.NewEncoder(),
		"LATIN-9": charmap.ISO8859_9.NewEncoder(),
		"KOI8-R":  charmap.KOI8R.NewEncoder(),
	}
)

// DecodeCharmap decode string from charmap
func DecodeCharmap(s string, c string) string {
	var dec *encoding.Decoder
	switch chrs := strings.ToUpper(c); chrs {
	case "CP866", "+7_FIDO", "+7", "IBM866", "CP850", "CP852", "CP848", "CP1250", "CP1251", "CP1252", "CP10000", "CP437", "CP895", "IBMPC", "LATIN-2", "LATIN-5", "LATIN-9", "KOI8-R":
		dec = cDecoder[chrs]
	case "UTF-8":
		return s
	default:
		dec = cDecoder["LATIN-1"]
	}
	b, err := dec.String(s)
	if err != nil {
		return s
	}
	return b
}

// EncodeCharmap encode string to charmap
func EncodeCharmap(s string, c string) string {
	var enc *encoding.Encoder
	switch c {
	case "CP866", "+7_FIDO", "+7", "IBM866", "CP850", "CP852", "CP848", "CP1250", "CP1251", "CP1252", "CP10000", "CP437", "CP895", "IBMPC", "LATIN-2", "LATIN-5", "LATIN-9", "KOI8-R":
		enc = cEncoder[c]
	case "UTF-8":
		return s
	default:
		enc = cEncoder["LATIN-1"]
	}
	out, err := encoding.ReplaceUnsupported(enc).String(s)
	if err != nil {
		panic(err)
	}
	return out
}
