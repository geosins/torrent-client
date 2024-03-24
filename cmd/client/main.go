package main

import (
	"fmt"

	"torrent/pkg/torrentFile"
)

func main() {
	// file := torrentFile.Read("../../singlefile_example.torrent")
	// file := torrentFile.Read("../../multyfile_example.torrent")
	// file := torrentFile.Read("../../video_example.torrent")
	// file := torrentFile.Read("../../structured_example.torrent")
	file := torrentFile.Read("../../other.torrent")
	fmt.Println(file)
}
