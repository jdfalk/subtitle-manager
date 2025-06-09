package providers

import (
	"fmt"

	"subtitle-manager/pkg/providers/addic7ed"
	"subtitle-manager/pkg/providers/animekalesi"
	"subtitle-manager/pkg/providers/animetosho"
	"subtitle-manager/pkg/providers/assrt"
	"subtitle-manager/pkg/providers/avistaz"
	"subtitle-manager/pkg/providers/betaseries"
	"subtitle-manager/pkg/providers/bsplayer"
	"subtitle-manager/pkg/providers/embedded"
	"subtitle-manager/pkg/providers/generic"
	"subtitle-manager/pkg/providers/gestdown"
	"subtitle-manager/pkg/providers/greeksubs"
	"subtitle-manager/pkg/providers/greeksubtitles"
	"subtitle-manager/pkg/providers/hdbits"
	"subtitle-manager/pkg/providers/hosszupuska"
	"subtitle-manager/pkg/providers/karagarga"
	"subtitle-manager/pkg/providers/ktuvit"
	"subtitle-manager/pkg/providers/legendasdivx"
	"subtitle-manager/pkg/providers/legendasnet"
	"subtitle-manager/pkg/providers/napiprojekt"
	"subtitle-manager/pkg/providers/napisy24"
	"subtitle-manager/pkg/providers/nekur"
	"subtitle-manager/pkg/providers/opensubtitles"
	"subtitle-manager/pkg/providers/opensubtitlescom"
	"subtitle-manager/pkg/providers/opensubtitlesvip"
	"subtitle-manager/pkg/providers/podnapisi"
	"subtitle-manager/pkg/providers/regielive"
	"subtitle-manager/pkg/providers/soustitres"
	"subtitle-manager/pkg/providers/subdivx"
	"subtitle-manager/pkg/providers/subf2m"
	"subtitle-manager/pkg/providers/subs4free"
	"subtitle-manager/pkg/providers/subs4series"
	"subtitle-manager/pkg/providers/subscene"
	"subtitle-manager/pkg/providers/subscenter"
	"subtitle-manager/pkg/providers/subssabbz"
	"subtitle-manager/pkg/providers/subsunacs"
	"subtitle-manager/pkg/providers/subsynchro"
	"subtitle-manager/pkg/providers/subtitrarinoi"
	"subtitle-manager/pkg/providers/subtitriidlv"
	"subtitle-manager/pkg/providers/subtitulamos"
	"subtitle-manager/pkg/providers/supersubtitles"
	"subtitle-manager/pkg/providers/titlovi"
	"subtitle-manager/pkg/providers/titrariro"
	"subtitle-manager/pkg/providers/titulky"
	"subtitle-manager/pkg/providers/turkcealtyazi"
	"subtitle-manager/pkg/providers/tusubtitulo"
	"subtitle-manager/pkg/providers/tvsubtitles"
	"subtitle-manager/pkg/providers/whisper"
	"subtitle-manager/pkg/providers/wizdom"
	"subtitle-manager/pkg/providers/xsubs"
	"subtitle-manager/pkg/providers/yavka"
	"subtitle-manager/pkg/providers/yifysubtitles"
	"subtitle-manager/pkg/providers/zimuku"
)

var factories = map[string]func() Provider{
	"addic7ed":         func() Provider { return addic7ed.New() },
	"animekalesi":      func() Provider { return animekalesi.New() },
	"animetosho":       func() Provider { return animetosho.New() },
	"assrt":            func() Provider { return assrt.New() },
	"avistaz":          func() Provider { return avistaz.New() },
	"betaseries":       func() Provider { return betaseries.New() },
	"bsplayer":         func() Provider { return bsplayer.New() },
	"embedded":         func() Provider { return embedded.New() },
	"gestdown":         func() Provider { return gestdown.New() },
	"generic":          func() Provider { return generic.New() },
	"greeksubs":        func() Provider { return greeksubs.New() },
	"greeksubtitles":   func() Provider { return greeksubtitles.New() },
	"hdbits":           func() Provider { return hdbits.New() },
	"hosszupuska":      func() Provider { return hosszupuska.New() },
	"karagarga":        func() Provider { return karagarga.New() },
	"ktuvit":           func() Provider { return ktuvit.New() },
	"legendasdivx":     func() Provider { return legendasdivx.New() },
	"legendasnet":      func() Provider { return legendasnet.New() },
	"napiprojekt":      func() Provider { return napiprojekt.New() },
	"napisy24":         func() Provider { return napisy24.New() },
	"nekur":            func() Provider { return nekur.New() },
	"opensubtitlescom": func() Provider { return opensubtitlescom.New() },
	"opensubtitlesvip": func() Provider { return opensubtitlesvip.New() },
	"podnapisi":        func() Provider { return podnapisi.New() },
	"regielive":        func() Provider { return regielive.New() },
	"soustitres":       func() Provider { return soustitres.New() },
	"subdivx":          func() Provider { return subdivx.New() },
	"subf2m":           func() Provider { return subf2m.New() },
	"subscene":         func() Provider { return subscene.New() },
	"subs4free":        func() Provider { return subs4free.New() },
	"subs4series":      func() Provider { return subs4series.New() },
	"subscenter":       func() Provider { return subscenter.New() },
	"subssabbz":        func() Provider { return subssabbz.New() },
	"subsunacs":        func() Provider { return subsunacs.New() },
	"subsynchro":       func() Provider { return subsynchro.New() },
	"subtitrarinoi":    func() Provider { return subtitrarinoi.New() },
	"subtitriidlv":     func() Provider { return subtitriidlv.New() },
	"subtitulamos":     func() Provider { return subtitulamos.New() },
	"supersubtitles":   func() Provider { return supersubtitles.New() },
	"titlovi":          func() Provider { return titlovi.New() },
	"titrariro":        func() Provider { return titrariro.New() },
	"titulky":          func() Provider { return titulky.New() },
	"turkcealtyazi":    func() Provider { return turkcealtyazi.New() },
	"tusubtitulo":      func() Provider { return tusubtitulo.New() },
	"tvsubtitles":      func() Provider { return tvsubtitles.New() },
	"whisper":          func() Provider { return whisper.New() },
	"wizdom":           func() Provider { return wizdom.New() },
	"xsubs":            func() Provider { return xsubs.New() },
	"yavka":            func() Provider { return yavka.New() },
	"yifysubtitles":    func() Provider { return yifysubtitles.New() },
	"zimuku":           func() Provider { return zimuku.New() },
}

// Get returns a provider by name. OpenSubtitles requires an API key.
func Get(name, openSubtitlesKey string) (Provider, error) {
	if name == "opensubtitles" {
		return opensubtitles.New(openSubtitlesKey), nil
	}
	if f, ok := factories[name]; ok {
		return f(), nil
	}
	return nil, fmt.Errorf("unknown provider %s", name)
}
