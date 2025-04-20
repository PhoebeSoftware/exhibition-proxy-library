package igdb

import (
	"fmt"
)

func (metaData *Metadata) GetArtworkURLS() []string {
	var artworkUrlList []string
	for _, image := range metaData.Artworks {
		artworkUrlList = append(
			artworkUrlList,
			fmt.Sprintf(
				"https://images.igdb.com/igdb/image/upload/t_1080p/%s.jpg",
				image.ImageID,
			),
		)
	}
	return artworkUrlList
}