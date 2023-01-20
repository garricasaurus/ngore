package activity

import (
	"git.okki.hu/garric/ngore/parse"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseResponse(t *testing.T) {

	t.Run("parse table data", func(t *testing.T) {
		doc := parse.MustParse(t, `
		<div class="fobox_tartalom"></div>
		<div class="fobox_tartalom">
			<div style="width:36%;float:left;">
				<div class="dt" style="border:0;width:200px;">Napi helyezés:</div>
				<div class="dd" style="border:0;">23235</div>
				<div class="dt" style="border:0;width:200px;">Heti helyezés:</div>
				<div class="dd" style="border:0;">23824</div>
				<div class="dt" style="border:0;width:200px;">Havi helyezés:</div>
				<div class="dd" style="border:0;">22752</div>
				<div class="dt" style="border:0;width:200px;">Elmúlt havi helyezés:</div>
				<div class="dd" style="border:0;">30666</div>
				<div class="dt" style="border:0;width:200px;">Jelenleg letölthet:</div>
				<div class="dd" style="border:0;">igen</div>
				<div class="dt" style="border:0;width:200px;">Lehetséges hit'n'runok száma:</div>
				<div class="dd" style="border:0;">3</div>
				<div class="dt" style="border:0;width:200px;">Hit'n'runolható torrentek száma:</div>
				<div class="dd" style="border:0;">korlátlan</div>
				<div class="dt" style="border:0;width:200px;">Hit'n'runolt hónapok száma:</div>
				<div class="dd" style="border:0;">5</div>
				<div class="dt" style="border:0;width:200px;">Hit'n'runolt torrentek száma:</div>
				<div class="dd" style="border:0;">0</div>
			</div>
		</div>
		`)

		t.Run("parse rank data", func(t *testing.T) {
			info := ParseResponse(doc)
			expectedRank := Rank{
				Daily:     "23235",
				Weekly:    "23824",
				Monthly:   "22752",
				PrevMonth: "30666",
			}
			assert.Equal(t, expectedRank, info.Rank)
		})

		t.Run("parse can download", func(t *testing.T) {
			info := ParseResponse(doc)
			assert.Equal(t, "igen", info.CanDownload)
		})

		t.Run("parse stats data", func(t *testing.T) {
			info := ParseResponse(doc)
			expectedStats := Stats{
				Current:     "3",
				Allowed:     "korlátlan",
				PenMonths:   "5",
				PenTorrents: "0",
			}
			assert.Equal(t, expectedStats, info.Stats)
		})

	})

	t.Run("missing table element", func(t *testing.T) {
		doc := parse.MustParse(t, "")
		info := ParseResponse(doc)
		expectedStats := Stats{
			Current:     "",
			Allowed:     "",
			PenMonths:   "",
			PenTorrents: "",
		}
		assert.Equal(t, expectedStats, info.Stats)
	})

	t.Run("invalid table length", func(t *testing.T) {
		doc := parse.MustParse(t, `
		<div class="fobox_tartalom">
			<div class="dd" style="border:0;">23235</div>
			<div class="dd" style="border:0;">23824</div>
		</div>
		`)
		info := ParseResponse(doc)
		expectedStats := Stats{
			Current:     "",
			Allowed:     "",
			PenMonths:   "",
			PenTorrents: "",
		}
		assert.Equal(t, expectedStats, info.Stats)
	})

	t.Run("parse history", func(t *testing.T) {
		doc := parse.MustParse(t, `		
		<div class="hnr_torrents">
			<div class="hnr_all2">
				<div class="hnr_tname">
					<a href="#" title="Test Title"></a>
				</div>
				<div class="hnr_tstart">1 órája</div>
				<div class="hnr_tlastactive">6 perce</div>
				<div class="hnr_tseed"><span class="stopped">Seed</span></div>
				<div class="hnr_tup">0 B</div>
				<div class="hnr_tdown">2.69 GiB</div>
				<div class="hnr_ttimespent"><span class="stopped">46ó 41p</span></div>
				<div class="hnr_tratio"><span class="stopped">0.000</span></div>
			</div>
		</div>
		<div style="clear:both;"></div>
		<div class="hnr_torrents">
			<div class="hnr_all">
				<div class="hnr_tname">
					<a href="#" title="Test Title 2"></a>
				</div>
				<div class="hnr_tstart">1 órája</div>
				<div class="hnr_tlastactive">19 perce</div>
				<div class="hnr_tseed"><span class="stopped">Seed</span></div>
				<div class="hnr_tup">0 B</div>
				<div class="hnr_tdown">3.19 GiB</div>
				<div class="hnr_ttimespent"><span class="stopped">47ó 6p</span></div>
				<div class="hnr_tratio"><span class="stopped">0.000</span></div>
			</div>
		</div>
		<div style="clear:both;"></div>
		<div class="hnr_torrents">
			<div class="hnr_all2">
				<div class="hnr_tname">
					<a href="#" title="Test Title 3"></a>
				</div>
				<div class="hnr_tstart">1 órája</div>
				<div class="hnr_tlastactive">10 perce</div>
				<div class="hnr_tseed"><span class="stopped">Seed</span></div>
				<div class="hnr_tup">0 B</div>
				<div class="hnr_tdown">1.77 GiB</div>
				<div class="hnr_ttimespent"><span class="stopped">46ó 39p</span></div>
				<div class="hnr_tratio"><span class="stopped">0.000</span></div>
			</div>
		</div>
		<div style="clear:both;"></div>	 
		<div class="lista_lab"></div>	
		`)

		info := ParseResponse(doc)
		history := info.History

		assert.Contains(t, history, TorrentActivity{
			Name:      "Test Title",
			Start:     "1 órája",
			Updated:   "6 perce",
			Status:    "Seed",
			Up:        "0 B",
			Down:      "2.69 GiB",
			Remaining: "46ó 41p",
			Ratio:     "0.000",
		})

		assert.Contains(t, history, TorrentActivity{
			Name:      "Test Title 2",
			Start:     "1 órája",
			Updated:   "19 perce",
			Status:    "Seed",
			Up:        "0 B",
			Down:      "3.19 GiB",
			Remaining: "47ó 6p",
			Ratio:     "0.000",
		})

		assert.Contains(t, history, TorrentActivity{
			Name:      "Test Title 3",
			Start:     "1 órája",
			Updated:   "10 perce",
			Status:    "Seed",
			Up:        "0 B",
			Down:      "1.77 GiB",
			Remaining: "46ó 39p",
			Ratio:     "0.000",
		})

	})

	t.Run("missing history element", func(t *testing.T) {
		doc := parse.MustParse(t, ``)
		info := ParseResponse(doc)
		assert.NotNil(t, info.History)
		assert.Len(t, info.History, 0)
	})

	t.Run("missing history attributes", func(t *testing.T) {
		doc := parse.MustParse(t, `
		<div class="hnr_torrents">
			<div class="hnr_all" />
		</div>		
		`)

		info := ParseResponse(doc)
		history := info.History

		assert.Len(t, history, 1)
		assert.Equal(t, TorrentActivity{}, history[0])
	})

	t.Run("missing status span", func(t *testing.T) {
		doc := parse.MustParse(t, `
		<div class="hnr_torrents">
			<div class="hnr_all">
				<div class="hnr_tseed"></div>				
			</div>
		</div>
		`)

		info := ParseResponse(doc)
		history := info.History

		assert.Len(t, history, 1)
		assert.Equal(t, "", history[0].Status)
	})
}
