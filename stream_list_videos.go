package bunny

import (
	"context"
	"fmt"
)

// StreamVideolist represents the response of the List Videos of the Stream Manage Videos API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/video_list

type StreamVideoList struct {
	TotalItems   int64         `json:"totalItems,omitempty"`
	CurrentPage  int64         `json:"currentPage,omitempty"`
	ItemsPerPage int32         `json:"itemsPerPage,omitempty"`
	Items        []StreamVideo `json:"items,omitempty"`
}

// Lists the Videos of the given Library Id
//
// Bunny.net API docs: https://docs.bunny.net/reference/video_list
func (s *StreamService) List(ctx context.Context, videoLibraryId int64) (*StreamVideo, error) {
	options := fmt.Sprintf("?page=%d&itemsPerPage=%d&orderBy=date", DefaultPaginationPage, DefaultPaginationPerPage)
	path := fmt.Sprintf("library/%d/videos%s", videoLibraryId, options)

	return resourceGet[StreamVideo](ctx, s.client, path, nil)
}
