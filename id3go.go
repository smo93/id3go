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

type ID3v2Frame struct {
	FrameID string
	Size    int
	Flags   [2]byte
	Data    string
}

type ID3 struct {
	fileName string
	
	Header ID3v2Header
	
	Title    ID3v2Frame
	Artist   ID3v2Frame
	Album    ID3v2Frame
	Year     ID3v2Frame
	Comments ID3v2Frame
	Track    ID3v2Frame
	Disc     ID3v2Frame
	Genre    ID3v2Frame
	Length   ID3v2Frame
}

func parseV2Header(reader *bufio.Reader, tag *ID3) {
	
}

func isTag(reader *bufio.Reader) bool {
	
}

func ParseTag(reader io.Reader) *ID3 {
	
}
