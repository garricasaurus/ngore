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

	})

}
