// In this file there are some helper functions used in the lib.
package id3go

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

// Skipper buffer for the skipB function
var skipper []byte = make([]byte, 1024)

// hasFrame checks if the next 'n' bytes form a frame ID.
// It takes 2 arguments, a bufio.Reader and a integer.
func hasFrame(reader *bufio.Reader, n int) bool {
	frameData, err := reader.Peek(n)
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

// hasFooter looks for footer identifier ("3DI")
func hasFooter(reader *bufio.Reader) bool {
	footerData, err := reader.Peek(3)
	if err != nil {
		return false
	}

	if string(footerData) == "3DI" {
		return true;
	}
	return false;
}

// Turns 32 bit syncsafe integer into int.
func parseSize(data []byte) int {
	var size byte
	for _, val := range data {
		size = size<<7 | val
	}

	return int(size)
}

// TODO
func parseString(reader *bufio.Reader, n int) string {
	return "TODODODO"
}

// parseGenre takes a byte slice as argument and returns
// a string, describing the file's genre.
func parseGenre(data []byte) string {
	if len(data) == 1 {
		return ConvGenInd(int(data[0]))
	}

	genre := string(data)
	if genre == "RX" || strings.HasPrefix(genre, "(RX)") {
		return "Remix"
	}
	if genre == "CR" || strings.HasPrefix(genre, "(CV)") {
		return "Cover"
	}

	index, err := strconv.Atoi(genre)
	if err == nil {
		if index >= 0 && index < len(genres) {
			return ConvGenInd(index)
		}
		return "Unknown"
	}

	index = -1
	_, err = fmt.Sscanf(genre, "(%d)", &index)
	if err == nil {
		if index >= 0 && index < len(genres) {
			return ConvGenInd(index)
		}
		return "Unknown"
	}

	return genre
}

// readB reads the next 'n' bytes from the given bufio.Reader
// and returns a byte slice.
func readB(reader *bufio.Reader, n int) []byte {
	buf := make([]byte, n)
	
	_, err := reader.Read(buf)
	if err != nil {
		panic(err)
	}
	
	return buf
}

// skipB skips the next 'n' bytes from the given bufio.Reader.
func skipB(reader *bufio.Reader, n int) {
	pos := 0
	for pos < n {
		left := n - pos
		if left > len(skipper) {
			left = len(skipper)
		}
		
		i, err := reader.Read(skipper[:left])
		if err != nil {
			panic(err)
		}
		
		pos += i
	}
}

// isV2Tag checks if ID3 v2.x tag is present.
func isV2Tag(reader *bufio.Reader) bool {
	data, err := reader.Peek(3)
	if err != nil {
		return false
	}
	
	if data[0] == 'I' && data[1] == 'D' && data[2] == '3' {
		return true
	}
	
	return false
}

// isV1Tag checks if ID3 v1.x tag is present.
func isV1Tag(file *os.File) bool {
	data := make([]byte, 3)
	f, err := file.Stat();
	if err != nil {
		panic(fmt.Sprintf("Unable to read file stats."))
	}
	file.ReadAt(data, f.Size() - 128)

	if data[0] == 'T' && data[1] == 'A' && data[2] == 'G' {
		return true
	}
	
	return false
}