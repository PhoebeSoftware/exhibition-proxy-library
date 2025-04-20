package igdb

import "fmt"

func (metaData *Metadata) GetCoverURL() string {
	return fmt.Sprintf("https://images.igdb.com/igdb/image/upload/t_cover_big/%s.jpg", metaData.Cover.ImageID)
}
