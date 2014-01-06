package id3go

import (
	"bufio"
	"fmt"
)

var skipper []byte = make([]byte, 1024)

func hasFrame(reader *bufio.Reader, size int) bool {
	frameData, err := reader.Peek(size)
	if err != nil {
		return false
	}
	
	for _, ch := range frameData {
		if (ch < 'A' || ch > 'Z') && (ch < '0' || ch > '9') {
			return false
		}
	}
	
	return true
}

func parseSize(data []byte) int {
	return int(data[0] << 21 | data[1] << 14 | data[2] << 7 | data[3])
}

func parseString(reader *bufio.Reader, n int) string {
	
}

func parseGenre(reader *bufio.Reader, n int) string {
	return convV1Genre(parseString(readB(reader, n)))
}

func readB(reader *bufio.Reader, n int) []byte {
	buf := make([]byte, n)
	
	i, err := reader.Read(buf)
	if err != nil {
		panic(err)
	}
	
	return buf
}

func skipB(reader *bufio.Reader, n int) {
	pos := 0
	for pos < n {
		left := n - pos
		if end > len(skipper) {
			end = len(skipper)
		}
		
		i, err := reader.Read(skipper[0:end]
		if err != nil {
			panic(err)
		}
		
		pos += i
	}
}
