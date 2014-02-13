// The basic structure of the ID3 v2.4 tag
// and functions handling it's parsing and writing
// back in the mp3 file

package id3go

import (
	"bufio"
	"fmt"
	"encoding/json"
	"strings"
)

// The structure holding a ID3 v2.4 tag header.
type ID3v24Header struct {
	Version           [2]int
	Unsynchronisation bool
	ExtendedHeader    bool
	ExperimentalInd   bool
	FooterPresent     bool
	Size              int
}

// The structure holding a ID3 v2.4 tag footer.
type ID3v24Footer struct {
	Version           [2]int
	Unsynchronisation bool
	ExtendedHeader    bool
	ExperimentalInd   bool
	FooterPresent     bool
	Size              int
}

// The structure holding a ID3 v2.4 tag frame.
type ID3v24Frame struct {
	ID                    string
	Size                  int
	TagAlterPreservation  bool
	FileAlterPreservation bool
	ReadOnly              bool
	GroupingIdentity      bool
	Compression           bool
	Encryption            bool
	Unsynchronisation     bool
	DataLengthIdicator    bool
	TextEncoding          byte
	Data                  string
}

// The structure holding a ID3 v2.4 tag.
type ID3v24 struct {
	Header ID3v24Header
	Frames []ID3v24Frame
	Footer *ID3v24Footer
}

// Reads size data saved as synchsafe integer from files
// containing v2.4 tag
func parseV24Size(reader *bufio.Reader) int {
	return parseSize(readB(reader, 4))
}

// Parses ID3 v2.2 tag from mp3 file, passed as 
// bufio.Reader and returns the ID3v24 tag JSON data.
func parseV24Tag(reader *bufio.Reader) []byte {
	tag := new(ID3v24)
	parseV24Header(reader, tag)
	if tag.Header.ExtendedHeader {
		skipB(reader, parseV24Size(reader) - 4)
	}

	for hasFrame(reader, 4) {
		parseV24Frame(reader, tag)
	}

	if tag.Header.FooterPresent {
		for hasFooter(reader) == false {
			skipB(reader, 1)
		}
		parseV24Footer(reader, tag)
	} else {
		tag.Footer = nil
	}

	result, err := json.Marshal(tag)
	if err != nil {
		panic(fmt.Sprintf("Unable to convert tag to JSON data."))
	}

	return result;
}

// Parses the ID3 v2.4 tag header to the Header component
// of the ID3v24 tag.
func parseV24Header(reader *bufio.Reader, tag *ID3v24) {
	skipB(reader, 3)
	tag.Header.Version[0] = int(readB(reader, 1)[0])
	tag.Header.Version[1] = int(readB(reader, 1)[0])
	flags := readB(reader, 1)
	tag.Header.Unsynchronisation = flags[0] & 1<<7 != 0
	tag.Header.ExtendedHeader = flags[0] & 1<<6 != 0
	tag.Header.ExperimentalInd = flags[0] & 1<<5 != 0
	tag.Header.FooterPresent = flags[0] & 1<<4 != 0
	tag.Header.Size = parseV24Size(reader)
}

// Parses the next frame and appends it to the tag.
// Takes 2 arguments, bufio.Reader and ID3v24 tag.
func parseV24Frame(reader *bufio.Reader, tag *ID3v24) {
	var frame ID3v24Frame

	frame.ID = string(readB(reader, 4))
	frame.Size = parseV24Size(reader)
	flags := readB(reader, 2)
	frame.TagAlterPreservation = flags[0] & 1<<6 != 0
	frame.FileAlterPreservation = flags[0] & 1<<5 != 0
	frame.ReadOnly = flags[0] & 1<<4 != 0
	frame.GroupingIdentity = flags[1] & 1<<6 != 0
	frame.Compression = flags[1] & 1<<3 != 0
	frame.Encryption = flags[1] & 1<<2 != 0
	frame.Unsynchronisation = flags[1] & 1<<1 != 0
	frame.DataLengthIdicator = flags[1] & 1 != 0

	if frame.ID[0] != 'T' && frame.ID[0] != 'W'  {
		skipB(reader, frame.Size)
		return
	}

	frame.TextEncoding = readB(reader, 1)[0]
	if frame.ID == "TCON" {
		frame.Data = parseGenre(readB(reader, frame.Size - 1))
	} else {
		frame.Data = strings.TrimSpace(string(readB(reader, frame.Size - 1)))
	}
	tag.Frames = append(tag.Frames, frame)
}

// Parses the ID3 v2.4 tag footer to the Footer component
// of the ID3v24 tag.
func parseV24Footer(reader *bufio.Reader, tag *ID3v24) {
	tag.Footer = new(ID3v24Footer)
	tag.Footer.Version = tag.Header.Version
	tag.Footer.Unsynchronisation = tag.Header.Unsynchronisation
	tag.Footer.ExtendedHeader = tag.Header.ExtendedHeader
	tag.Footer.ExperimentalInd = tag.Header.ExperimentalInd
	tag.Footer.FooterPresent = tag.Header.FooterPresent
	tag.Footer.Size = tag.Header.Size
}
