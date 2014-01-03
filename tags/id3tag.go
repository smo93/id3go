package id3go

import (
	"io"
	"bufio"
)

type ID3v2Header struct {
	Version           [2]int
	Size              int32
	Unsynchronization bool
	Extended          bool
	Experimental      bool
	Footer            bool
}

type ID3 struct {
	fileName string
	
	Header ID3v2Header
	
	Name   string
	Artist string
	Album  string
	Year   string
	Track  string
	Disc   string
	Genre  string
	Length string
}

func parseV2Header(reader *bufio.Reader, tag *ID3) {
	
}

func isTag(reader *bufio.Reader) bool {
	
}

func ParseTag(reader io.Reader) {
	
}
