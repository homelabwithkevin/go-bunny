package bunny

import (
	"context"
	"fmt"
)

// StreamCreateVideoOptions are the request parameters for the Stream Create Video API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/video_deletevide
type StreamDeleteVideoOptions struct {
	VideoLibraryId int64  `json:"VideoLibraryId,omitempty"`
	VideoId        string `json:"VideoId,omitempty"`
}

// Create creates a new Video
//
// Bunny.net API docs: https://docs.bunny.net/reference/video_deletevideo
func (s *StreamService) Delete(ctx context.Context, opts *StreamDeleteVideoOptions) error {
	videoIdUrl := fmt.Sprintf("/library/%d/videos/%s", opts.VideoLibraryId, opts.VideoId)

	return resourceDelete(ctx, s.client, videoIdUrl, opts)
}