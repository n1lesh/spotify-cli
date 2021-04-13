package pause

import (
	"spotify/pkg"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPauseCommand(t *testing.T) {
	api := new(pkg.MockSpotifyAPI)
	api.On("Pause").Return(nil)

	err := pause(api)
	require.NoError(t, err)
}
