package details

import (
	"git.okki.hu/garric/ngore/parse"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseDetails(t *testing.T) {

	doc := parse.MustParse(t, `
<div class="torrent_leiras">
    <table width="100%">
        <tbody>
        <tr>
            <td align="center" class="inforbar_img"><img
                    alt="Borító" src="https://nc-img.cdn.l7cache.com/covers/7Y_Qz2gRTkxFNBdn?27818835"></td>
            <td class="inforbar_txt">
                <div class="infobar_title"> Black Adam</div>
                <table>
                    <tbody>
                    <tr>
                        <td style="vertical-align:top;width:110px;">Megjelenés éve:</td>
                        <td>2022</td>
                    </tr>
                    <tr>
                        <td style="vertical-align:top;width:110px;">Rendező:</td>
                        <td>Jaume Collet-Serra</td>
                    </tr>
                    <tr>
                        <td style="vertical-align:top;width:110px;">Szereplők:</td>
                        <td>Dwayne Johnson, Aldis Hodge, Pierce Brosnan, Noah Centineo, Sarah Shahi</td>
                    </tr>
                    <tr>
                        <td style="vertical-align:top;width:110px;">Ország:</td>
                        <td>United States, Canada, New Zealand, Hungary</td>
                    </tr>
                    <tr>
                        <td style="vertical-align:top;width:110px;">Hossz:</td>
                        <td>125 perc</td>
                    </tr>
                    <tr>
                        <td style="vertical-align:top;width:110px;">Címkék:</td>
                        <td><a href="torrents.php?tags=akci%C3%B3" style="border-bottom:1px dotted white;">akció</a>, <a
                                href="torrents.php?tags=kaland" style="border-bottom:1px dotted white;">kaland</a>, <a
                                href="torrents.php?tags=fantasy" style="border-bottom:1px dotted white;">fantasy</a>
                        </td>
                    </tr>
                    <tr>
                        <td style="vertical-align:top;width:110px;">IMDb értékelés:</td>
                        <td>6.9</td>
                    </tr>
                    <tr>
                        <td style="vertical-align:top;width:110px;">IMDb link:</td>
                        <td><a href="https://dereferer.link/?https://imdb.com/title/tt6443346" target="_blank">https://imdb.com/title/tt6443346</a>
                        </td>
                    </tr>
                    <tr>
                        <td style="vertical-align:top;width:110px;">Egyéb link:</td>
                        <td><a href="https://dereferer.link/?https://www.themoviedb.org/movie/436270" target="_blank">https://www.themoviedb.org/mov...</a>
                        </td>
                    </tr>
                    </tbody>
                </table>
            </td>
        </tr>
        </tbody>
    </table>
</div>
	`)

	details := ParseDetails(doc)

	t.Run("title", func(t *testing.T) {
		assert.Equal(t, "Black Adam", details.Title)
	})

	t.Run("release year", func(t *testing.T) {
		assert.Equal(t, "2022", details.ReleaseYear)
	})

	t.Run("director", func(t *testing.T) {
		assert.Equal(t, "Jaume Collet-Serra", details.Director)
	})

	t.Run("actors", func(t *testing.T) {
		assert.Equal(t, "Dwayne Johnson, Aldis Hodge, Pierce Brosnan, Noah Centineo, Sarah Shahi", details.Actors)
	})

	t.Run("country", func(t *testing.T) {
		assert.Equal(t, "United States, Canada, New Zealand, Hungary", details.Country)
	})

	t.Run("labels", func(t *testing.T) {
		assert.Equal(t, "akció kaland fantasy", details.Labels)
	})

	t.Run("imdb rating", func(t *testing.T) {
		assert.Equal(t, "6.9", details.ImdbRating)
	})

	t.Run("imdb label", func(t *testing.T) {
		assert.Equal(t, "https://imdb.com/title/tt6443346", details.ImdbLink)
	})

	t.Run("length", func(t *testing.T) {
		assert.Equal(t, "125 perc", details.Length)
	})

	t.Run("other link", func(t *testing.T) {
		assert.Equal(t, "https://www.themoviedb.org/movie/436270", details.OtherLink)
	})

}
