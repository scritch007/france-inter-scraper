package fiscrap

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"path"
	"testing"
	"github.com/stretchr/testify/mock"
)

type fetchMock struct {
	mock.Mock
}

func (f *fetchMock) Do(url string) (io.ReadCloser, error) {
	called := f.Called(url)
	res, _ := called.Get(0).(io.ReadCloser)
	return res, called.Error(1)
}

func TestScrapper_parseList(t *testing.T) {
	fm := &fetchMock{}
	dom := fm.On("Do", mock.Anything)
	dom.Run(func(args mock.Arguments) {
		f, err := os.Open(path.Join("testdata", args[0].(string)))
		dom.ReturnArguments = mock.Arguments{io.ReadCloser(f), err}
	})
	s := Scrapper{
		fetch: fm,
	}
	list, err := s.parseList("odysses.html")
	require.NoError(t, err)

	assert.Equal(t, []string{
		"https://www.franceinter.fr/emissions/les-odyssees/jeanne-barret",
		"https://www.franceinter.fr/emissions/les-odyssees/victorine-brocher-heroine-de-la-commune",
		"https://www.franceinter.fr/emissions/les-odyssees/stanley-sur-la-piste-de-david-livingstone",
		"https://www.franceinter.fr/emissions/les-odyssees/les-evades-d-alcatraz",
		"https://www.franceinter.fr/emissions/les-odyssees/l-iliade-episode-1",
		"https://www.franceinter.fr/emissions/les-odyssees/l-iliade-episode-2",
		"https://www.franceinter.fr/emissions/les-odyssees/l-iliade-episode-3",
		"https://www.franceinter.fr/emissions/les-odyssees/odyssee-ulysse-prisonnier-de-la-nymphe-calypso",
		"https://www.franceinter.fr/emissions/les-odyssees/odyssee-episode-2",
		"https://www.franceinter.fr/emissions/les-odyssees/odyssee-episode-3-ulysse-et-circe-la-magicienne",
		"https://www.franceinter.fr/emissions/les-odyssees/alice-guy-la-premiere-realisatrice-de-l-histoire-du-cinema",
		"https://www.franceinter.fr/emissions/les-odyssees/orson-welles-et-la-guerre-des-mondes-a-t-il-vraiment-panique-l-amerique",
		"https://www.franceinter.fr/emissions/les-odyssees/plongee-dans-les-abysses-lorsque-le-sous-marin-alvin-decouvre-les-sources-hydrothermales",
		"https://www.franceinter.fr/emissions/les-odyssees/napoleon-episode-1",
		"https://www.franceinter.fr/emissions/les-odyssees/napoleon-episode-2",
		"https://www.franceinter.fr/emissions/les-odyssees/le-ballon-de-la-liberte",
		"https://www.franceinter.fr/emissions/les-odyssees/jeanne-d-arc",
		"https://www.franceinter.fr/emissions/les-odyssees/jeanne-d-arc-episode-2",
		"https://www.franceinter.fr/emissions/les-odyssees/mary-anning",
		"https://www.franceinter.fr/emissions/les-odyssees/la-decouverte-de-la-tombe-de-gengis-khan",
	}, list)
}
