package fiscrap

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"path"
	"testing"
)

func TestScrapper_parsePodcastPage(t *testing.T) {
	fm := &fetchMock{}
	dom := fm.On("Do", mock.Anything)
	dom.Run(func(args mock.Arguments) {
		f, err := os.Open(path.Join("testdata/jeanne-barret.html"))
		dom.ReturnArguments = mock.Arguments{io.ReadCloser(f), err}
	})
	s := Scrapper{
		fetch: fm,
	}
	a, b, err := s.parsePodcastPage("https://www.franceinter.fr/emissions/les-odyssees/jeanne-barret")
	require.NoError(t, err)
	assert.Equal(t, "https://cdn.radiofrance.fr/s3/cruiser-production/2021/07/2ae5451e-bf4f-482d-8b77-1aaecc276f8e/wl-net_mfi_a9cf7f77-3281-4579-831a-7e39fee30b0d-10.0121029927.mp3", a)
	assert.Equal(t, "Écouter Jeanne Barret : la première femme qui a fait le tour du monde", b)
}

