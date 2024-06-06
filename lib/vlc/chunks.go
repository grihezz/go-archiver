package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type encodingTable map[rune]string
type BinaryChunk string
type BinaryChunks []BinaryChunk

const chunkSize = 8
const sep = " "

func Encode(str string) []byte {
	str = prepareText(str)

	chunks := splitByChunks(encodeBin(str), chunkSize)

	return chunks.Bytes()
}

func (bcs BinaryChunks) Bytes() []byte {
	res := make([]byte, 0, len(bcs))
	for _, bc := range bcs {
		res = append(res, bc.Byte())
	}
	return res
}

func (bc BinaryChunk) Byte() byte {
	num, err := strconv.ParseUint(string(bc), 2, chunkSize)
	if err != nil {
		panic("cant parce binary chunk: " + err.Error())
	}
	return byte(num)
}

func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)
	chunkCount := strLen / chunkSize

	if strLen/chunkCount != 0 {
		chunkCount++
	}

	res := make(BinaryChunks, 0, chunkCount)

	var buf strings.Builder

	for i, ch := range bStr {
		buf.WriteString(string(ch))
		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))

		res = append(res, BinaryChunk(lastChunk))
	}
	return res
}

func NewBinChunks(data []byte) BinaryChunks {
	res := make(BinaryChunks, 0, len(data))

	for _, part := range data {
		res = append(res, NewBinChunk(part))
	}

	return res
}
func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}

func (bcs BinaryChunks) Join() string {
	var buf strings.Builder

	for _, bc := range bcs {
		buf.WriteString(string(bc))
	}
	return buf.String()
}
