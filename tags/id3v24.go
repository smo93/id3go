package id3go

import (
	"bufio"
)

func parseV24Size(reader *bufio.Reader) int {
	return parseSize(readBytes(reader, 4))
}

func parseV24Tag(reader *bufio.Reader, tag *ID3) {
	for hasFrame(reader, 4) {
		parseV24Frame(reader, tag)
	}
}

func parseV24Frame(reader *bufio.Reader, tag *ID3) {
	id := string(readB(reader, 4))
	size := parseV24Size(reader)
	flags := readB(reader, 2)
	
	switch id {
	case "TRCK":
		tag.Track.FrameID = id
		tag.Track.Size = size
		tag.Track.Data = parseString(reader, size)
	case "TPE1":
		tag.Artist.FrameID = id
		tag.Artist.Size = size
		tag.Artist.Data = parseString(reader, size)
	case "TIT2":
		tag.Title.FrameID = id
		tag.Title.Size = size
		tag.Title.Data = parseString(reader, size)
	case "TPOS":
		tag.Disk.FrameID = id
		tag.Disk.Size = size
		tag.Disk.Data = parseString(reader, size)
	case "TALB":
		tag.Album.FrameID = id
		tag.Album.Size = size
		tag.Album.Data = parseString(reader, size)
	case "TLEN":
		tag.Length.FrameID = id
		tag.Length.Size = size
		tag.Length.Data = parseString(reader, size)
	case "TDRC":
		tag.Year.FrameID = id
		tag.Year.Size = size
		tag.Year.Data = parseString(reader, size)
	case "TCON":
		tag.Genre.FrameID = id
		tag.Genre.Size = size
		tag.Genre.Data = parseString(reader, size)
	default:
		skipB(reader, size)
		}
}
