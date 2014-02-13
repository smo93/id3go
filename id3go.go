package id3go

import (
	"os"
	"bufio"
	"fmt"
)

// ParseTag takes the path to the mp3 file as a string, checks the tag
// version and returns the tag's JSON data.
func ParseTag(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	if isV2Tag(reader) {
		//parse v2 tag
		version, err := reader.Peek(5)
		if err != nil {
			panic(err)
		}
		switch int(version[3]) {
			case 2:
				return parseV22Tag(reader)
			case 3:
				return parseV23Tag(reader)
			case 4:
				return parseV24Tag(reader)
		}
	}

	if isV1Tag(file) {
		//parse v1 tag
		data := make([]byte, 128)
		f, err := file.Stat();
		if err != nil {
			panic(fmt.Sprintln("Unable to read file stats."))
		}
		file.ReadAt(data, f.Size() - 128)
		return parseV1(data)
	}

	//no tag in the file
	return nil
}