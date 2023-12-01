package bunny_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/simplesurance/bunny-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStreamListVideosCRUD(t *testing.T) {
	// Stream Video Has Separate API Key than Account!
	// https://docs.bunny.net/reference/api-overview
	clt := bunny.NewClient("", "stream")

	// Personal values, use your own
	const videoLibraryId int64 = 179632

	listVideos, err := clt.Stream.List(context.Background(), videoLibraryId, nil)
	fmt.Print(listVideos)
	require.NoError(t, err, "video get failed")
	assert.NotNil(t, listVideos)

	// Assumes you have one video.
	assert.Equal(
		t,
		listVideos.Items[0].VideoLibraryId,
		videoLibraryId,
	)

	actuallyListVideos := false
	if actuallyListVideos {
		for _, videos := range listVideos.Items {
			fmt.Print(videos)
		}
	}
}
