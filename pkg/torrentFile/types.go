package torrentFile

import (
	"fmt"
)

type TorrentFile struct {
    announce string `example:"udp://opentor.net:6969, http://bt2.t-ru.org/ann?pk=82ed09ca70ed0a026f5d31dda25cd0e0, http://bt02.nnm-club.cc:2710/announce"`
    announceList []string `example:"[udp://opentor.net:6969, http://bt2.t-ru.org/ann?pk=82ed09ca70ed0a026f5d31dda25cd0e0]"`
    comment string `example:"http://rutor.is/torrent/964171, https://rutracker.org/forum/viewtopic.php?t=4321272, http://nnmclub.to/forum/viewtopic.php?p=1862558"`
    creationDate uint `example:"1278018624"`
    createdBy string `example:"uTorrent/1.8.2, BitTorrent/7730"`
    encoding string `example:"UTF-8"`
    info TorrentFileInfo
}

type TorrentFileInfo struct {
    name string
    pieceLength uint
    length int64 // empty in multi mode
    pieces []byte
    files []TorrentFileInfoFile  // empty in single mode
}

type TorrentFileInfoFile struct {
    length int64
    path []string // For example, a the file "dir1/dir2/file.ext" would consist string array: ["dir1", "dir2", "file.ext"]
}

func (file *TorrentFile) String() string {
    return fmt.Sprintf(
        "{\n\tannounce: %v.\n\tannounceList: %v,\n\tcomment: %v,\n\tcreationDate: %v,\n\tcreatedBy: %v,\n\tencoding: %v,\n\tinfo: " +
        "{\n\t\tname: %v,\n\t\tpieceLength: %v,\n\t\tlength: %v,\n\t\tpieces: ...,\n\t\tfiles: %v\n\t}\n}",
        file.announce,
        file.announceList,
        file.comment,
        file.creationDate,
        file.createdBy,
        file.encoding,
        file.info.name,
        file.info.pieceLength,
        file.info.length,
        file.info.files,
    )
}

func New(bencode interface{}) *TorrentFile {
    defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

    return newTorrentFile(bencode)
}

func newTorrentFile(bencode interface{}) *TorrentFile {
    torrentFile := TorrentFile{}
    
    for key, value := range bencode.(map[string]interface{}) {
        switch key {
        case "announce":
            torrentFile.announce = string(value.([]byte))
        case "comment": 
            torrentFile.comment = string(value.([]byte))
        case "created by":
            torrentFile.createdBy = string(value.([]byte))
        case "creation date":
            torrentFile.creationDate = uint(value.(int64))
        case "encoding":
            torrentFile.encoding = string(value.([]byte))
        case "announce-list":
            for _, arr := range value.([]interface{}) {
                url := arr.([]interface{})[0]
                torrentFile.announceList = append(torrentFile.announceList, string(url.([]byte)))
            }
        case "info":
            for infoKey, infoValue := range value.(map[string]interface{}) {
                switch infoKey {
                    case "name":
                        torrentFile.info.name = string(infoValue.([]byte))
                    case "piece length":
                        torrentFile.info.pieceLength = uint(infoValue.(int64))
                    case "length":
                        torrentFile.info.length = infoValue.(int64)
                    case "pieces":
                        torrentFile.info.pieces = infoValue.([]byte)
                    case "files":
                        for _, infoFile := range infoValue.([]interface{}) {
                            file := newTorrentFileInfoFile(infoFile)
                            torrentFile.info.files = append(torrentFile.info.files, *file)
                        }
                }
            }
        }
    }

    return &torrentFile
}

func newTorrentFileInfoFile(infoFile interface{}) *TorrentFileInfoFile {
    file := TorrentFileInfoFile{}

    for key, value := range infoFile.(map[string]interface{}) {
        switch key {
        case "length":
            file.length = value.(int64)
        case "path":
            for _, pathPart := range value.([]interface{}) {
                file.path = append(file.path, string(pathPart.([]byte)))
            }
        }
    }

    return &file
}
