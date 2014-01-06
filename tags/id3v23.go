package id3go

import (
	"bufio"
)

func parseV23Size(reader *bufio.Reader) int {
	return parseSize(readBytes(reader, 4))
}

func parseV23Tag(reader *bufio.Reader, tag *ID3) {

}

func parseV23Frame(reader *bufio.Reader, tag *ID3) {
	
}
