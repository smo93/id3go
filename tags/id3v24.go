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
		parseV24FrameFlags(flags, tag.Track)
		tag.Track.Data = parseString(reader, size)
	case "TPE1":
		tag.Artist.FrameID = id
		tag.Artist.Size = size
		parseV24FrameFlags(flags, tag.Artist)
		tag.Artist.Data = parseString(reader, size)
	case "TIT2":
		tag.Title.FrameID = id
		tag.Title.Size = size
		parseV24FrameFlags(flags, tag.Title)
		tag.Title.Data = parseString(reader, size)
	case "TPOS":
		tag.Disk.FrameID = id
		tag.Disk.Size = size
		parseV24FrameFlags(flags, tag.Disk)
		tag.Disk.Data = parseString(reader, size)
	case "TALB":
		tag.Album.FrameID = id
		tag.Album.Size = size
		parseV24FrameFlags(flags, tag.Album)
		tag.Album.Data = parseString(reader, size)
	case "TLEN":
		tag.Length.FrameID = id
		tag.Length.Size = size
		parseV24FrameFlags(flags, tag.Length)
		tag.Length.Data = parseString(reader, size)
	case "TDRC":
		tag.Year.FrameID = id
		tag.Year.Size = size
		parseV24FrameFlags(flags, tag.Year)
		tag.Year.Data = parseString(reader, size)
	case "TCON":
		tag.Genre.FrameID = id
		tag.Genre.Size = size
		parseV24FrameFlags(flags, tag.Genre)
		tag.Genre.Data = parseString(reader, size)
	default:
		skipB(reader, size)
		}
}

func parseV24FrameFlags(data []byte, frame *ID3v2Frame) {
	frame.TagAlterPreservation = data[0] & (1 << 6) != 0
	frame.FileAlterPreservation = data[0] & (1 << 5) != 0
	frame.ReadOnly = data[0] & (1 << 4) != 0
	frame.GroupingIdentity = data[1] & (1 << 6) != 0
	frame.Compression = data[1] & (1 << 3) != 0
	frame.Encryption = data[1] & (1 << 2) != 0
	frame.Unsynchronisation = data[1] & (1 << 1) != 0
	frame.DataLengthIndicator = data[1] & 1 != 0
}

