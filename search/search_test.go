package search

import (
	"git.okki.hu/garric/ngore/parse"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApi_Search(t *testing.T) {

	t.Run("parse single match", func(t *testing.T) {
		doc := parse.MustParse(t, `
		<div class="box_torrent">
			<div class="box_alap_img">
				<a href="/torrents.php?tipus=xvid_hun">
					<img src="https://static.ncore.pro/styles/brutecore/ico/ico_xvid_hun.png"
						 class="categ_link" alt="SD/HU"
						 title="Filmek tömörített formátumban, magyarul.">
				</a>
			</div>
			<div class="box_nagy">
				<div class="box_nev2">
					<div style='display:none;' id='borito3194285'></div>
					<div class="tabla_szoveg">
						<div style="cursor:pointer" onclick="konyvjelzo('3194285');" class="torrent_konyvjelzo2"></div>
						<div class="torrent_txt">
							<a href="torrents.php?action=details&id=3194285" onclick="torrent(3194285); return false;"
							   title="A másik Göring - megosztott testvériség">
								<nobr>A másik Göring - megosztott testvériség</nobr>
							</a>
							<div class='torrent_txt_also'>
								<div class="infobar">
									<img onmouseout="elrejt('borito3194285')"
										 onmouseover="mutat('https://nc-img.cdn.l7cache.com/covers/L9_kMzZ3fwZFl_Zl?27055080', '300', 'borito3194285', this)"
										 border="0"
										 src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"
										 class="infobar_ico">
								</div>
								<div class="siterank"><span title="The Other Goering - A Divided Brotherhood">The Other Goering - A Divided ...</span>
								</div>
							</div>
						</div>
					</div>
		
					<div class="torrent_ok" title="A torrent megfelel a szabályoknak"></div>
				</div>
				<div class="users_box_sepa"></div>
				<div class="box_feltoltve2">2021-06-10<br>08:00:19</div>
				<div class="users_box_sepa"></div>
				<div class="box_meret2">699.82 MiB</div>
				<div class="users_box_sepa"></div>
				<div class="box_d2">++</div>
				<div class="users_box_sepa"></div>
				<div class="box_s2"><a class="torrent" href="torrents.php?action=details&id=3194285&peers=1#peers">6</a></div>
				<div class="users_box_sepa"></div>
				<div class="box_l2"><a class="torrent" href="torrents.php?action=details&id=3194285&peers=1#peers">0</a></div>
				<div class="users_box_sepa"></div>
				<div class="box_feltolto2"><a href="wiki.php?action=read&id=391"><span
						class="feltolto_szin">Anonymous</span></a></div>
			</div>
		</div>`)
		results := ParseResponse(doc)
		expected := []*Torrent{
			{
				Title:    "A másik Göring - megosztott testvériség",
				AltTitle: "The Other Goering - A Divided Brotherhood",
				Health:   "++",
				Peers:    "0",
				Seeds:    "6",
				Size:     "699.82 MiB",
				Uploaded: "2021-06-10 08:00:19",
				Uploader: "Anonymous",
			},
		}
		assert.Equal(t, expected, results.Torrents)
	})

	t.Run("no torrent boxes", func(t *testing.T) {
		doc := parse.MustParse(t, `<div />`)
		results := ParseResponse(doc)
		assert.Empty(t, results.Torrents)
	})

	t.Run("no txt node", func(t *testing.T) {
		doc := parse.MustParse(t, `
		<div class="box_torrent">
			<div class="box_nagy"></div>
		</div>`)
		results := ParseResponse(doc)
		assert.Equal(t, 1, len(results.Torrents))
		assert.Equal(t, "", results.Torrents[0].Title)
		assert.Equal(t, "", results.Torrents[0].AltTitle)
	})

	t.Run("missing title", func(t *testing.T) {
		doc := parse.MustParse(t, `
		<div class="box_torrent">
			<div class="box_nagy">
				<div class="box_nev2">
					<div class="tabla_szoveg">
						<div class="torrent_txt">
							<a comment="no title attribute" />							
						</div>
					</div>					
				</div>				
			</div>
		</div>`)
		results := ParseResponse(doc)
		assert.Equal(t, 1, len(results.Torrents))
		assert.Equal(t, "", results.Torrents[0].Title)
	})

	t.Run("alt-title", func(t *testing.T) {

		t.Run("missing span", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_nagy">
					<div class="box_nev2">
						<div class="tabla_szoveg">
							<div class="torrent_txt">
								<div class='torrent_txt_also'>
									<div class="siterank" />
								</div>
							</div>
						</div>					
					</div>				
				</div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Title)
		})

		t.Run("missing title attribute", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_nagy">
					<div class="box_nev2">
						<div class="tabla_szoveg">
							<div class="torrent_txt">
								<div class='torrent_txt_also'>
									<div class="siterank"><span comment="no title attribute"></span>
								</div>
							</div>
						</div>					
					</div>				
				</div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Title)
		})

	})

	t.Run("health", func(t *testing.T) {

		t.Run("health data", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_d2">test</div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "test", results.Torrents[0].Health)
		})

		t.Run("missing node", func(t *testing.T) {
			doc := parse.MustParse(t, `<div class="box_torrent" />`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Health)
		})

		t.Run("missing content", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_d2" />
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Health)
		})

	})

	t.Run("peers", func(t *testing.T) {

		t.Run("peers data", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_l2"><a href="#">test</a></div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "test", results.Torrents[0].Peers)
		})

		t.Run("missing node", func(t *testing.T) {
			doc := parse.MustParse(t, `<div class="box_torrent" />`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Peers)
		})

		t.Run("missing href node", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_l2"></div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Peers)
		})

		t.Run("missing content", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_l2"><a href="#"></a></div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Peers)
		})

	})

	t.Run("seeds", func(t *testing.T) {

		t.Run("seeds data", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_s2"><a href="#">test</a></div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "test", results.Torrents[0].Seeds)
		})

		t.Run("missing node", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">				
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Seeds)
		})

		t.Run("missing href node", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_s2"></div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Seeds)
		})

		t.Run("missing data", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_s2"><a href="#"></a></div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Seeds)
		})

	})

	t.Run("size", func(t *testing.T) {

		t.Run("size data", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_meret2">test</div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "test", results.Torrents[0].Size)
		})

		t.Run("missing node", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">				
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Size)
		})

		t.Run("missing data", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_meret2"></div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Size)
		})

	})

	t.Run("uploaded", func(t *testing.T) {

		t.Run("uploaded data", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_feltoltve2">test</div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "test", results.Torrents[0].Uploaded)
		})

		t.Run("tags filtered from data", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_feltoltve2">test1<br/>test2</div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "test1 test2", results.Torrents[0].Uploaded)
		})

		t.Run("missing node", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Uploaded)
		})

		t.Run("missing data", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_feltoltve2"></div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Uploaded)
		})

	})

	t.Run("uploader", func(t *testing.T) {

		t.Run("uploader data", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_feltolto2"><span class="feltolto_szin">test</span></div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "test", results.Torrents[0].Uploader)
		})

		t.Run("missing node", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Uploader)
		})

		t.Run("missing span", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_feltolto2"></div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Uploader)
		})

		t.Run("missing data", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div class="box_torrent">
				<div class="box_feltolto2"><span class="feltolto_szin"></span></div>
			</div>`)
			results := ParseResponse(doc)
			assert.Equal(t, 1, len(results.Torrents))
			assert.Equal(t, "", results.Torrents[0].Uploader)
		})

	})

	t.Run("paging", func(t *testing.T) {

		t.Run("parse page info", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div id="pager_bottom">
				<a href="#"><strong>Első</strong></a>
				| <a href="#"><strong>1-25</strong></a>
				| <a href="#" id="pPa"><strong>26-50</strong></a>
				| <span class="active_link"><strong>51-75</strong></span> 
				| <a href="#" id="nPa"><strong>76-96</strong></a>
				| <a href="#"><strong>Utolsó</strong></a>
			</div>
			`)
			expected := &PageInfo{
				Current: 3,
				Prev:    2,
				Next:    4,
			}
			results := ParseResponse(doc)
			assert.Equal(t, expected, results.Page)
		})

		t.Run("first page", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div id="pager_bottom">
				<span class="active_link"><strong>1-25</strong></span>
				| <a href="#" id="nPa"><strong>26-50</strong></a>
				| <a href="#"><strong>51-75</strong></a>
				| <a href="#"><strong>76-96</strong></a>
				| <a href="#"><strong>Utolsó</strong></a>
			</div>
			`)
			results := ParseResponse(doc)
			expected := &PageInfo{
				Current: 1,
				Prev:    1,
				Next:    2,
			}
			assert.Equal(t, expected, results.Page)
		})

		t.Run("last page", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div id="pager_bottom">
				<a href="#"><strong>Első</strong></a> 
				| <a href="#"><strong>1-25</strong></a> 
				| <a href="#"><strong>26-50</strong></a> 
				| <a href="#" id="pPa"><strong>51-75</strong></a> 
				| <span class="active_link"><strong>76-96</strong></span>
			</div>
			`)
			expected := &PageInfo{
				Current: 4,
				Prev:    3,
				Next:    4,
			}
			results := ParseResponse(doc)
			assert.Equal(t, expected, results.Page)
		})

		t.Run("invalid range separator", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div id="pager_bottom">
				<a href="#"><strong>Első</strong></a> 
				| <span class="active_link"><strong>76..96</strong></span>
			</div>
			`)
			results := ParseResponse(doc)
			assert.Equal(t, defaultPage, results.Page.Current)
		})

		t.Run("range not a number", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div id="pager_bottom">
				<a href="#"><strong>Első</strong></a> 
				| <span class="active_link"><strong>foo-bar</strong></span>
			</div>
			`)
			results := ParseResponse(doc)
			assert.Equal(t, defaultPage, results.Page.Current)
		})

		t.Run("missing active span", func(t *testing.T) {
			doc := parse.MustParse(t, `
			<div id="pager_bottom">
				<a href="#"><strong>Első</strong></a>
			</div>
			`)
			results := ParseResponse(doc)
			assert.Equal(t, defaultPage, results.Page.Current)
		})

	})

}
