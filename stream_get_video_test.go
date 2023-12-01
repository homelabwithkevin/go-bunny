package bunny_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/simplesurance/bunny-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStreamGetVideoCRUD(t *testing.T) {
	clt := bunny.NewClient("", "stream")

	const videoLibraryId int64 = 179632
	videoId := "e8117722-7af4-4723-aacf-36fb5f81496d"

	videoInformation, err := clt.Stream.Get(context.Background(), videoLibraryId, videoId)
	fmt.Print(videoInformation)
	require.NoError(t, err, "video get failed")
	assert.NotNil(t, videoInformation)
}
