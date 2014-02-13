package id3go

import (
	"testing"
)

type fileTest struct {
	path string
	tag string
}

func testFile(t *testing.T, expected fileTest) {
	actual := string(ParseTag(expected.path))

	if actual != expected.tag {
		t.Errorf("JSON data: expected %s\ngot %s", expected.tag, actual)
	}
}

func TestEmpty(t *testing.T) {
	testFile(t, fileTest{"test_empty.mp3", ""})
}

func TestV100(t *testing.T) {
	testFile(t, fileTest{"test_100.mp3", "{\"Version\":[1,0],\"SongTitle\":\"TITLE1234567890123456789012345\",\"Artist\":\"ARTIST123456789012345678901234\",\"Album\":\"\",\"Year\":\"2001\",\"Comment\":\"COMMENT123456789012345678901\\u0000\\u0001\",\"AlbumTrack\":0,\"Genre\":\"Classic Rock\"}"})
}

func TestV220(t *testing.T) {
	testFile(t, fileTest{"test_220.mp3", "{\"Header\":{\"Version\":[2,0],\"Unsynchronisation\":false,\"Compression\":false,\"Size\":181},\"Frames\":[{\"ID\":\"TT2\",\"Size\":12,\"TextEncoding\":0,\"Data\":\"There There\"},{\"ID\":\"TP1\",\"Size\":10,\"TextEncoding\":0,\"Data\":\"Radiohead\"},{\"ID\":\"TYE\",\"Size\":5,\"TextEncoding\":0,\"Data\":\"2003\"},{\"ID\":\"TCO\",\"Size\":5,\"TextEncoding\":0,\"Data\":\"Alternative\"}]}"})
}

func TestV230(t *testing.T) {
	testFile(t, fileTest{"test_230.mp3", "{\"Header\":{\"Version\":[3,0],\"Unsynchronisation\":false,\"ExtendedHeader\":false,\"ExperimentalIndicator\":false,\"Size\":189},\"Frames\":[{\"ID\":\"TIT2\",\"Size\":30,\"TagAlterPreserv\":false,\"FileAlterPreserv\":false,\"ReadOnly\":false,\"Compression\":false,\"Encryption\":false,\"GroupingIdentity\":false,\"TextEncoding\":0,\"Data\":\"Everything In Its Right Place\"},{\"ID\":\"TPE1\",\"Size\":10,\"TagAlterPreserv\":false,\"FileAlterPreserv\":false,\"ReadOnly\":false,\"Compression\":false,\"Encryption\":false,\"GroupingIdentity\":false,\"TextEncoding\":0,\"Data\":\"Radiohead\"},{\"ID\":\"TALB\",\"Size\":6,\"TagAlterPreserv\":false,\"FileAlterPreserv\":false,\"ReadOnly\":false,\"Compression\":false,\"Encryption\":false,\"GroupingIdentity\":false,\"TextEncoding\":0,\"Data\":\"Kid A\"},{\"ID\":\"TYER\",\"Size\":5,\"TagAlterPreserv\":false,\"FileAlterPreserv\":false,\"ReadOnly\":false,\"Compression\":false,\"Encryption\":false,\"GroupingIdentity\":false,\"TextEncoding\":0,\"Data\":\"2000\"},{\"ID\":\"TCON\",\"Size\":12,\"TagAlterPreserv\":false,\"FileAlterPreserv\":false,\"ReadOnly\":false,\"Compression\":false,\"Encryption\":false,\"GroupingIdentity\":false,\"TextEncoding\":0,\"Data\":\"Alternative\"},{\"ID\":\"TRCK\",\"Size\":2,\"TagAlterPreserv\":false,\"FileAlterPreserv\":false,\"ReadOnly\":false,\"Compression\":false,\"Encryption\":false,\"GroupingIdentity\":false,\"TextEncoding\":0,\"Data\":\"1\"}]}"})
}

func TestV240(t *testing.T) {
	testFile(t, fileTest{"test_240.mp3", "{\"Header\":{\"Version\":[4,0],\"Unsynchronisation\":false,\"ExtendedHeader\":false,\"ExperimentalInd\":false,\"FooterPresent\":false,\"Size\":6},\"Frames\":[{\"ID\":\"TDRC\",\"Size\":6,\"TagAlterPreservation\":false,\"FileAlterPreservation\":false,\"ReadOnly\":false,\"GroupingIdentity\":false,\"Compression\":false,\"Encryption\":false,\"Unsynchronisation\":false,\"DataLengthIdicator\":false,\"TextEncoding\":0,\"Data\":\"2011\\u0000\"},{\"ID\":\"TRCK\",\"Size\":7,\"TagAlterPreservation\":false,\"FileAlterPreservation\":false,\"ReadOnly\":false,\"GroupingIdentity\":false,\"Compression\":false,\"Encryption\":false,\"Unsynchronisation\":false,\"DataLengthIdicator\":false,\"TextEncoding\":0,\"Data\":\"07/08\\u0000\"},{\"ID\":\"TPOS\",\"Size\":5,\"TagAlterPreservation\":false,\"FileAlterPreservation\":false,\"ReadOnly\":false,\"GroupingIdentity\":false,\"Compression\":false,\"Encryption\":false,\"Unsynchronisation\":false,\"DataLengthIdicator\":false,\"TextEncoding\":0,\"Data\":\"1/1\\u0000\"},{\"ID\":\"TPE1\",\"Size\":10,\"TagAlterPreservation\":false,\"FileAlterPreservation\":false,\"ReadOnly\":false,\"GroupingIdentity\":false,\"Compression\":false,\"Encryption\":false,\"Unsynchronisation\":false,\"DataLengthIdicator\":false,\"TextEncoding\":0,\"Data\":\"Radiohead\"},{\"ID\":\"TALB\",\"Size\":18,\"TagAlterPreservation\":false,\"FileAlterPreservation\":false,\"ReadOnly\":false,\"GroupingIdentity\":false,\"Compression\":false,\"Encryption\":false,\"Unsynchronisation\":false,\"DataLengthIdicator\":false,\"TextEncoding\":0,\"Data\":\"The King Of Limbs\"},{\"ID\":\"TPE2\",\"Size\":10,\"TagAlterPreservation\":false,\"FileAlterPreservation\":false,\"ReadOnly\":false,\"GroupingIdentity\":false,\"Compression\":false,\"Encryption\":false,\"Unsynchronisation\":false,\"DataLengthIdicator\":false,\"TextEncoding\":0,\"Data\":\"Radiohead\"},{\"ID\":\"TIT2\",\"Size\":18,\"TagAlterPreservation\":false,\"FileAlterPreservation\":false,\"ReadOnly\":false,\"GroupingIdentity\":false,\"Compression\":false,\"Encryption\":false,\"Unsynchronisation\":false,\"DataLengthIdicator\":false,\"TextEncoding\":0,\"Data\":\"Give Up The Ghost\"},{\"ID\":\"TSRC\",\"Size\":13,\"TagAlterPreservation\":false,\"FileAlterPreservation\":false,\"ReadOnly\":false,\"GroupingIdentity\":false,\"Compression\":false,\"Encryption\":false,\"Unsynchronisation\":false,\"DataLengthIdicator\":false,\"TextEncoding\":0,\"Data\":\"GBU4B1100009\"},{\"ID\":\"TCON\",\"Size\":12,\"TagAlterPreservation\":false,\"FileAlterPreservation\":false,\"ReadOnly\":false,\"GroupingIdentity\":false,\"Compression\":false,\"Encryption\":false,\"Unsynchronisation\":false,\"DataLengthIdicator\":false,\"TextEncoding\":0,\"Data\":\"Alternative\"}],\"Footer\":null}"})
}

