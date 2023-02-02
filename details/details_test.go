package details

import (
	"git.okki.hu/garric/ngore/parse"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseDetails(t *testing.T) {

	doc := parse.MustParse(t, detailsHtml)

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

	t.Run("cover image", func(t *testing.T) {
		assert.Equal(t, "https://nc-img.cdn.l7cache.com/covers/7Y_Qz2gRTkxFNBdn?27818835", details.CoverImage)
	})

	t.Run("other images", func(t *testing.T) {
		assert.Equal(t, []string{
			"https://nc-img.cdn.l7cache.com/p0J_1xb3wtqCWB57?27818519",
			"https://nc-img.cdn.l7cache.com/pdGeyz8R7TWfDXx8?27818519",
			"https://nc-img.cdn.l7cache.com/rlLXo-m87t2hAevW?27818519",
		}, details.OtherImages)
	})

}

const detailsHtml = `
<div class="fobox_all" id="details1">
    <div class="fobox_fej_t"></div>
    <div class="fobox_fej">Torrent adatai</div>
    <div class="fobox_fej_b"></div>
    <div class="fobox_tartalom_t"></div>
    <div class="fobox_tartalom">
        <div class="torrent_reszletek_cim">Black.Adam.2022.RERiP.MA.WEBRip.x264.HuN-No1</div>
        <div class="torrent_reszletek_konyvjelzo">
            <a href="javascript:void(0);" id="konyvjelzo_ajax" onclick="konyvjelzo('3437646');"
               style="font-weight:normal;" title="Hozzáadás a könyvjelzőkhöz">[könyvjelzőkhöz adás]</a>
            <a href="torrents.php?action=addnews&amp;id=3437646&amp;getunique=8ffe570cf24f14b749ee44e32587c118"
               style="font-weight:normal;"
               title="A torrent ajánlása a hírfolyamon!">[hírfolyamra]</a>
            <a href="torrents.php?action=notification&amp;id=3437646" style="font-weight:normal;">[bejelentés]</a>
        </div>
        <center>
            <table border="0" cellspacing="5">
                <tbody>
                <tr>
                    <td><a class="fancy_groups" href="https://nc-img.cdn.l7cache.com/WlZe0pbrLi1CM4md?27818519"
                           rel="g3437646"><img
                            border="0" src="https://nc-img.cdn.l7cache.com/p0J_1xb3wtqCWB57?27818519"></a></td>
                    <td><a class="fancy_groups" href="https://nc-img.cdn.l7cache.com/p0J_1xb3wtDf9B57?27818519"
                           rel="g3437646"><img
                            border="0" src="https://nc-img.cdn.l7cache.com/pdGeyz8R7TWfDXx8?27818519"></a></td>
                    <td><a class="fancy_groups" href="https://nc-img.cdn.l7cache.com/pdGeyz8R7TNhqXx8?27818519"
                           rel="g3437646"><img
                            border="0" src="https://nc-img.cdn.l7cache.com/rlLXo-m87t2hAevW?27818519"></a></td>
                </tr>
                <tr>
                    <td align="center"><em>(720x302)</em></td>
                    <td align="center"><em>(720x302)</em></td>
                    <td align="center"><em>(720x302)</em></td>
                </tr>
                </tbody>
            </table>
        </center>
        <br>
        <center id="detailsad">
            <div class="banner" data-banner-id="279" data-id="281" data-zone-id="1,2,3,4"><a
                    href="http://www.ecigionline.com/?utm_source=ncore&amp;utm_medium=banner&amp;utm_campaign=2014JUN"
                    target="_blank"><img src="https://nc-img.cdn.l7cache.com/bimages/f449d9ba7fa0118f74973bd3d91d93cc.png"
                                         style="width:800px;height:100px;border:0;"></a>
            </div>
        </center>
        <!--<script type="text/javascript">
            $(document).ready(function(){
                $('#detailsad').html('<iframe id="detailsadb" src="http://yayo.hu/api/banner/800" style="width:800px;border:0;height:116px"></iframe>');
            });
        </script>-->
        <br>
        <div class="torrent_reszletek">
            <div class="torrent_col1">
                <div class="dt">Típus:</div>
                <div class="dd"><a href="torrents.php?csoport_listazas=osszes_film"
                                   title="Még több ebből a kategóriából">Film</a> &gt; <a
                        href="torrents.php?tipus=xvid_hun" title="Még több ebből a kategóriából">SD/HU</a></div>
                <div class="dt">Feltöltve:</div>
                <div class="dd">2022-11-22 10:59:51</div>
                <div class="dt">Feltöltő:</div>
                <div class="dd">
                    <span class="feltolto_szin">Anonymous</span>
                </div>
                <div class="dt">Hozzászólás:</div>
                <div class="dd">106 db</div>
            </div>
            <div class="torrent_col2">
                <div class="dt">Seederek:</div>
                <div class="dd"><a href="javascript:void(0);" onclick="peers('peers', '3437646', 1);" title="Peerlista">6855</a>
                </div>
                <div class="dt">Leecherek:</div>
                <div class="dd"><a href="javascript:void(0);" onclick="peers('peers', '3437646', 1);" title="Peerlista">556</a>
                </div>
                <div class="dt">Letöltve:</div>
                <div class="dd">+++++</div>
                <div class="dt">Sebesség:</div>
                <div class="dd">2.52 MiB/s (becsült)</div>
                <div class="dt">Méret:</div>
                <div class="dd">2.03 GiB (2177241852 bájt)</div>
                <div class="dt">NFO:</div>
                <div class="dd"><a href="javascript:void(0);" onclick="getNfo('nfocontent', '3437646', 1);"
                                   title="Mutat">Mutat</a></div>
                <div class="dt">Fájlok:</div>
                <div class="dd"><a href="javascript:void(0);" onclick="files('filelist', '3437646', 1);" title="Fájlok">3</a>
                </div>
                <!--<div class="dt">Mafab.hu</div>
        <div class="dd"><a onclick="getMafab('mafab', 'tt6443346', 0, 1);" href="javascript:void(0);">Mutat</a></div>-->
                <div class="dt">Más verziók:</div>
                <div class="dd"><a href="javascript:void(0);" onclick="other_versions('otherlist', '6443346', '3437646', 1);"
                                   title="Más verziók">Mutat</a></div>
            </div>
        </div>
        <div style="clear:both;"></div>
        <div class="download"><a href="torrents.php?action=download&amp;id=3437646&amp;key=7b200e994f68aa46bc4b55c9c85d9fc1"
                                 style="font-size:11px;text-decoration:underline;text-transform:uppercase;">Torrent
            letöltése</a></div>
        <div id="filelist"></div>
        <!-- <div id="mafab"></div>-->
        <div id="otherlist"></div>
        <div id="nfocontent"></div>
        <a name="peers"></a>
        <div id="peers"></div>
        <div id="subtitles"></div>
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
                                <td><a href="torrents.php?tags=akci%C3%B3"
                                       style="border-bottom:1px dotted white;">akció</a>, <a
                                        href="torrents.php?tags=kaland"
                                        style="border-bottom:1px dotted white;">kaland</a>, <a
                                        href="torrents.php?tags=fantasy"
                                        style="border-bottom:1px dotted white;">fantasy</a></td>
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
                                <td><a href="https://dereferer.link/?https://www.themoviedb.org/movie/436270"
                                       target="_blank">https://www.themoviedb.org/mov...</a></td>
                            </tr>
                            </tbody>
                        </table>
                    </td>
                </tr>
                </tbody>
            </table>
        </div>
        <div class="torrent_leiras proba42"></div>
        <div class="torrent_leiras"></div>
        <div class="download"><a href="torrents.php?action=download&amp;id=3437646&amp;key=7b200e994f68aa46bc4b55c9c85d9fc1"
                                 style="font-size:11px;text-decoration:underline;text-transform:uppercase;">Torrent
            letöltése</a></div>
    </div>
    <div class="fobox_lab"></div>
</div>
`
