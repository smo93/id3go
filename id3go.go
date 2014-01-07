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
	
	Data    string
}

type ID3 struct {
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
	data := readB(reader, 10)
	tag.Header.Version[0] = int(data[3])
	tag.Header.Version[1] = int(data[4])
	tag.Header.Unsynchronization = data[5] & (1 << 7) != 0
	tag.Header.Extended = data[5] & (1 << 6) != 0
	tag.Header.Experimental = data[5] & (1 << 5) != 0
	tag.Header.Footer = data[5] & (1 << 4) != 0
	tag.Header.Size = parseSize(data[6:])
}

func isTag(reader *bufio.Reader) bool {
	data, err := reader.Peek(3)
	if err != nil {
		return false
	}
	
	if data[0] == 'I' && data[1] == 'D' && data[2] == '3' {
		return true
	}
	
	return false
}

func ParseTag(reader io.Reader) *ID3 {
	
}

func (frame *ID3v2Frame) EditFrame(data String) {
	
}

func (header *ID3v2Header) EditHeader(unsynch, ext, exp, foo Bool) {
	
} 

func WriteToFile(writer io.Writer, tag *ID3) {
	
}
