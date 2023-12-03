package bunny

import (
	"context"
	"fmt"
)

// StreamVideoStatistics represents the response of the Stream Statistics of the Stream Manage Videos API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/video_getvideostatistics
type StreamVideoStatistics struct {
	ViewsChart        []any `json:"viewsChart,omitempty"`
	WatchTimeChart    []any `json:"watchTimeChart,omitempty"`
	CountryViewCounts []any `json:"countryViewCounts,omitempty"`
	CountryWatchTime  []any `json:"countryWatchTime,omitempty"`
	EngagementScore   int32 `json:"engagementScore,omitempty"`
}

// StreamVideoStatisticsOptions represents the options for Stream Video
//
// Bunny.net API docs: https://docs.bunny.net/reference/video_getvideostatistics
type StreamVideoStatisticsOptions struct {
	VideoLibraryId int64  `json:"libraryId,omitempty"`
	DateFrom       any    `json:"dateFrom,omitempty"`
	DateTo         any    `json:"dateTo,omitempty"`
	Hourly         bool   `json:"hourly,omitempty"`
	VideoGuid      string `json:"videoGuid,omitempty"`
}

// Gets the Video Statistics with the given Library Id and Video Id
//
// Bunny.net API docs: https://docs.bunny.net/reference/video_getvideostatistics
func (s *StreamService) GetVideoStatistics(ctx context.Context, opts *StreamVideoStatisticsOptions) (*StreamVideoStatistics, error) {
	path := fmt.Sprintf("library/%d/statistics", opts.VideoLibraryId)
	return resourceGet[StreamVideoStatistics](ctx, s.client, path, nil)
}
