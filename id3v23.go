// The basic structure of the ID3 v2.3 tag
// and functions handling it's parsing and writing
// back in the mp3 file

package id3go

import (
	"bufio"
	"fmt"
	"encoding/json"
)

// The structure holding a ID3 v2.3 tag header.
type ID3v23Header struct {
	Version               [2]int
	Unsynchronisation     bool
	ExtendedHeader        bool
	ExperimentalIndicator bool
	Size                  int
}

// The structure holding a ID3 v2.3 tag frame.
type ID3v23Frame struct {
	ID               string
	Size             int
	TagAlterPreserv  bool
	FileAlterPreserv bool
	ReadOnly         bool
	Compression      bool
	Encryption       bool
	GroupingIdentity bool
	TextEncoding     byte
	Data             string
}

// The structure holding a ID3 v2.3 tag.
type ID3v23 struct {
	Header ID3v23Header
	Frames []ID3v23Frame
}

// Reads size data saved as synchsafe integer from files
// containing v2.3 tag
func parseV23Size(reader *bufio.Reader) int {
	return parseSize(readB(reader, 4))
}

// Parses ID3 v2.3 tag from mp3 file, passed as 
// bufio.Reader and returns the ID3v23 tag JSON data
func parseV23Tag(reader *bufio.Reader) []byte {
	tag := new(ID3v23)
	parseV23Header(reader, tag)
	if tag.Header.ExtendedHeader {
		skipB(reader, parseV23Size(reader) - 4)
	}

	for hasFrame(reader, 4) {
		parseV23Frame(reader, tag)
	}

	result, err := json.Marshal(tag)
	if err != nil {
		panic(fmt.Sprintf("Unable to convert tag to JSON data."))
	}

	return result;
}

// Parses the ID3 v2.3 tag header to the Header component
// of the ID3v23 tag.
func parseV23Header(reader *bufio.Reader, tag *ID3v23) {
	skipB(reader, 3)
	tag.Header.Version[0] = int(readB(reader, 1)[0])
	tag.Header.Version[1] = int(readB(reader, 1)[0])
	flags := readB(reader, 1)
	tag.Header.Unsynchronisation = flags[0] & 1<<7 != 0
	tag.Header.ExtendedHeader = flags[0] & 1<<6 != 0
	tag.Header.ExperimentalIndicator = flags[0] & 1<<5 != 0
	tag.Header.Size = parseV23Size(reader)
}

// Parses the next frame and appends it to the tag.
// Takes 2 arguments, bufio.Reader and ID3v23 tag.
func parseV23Frame(reader *bufio.Reader, tag *ID3v23) {
	var frame ID3v23Frame
	frame.ID = string(readB(reader, 4))
	frame.Size = parseV23Size(reader)
	flags := readB(reader, 2)
	frame.TagAlterPreserv = flags[0] & 1<<7 != 0
	frame.FileAlterPreserv = flags[0] & 1<<6 != 0
	frame.ReadOnly = flags[0] & 1<<5 != 0
	frame.Compression = flags[1] & 1<<7 != 0
	frame.Encryption = flags[1] & 1<<6 != 0
	frame.GroupingIdentity = flags[1] & 1<<5 != 0

	if frame.ID[0] != 'T' && frame.ID[0] != 'W'  {
		skipB(reader, frame.Size)
		return
	}

	frame.TextEncoding = readB(reader, 1)[0]
	if frame.ID == "TCON" {
		frame.Data = parseGenre(readB(reader, frame.Size - 1))
	} else {
		frame.Data = string(readB(reader, frame.Size - 1))
	}
	tag.Frames = append(tag.Frames, frame)
}
