package torrentFile

import (
	"io"
	"os"
)

func Read(filePath string) *TorrentFile {
    content, err := readFile(filePath)

    if (err != nil) {
        panic("Can't read file")
    }

    object := ParseBencod(content)

    torrentFile := New(object)

    if (torrentFile == nil) {
        panic("Wrong torrent file format")
    }

    return torrentFile
}

func readFile(filePath string) ([]byte, error) {
    var content []byte

    file, err := os.Open(filePath)
    if (err != nil) {
        return content, err
    }

    content, err = io.ReadAll(file)
    if (err != nil) {
        return content, err
    }

    file.Close()

    return content, nil
}
