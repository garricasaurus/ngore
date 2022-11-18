package ncgore

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApi_Search(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		html := `
			<div class="torrent_txt">
				<a href="torrents.php?action=details&id=3194285" onclick="torrent(3194285); return false;" title="A másik Göring - megosztott testvériség"><nobr>A másik Göring - megosztott testvériség</nobr></a>
				<div class='torrent_txt_also'><div class="infobar"><img onmouseout="elrejt('borito3194285')" onmouseover="mutat('https://nc-img.cdn.l7cache.com/covers/L9_kMzZ3fwZFl_Zl?27055080', '300', 'borito3194285', this)" border="0" src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7" class="infobar_ico"></div><div class="siterank"> <span title="The Other Goering - A Divided Brotherhood">The Other Goering - A Divided ...</span> </div></div></div>
			</div>`
		_, err := w.Write([]byte(html))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()
	api := apiWithMockClient(server)
	results, err := api.Search(&SearchParams{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(results))
	assert.Equal(t, "A másik Göring - megosztott testvériség", results[0].Title)
}
