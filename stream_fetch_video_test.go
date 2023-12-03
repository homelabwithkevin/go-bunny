package bunny_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/simplesurance/bunny-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStreamFetchVideoCRUD(t *testing.T) {
	// Stream Video Has Separate API Key than Account!
	// https://docs.bunny.net/reference/api-overview
	clt := bunny.NewClient("", "stream")

	// Personal values, use your own
	videoOptions := bunny.StreamFetchVideoOptions{
		VideoLibraryId: 179632,
		CollectionId:   "",
		ThumbnailTime:  0,
		Url:            "https://download.blender.org/peach/bigbuckbunny_movies/BigBuckBunny_320x180.mp4",
		Title:          "fetch-testing",
	}

	result, err := clt.Stream.Fetch(context.Background(), &videoOptions)
	fmt.Print(result)
	require.NoError(t, err, "video create failed")
	assert.Equal(
		t,
		result.Success,
		true,
		"fetch was not successful",
	)
}
