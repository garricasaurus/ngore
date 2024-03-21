package recommended

import (
	"testing"

	"github.com/gar-r/ngore/parse"
	"github.com/stretchr/testify/assert"
)

func TestParseRecommendations(t *testing.T) {
	doc := parse.MustParse(t, example)
	rec := ParseRecommendations(doc)

	t.Run("movie recommendations", func(t *testing.T) {
		t.Run("staff recommendations", func(t *testing.T) {
			assert.Equal(t, &RecommendationWithImage{
				Recommendation: Recommendation{
					Id:   "1980392",
					Name: "Saul.fia.2015.1080p.HUN.Blu-ray.AVC.DTS-HD.MA.5.1-HyperX",
				},
				Image: "https://nc-img.cdn.l7cache.com/covers/-9ebvpwwI2vSwemp?24370470",
			}, rec.Movies.Staff[0])
			assert.Equal(t, &RecommendationWithImage{
				Recommendation: Recommendation{
					Id:   "2246144",
					Name: "A hihetetlen Hulk BD50",
				},
				Image: "https://nc-img.cdn.l7cache.com/covers/7DBrgkWLcbnC7_vV?25390065",
			}, rec.Movies.Staff[1])
		})

		t.Run("active items", func(t *testing.T) {
			assert.Equal(t, &Recommendation{
				Id:   "3478636",
				Name: "Ticket.to.Paradise.2022.BDRip.x264.HuN-No1",
			}, rec.Movies.Active[0])
			assert.Equal(t, &Recommendation{
				Id:   "3476470",
				Name: "Infinite.2021.BDRip.x264.HuN-No1",
			}, rec.Movies.Active[1])
		})
	})

	t.Run("series recommendations", func(t *testing.T) {
		t.Run("staff recommendations", func(t *testing.T) {
			assert.Equal(t, &RecommendationWithImage{
				Recommendation: Recommendation{
					Id:   "2887592",
					Name: "Szerelem zálogba S01",
				},
				Image: "https://nc-img.cdn.l7cache.com/covers/-9eboLv2c2vSwBmp?26130788",
			}, rec.Series.Staff[0])
			assert.Equal(t, &RecommendationWithImage{
				Recommendation: Recommendation{
					Id:   "3206598",
					Name: "Tom.Clancys.Jack.Ryan.S01.REPACK.720p.WEBRip.DD+5.1.x264.HuN-No1",
				},
				Image: "https://nc-img.cdn.l7cache.com/covers/jg_mp8ZKcRwfle8z?27105406",
			}, rec.Series.Staff[1])
		})

		t.Run("active items", func(t *testing.T) {
			assert.Equal(t, &Recommendation{
				Id:   "3481532",
				Name: "The.Last.of.Us.S01E05.CRAV.WEBRip.x264.HUN.ENG-FULCRUM",
			}, rec.Series.Active[0])
			assert.Equal(t, &Recommendation{
				Id:   "3481378",
				Name: "The.Last.of.Us.S01E05.1080p.AMZN.WEB-DL.DDP5.1.Atmos.H.264.HUN.ENG-PTHD",
			}, rec.Series.Active[1])
		})
	})

}

const example string = `
<div id="main_all">
	<div id="main">
			<div class="infocsik" style="background:#BB0000;"><a href="index.php" style="color:#fff;">Új hír van a főoldalon!</a> / <a href="promo.php" style="color:#fff;">Új promóciós hír lett kiírva</a></div>
	<div id="main_tartalom">
<div class="fobox_all">
	<div class="fobox_fej_t"></div>
	<div class="fobox_fej">Kategóriák</div>
	<div class="fobox_fej_b"></div>
	<div class="fobox_tartalom_t"></div>
	<div class="fobox_tartalom">
		<br>
		<div style="text-align:center"><a href="#film"> Film </a> | <a href="#sorozat"> Sorozat </a> | <a href="#jatek"> Játék </a> | <a href="#zene"> Zene </a> | <a href="#program"> Program </a> | <a href="#ebook"> eBook </a><br><br>
	</div><br>        </div>
	<div class="fobox_lab"></div>
</div>
<a name="film">&nbsp;</a>    <div class="fobox_all">
	<div class="fobox_fej_t"></div>
	<div class="fobox_fej">Film</div>
	<div class="fobox_fej_b"></div>
	<div class="fobox_tartalom_t"></div>
	<div class="fobox_tartalom">
				<div style="display:flex; flex-direction:row;">
			<div style="border-right: 4px solid #e0e0e0;width: 50%;">
				Staff által jelölt
				<br><br><br>
				<div style="position:relative; overflow:hidden; text-align:center">
												<div style="float:left; width:33%">
							<a href="torrents.php?action=details&amp;id=1980392" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/-9ebvpwwI2vSwemp?24370470" title="Saul.fia.2015.1080p.HUN.Blu-ray.AVC.DTS-HD.MA.5.1-HyperX" width="120" height="176" border="1"></a>
																<br><br>
							<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;HD: 43.17 GiB<br>
							Seederek: 5<br>								Leecherek: 1							</div>
													<div style="float:left; width:33%">
							<a href="torrents.php?action=details&amp;id=2246144" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/7DBrgkWLcbnC7_vV?25390065" title="A hihetetlen Hulk BD50" width="120" height="176" border="1"></a>
																<br><br>
							<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;HD: 32.52 GiB<br>
							Seederek: 5<br>								Leecherek: 0							</div>
													<div style="float:left; width:33%">
							<a href="torrents.php?action=details&amp;id=1662956" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/7k_j3lZNsAwS2_wD" title="Hundraaringen.som.klev.ut.genom.fonstret.och.forsvann.2013.720p.BluRay.DD5.1.x264.HuN-TRiNiTY" width="120" height="176" border="1"></a>
																<br><br>
							<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;HD: 5.63 GiB<br>
							Seederek: 181<br>								Leecherek: 13							</div>
													<div style="float:left; width:33%">
							<br><br><a href="torrents.php?action=details&amp;id=3484033" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/jAeRd39bh9vUOB-7?27941723" title="Napszállta BD50" width="120" height="176" border="1"></a>
																<br><br>
							<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;HD: 30.78 GiB<br>
							Seederek: 36<br>								Leecherek: 0							</div>
													<div style="float:left; width:33%">
							<br><br><a href="torrents.php?action=details&amp;id=3481732" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/qlX6loG-HKkuRedE?27935516" title="A csend kódja" width="120" height="176" border="1"></a>
																<br><br>
							<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;DVD: 4.35 GiB<br>
							Seederek: 28<br>								Leecherek: 1							</div>
													<div style="float:left; width:33%">
							<br><br><a href="torrents.php?action=details&amp;id=1098865" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/oWX8vn9bSgdFwBl8" title="Titanic BD50" width="120" height="176" border="1"></a>
																<br><br>
							<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;HD: 44.52 GiB<br>
							Seederek: 11<br>								Leecherek: 1							</div>
												</div>
					<br><br><br>
					Legaktívabb torrentek						<br><br><br>
					<div style="position:relative; overflow:hidden; text-align:center">
												<div style="float:left; width:33%">
							<a href="torrents.php?action=details&amp;id=3420516" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/5nXKrMYji37h9eA6?27774150" title="Bullet.Train.2022.BDRip.x264.HuN-CRLS" width="120" height="176" border="1"></a>
																<br><br>
							<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;SD: 2.07 GiB<br>
							Seederek: 9967<br>								Leecherek: 930							</div>
													<div style="float:left; width:33%">
							<a href="torrents.php?action=details&amp;id=3478636" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/0Je1xDLxCDgC9X57?27928028" title="Ticket.to.Paradise.2022.BDRip.x264.HuN-No1" width="120" height="176" border="1"></a>
																<br><br>
							<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;SD: 1.53 GiB<br>
							Seederek: 6721<br>								Leecherek: 782							</div>
													<div style="float:left; width:33%">
							<a href="torrents.php?action=details&amp;id=3389970" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/kG4OvjANt8RH1BWM?27687293" title="Top.Gun.Maverick.2022.IMAX.WEBRip.x264.HuN-No1" width="120" height="176" border="1"></a>
																<br><br>
							<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;SD: 2.07 GiB<br>
							Seederek: 9154<br>								Leecherek: 737							</div>
											</div>
			</div>
			<div style="width:50%; padding-left: 16px;">
										<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3478636" target="_blank">Ticket.to.Paradise.2022.BDRip.x264.HuN-No1</a>
						<div style="min-width: 150px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3476470" target="_blank">Infinite.2021.BDRip.x264.HuN-No1</a>
						<div style="min-width: 67.367434386717px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3475331" target="_blank">Black.Panther.Wakanda.Forever.2022.BDRip.x264.HuN-No1</a>
						<div style="min-width: 60.672201392608px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3477453" target="_blank">Good.Luck.to.You.Leo.Grande.2022.BDRip.DD2.0.x264.HuN-Essence</a>
						<div style="min-width: 45.460632029995px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3420516" target="_blank">Bullet.Train.2022.BDRip.x264.HuN-CRLS</a>
						<div style="min-width: 42.554900910552px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482790" target="_blank">Halloween.Ends.2022.BDRiP.DD5.1.x264.HUN-Gianni</a>
						<div style="min-width: 40.921264059989px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3389970" target="_blank">Top.Gun.Maverick.2022.IMAX.WEBRip.x264.HuN-No1</a>
						<div style="min-width: 37.640599892876px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3478617" target="_blank">Ticket.to.Paradise.2022.720p.BluRay.DD5.1.x264.HUN.ENG-PTHD</a>
						<div style="min-width: 34.493840385645px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3459645" target="_blank">The.Menu.2022.MA.WEBRip.x264.HUN-FULCRUM</a>
						<div style="min-width: 34.480449919657px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481045" target="_blank">Detective.Knight.Rogue.2022.480p.BluRay.DD2.0.x264.HuN-BaKeR</a>
						<div style="min-width: 33.355650776647px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3475549" target="_blank">Black.Panther.Wakanda.Forever.2022.IMAX.WEBRip.x264.HuN-No1</a>
						<div style="min-width: 29.35190144617px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3413224" target="_blank">Jurassic.World.Dominion.2022.Theatrical.Cut.BDRip.x264.HuN-No1</a>
						<div style="min-width: 29.258168184253px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3480888" target="_blank">Your.Place.or.Mine.2023.720p.NF.WEB-DL.DD+5.1.Atmos.H.264.HuN-No1</a>
						<div style="min-width: 28.146759507231px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3437646" target="_blank">Black.Adam.2022.RERiP.MA.WEBRip.x264.HuN-No1</a>
						<div style="min-width: 27.021960364221px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3478618" target="_blank">Ticket.to.Paradise.2022.1080p.BluRay.DDP7.1.x264.HUN.ENG-PTHD</a>
						<div style="min-width: 26.807712908409px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3475512" target="_blank">Black.Panther.Wakanda.Forever.2022.READ.NFO.IMAX.1080p.DSNP.WEB-DL.DD+7.1.H.264.HuN-No1</a>
						<div style="min-width: 26.419389394751px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3332541" target="_blank">Uncharted.2022.BDRip.x264.HuN-No1</a>
						<div style="min-width: 22.281735404392px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3427099" target="_blank">Where.the.Crawdads.Sing.2022.MA.WEBRip.x264.HuN-No1</a>
						<div style="min-width: 21.880021424746px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3479754" target="_blank">Detective.Knight.Redemption.2022.480p.BluRay.DD2.0.x264.HuN-BaKeR</a>
						<div style="min-width: 21.679164434922px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3425145" target="_blank">Minions.The.Rise.of.Gru.2022.BDRip.x264.HuN-No1</a>
						<div style="min-width: 21.143545795394px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3393215" target="_blank">The.Northman.2022.BDRiP.DD5.1.x264.HUN-Gianni</a>
						<div style="min-width: 17.099625066952px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3449458" target="_blank">The.Banshees.of.Inisherin.2022.MA.WEBRip.x264.HUN-FULCRUM</a>
						<div style="min-width: 16.510444563471px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3397800" target="_blank">Thor.Love.and.Thunder.2022.IMAX.DSNP.WEBRip.x264.HuN-No1</a>
						<div style="min-width: 16.3631494376px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3257496" target="_blank">Free.Guy.2021.BDRip.x264.HuN-No1</a>
						<div style="min-width: 15.867702196036px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3196776" target="_blank">Nobody.2021.BDRip.x264.HuN-No1</a>
						<div style="min-width: 15.251740760578px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3273903" target="_blank">Wrath.of.Man.2021.BDRip.x264.HuN-No1</a>
						<div style="min-width: 15.010712372791px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3363345" target="_blank">Doctor.Strange.in.the.Multiverse.of.Madness.2022.IMAX.DSNP.WEBRip.x264.HUN-FULCRUM</a>
						<div style="min-width: 14.635779325121px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482592" target="_blank">Halloween.Ends.2022.720p.BluRay.DD5.1.x264.HUN.ENG-PTHD</a>
						<div style="min-width: 14.367970005356px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3480133" target="_blank">Intet.2022.WEBRip.x264.HuN-No1</a>
						<div style="min-width: 14.26084627745px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3480334" target="_blank">After.Chernobyl.2021.WEBRip.x264.HuN-No1</a>
						<div style="min-width: 13.818960899839px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3381631" target="_blank">Elvis.2022.WEBRip.x264.HuN-No1</a>
						<div style="min-width: 13.671665773969px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3476456" target="_blank">Infinite.2021.720p.BluRay.DD5.1.x264.HUN.ENG-PTHD</a>
						<div style="min-width: 13.350294590252px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482593" target="_blank">Halloween.Ends.2022.1080p.BluRay.DDP7.1.x264.HUN.ENG-PTHD</a>
						<div style="min-width: 13.269951794322px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3267810" target="_blank">Venom.Let.There.Be.Carnage.2021.BDRip.x264.HuN-No1</a>
						<div style="min-width: 13.015532940546px;background: #43A047;height: 16px;"></div>
					</div>
									</div>
		</div>
			</div>
	<div class="fobox_lab"></div>
</div>
<a name="sorozat">&nbsp;</a>    <div class="fobox_all">
	<div class="fobox_fej_t"></div>
	<div class="fobox_fej">Sorozat | <a href="#top" class="premium">A lap tetejére</a></div>
	<div class="fobox_fej_b"></div>
	<div class="fobox_tartalom_t"></div>
	<div class="fobox_tartalom">
				<div style="display:flex; flex-direction:row;">
			<div style="border-right: 4px solid #e0e0e0;width: 50%;">
				Staff által jelölt
				<br><br><br>
				<div style="position:relative; overflow:hidden; text-align:center">
								<div style="float:left; width:33%">
					<a href="torrents.php?action=details&amp;id=2887592" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/-9eboLv2c2vSwBmp?26130788" title="Szerelem zálogba S01" width="120" height="176" border="1"></a>
												<br><br>
					<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;SD: 92.04 GiB<br>
					Seederek: 9<br>						Leecherek: 4					</div>
									<div style="float:left; width:33%">
					<a href="torrents.php?action=details&amp;id=3206598" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/jg_mp8ZKcRwfle8z?27105406" title="Tom.Clancys.Jack.Ryan.S01.REPACK.720p.WEBRip.DD+5.1.x264.HuN-No1" width="120" height="176" border="1"></a>
												<br><br>
					<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;HD: 22.95 GiB<br>
					Seederek: 299<br>						Leecherek: 63					</div>
									<div style="float:left; width:33%">
					<a href="torrents.php?action=details&amp;id=164696" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/EdeY9J5gcDrCLB-x?27886680" title="Batman - A rajzfilmsorozat S01E01-E14" width="120" height="176" border="1"></a>
												<br><br>
					<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;SD: 4.36 GiB<br>
					Seederek: 37<br>						Leecherek: 3					</div>
									<div style="float:left; width:33%">
					<br><br><a href="torrents.php?action=details&amp;id=3424202" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/7YBQzLbNHkxFN4dn?27782471" title="Guillermo.del.Toros.Cabinet.of.Curiosities.S01.720p.NF.WEB-DL.DD+5.1.Atmos.H.264.HuN-No1" width="120" height="176" border="1"></a>
												<br><br>
					<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;HD: 10.58 GiB<br>
					Seederek: 347<br>						Leecherek: 45					</div>
									<div style="float:left; width:33%">
					<br><br><a href="torrents.php?action=details&amp;id=2280703" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/1zBPkNK5i89HLeJn?25531855" title="FIFA.World.Cup.2018.HDTV.HUN.x264-JVC" width="120" height="176" border="1"></a>
												<br><br>
					<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;SD: 115.94 GiB<br>
					Seederek: 5<br>						Leecherek: 2					</div>
									<div style="float:left; width:33%">
					<br><br><a href="torrents.php?action=details&amp;id=3250394" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/PxXzqkdpfMlT-_bM?27261408" title="Charlies.Angels.S01-S05.COMPLETE.1080p.BluRay.DD2.0.x264.HuN-TRiNiTY" width="120" height="176" border="1"></a>
												<br><br>
					<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;HD: 465.39 GiB<br>
					Seederek: 9<br>						Leecherek: 6					</div>
								</div>
			<br><br><br>
			Legaktívabb torrentek					<br><br><br>
			<div style="position:relative; overflow:hidden; text-align:center">
									<div style="float:left; width:33%">
					<a href="torrents.php?action=details&amp;id=3466693" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/K9Xg6-8QCZduA_D2?27900002" title="A.Kiraly.S01.HUN.WEB-DL.H264-LEGION" width="120" height="176" border="1"></a>
												<br><br>
					<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;SD: 5.79 GiB<br>
					Seederek: 3409<br>						Leecherek: 1324						</div>
										<div style="float:left; width:33%">
					<a href="torrents.php?action=details&amp;id=3416201" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/ZV_3KgY8IGkHwXpM?27762352" title="The.Lord.of.the.Rings.The.Rings.of.Power.S01.AMZN.WEBRip.x264.HUN.ENG-FULCRUM" width="120" height="176" border="1"></a>
												<br><br>
					<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;SD: 7.16 GiB<br>
					Seederek: 3018<br>						Leecherek: 948						</div>
										<div style="float:left; width:33%">
					<a href="torrents.php?action=details&amp;id=3466958" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/6wBxgnLltNAtkXpy?27900952" title="A.Kiraly.S01.HUN.WEB-DL.1080p.H264-LEGION" width="120" height="176" border="1"></a>
												<br><br>
					<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;HD: 18.5 GiB<br>
					Seederek: 1964<br>						Leecherek: 771						</div>
										</div>
			</div>
			<div style="width:50%; padding-left: 16px;">
										<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481532" target="_blank">The.Last.of.Us.S01E05.CRAV.WEBRip.x264.HUN.ENG-FULCRUM</a>
						<div style="min-width: 150px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481378" target="_blank">The.Last.of.Us.S01E05.1080p.AMZN.WEB-DL.DDP5.1.Atmos.H.264.HUN.ENG-PTHD</a>
						<div style="min-width: 133.48769193273px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481377" target="_blank">The.Last.of.Us.S01E05.720p.AMZN.WEB-DL.DDP5.1.Atmos.H.264.HUN.ENG-PTHD</a>
						<div style="min-width: 122.05703144041px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481380" target="_blank">The.Last.of.Us.S01E05.2160p.HMAX.WEB-DL.DDP5.1.Atmos.DV.HDR.H.265.HUN.ENG-PTHD</a>
						<div style="min-width: 54.64294418718px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482015" target="_blank">Golkiralysag.S01E06.RTLM.WEB-DL.H.264.HUN-FULCRUM</a>
						<div style="min-width: 39.178649768462px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3478645" target="_blank">The.Last.of.Us.S01E04.CRAV.WEBRip.x264.HUN.ENG-FULCRUM</a>
						<div style="min-width: 38.971484279795px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482823" target="_blank">NFL - Super Bowl LVII - Philadelphia Eagles vs. Kansas City Chiefs 2023.02.13. 1080p</a>
						<div style="min-width: 27.516451377041px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3478571" target="_blank">The.Last.of.Us.S01E04.1080p.AMZN.WEB-DL.DDP5.1.Atmos.H.264.HUN.ENG-PTHD</a>
						<div style="min-width: 25.65196197904px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3474190" target="_blank">The.Last.of.Us.S01E03.CRAV.WEBRip.x264.HUN.ENG-FULCRUM</a>
						<div style="min-width: 23.190348525469px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482530" target="_blank">Capak.kozott.S05E06.RTLM.WEB-DL.H.264.HUN-FULCRUM</a>
						<div style="min-width: 22.861320984645px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3466098" target="_blank">The.Last.of.Us.S01E01.PROPER.CRAV.WEBRip.x264.HUN.ENG-FULCRUM</a>
						<div style="min-width: 21.325859127468px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3478570" target="_blank">The.Last.of.Us.S01E04.720p.AMZN.WEB-DL.DDP5.1.Atmos.H.264.HUN.ENG-PTHD</a>
						<div style="min-width: 21.130879844017px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482017" target="_blank">Golkiralysag.S01E06.1080p.RTLM.WEB-DL.AAC2.0.H.264.HUN-FULCRUM</a>
						<div style="min-width: 20.533755788447px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482532" target="_blank">Capak.kozott.S05E06.1080p.RTLM.WEB-DL.AAC2.0.H.264.HUN-FULCRUM</a>
						<div style="min-width: 20.192542042408px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481379" target="_blank">The.Last.of.Us.S01E05.2160p.CRAV.WEB-DL.DDP5.1.Atmos.H.265.HUN.ENG-PTHD</a>
						<div style="min-width: 19.900073117231px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3469899" target="_blank">The.Last.of.Us.S01E02.CRAV.WEBRip.x264.HUN.ENG-FULCRUM</a>
						<div style="min-width: 19.449183524251px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483690" target="_blank">Apatigris.S03E05.HUN.WEB-DL.H264-LEGION</a>
						<div style="min-width: 16.756032171582px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3466693" target="_blank">A.Kiraly.S01.HUN.WEB-DL.H264-LEGION</a>
						<div style="min-width: 16.427004630758px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482531" target="_blank">Capak.kozott.S05E06.720p.RTLM.WEB-DL.AAC2.0.H.264.HUN-FULCRUM</a>
						<div style="min-width: 15.951742627346px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482016" target="_blank">Golkiralysag.S01E06.720p.RTLM.WEB-DL.AAC2.0.H.264.HUN-FULCRUM</a>
						<div style="min-width: 15.415549597855px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481121" target="_blank">A.Spy.Among.Friends.S01.WEBRiP.DD2.0.x264.Hun-ARROW</a>
						<div style="min-width: 14.818425542286px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483817" target="_blank">Star.Wars.The.Bad.Batch.S02E09.720p.DSNP.WEB-DL.DDP5.1.H.264.HUN.ENG-PTHD</a>
						<div style="min-width: 14.525956617109px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481508" target="_blank">FBI.S05E05.HUN.WEBRip.x264-HNZ</a>
						<div style="min-width: 13.794784304168px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3480995" target="_blank">The.Legend.of.Vox.Machina.S02.720p.AMZN.WEB-DL.DDP5.1.H.264.HUNSUB-PTHD</a>
						<div style="min-width: 12.990494759932px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3438637" target="_blank">Wednesday.S01.NF.WEBRip.x264.Hun.Eng-MaMMuT</a>
						<div style="min-width: 12.563977577382px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3479478" target="_blank">Apatigris.S03E04.HUN.WEB-DL.H264-LEGION</a>
						<div style="min-width: 12.527418961735px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3478723" target="_blank">Star.Trek.Strange.New.Worlds.S01.720p.AMZN.WEB-DL.DD+5.1.H.264.HuN-No1</a>
						<div style="min-width: 12.052156958323px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483818" target="_blank">Star.Wars.The.Bad.Batch.S02E09.1080p.DSNP.WEB-DL.DDP5.1.H.264.HUN.ENG-PTHD</a>
						<div style="min-width: 11.979039727029px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3480996" target="_blank">The.Legend.of.Vox.Machina.S02.1080p.AMZN.WEB-DL.DDP5.1.H.264.HUNSUB-PTHD</a>
						<div style="min-width: 11.832805264441px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3474072" target="_blank">The.Last.of.Us.S01E03.720p.AMZN.WEB-DL.DDP5.1.Atmos.H.264.HUN.ENG-PTHD</a>
						<div style="min-width: 11.796246648794px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483692" target="_blank">Apatigris.S03E05.HUN.WEB-DL.1080p.H264-LEGION</a>
						<div style="min-width: 11.308798440166px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483217" target="_blank">FBI.S05E04.HUN.WEBRip.x264-HNZ</a>
						<div style="min-width: 11.162563977577px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3465975" target="_blank">The.Last.of.Us.S01E01.720p.AMZN.WEB-DL.DDP5.1.Atmos.H.264.HUN.ENG-PTHD</a>
						<div style="min-width: 10.931026078479px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3474073" target="_blank">The.Last.of.Us.S01E03.1080p.AMZN.WEB-DL.DDP5.1.Atmos.H.264.HUN.ENG-PTHD</a>
						<div style="min-width: 10.906653668048px;background: #43A047;height: 16px;"></div>
					</div>
									</div>
		</div>
			</div>
	<div class="fobox_lab"></div>
</div>
<a name="jatek">&nbsp;</a>    <div class="fobox_all">
	<div class="fobox_fej_t"></div>
	<div class="fobox_fej">Játék | <a href="#top" class="premium">A lap tetejére</a></div>
	<div class="fobox_fej_b"></div>
	<div class="fobox_tartalom_t"></div>
	<div class="fobox_tartalom">
				<div style="display:flex; flex-direction:row;">
			<div style="border-right: 4px solid #e0e0e0;width: 50%;">
				Staff által jelölt
				<br><br><br>
				<div style="position:relative; overflow:hidden; text-align:center">
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3435482" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/lLBo-76KfPpSw_vW?27813153" title="Marvels_Spider-Man_Miles_Morales-FLT" width="120" height="176" border="1"></a>
										<br><br>
				PC ISO: 39.12 GiB<br>
				Seederek: 580<br>					Leecherek: 6				</div>
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3369709" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/WA4Aw772iEMi2Bjz?27621483" title="F1_22-Razor1911" width="120" height="176" border="1"></a>
										<br><br>
				PC ISO: 44.17 GiB<br>
				Seederek: 360<br>					Leecherek: 3				</div>
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3101681" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/dG_ylQ-LaWGUq4x8?26760596" title="LEGO.Harry.Potter.Years.1-7.GoG.Classic-I_KnoW" width="120" height="176" border="1"></a>
										<br><br>
				PC ISO: 13.21 GiB<br>
				Seederek: 129<br>					Leecherek: 0				</div>
							<div style="float:left; width:33%">
				<br><br><a href="torrents.php?action=details&amp;id=3362472" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/DQ_2YlJDT69ioXyW?27596533" title="Depth of Extinction v55.3.0.56562 GOG" width="120" height="176" border="1"></a>
										<br><br>
				PC RIP: 303.05 MiB<br>
				Seederek: 4<br>					Leecherek: 0				</div>
							<div style="float:left; width:33%">
				<br><br><a href="torrents.php?action=details&amp;id=3163030" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/vjBNlmM6SRyf-BNx?26944154" title="Empire Earth Gold Edition v2.0.0.2974.25522 GOG" width="120" height="176" border="1"></a>
										<br><br>
				PC RIP: 658.77 MiB<br>
				Seederek: 110<br>					Leecherek: 1				</div>
							<div style="float:left; width:33%">
				<br><br><a href="torrents.php?action=details&amp;id=3176798" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/jg_mpnRviRwfle8z?26990706" title="Baldur's Gate - Enhanced Edition v2.6.5.0.45309 GOG" width="120" height="176" border="1"></a>
										<br><br>
				PC RIP: 5.15 GiB<br>
				Seederek: 18<br>					Leecherek: 1				</div>
						</div>
		<br><br><br>
		Legaktívabb torrentek				<br><br><br>
		<div style="position:relative; overflow:hidden; text-align:center">
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3483995" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/YQed5lEJFxmFbBVg?27941598" title="Returnal-FLT" width="120" height="176" border="1"></a>
										<br><br>
				PC ISO: 52.88 GiB<br>
				Seederek: 1028<br>					Leecherek: 56						</div>
									<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3350959" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/l-XWzZRlfP1SvXOx?27570547" title="Far.Cry.6-EMPRESS" width="120" height="176" border="1"></a>
										<br><br>
				PC ISO: 78.42 GiB<br>
				Seederek: 1065<br>					Leecherek: 33						</div>
									<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3377457" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/AQ4p182MSR7fnXxw?27651913" title="Diablo II - Resurrected v1.3.70409 - FitGirl Repack" width="120" height="176" border="1"></a>
										<br><br>
				PC RIP: 30.33 GiB<br>
				Seederek: 447<br>					Leecherek: 27						</div>
										</div>
			</div>
			<div style="width:50%; padding-left: 16px;">
										<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483995" target="_blank">Returnal-FLT</a>
						<div style="min-width: 150px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483279" target="_blank">Wanted.Dead-FLT</a>
						<div style="min-width: 38.291457286432px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483981" target="_blank">Pharaoh_A_New_Era-Razor1911</a>
						<div style="min-width: 34.974874371859px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3453395" target="_blank">ELDEN RING Deluxe Edition v1.08.1</a>
						<div style="min-width: 29.170854271357px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=1645689" target="_blank">Minecraft v1.8</a>
						<div style="min-width: 26.532663316583px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3350959" target="_blank">Far.Cry.6-EMPRESS</a>
						<div style="min-width: 25.628140703518px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3299146" target="_blank">Red.Dead.Redemption.2.Build.1436.28-EMPRESS</a>
						<div style="min-width: 24.723618090452px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3473719" target="_blank">The Sims 4 2014 Deluxe Edition v1.94.147.1030</a>
						<div style="min-width: 22.613065326633px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3479930" target="_blank">UNCHARTED - Legacy of Thieves Collection v1.3.20900</a>
						<div style="min-width: 22.537688442211px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482753" target="_blank">Euro Truck Simulator 2 v1.46.2.17s</a>
						<div style="min-width: 22.462311557789px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=1983859" target="_blank">Grand.Theft.Auto.San.Andreas-HOODLUM</a>
						<div style="min-width: 21.859296482412px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3358122" target="_blank">God_of_War_v1.0.12-Razor1911</a>
						<div style="min-width: 21.708542713568px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3389740" target="_blank">The.Sims.4.Deluxe.Edition.2014.v1.90.375.1020.Multi18.HUN-Jocoka</a>
						<div style="min-width: 21.482412060302px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=1332393" target="_blank">Heroes.of.Might.and.Magic.3.Complete.Edition.GOG.Classic.ISO-RAiN</a>
						<div style="min-width: 21.256281407035px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3478638" target="_blank">The Quarry - Deluxe Edition v2023.01.31.</a>
						<div style="min-width: 20.72864321608px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=2020256" target="_blank">Need for Speed Most Wanted v1.3</a>
						<div style="min-width: 20.276381909548px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3410373" target="_blank">Cities Skylines - Deluxe Edition v1.15.0.F7</a>
						<div style="min-width: 20.201005025126px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=2128123" target="_blank">Need for Speed Underground 2 v1.2</a>
						<div style="min-width: 20.201005025126px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3384294" target="_blank">Marvels_Spider-Man_Remastered-FLT</a>
						<div style="min-width: 19.824120603015px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3484364" target="_blank">Pharaoh - A New Era v62416 GOG</a>
						<div style="min-width: 19.597989949749px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=2026645" target="_blank">Call of Duty 2 v1.3</a>
						<div style="min-width: 19.447236180905px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3476928" target="_blank">Cyberpunk 2077 v1.61.dlss3.61835 GOG</a>
						<div style="min-width: 17.788944723618px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3252894" target="_blank">Age.of.Empires.IV-CODEX</a>
						<div style="min-width: 17.261306532663px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3479470" target="_blank">HOT_WHEELS_UNLEASHED_Game_of_the_Year_Edition-Razor1911</a>
						<div style="min-width: 14.698492462312px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3149259" target="_blank">Call of Duty 4 - Modern Warfare</a>
						<div style="min-width: 14.246231155779px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3162706" target="_blank">Assassins.Creed.Valhalla.Repack-EMPRESS</a>
						<div style="min-width: 13.492462311558px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3380631" target="_blank">Dying.Light.2.Stay.Human-EMPRESS</a>
						<div style="min-width: 13.492462311558px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3375350" target="_blank">Forza Horizon 5 - Premium Edition v1.1.484.939.0 - FitGirl Repack</a>
						<div style="min-width: 13.417085427136px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=2159835" target="_blank">The Sims 3 - Complete Edition</a>
						<div style="min-width: 13.341708542714px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3434618" target="_blank">Farming Simulator 22 v1.8.2.0</a>
						<div style="min-width: 12.964824120603px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3435482" target="_blank">Marvels_Spider-Man_Miles_Morales-FLT</a>
						<div style="min-width: 12.889447236181px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482721" target="_blank">Stray v1.4.227</a>
						<div style="min-width: 12.738693467337px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=1815375" target="_blank">Grand.Theft.Auto.V-RELOADED</a>
						<div style="min-width: 12.13567839196px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3382349" target="_blank">Stray_v1.4-Razor1911</a>
						<div style="min-width: 11.984924623116px;background: #43A047;height: 16px;"></div>
					</div>
									</div>
		</div>
			</div>
	<div class="fobox_lab"></div>
</div>
<a name="zene">&nbsp;</a>    <div class="fobox_all">
	<div class="fobox_fej_t"></div>
	<div class="fobox_fej">Zene | <a href="#top" class="premium">A lap tetejére</a></div>
	<div class="fobox_fej_b"></div>
	<div class="fobox_tartalom_t"></div>
	<div class="fobox_tartalom">
				<div style="display:flex; flex-direction:row;">
			<div style="border-right: 4px solid #e0e0e0;width: 50%;">
				Staff által jelölt
				<br><br><br>
				<div style="position:relative; overflow:hidden; text-align:center">
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3459101" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/ZV_3Kor8IGkHwXpM?27877874" title="Faceless Man - Road to Darkness" width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;MP3: 109.26 MiB<br>
				Seederek: 56<br>					Leecherek: 0				</div>
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3315619" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/jO_lmy9RUd9FmBPY?27480627" title="Red Hot Chili Peppers - Unlimited Love" width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/en.gif" height="10">&nbsp;MP3: 169.05 MiB<br>
				Seederek: 114<br>					Leecherek: 1				</div>
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3309259" target="_blank"><img src="https://static.ncore.pro/static/images/nincs_borito.png" title="Nightwish-Walking_In_The_Air_-_The_Greatest_Ballads-(SPI409CD)-REMASTERED-CD-FLAC-2011-c05" width="120" height="176"></a>					<br><br>
				<img src="https://static.ncore.pro/static/images/flags/en.gif" height="10">&nbsp;Lossless: 467.44 MiB<br>
				Seederek: 17<br>					Leecherek: 2				</div>
							<div style="float:left; width:33%">
				<br><br><a href="torrents.php?action=details&amp;id=762011" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/7NBMyzo0aWkU0_Pm" title="Neo - Kontroll OST" width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;Lossless: 241.38 MiB<br>
				Seederek: 20<br>					Leecherek: 2				</div>
							<div style="float:left; width:33%">
				<br><br><a href="torrents.php?action=details&amp;id=1475272" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/DQ_2ZqxWT69io_yW" title="Linkin Park - Greatest Hits" width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/en.gif" height="10">&nbsp;MP3: 366.12 MiB<br>
				Seederek: 199<br>					Leecherek: 2				</div>
							<div style="float:left; width:33%">
				<br><br><a href="torrents.php?action=details&amp;id=3211830" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/vjBNl-LghRyf-BNx?27126213" title="Dua Lipa - Future Nostalgia - The Moonlight Edition" width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/en.gif" height="10">&nbsp;Lossless: 347.03 MiB<br>
				Seederek: 75<br>					Leecherek: 0				</div>
						</div>
		<br><br><br>
		Legaktívabb torrentek				<br><br><br>
		<div style="position:relative; overflow:hidden; text-align:center">
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3458793" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/K9Xg6o20uZduA_D2?27876650" title="Rádiókabaré - Buek 2023" width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;MP3: 492.17 MiB<br>
				Seederek: 448<br>					Leecherek: 42						</div>
									<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3451203" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/1zBPgjboF89HLXJn?27854889" title="Rádiókabaré 2022" width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;MP3: 2.76 GiB<br>
				Seederek: 192<br>					Leecherek: 32						</div>
									<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3479374" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/vj_NPlo3aRyf-4Nx?27929907" title="Rádió 1 - Balázsék 2023.02.07." width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;MP3: 337.94 MiB<br>
				Seederek: 50<br>					Leecherek: 26						</div>
										</div>
			</div>
			<div style="width:50%; padding-left: 16px;">
										<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482341" target="_blank">Rádiókabaré - 2023.02.11.</a>
						<div style="min-width: 150px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3484013" target="_blank">Kárpátia - Napkelettől napnyugatig</a>
						<div style="min-width: 146.56488549618px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482339" target="_blank">Rádiókabaré - 2023.02.04.</a>
						<div style="min-width: 143.51145038168px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481668" target="_blank">Hungarica - Szavazz magadra!</a>
						<div style="min-width: 100.38167938931px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481227" target="_blank">Rádió 1 - Balázsék 2023.02.10.</a>
						<div style="min-width: 97.328244274809px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481740" target="_blank">Triász - Triász V. - Elveszett évek</a>
						<div style="min-width: 94.656488549618px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483929" target="_blank">Kárpátia - Napkelettől napnyugatig (2023)[FLAC]-Naftamusic</a>
						<div style="min-width: 91.984732824427px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3351779" target="_blank">VA_-_Shazam_Summer_Top_50_Hungary_Chart-Web-2022-JKoop</a>
						<div style="min-width: 90.076335877863px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482381" target="_blank">Depeche Mode - Ghosts Again (Promo)</a>
						<div style="min-width: 75.572519083969px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3480080" target="_blank">In Flames - Foregone - Limited Edition</a>
						<div style="min-width: 72.519083969466px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482754" target="_blank">VA_-_DJs_Top_100_Based_On_The_DJ_Mag_Winter_Chart-Web-2023-JKoop</a>
						<div style="min-width: 71.374045801527px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481827" target="_blank">Ámokfutók - Szállj velem!</a>
						<div style="min-width: 69.847328244275px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482805" target="_blank">Rádió 1 - Balázsék 2023.02.13.</a>
						<div style="min-width: 65.267175572519px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481120" target="_blank">Ámokfutók - Szállj velem! (2023)[FLAC]-Naftamusic</a>
						<div style="min-width: 62.977099236641px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483393" target="_blank">Rádió 1 - Balázsék 2023.02.14.</a>
						<div style="min-width: 61.450381679389px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3443712" target="_blank">VA_-_Shazam_Winter_Top_50_Hungary_Chart-Web-2022-JKoop</a>
						<div style="min-width: 60.687022900763px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483922" target="_blank">Rádió 1 - Balázsék 2023.02.15.</a>
						<div style="min-width: 57.251908396947px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3480490" target="_blank">Rádió 1 - Balázsék 2023.02.09.</a>
						<div style="min-width: 53.81679389313px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481395" target="_blank">Linkin_Park-Lost-SINGLE-WEB-2023-ENRiCH</a>
						<div style="min-width: 53.053435114504px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481010" target="_blank">VA_-_Bravo_Hits_Vol.120-2CD-2023-MOD</a>
						<div style="min-width: 48.473282442748px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=2355122" target="_blank">Bagossy.Brothers.Company-Veled.Utazom.2019.WEB.RetReiveR</a>
						<div style="min-width: 47.709923664122px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3480511" target="_blank">Depeche_Mode-Ghosts_Again-SINGLE-WEB-2023-ENRiCH</a>
						<div style="min-width: 47.328244274809px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3480607" target="_blank">Depeche Mode - Ghosts Again - Single</a>
						<div style="min-width: 42.36641221374px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482521" target="_blank">Ennio Morricone - Love Songs &amp; Romantic Scores OST</a>
						<div style="min-width: 42.36641221374px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3479826" target="_blank">Rádió 1 - Balázsék 2023.02.08.</a>
						<div style="min-width: 41.221374045802px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3302061" target="_blank">VA_-_Shazam_Spring_Top_50_Hungary_Chart-Web-2022-JKoop</a>
						<div style="min-width: 40.458015267176px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481813" target="_blank">Counter Clockwise - Szó Szegik, Vagy Bennszakad</a>
						<div style="min-width: 40.076335877863px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3479626" target="_blank">VA_-_Bravo_Hits_Vol._120-Web-2023-JKoop</a>
						<div style="min-width: 39.69465648855px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=64225" target="_blank">VA - A nagy retro mix album</a>
						<div style="min-width: 39.69465648855px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482918" target="_blank">From The Sky - Antarktika</a>
						<div style="min-width: 37.786259541985px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481239" target="_blank">Retro Rádió - Bochkor 2023.02.10.</a>
						<div style="min-width: 37.022900763359px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3484497" target="_blank">Rádió 1 - Balázsék 2023.02.16.</a>
						<div style="min-width: 35.87786259542px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3152042" target="_blank">Bagossy Brothers Company - Fordul a világ</a>
						<div style="min-width: 35.114503816794px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3479374" target="_blank">Rádió 1 - Balázsék 2023.02.07.</a>
						<div style="min-width: 33.969465648855px;background: #43A047;height: 16px;"></div>
					</div>
									</div>
		</div>
			</div>
	<div class="fobox_lab"></div>
</div>
<a name="program">&nbsp;</a>    <div class="fobox_all">
	<div class="fobox_fej_t"></div>
	<div class="fobox_fej">Program | <a href="#top" class="premium">A lap tetejére</a></div>
	<div class="fobox_fej_b"></div>
	<div class="fobox_tartalom_t"></div>
	<div class="fobox_tartalom">
				<div style="display:flex; flex-direction:row;">
			<div style="border-right: 4px solid #e0e0e0;width: 50%;">
				Staff által jelölt
				<br><br><br>
				<div style="position:relative; overflow:hidden; text-align:center">
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3272912" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/7kBjYGnRIAwS2XwD?27334857" title="VidiCable.v1.10-F4CG" width="120" height="176" border="1"></a>
										<br><br>
				PC RIP: 110.34 MiB<br>
				Seederek: 33<br>					Leecherek: 0				</div>
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3166166" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/j2eEmLGpHPmSrezv?26955006" title="AVS4YOU.Software.AIO.Installation.Package.v5.0.5.167-F4CG" width="120" height="176" border="1"></a>
										<br><br>
				PC RIP: 288.79 MiB<br>
				Seederek: 32<br>					Leecherek: 0				</div>
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3065583" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/vwBZq-2OtN3tzXNl?26620275" title="[Android] YAZIO PRO v6.7.0" width="120" height="176" border="1"></a>
										<br><br>
				Mobil: 10.38 MiB<br>
				Seederek: 36<br>					Leecherek: 0				</div>
							<div style="float:left; width:33%">
				<br><br><a href="torrents.php?action=details&amp;id=3065126" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/kGeOjLdmi8RH1_WM?26618541" title="FL Studio Producer Edition v20.7.2" width="120" height="176" border="1"></a>
										<br><br>
				PC RIP: 1.01 GiB<br>
				Seederek: 399<br>					Leecherek: 0				</div>
							<div style="float:left; width:33%">
				<br><br><a href="torrents.php?action=details&amp;id=3060285" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/L9_kq1oZSwZFlBZl?26602144" title="[Android] Bud Spencer and Terence Hill - Slaps And Beans v1.05" width="120" height="176" border="1"></a>
										<br><br>
				Mobil: 175.85 MiB<br>
				Seederek: 47<br>					Leecherek: 0				</div>
							<div style="float:left; width:33%">
				<br><br><a href="torrents.php?action=details&amp;id=3059793" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/K9Xgq6N9fZduAXD2?26600195" title="[Android] A Bosszú v1.1.0" width="120" height="176" border="1"></a>
										<br><br>
				Mobil: 2.63 MiB<br>
				Seederek: 16<br>					Leecherek: 0				</div>
						</div>
		<br><br><br>
		Legaktívabb torrentek				<br><br><br>
		<div style="position:relative; overflow:hidden; text-align:center">
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=1879446" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/7Y_Q50NbHkxFN4dn" title="Microsoft Office 2016 Professional Plus  x86 - x64 MSDN HUN VLSC" width="120" height="176" border="1"></a>
										<br><br>
				PC ISO: 1.73 GiB<br>
				Seederek: 3028<br>					Leecherek: 283						</div>
									<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=1808234" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/9zeV8yqys15iQeKb" title="Rosetta Stone TOTALe all languages packs 2015.04.26." width="120" height="176" border="1"></a>
										<br><br>
				PC RIP: 37.63 GiB<br>
				Seederek: 131<br>					Leecherek: 132						</div>
									<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3392178" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/9z4VgG7LS15iQXKb?27693435" title="iGO HERE 2022 Q2 Map" width="120" height="176" border="1"></a>
										<br><br>
				PC RIP: 6.93 GiB<br>
				Seederek: 801<br>					Leecherek: 82						</div>
										</div>
			</div>
			<div style="width:50%; padding-left: 16px;">
										<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=1879446" target="_blank">Microsoft Office 2016 Professional Plus  x86 - x64 MSDN HUN VLSC</a>
						<div style="min-width: 150px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3298528" target="_blank">Microsoft Office Professional Plus 2021 VL v2108 HUN x86-x64</a>
						<div style="min-width: 131.53266331658px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3467023" target="_blank">Microsoft Windows 10 Consumer Editions v22H2 HUN x64</a>
						<div style="min-width: 79.020100502513px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3300318" target="_blank">Microsoft.Windows.10.AIO.64bit.Hun.Build.19044.1566-fatebringer</a>
						<div style="min-width: 71.356783919598px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=2360275" target="_blank">Adobe.Photoshop.CC.2019.v20.0.4.Pre-Activated.HUN.x64-D.G</a>
						<div style="min-width: 65.577889447236px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3472658" target="_blank">Microsoft.Office.2021.Professional.Plus.VL.LTSC.2023.January.V3-fatebringer</a>
						<div style="min-width: 64.070351758794px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3470013" target="_blank">Adobe.Photoshop.CC.2023.24.1.1.238.Repack.Multilingual.x64-D.G</a>
						<div style="min-width: 61.05527638191px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3308077" target="_blank">Adobe.Acrobat.Pro.DC.2022.001.20085.Multilingual.x64-D.G</a>
						<div style="min-width: 60.175879396985px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3245680" target="_blank">DAEMON Tools Ultra v6.1.0.1723 x64 HUN</a>
						<div style="min-width: 54.773869346734px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=2362229" target="_blank">DAEMON Tools Ultra v5.4.0.894 x64 HUN</a>
						<div style="min-width: 49.120603015075px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3467027" target="_blank">Microsoft Windows 11 Consumer Editions v22H2 HUN x64</a>
						<div style="min-width: 48.366834170854px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3257937" target="_blank">Microsoft Windows 10 Enterprise LTSC 2021 HUN x64</a>
						<div style="min-width: 45.603015075377px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3457075" target="_blank">Total Commander v10.52 x86-x64</a>
						<div style="min-width: 39.447236180905px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3471095" target="_blank">Adobe.Acrobat.Pro.DC.2022.003.20314.Multilingual.x64-D.G</a>
						<div style="min-width: 37.060301507538px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3462337" target="_blank">IObit Driver Booster Pro v10.1.0.86 HUN</a>
						<div style="min-width: 36.683417085427px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3298527" target="_blank">Microsoft Office Professional Plus 2019 VL v1808 HUN x86-x64</a>
						<div style="min-width: 36.180904522613px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3451306" target="_blank">Adobe.Photoshop.CC.2023.24.1.0.166.Multilingual.x64-D.G</a>
						<div style="min-width: 33.417085427136px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=2998775" target="_blank">Microsoft Office 2016 Standard VLSC HUN x64</a>
						<div style="min-width: 32.286432160804px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3451242" target="_blank">Adobe.Lightroom.Classic.CC.2023.12.1.0.202212072312.Multilingual.x64-D.G</a>
						<div style="min-width: 32.1608040201px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3397073" target="_blank">CCleaner.Professional.Edition.v6.03.0.10002-CaviaR</a>
						<div style="min-width: 32.035175879397px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3451737" target="_blank">Adobe.Premiere.Pro.CC.2023.23.1.0.86.Multilingual.x64-D.G</a>
						<div style="min-width: 31.78391959799px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3164939" target="_blank">Autodesk AutoCAD v2022 x64 HUN</a>
						<div style="min-width: 31.532663316583px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3462815" target="_blank">Microsoft.Office.2016.Professional.Plus.VL.32bit.Hun.2023.January-fatebringer</a>
						<div style="min-width: 31.281407035176px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3392178" target="_blank">iGO HERE 2022 Q2 Map</a>
						<div style="min-width: 31.155778894472px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3484505" target="_blank">Adobe.Acrobat.Pro.DC.2022.003.20322.Multilingual.x64-D.G</a>
						<div style="min-width: 29.020100502513px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3339910" target="_blank">IObit Driver Booster Pro v9.3.0.209 HUN</a>
						<div style="min-width: 27.763819095477px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3102616" target="_blank">SketchUp Pro 2021 v21.0.339 x64</a>
						<div style="min-width: 27.261306532663px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3469644" target="_blank">Adobe.Illustrator.CC.2023.27.2.0.339.Multilingual.x64-D.G</a>
						<div style="min-width: 27.010050251256px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3194806" target="_blank">Total Commander v10.0 HUN-ENG x86-x64</a>
						<div style="min-width: 25.753768844221px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3398243" target="_blank">[MAC] Microsoft Office 2021 v16.64 VL</a>
						<div style="min-width: 25.376884422111px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3308448" target="_blank">Adobe.Photoshop.CC.2022.23.2.2.325.Multilingual.x64-D.G</a>
						<div style="min-width: 25.376884422111px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3300320" target="_blank">Microsoft.Windows.11.AIO.64bit.Hun.Build.22000.527-fatebringer</a>
						<div style="min-width: 23.743718592965px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3435735" target="_blank">Microsoft.Office.365.v2211.Build.16.0.15726.20202.x64-BTCRiSO</a>
						<div style="min-width: 23.492462311558px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3470817" target="_blank">Rarlab.Winrar.6.20.ENG-HUN.x86-x64-D.G</a>
						<div style="min-width: 22.738693467337px;background: #43A047;height: 16px;"></div>
					</div>
									</div>
		</div>
			</div>
	<div class="fobox_lab"></div>
</div>
<a name="ebook">&nbsp;</a>    <div class="fobox_all">
	<div class="fobox_fej_t"></div>
	<div class="fobox_fej">eBook | <a href="#top" class="premium">A lap tetejére</a></div>
	<div class="fobox_fej_b"></div>
	<div class="fobox_tartalom_t"></div>
	<div class="fobox_tartalom">
				<div style="display:flex; flex-direction:row;">
			<div style="border-right: 4px solid #e0e0e0;width: 50%;">
				Staff által jelölt
				<br><br><br>
				<div style="position:relative; overflow:hidden; text-align:center">
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3443594" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/PxXzdPAESMlT-4bM?27834681" title="BBC Wildlife 2022" width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/en.gif" height="10">&nbsp;&nbsp;eBook: 1.39 GiB<br>
				Seederek: 14<br>					Leecherek: 0				</div>
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3263198" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/jg_mMZAriRwfle8z?27303609" title="The Fishkeeper ZA 2021" width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/en.gif" height="10">&nbsp;&nbsp;eBook: 70.65 MiB<br>
				Seederek: 9<br>					Leecherek: 0				</div>
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=2873749" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/K9_g9zGQTZduAXD2?26113537" title="Dezső András - Maffiózók mackónadrágban" width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;eBook: 1.55 MiB<br>
				Seederek: 266<br>					Leecherek: 0				</div>
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3053125" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/o0B58o3AsoAFKX5y?26577620" title="Wilhelm Blatzheim - 100% Wurst" width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/en.gif" height="10">&nbsp;&nbsp;eBook: 14.18 MiB<br>
				Seederek: 6<br>					Leecherek: 0				</div>
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=2993148" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/-9XbqMG2S2vSwXmp?26382798" title="Paul Bocuse - A főzés magasiskolája" width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;eBook: 926.76 MiB<br>
				Seederek: 85<br>					Leecherek: 0				</div>
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=2971679" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/lZ_0268dazbFM4md?26302252" title="PC World 2019" width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;eBook: 410.66 MiB<br>
				Seederek: 101<br>					Leecherek: 0				</div>
						</div>
		<br><br><br>
		Legaktívabb torrentek				<br><br><br>
		<div style="position:relative; overflow:hidden; text-align:center">
							<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=587868" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/KoXq6y88uNqty_-g" title="Benny Lewis - Language Hacking Guide" width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;eBook: 77.43 MiB<br>
				Seederek: 254<br>					Leecherek: 19						</div>
									<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=3276864" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/DpeGM86qTK3ugX6Y?27347771" title="Ludas Matyi 1867-1993." width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;eBook: 58.47 GiB<br>
				Seederek: 88<br>					Leecherek: 12						</div>
									<div style="float:left; width:33%">
				<a href="torrents.php?action=details&amp;id=1031296" target="_blank"><img src="https://nc-img.cdn.l7cache.com/covers/EdeYZgjvCDrCLX-x" title="George R. R. Martin - Királyok csatája - A Tűz és Jég dala II." width="120" height="176" border="1"></a>
										<br><br>
				<img src="https://static.ncore.pro/static/images/flags/hu.gif" height="10">&nbsp;&nbsp;eBook: 10.75 MiB<br>
				Seederek: 172<br>					Leecherek: 9						</div>
										</div>
			</div>
			<div style="width:50%; padding-left: 16px;">
										<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481614" target="_blank">PC World 2023 - 02.</a>
						<div style="min-width: 150px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483394" target="_blank">Jean-Francois Marmion - A hülyeség pszichológiája</a>
						<div style="min-width: 93.597063621533px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481774" target="_blank">Dr. Krasznay Csaba - Kiberbiztonság a XXI. században</a>
						<div style="min-width: 92.862969004894px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482551" target="_blank">Natalie Barelli - Az ​álláshirdetés</a>
						<div style="min-width: 70.962479608483px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481612" target="_blank">Various Authors - Életbátorság</a>
						<div style="min-width: 63.621533442088px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481839" target="_blank">Michael Lewis - Harcos ​pszichológusok</a>
						<div style="min-width: 52.120717781403px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483077" target="_blank">Sarah Pearse - A pihenőhely</a>
						<div style="min-width: 49.184339314845px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483118" target="_blank">L. J. Shen - A szélhámos</a>
						<div style="min-width: 48.205546492659px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482110" target="_blank">VA - Galaktika 2022</a>
						<div style="min-width: 46.492659053834px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482961" target="_blank">Janikovszky Éva - Naplóm</a>
						<div style="min-width: 44.535073409462px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482989" target="_blank">Sven Felix Kellerhoff - NSDAP</a>
						<div style="min-width: 43.433931484502px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481851" target="_blank">Aurélie Godefroy - Élj ​úgy, mint Coco Chanel</a>
						<div style="min-width: 41.84339314845px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482980" target="_blank">Su Tong - Rizs</a>
						<div style="min-width: 41.109298531811px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481220" target="_blank">Nora Roberts - Az ​Oromház rejtélye</a>
						<div style="min-width: 40.497553017945px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482542" target="_blank">Jennifer Lynn Barnes - Az ​örökség kulcsa</a>
						<div style="min-width: 40.008156606852px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483285" target="_blank">Megan Miranda - Az ​utolsó vendég</a>
						<div style="min-width: 40.008156606852px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482387" target="_blank">Maya Banks - Szürke árnyékok</a>
						<div style="min-width: 38.784665579119px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3480986" target="_blank">Corinne Michaels - Vár ​ránk a pillanat</a>
						<div style="min-width: 38.662316476346px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482789" target="_blank">Royd Tolkien - Lyukas ​bakancslista</a>
						<div style="min-width: 38.295269168026px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483328" target="_blank">Ambrose Parry - A ​halál művészete</a>
						<div style="min-width: 37.805872756933px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482524" target="_blank">Rachael English - Papírkarkötő</a>
						<div style="min-width: 36.827079934747px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481842" target="_blank">Corinne Michaels - Vár ​ránk a pillanat</a>
						<div style="min-width: 36.827079934747px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482901" target="_blank">Papp Gergő - Pimaszúr, Kajdi Csaba - A nagy Cyla-sztori</a>
						<div style="min-width: 34.991843393148px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483348" target="_blank">Sharon Cameron - Kékmadár</a>
						<div style="min-width: 33.646003262643px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482048" target="_blank">Javier Marías - A ​szívem fehér</a>
						<div style="min-width: 33.646003262643px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482784" target="_blank">Bridget Asher - A ​férjem szeretői</a>
						<div style="min-width: 33.278955954323px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483899" target="_blank">Andrea Bartz - Itt sem voltunk</a>
						<div style="min-width: 32.911908646003px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483986" target="_blank">Zachary Reed - Phoenix ​– A bandavezér</a>
						<div style="min-width: 32.667210440457px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3484002" target="_blank">Lorraine Heath - Hercegné ​kerestetik</a>
						<div style="min-width: 31.566068515498px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3482515" target="_blank">Bengt Börjeson - Terápia</a>
						<div style="min-width: 31.076672104405px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483376" target="_blank">Kellei György - Az ​elátkozott villa</a>
						<div style="min-width: 30.831973898858px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3483873" target="_blank">Robin Sloan - Penumbra ​úr nonstop könyvesboltja</a>
						<div style="min-width: 30.831973898858px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481517" target="_blank">Tabea Koenig - A párizsi maszkkészítő</a>
						<div style="min-width: 30.831973898858px;background: #43A047;height: 16px;"></div>
					</div>
											<div style="height:24px;display:flex;flex-direction:row;align-items: center;justify-content: space-between;">
						<a style="line-height: 24px;text-overflow: ellipsis;overflow: hidden;white-space: nowrap;margin-right: 16px;" href="torrents.php?action=details&amp;id=3481592" target="_blank">Ferdinand von Schirach - Collini ​nem beszél</a>
						<div style="min-width: 30.709624796085px;background: #43A047;height: 16px;"></div>
					</div>
									</div>
		</div>
			</div>
	<div class="fobox_lab"></div>
</div>
	</div>
</div>
</div>
`
