package adexp

type AdexpPsg struct {
	Title string
	Addr  []string
}

type Fpl struct {
	AdexpPsg
	Adep    string
	Ades    string
	Arcid   string
	Arctyp  string
	Ceqpt   string
	Eobd    string
	Eobt    string
	Filtim  string
	IfplId  string
	Origin  string
	Seqpt   string
	Wktrc   string
	Opr     string
	Pnb     string
	Reg     string
	Rmk     string
	Rvr     int
	Sel     string
	Src     string
	Ttleet  string
	Rfl     string
	Speed   string
	Fltrul  string
	Fltyp   string
	Route   []string
	Altrnt1 string
	Eetfir  []string
	Rtepts  []string
	Atsrt   []string
}
