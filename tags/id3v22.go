package id3go

import (
	"bufio"
)

func parseV22Size(reader *bufio.Reader) int {
	size := readBytes(reader, 3))
	return int(size[0] << 14 | size[1] << 7 | size[2])
}

func parseV22Tag(reader *bufio.Reader, tag *ID3) {

}

func parseV22Frame(reader *bufio.Reader, tag *ID3) {
	
}
