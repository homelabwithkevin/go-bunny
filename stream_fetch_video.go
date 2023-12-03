package bunny

import (
	"context"
	"fmt"
)

// StreamCreateVideoOptions are the request parameters for the Stream Create Video API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/video_fetchnewvideo
type StreamFetchVideoOptions struct {
	VideoLibraryId int64  `json:"VideoLibraryId,omitempty"`
	CollectionId   string `json:"CollectionId,omitempty"`
	ThumbnailTime  int32  `json:"ThumbnailTime,omitempty"`
	Url            string `json:"Url,omitempty"`
	Title          string `json:"Title,omitempty"`
}

type StreamVideoFetch struct {
	Success    bool   `json:"success,omitempty"`
	Message    string `json:"message,omitempty"`
	StatusCode int32  `json:"statusCode,omitempty"`
}

// Create creates a new Video
//
// Bunny.net API docs: https://docs.bunny.net/reference/video_fetchnewvideo
func (s *StreamService) Fetch(ctx context.Context, opts *StreamFetchVideoOptions) (*StreamVideoFetch, error) {
	videoLibraryId := fmt.Sprintf("/library/%d/videos/fetch", opts.VideoLibraryId)

	return resourcePostWithResponse[StreamVideoFetch](
		ctx,
		s.client,
		videoLibraryId,
		opts,
	)
}
