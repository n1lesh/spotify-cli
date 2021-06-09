package next

import (
	"spotify/internal"
	"testing"

	"github.com/brianstrauch/spotify"
	"github.com/brianstrauch/spotify/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNextCommand(t *testing.T) {
	api := new(spotify.MockAPI)

	playback1 := &model.Playback{
		IsPlaying:  true,
		ProgressMs: 0,
		Item: model.Item{
			ID:   "0",
			Type: "track",
			Name: "Song",
			Artists: []model.Artist{
				{Name: "Artist"},
			},
			DurationMs: 1000,
		},
	}

	playback2 := new(model.Playback)
	*playback2 = *playback1
	playback2.Item.ID = "1"

	api.On("Status").Return(playback1, nil)
	api.On("WaitForUpdatedPlayback", mock.AnythingOfType("func(*model.Playback) bool")).Return(playback2, nil)
	api.On("Next").Return(nil)

	status, err := next(api)
	require.Equal(t, "🎵 Song\n🎤 Artist\n▶️  0:00 [                ] 0:01\n", status)
	require.NoError(t, err)
}

func TestNoActiveDeviceErr(t *testing.T) {
	api := new(spotify.MockAPI)
	api.On("Status").Return(nil, nil)

	_, err := next(api)
	require.Equal(t, internal.NoActiveDeviceErr, err.Error())
}
