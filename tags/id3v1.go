package id3go

import (
	"fmt"
	"strings"
)

type ID3v1 struct {
	ID3
	title    	string
	artist   	string
	album  	    string
	year        int
	comments    string
	albumtTrack byte
	genre       byte
}

func NewV1Tag(tagBytes []byte) ID3v1 {
	
}

func (tag ID3v1) WriteToFile {
	
}

func (tag ID3v1) GetTitle() string {
	return tag.title
}

func (tag ID3v1) GetArtist() string {
	return tag.artist
}

func (tag ID3v1) GetAlbum() string {
	return tag.album
}

func (tag ID3v1) GetYear() int {
	return tag.year
}

func (tag ID3v1) GetComments() string {
	return tag.comments
}

func (tag ID3v1) GetAlbumTrack() byte {
	return tag.albumTrack
}

func (tag ID3v1) GetGenre() byte {
	return tag.genre
} 

func (tag ID3v1) SetTitle(title string) {
	tag.title = title[:30]
}

func (tag ID3v1) SetArtist(artist string) {
	tag.artist = artist[:30]
}

func (tag ID3v1) SetAlbum(album string) {
	tag.album = album[:30]
}

func (tag ID3v1) SetYear(year int) {
	if year > 9999 {
		tag.year = 0
	} else {
		tag.year = year
	}
}

func (tag ID3v1) SetComments(comments string) {
	if tag.version[1] == 0 {
		tag.comments = comments[:30]
	} else {
		tag.comments = comments[:28]
	}
}

func (tag ID3v1) SetAlbumTrack(albumTrack byte) {
	tag.albumTrack = albumTrack
}

func (tag ID3v1) SetGenre(genre byte) {
	tag.genre = genre
}

func (tag ID3v1) ConvertToV2() ID3v2 {
	
}
