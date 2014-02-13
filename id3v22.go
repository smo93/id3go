// The basic structure of the ID3 v2.2 tag
// and functions handling it's parsing and writing
// back in the mp3 file

package id3go

import (
	"bufio"
	"fmt"
	"encoding/json"
)

// The structure holding a ID3 v2.2 tag header.
type ID3v22Header struct {
	Version           [2]int
	Unsynchronisation bool
	Compression       bool
	Size              int
}

// The structure holding a ID3 v2.3 tag frame.
type ID3v22Frame struct {
	ID   string
	Size int
	TextEncoding byte
	Data string
}

// The structure holding a ID3 v2.3 tag.
type ID3v22 struct {
	Header ID3v22Header
	Frames []ID3v22Frame
}

// Reads ID3 v2.2 tag size saved as synchsafe integer
func parseV22Size(reader *bufio.Reader) int {
	return parseSize(readB(reader, 4))
}

// Reads size data saved as synchsafe integer from ID3 v2.2
// tag frames
func parseV22FrameSize(reader *bufio.Reader) int {
	size := readB(reader, 3)
	return int(size[0] << 14 | size[1] << 7 | size[2])
}

// Parses ID3 v2.2 tag from mp3 file, passed as 
// bufio.Reader and returns the ID3v22 tag JSON data.
func parseV22Tag(reader *bufio.Reader) []byte {
	tag := new(ID3v22)
	parseV22Header(reader, tag)

	for hasFrame(reader, 3) {
		parseV22Frame(reader, tag)
	}
	result, err := json.Marshal(tag)
	if err != nil {
		panic(fmt.Sprintf(fmt.Sprintln(err)))
	}

	return result;
}

// Parses the ID3 v2.2 tag header to the Header component
// of the ID3v22 tag.
func parseV22Header(reader *bufio.Reader, tag *ID3v22) {
	skipB(reader, 3)
	tag.Header.Version[0] = int(readB(reader, 1)[0])
	tag.Header.Version[1] = int(readB(reader, 1)[0])
	flags := readB(reader, 1)
	tag.Header.Unsynchronisation = flags[0] & 1<<7 != 0
	tag.Header.Compression = flags[0] & 1<<6 != 0
	tag.Header.Size = parseV22Size(reader)
}

// Parses the next frame and appends it to the tag.
// Takes 2 arguments, bufio.Reader and ID3v22 tag.
func parseV22Frame(reader *bufio.Reader, tag *ID3v22) {
	var frame ID3v22Frame
	frame.ID = string(readB(reader, 3))
	//fmt.Println(frame.ID)
	frame.Size = parseV22FrameSize(reader)
	if frame.ID[0] != 'T' && frame.ID[0] != 'W'  {
		skipB(reader, frame.Size)
		return
	}
	frame.TextEncoding = readB(reader, 1)[0]
	if frame.ID == "TCO" {
		frame.Data = parseGenre(readB(reader, frame.Size - 1))
	} else {
		frame.Data = string(readB(reader, frame.Size - 1))
	}
	tag.Frames = append(tag.Frames, frame)
}
