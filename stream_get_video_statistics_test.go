package bunny_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/simplesurance/bunny-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStreamGetVideoStatisticsCRUD(t *testing.T) {
	// Stream Video Has Separate API Key than Account!
	// https://docs.bunny.net/reference/api-overview
	clt := bunny.NewClient("", "stream")

	// Personal values, use your own
	const videoLibraryId int64 = 179632
	videoId := "e8117722-7af4-4723-aacf-36fb5f81496d"

	// Get Video Statistics
	videoOptions := bunny.StreamVideoStatisticsOptions{
		VideoLibraryId: videoLibraryId,
		DateFrom:       nil,
		DateTo:         nil,
		Hourly:         false,
		VideoGuid:      videoId,
	}

	result, err := clt.Stream.GetVideoStatistics(context.Background(), &videoOptions)
	fmt.Print(result)
	require.NoError(t, err, "video get failed")
	assert.NotNil(t, result)
}
