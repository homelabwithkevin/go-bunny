package bunny_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/simplesurance/bunny-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStreamDeleteVideoCRUD(t *testing.T) {
	// Stream Video Has Separate API Key than Account!
	// https://docs.bunny.net/reference/api-overview
	clt := bunny.NewClient("", "stream")

	// Personal values, use your own
	const videoLibraryId int64 = 179632
	videoId = "1234"

	videoOptions := bunny.StreamDeleteVideoOptions{
		VideoLibraryId: videoLibraryId,
		VideoId: videoId,
	}

	result, err := clt.Stream.Delete(context.Background(), &videoOptions)
	fmt.Print(result)
	require.NoError(t, err, "video delete failed")
}