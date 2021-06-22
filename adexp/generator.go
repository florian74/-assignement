package adexp

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var Addresses = []string{
	"EHAAZQZX",
	"EGZYTTTE",
	"EGLLZTZP",
	"EGLLZTZR",
}

var Airports = []string{
	"EGLL",
	"EHAM",
	"LSZH",
	"LSGG",
}

var Aircrafts = []string{
	"A380",
	"A19N",
	"EFAN",
	"CONC",
	"B741",
}

var RoutePoints = []string{
	"N0402F270",
	"BPK",
	"UM185",
	"CLN",
	"UL620",
	"REDFA/N0390F230",
	"N0433F350",
	"LIFFY5A",
	"LIFFY",
	"UL975",
	"WAL",
	"M16",
	"DOLAS",
	"L603",
	"LAMSO",
}

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var digits = []rune("0123456789")

func randSeq(n int, runes []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func manyOf(slice []string, amount int) []string {
	many := make([]string, amount)
	for i := 0; i < amount; i++ {
		many[i] = slice[rand.Intn(len(slice))]
	}
	return many
}

func anyOf(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

func genAircraft() string {
	return randSeq(3, letters) + randSeq(3, digits)
}

func genIFplIpd() string {
	return randSeq(2, letters) + randSeq(8, digits)
}

func someDate() time.Time {
	return time.Date(2019+rand.Intn(2),
		time.Month(rand.Intn(11)+1),
		rand.Intn(20),
		rand.Intn(23),
		rand.Intn(59),
		0,
		0,
		time.Local)
}

func durationSince(departure time.Time) time.Duration {
	elapsed := time.Since(departure)
	if elapsed > 0 {
		return elapsed
	}
	return 0
}

type Generator struct {
	sync.RWMutex
	baseFLP *Fpl
}

func NewGenerator() *Generator {
	return &Generator{
		RWMutex: sync.RWMutex{},
		baseFLP: &Fpl{
			AdexpPsg: AdexpPsg{
				Title: "IFPL",
				Addr:  nil,
			},
		},
	}
}

func (g Generator) Next() []byte {
	g.randomizeFPL()
	return g.toJson()
}

func (g Generator) randomizeFPL() {
	g.Lock()
	defer g.Unlock()
	day := someDate()
	dayBefore := day.Add(-time.Second * 3600 * (24 + time.Duration(rand.Intn(5))))
	elapsed := durationSince(day)
	elapsedHours := time.Duration(elapsed.Seconds()) * time.Second / time.Hour
	elapsedMinutes := (time.Duration(elapsed.Seconds()) * time.Second % time.Hour) / time.Minute

	g.baseFLP.Addr = manyOf(Addresses, rand.Intn(9)+1)
	g.baseFLP.Ades = anyOf(Airports)
	g.baseFLP.Adep = anyOf(Airports)
	g.baseFLP.Arcid = genAircraft()
	g.baseFLP.Arctyp = anyOf(Aircrafts)
	g.baseFLP.Ceqpt = "SRGWY"
	g.baseFLP.Eobd = fmt.Sprintf("%02d%02d%02d", day.Year(), day.Month(), day.Day())
	g.baseFLP.Eobt = fmt.Sprintf("%02d%02d", day.Hour(), day.Minute())
	g.baseFLP.Filtim = fmt.Sprintf("%02d%02d%02d", dayBefore.Year(), dayBefore.Month(), dayBefore.Day())
	g.baseFLP.IfplId = genIFplIpd()
	g.baseFLP.Origin = "-NETWORKTYPE AFTN -FAC " + anyOf(Addresses)
	g.baseFLP.Seqpt = "C"
	g.baseFLP.Wktrc = "M"
	g.baseFLP.Opr = "ABC"
	g.baseFLP.Pnb = "B2"
	g.baseFLP.Reg = "GAAPO"
	g.baseFLP.Rmk = randSeq(rand.Intn(20), letters)
	g.baseFLP.Rvr = 200
	g.baseFLP.Sel = "DSGL"
	g.baseFLP.Src = "FPL"
	g.baseFLP.Ttleet = fmt.Sprintf("%02d%02d", elapsedHours, elapsedMinutes)
	g.baseFLP.Rfl = "F" + randSeq(3, digits)
	g.baseFLP.Speed = "N0" + randSeq(3, digits)
	g.baseFLP.Fltrul = "I"
	g.baseFLP.Fltyp = "S"
	g.baseFLP.Route = manyOf(RoutePoints, rand.Intn(5)+2)
}

func (g Generator) toJson() []byte {
	g.RLock()
	defer g.RLock()
	msg, err := json.Marshal(g.baseFLP)
	if err != nil {
		panic(err)
	}
	return msg
}
