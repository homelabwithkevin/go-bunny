package bunny

// StreamService communicates with the /{libraryId}/videos/{videoId} API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/video_getvideo
type StreamService struct {
	client *Client
}