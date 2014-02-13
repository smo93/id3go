# ID3Go
ID3Go is a small open-source Go package for reading ID3 tags in mp3 files.

### Setup
To use the package you'll need a working `Go v1.1.x` or later installation.  
To get the package run the following command:  
`go get github.com/smo93/id3go`

### How to use
To use the package in your code, you need to import it first. Add the following at the begining of your code: `import "github.com/smo93/id3go"`  

#### Getting the tag from a mp3 file  
To get the metadata of a mp3 file you will use the `ParseTag` function. It takes one argument of type `string`, which is the path to the mp3 file. The function returns a byte slice, containing the tag object, represented as JSON data.  
So the result of parsing a file containing v1 ID3 tag would be something like this (printed as `string`):  
```
jsonData := id3go.ParseTag("test_100.mp3")
fmt.Println(string(jsonData))

// Output:
// {"Version":[1,0],"SongTitle":"TITLE1234567890123456789012345","Artist":"ARTIST123456789012345678901234","Album":"","Year":"2001","Comment":"COMMENT123456789012345678901\u0000\u0001","AlbumTrack":0,"Genre":"Classic Rock"}
```

### TODO
- Edit and write tag to file functionality
- Recognition of all frames (currently only the text and url frames are recognized)

### LICENSE
The package is licensed under the [MIT License](LICENSE)