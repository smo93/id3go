package id3go

import (
	"bufio"
)

func parseV24Size(reader *bufio.Reader) int {
	return int(data[0] << 21 | data[1] << 14 | data[2] << 7 | data[3])
}

func parseV24Tag(reader *bufio.Reader, tag *ID3) {

}

func parseV24Frame(reader *bufio.Reader, tag *ID3) {
	
}
