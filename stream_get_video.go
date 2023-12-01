package bunny

import (
	"context"
	"fmt"
)

// Video represents the response of the the List and Get Stream Manage Videos API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/video_getvideo
type Captions struct {
	Srclang string `json:"srclang,omitempty"`
	Label   string `json:"label,omitempty"`
}

type StreamVideo struct {
	VideoLibraryId       int64      `json:"videoLibraryId,omitempty"`
	Guid                 string     `json:"guid,omitempty"`
	Title                string     `json:"title,omitempty"`
	DateUploaded         string     `json:"dateUploaded,omitempty"`
	Views                int64      `json:"views,omitempty"`
	IsPublic             bool       `json:"isPublic,omitempty"`
	Length               int32      `json:"length,omitempty"`
	Status               int        `json:"status,omitempty"`
	Framerate            int        `json:"Framerate,omitempty"`
	Rotation             int        `json:"rotation,omitempty"`
	Width                int32      `json:"width,omitempty"`
	Height               int32      `json:"height,omitempty"`
	AvailableResolutions string     `json:"availableResolutions,omitempty"`
	ThumbnailCount       int32      `json:"thumbnailCount,omitempty"`
	EncodeProgress       int32      `json:"encodeProgress,omitempty"`
	StorageSize          int64      `json:"storageSize,omitempty"`
	Captions             []Captions `json:"captions,omitempty"`
	HasMP4Fallback       bool       `json:"hasMP4Fallback,omitempty"`
	CollectionId         string     `json:"collectionId,omitempty"`
	ThumbnailFileName    string     `json:"thumbnailFileName,omitempty"`
	AverageWatchTime     int64      `json:"averageWatchTime,omitempty"`
	TotalWatchTime       int64      `json:"totalWatchTime,omitempty"`
	Category             string     `json:"category,omitempty"`
	Chapters             string     `json:"chapters,omitempty"`
	Moments              string     `json:"moments,omitempty"`
	MetaTags             string     `json:"metaTags,omitempty"`
	TranscodingMessages  string     `json:"transcodingMessages,omitempty"`
}

// Get retrieves the Storage Zone with the given id.
//
// Bunny.net API docs: https://docs.bunny.net/reference/storagezonepublic_index2
func (s *StorageZoneService) Get(ctx context.Context, id int64) (*StorageZone, error) {
	path := fmt.Sprintf("storagezone/%d", id)
	return resourceGet[StorageZone](ctx, s.client, path, nil)
}
