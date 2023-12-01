package bunny

import (
	"context"
	"fmt"
)

// StreamVideos represents the response of the List Videos of the Stream Manage Videos API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/video_list
type StreamVideos PaginationReply[StreamVideo]

// StreamVideoListListOpts represents both PaginationOptions and the other optional query parameters of the List endpoint.
type StreamVideoListListOpts struct {
	VideoLibraryGetOpts
	PaginationOptions
}

// List retrieves the Videos in the Stream's Library Id
// if opts.Page or or opts.PerPage is < 1, the related DefaultPagination values are used.e Stream Library, by Library Id
//
// Bunny.net API docs: https://docs.bunny.net/reference/video_list
func (s *StreamService) List(ctx context.Context, videoLibraryId int64, opts *StreamVideoListListOpts) (*StreamVideos, error) {
	path := fmt.Sprintf("/library/%d/videos", videoLibraryId)
	var res StreamVideos

	if opts == nil {
		opts = &StreamVideoListListOpts{
			PaginationOptions: PaginationOptions{
				Page:    DefaultPaginationPage,
				PerPage: DefaultPaginationPerPage,
			},
		}
	} else {
		opts.ensureConstraints()
	}

	req, err := s.client.newGetRequest(path, opts)
	if err != nil {
		return nil, err
	}

	if err := s.client.sendRequest(ctx, req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
