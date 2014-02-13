// The basic structure of the ID3 v1.0 and v1.1 tag
// and functions handling it's parsing and writing
// back in the mp3 file

package id3go

import (
	"encoding/json"
	"fmt"
)

// The structure holding a ID3 v1.0 and v1.1 tag.
type ID3v1 struct {
	Version    [2]int
	SongTitle  string
	Artist     string
	Album      string
	Year       string
	Comment    string
	AlbumTrack int
	Genre      string
}


//Reads the mp3 with v1 and v1.1 ID3 tag and returns a JSON data
//Takes one argument of type bufio.Reader
func parseV1(data []byte) []byte {
	fmt.Println("v1x")
	var tag ID3v1
	tag.SongTitle = string(data[3:33])
	tag.Artist = string(data[33:63])
	tag.Album = string(data[63:93])
	tag.Year = string(data[93:97])
	if data[122] == 0 && data[123] != 0 {
		tag.Version[0] = 1
		tag.Version[1] = 1
		tag.Comment = string(data[97:125])
		tag.AlbumTrack = int(data[126])
	} else {
		tag.Version[0] = 1
		tag.Version[1] = 0
		tag.Comment = string(data[97:127])
		tag.Album = ""
	}
	tag.Genre = parseGenre(data[126:127])

	result, err := json.Marshal(tag)
	if err != nil {
		panic(fmt.Sprintf("Unable to convert tag to JSON data."))
	}

	return result;
}

//Takes JSON data and returns and ID3v1 variable
func unparseV1(data []byte) ID3v1 {
	var tag ID3v1
	err := json.Unmarshal(data, tag)
	if err != nil {
		panic(fmt.Sprintf("Unable to convert JSON data to tag."))
	}

	return tag
}
