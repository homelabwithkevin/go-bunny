package bunny

import (
	"context"
	"fmt"
)

// StreamCreateVideoOptions are the request parameters for the Stream Create Video API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/video_createvideo
type StreamCreateVideoOptions struct {
	VideoLibraryId int64  `json:"VideoLibraryId,omitempty"`
	Title          string `json:"Title,omitempty"`
	CollectionId   string `json:"CollectionId,omitempty"`
	ThumbnailTime  *int32 `json:"ThumbnailTime,omitempty"`
}

// Create creates a new Video
//
// Bunny.net API docs: https://docs.bunny.net/reference/video_createvideo
func (s *StreamService) Create(ctx context.Context, opts *StreamCreateVideoOptions) (*StreamVideo, error) {
	videoLibraryId := fmt.Sprintf("/library/%d/videos", opts.VideoLibraryId)

	return resourcePostWithResponse[StreamVideo](
		ctx,
		s.client,
		videoLibraryId,
		opts,
	)
}
