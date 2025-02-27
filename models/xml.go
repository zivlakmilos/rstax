/* Copyright © 2025 Milos Zivlak

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package models

import (
	"encoding/xml"
)

type PodaciPoreskeDeklaracije struct {
	XMLName                   xml.Name                  `xml:"ns1:PodaciPoreskeDeklaracije"`
	XMLNs1                    string                    `xml:"xmlns:ns1,attr"`
	XMLXsi                    string                    `xml:"xmlns:xsi,attr"`
	PodaciOPrijavi            PodaciOPrijavi            `xml:"ns1:PodaciOPrijavi"`
	PodaciOPoreskomObvezniku  PodaciOPoreskomObvezniku  `xml:"ns1:PodaciOPoreskomObvezniku"`
	PodaciONacinuOstvarivanja PodaciONacinuOstvarivanja `xml:"ns1:PodaciONacinuOstvarivanjaPrihoda"`
	DeklarisaniPodaciOVrstama DeklarisaniPodaciOVrstama `xml:"ns1:DeklarisaniPodaciOVrstamaPrihoda"`
	Ukupno                    Ukupno                    `xml:"ns1:Ukupno"`
	Kamata                    Kamata                    `xml:"ns1:Kamata"`
	PodaciODodatnojKamati     string                    `xml:"ns1:PodaciODodatnojKamati"`
}

type PodaciOPrijavi struct {
	VrstaPrijave             int    `xml:"ns1:VrstaPrijave"`
	ObracunskiPeriod         string `xml:"ns1:ObracunskiPeriod"`
	DatumOstvarivanjaPrihoda string `xml:"ns1:DatumOstvarivanjaPrihoda"`
	Rok                      int    `xml:"ns1:Rok"`
	DatumDospelostiObaveze   string `xml:"ns1:DatumDospelostiObaveze"`
	DatumObracunaKamate      string `xml:"ns1:DatumObracunaKamate"`
}

type PodaciOPoreskomObvezniku struct {
	PoreskiIdentifikacioniBroj string `xml:"ns1:PoreskiIdentifikacioniBroj"`
	ImePrezimeObveznika        string `xml:"ns1:ImePrezimeObveznika"`
	UlicaBrojPoreskogObveznika string `xml:"ns1:UlicaBrojPoreskogObveznika"`
	PrebivalisteOpstina        int    `xml:"ns1:PrebivalisteOpstina"`
	JMBGPodnosiocaPrijave      string `xml:"ns1:JMBGPodnosiocaPrijave"`
	TelefonKontaktOsobe        string `xml:"ns1:TelefonKontaktOsobe"`
	ElectronPosta              string `xml:"ns1:ElektronskaPosta"`
	DrzavaRezidenstva          string `xml:"ns1:DrzavaRezidenstva"`
}

type PodaciONacinuOstvarivanja struct {
	NacinIsplate string `xml:"ns1:NacinIsplate"`
	TekuciRacun  string `xml:"ns1:TekuciRacun"`
}

type DeklarisaniPodaciOVrstama struct {
	PodaciOVrstama []PodaciOVrstama `xml:"ns1:PodaciOVrstamaPrihoda>ns1:PodaciOVrstama"`
}

type PodaciOVrstama struct {
	RedniBroj         int     `xml:"ns1:RedniBroj"`
	SifraVrstePrihoda string  `xml:"ns1:SifraVrstePrihoda"`
	BrutoPrihod       float64 `xml:"ns1:BrutoPrihod"`
	OsnovicaZaPorez   float64 `xml:"ns1:OsnovicaZaPorez"`
	ObracunatiPorez   float64 `xml:"ns1:ObracunatiPorez"`
	PorezZaUplatu     float64 `xml:"ns1:PorezZaUplatu"`
}

type Ukupno struct {
	FondSati                float64 `xml:"ns1:FondSati"`
	BrutoPrihod             float64 `xml:"ns1:BrutoPrihod"`
	OsnovicaZaPorez         float64 `xml:"ns1:OsnovicaZaPorez"`
	ObracunatiPorez         float64 `xml:"ns1:ObracunatiPorez"`
	PorezPlacenDrugojDrzavi float64 `xml:"ns1:PorezPlacenDrugojDrzavi"`
	PorezZaUplatu           float64 `xml:"ns1:PorezZaUplatu"`
	OsnovicaZaDoprinose     float64 `xml:"ns1:OsnovicaZaDoprinose"`
	PIO                     float64 `xml:"ns1:PIO"`
	Zdravstvo               float64 `xml:"ns1:ZDRAVSTVO"`
	Nezaposlenost           float64 `xml:"ns1:NEZAPOSLENOST"`
}

type Kamata struct {
	PorezZaUplatu       float64 `xml:"ns1:PorezZaUplatu"`
	OsnovicaZaDoprinose float64 `xml:"ns1:OsnovicaZaDoprinose"`
	PIO                 float64 `xml:"ns1:PIO"`
	Zdravstvo           float64 `xml:"ns1:ZDRAVSTVO"`
	Nezaposlenost       float64 `xml:"ns1:NEZAPOSLENOST"`
}
