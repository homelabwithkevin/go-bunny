package bunny_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/simplesurance/bunny-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStreamCreateVideoCRUD(t *testing.T) {
	// Stream Video Has Separate API Key than Account!
	// https://docs.bunny.net/reference/api-overview
	clt := bunny.NewClient("", "stream")

	// Personal values, use your own
	const videoLibraryId int64 = 179632

	videoOptions := bunny.StreamCreateVideoOptions{
		VideoLibraryId: videoLibraryId,
		Title:          "golang-testing",
		CollectionId:   "",
		ThumbnailTime:  nil,
	}

	result, err := clt.Stream.Create(context.Background(), &videoOptions)
	fmt.Print(result)
	fmt.Printf("Guid: %s", result.Guid)
	require.NoError(t, err, "video create failed")
	assert.Equal(
		t,
		result.VideoLibraryId,
		videoLibraryId,
		"library id is incorrect",
	)
}
