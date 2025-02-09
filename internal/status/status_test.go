package status

import (
	"spotify/internal"
	"testing"

	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
)

func TestStatus_Track(t *testing.T) {
	api := new(internal.MockAPI)

	playback := &spotify.Playback{
		IsPlaying:  true,
		ProgressMs: 0,
		Item: spotify.Item{
			Type: "track",
			Name: "Track",
			Artists: []spotify.Artist{
				{Name: "Artist"},
			},
			DurationMs: 1000,
		},
	}

	api.On("GetPlayback").Return(playback, nil)

	status, err := status(api)
	require.NoError(t, err)
	require.Equal(t, "   Track\r🎵\n   Artist\r🎤\n   0:00 [                ] 0:01\r▶️\n", status)
}

func TestStatus_Podcast(t *testing.T) {
	api := new(internal.MockAPI)

	playback := &spotify.Playback{
		IsPlaying:  true,
		ProgressMs: 0,
		Item: spotify.Item{
			Type: "episode",
			Name: "Episode",
			Show: spotify.Show{
				Name: "Podcast",
			},
			DurationMs: 1000,
		},
	}

	api.On("GetPlayback").Return(playback, nil)

	status, err := status(api)
	require.NoError(t, err)
	require.Equal(t, "   Episode\r🎵\n   Podcast\r🎤\n   0:00 [                ] 0:01\r▶️\n", status)
}

func TestStatus_ErrNoActiveDevice(t *testing.T) {
	api := new(internal.MockAPI)
	api.On("GetPlayback").Return(nil, nil)

	_, err := status(api)
	require.Error(t, err)
	require.Equal(t, internal.ErrNoActiveDevice, err.Error())
}

func TestJoinArtists(t *testing.T) {
	artists := []spotify.Artist{
		{Name: "Artist 1"},
		{Name: "Artist 2"},
	}
	require.Equal(t, "Artist 1, Artist 2", joinArtists(artists))
}

func TestFormatTime_OneMinute(t *testing.T) {
	require.Equal(t, "1:00", formatTime(60*1000))
}

func TestFormatTime_OneHour(t *testing.T) {
	require.Equal(t, "1:00:00", formatTime(60*60*1000))
}
