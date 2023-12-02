package bunny_test

import (
	"context"
	"testing"

	bunny "github.com/simplesurance/bunny-go"
	"github.com/stretchr/testify/require"
)

func TestStreamDeleteVideoCRUD(t *testing.T) {
	// Stream Video Has Separate API Key than Account!
	// https://docs.bunny.net/reference/api-overview
	clt := bunny.NewClient("", "stream")

	// Personal values, use your own
	const videoLibraryId int64 = 179632

	// Create Video
	createVideoOptions := bunny.StreamCreateVideoOptions{
		VideoLibraryId: videoLibraryId,
		Title:          "streamvideo",
		CollectionId:   "",
		ThumbnailTime:  nil,
	}

	resultCreateVideo, _ := clt.Stream.Create(context.Background(), &createVideoOptions)
	videoId := resultCreateVideo.Guid

	// Delete Video
	videoOptions := bunny.StreamDeleteVideoOptions{
		VideoLibraryId: videoLibraryId,
		VideoId:        videoId,
	}

	err := clt.Stream.Delete(context.Background(), &videoOptions)
	require.NoError(t, err, "video delete failed")
}
