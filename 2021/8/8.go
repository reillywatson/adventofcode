package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
0 := "abcefg"
1 := "cf"
2 := "acdeg"
3 := "acdfg"
4 := "bcdf"
5 := "abdfg"
6 := "abdefg"
7 := "acf"
8 := "abcdefg"
9 := "abcdfg"

 aaaa
b    c
b    c
 dddd
e    f
e    f
 gggg

 dddd
ef   ab
ef   ab
 efef
?    ab
?    ab
 ????


 acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
cdfeb fcadb cdfeb cdbaf

b = cd
e = cf
d = a

1 = gc
7 = cgb
4 = gecf
8 = fadegcb
*/

func perm(a []byte, f func([]byte), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

var nums = map[int]string{
	0: "abcefg",
	1: "cf",
	2: "acdeg",
	3: "acdfg",
	4: "bcdf",
	5: "abdfg",
	6: "abdefg",
	7: "acf",
	8: "abcdefg",
	9: "abcdfg",
}

func overlap(a, b string) int {
	n := 0
	for _, x := range a {
		for _, y := range b {
			if x == y {
				n++
			}
		}
	}
	return n
}

func main() {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		input := strings.Split(strings.Split(line, " | ")[0], " ")
		output := strings.Split(strings.Split(line, " | ")[1], " ")
		nums := map[int]string{}
		// first pass, try to extract numbers we know about
		for _, i := range input {
			num := -1
			switch len(i) {
			case 2:
				num = 1
			case 3:
				num = 7
			case 4:
				num = 4
			case 5:
				// 2
				// 3
				// 5
			case 6:
				// 0
				// 6
				// 9
			case 7:
				num = 8
			}
			if num > -1 {
				nums[num] = i
			}
		}
		for _, i := range input {
			num := -1
			switch len(i) {
			case 2:
				num = 1
			case 3:
				num = 7
			case 4:
				num = 4
			case 5:
				if overlap(i, nums[1]) == 1 && overlap(i, nums[4]) == 2 {
					num = 2
				} else if overlap(i, nums[1]) == 2 {
					num = 3
				} else {
					num = 5
				}
			case 6:
				if overlap(i, nums[1]) == 2 && overlap(i, nums[4]) == 3 {
					num = 0
				} else if overlap(i, nums[4]) == 3 {
					num = 6
				} else {
					num = 9
				}
			case 7:
				num = 8
			}
			if num > -1 {
				nums[num] = i
			}
		}
		s := ""
		for _, out := range output {
			found := false
			for k, v := range nums {
				if sortStr(v) == sortStr(out) {
					s += strconv.Itoa(k)
					found = true
					break
				}
			}
			if !found {
				panic(fmt.Sprintf("not found! out: %s nums: %v", out, nums))
			}
		}
		fmt.Println(output, s)
		n, _ := strconv.Atoi(s)
		sum += n
	}
	fmt.Println(sum)
}

func sortStr(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}

func main_partone() {
	nums := 0
	for _, line := range strings.Split(input, "\n") {
		output := strings.Split(strings.Split(line, " | ")[1], " ")
		for _, o := range output {
			var num int
			switch len(o) {
			case 2:
				num = 1
			case 3:
				num = 7
			case 4:
				num = 4
			case 5:
				// 2
				// 3
				// 5
			case 6:
				// 0
				// 6
				// 9
			case 7:
				num = 8
			}
			if num == 1 || num == 4 || num == 7 || num == 8 {
				nums++
			}
		}
	}
	fmt.Println(nums)
}

var testInput = `acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf`

var testInput2 = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

var input = `bceadfg ebcdaf gdaecf fgcaeb fdbea fcdea cbda begdf afb ba | acfbeg cbfgea gebacf ab
bg dfcba agedcb dfgab agb cafedgb gfdea ceafdg bfgade gbef | befg gafde dafbg cfbadeg
fg befdc afecdbg gbf gfcdab cfge cbdfeg efbdg baegd deafcb | bgf edfbac ebgdf fgb
agbec dbfac gfadeb dgaefcb febcda dcgf bcagd gda fgacdb gd | fgcd fadbce agd dag
bgfad cdfea cfedag bafedc efbda befcadg egbcdf bed be ecab | ecba bed edfab afecgd
dcebaf gf agbcfe agdbcf bfg eafbc gfea gbdce ceafgdb cegfb | faecbd ebcdaf fbg gf
cbgfad fgcae agfbd egd bdafecg ebad bgdcef ed egfda dgfaeb | eadb cagbfd ecgfa agcedfb
afb efadc ba cafbdge daeb cbgfad fbegc cgaedf ceafb fbcdea | fdcbea dfcbage daecf ba
cfbg efc ecdbf eafgbdc fc gfdeb fagdeb cbead ecgafd bdfecg | dbgef fdbceg adbce fegdba
faegdb daegcfb cbgeaf cgbfe fdg bcfda gd gced gcefdb dcbgf | fgd dg fbcad gcfeb
fcagbd ega aecbf eafgb fgcaedb dgeafb fgdab ge gadbec gfde | ge fdge ega egabf
bgcaf dfgea fdbga cafdbeg db adbc cgdabf edfgcb aecbfg bfd | cgdbaf gbadf gbcfae gcabf
bcgea baegdc efagb feacgdb edafb gefcdb fg gef cgfa febagc | gafeb efbga agcf gecbdaf
egbf fdeba fdg ecdafb ebgdfa cegad dbafceg gf agdfe bgacdf | dbceaf agdfe eadbf cfebgad
ged aebg dabegc gdeafc gacbd edfagbc agcdbf eg efbcd gbced | gdebc eabg geab gcadbe
acfe acbgfde gfbed bdcgaf gaebfc ebgaf cebadg cgfab eab ea | fgacb edagbc bgeaf ea
adbgc acbe aedcgb fgadce cb gcb dgcea ecdgfb dbefacg adbgf | fgbcde adcgefb bc eagcdfb
caeg fcbagd ce cgebda afebcdg dfbae acgbd begfdc ceb debca | febda ce ce cdgbfe
bad db agdfec fbacdg gedcbfa bcfd dgafb fdcga adgebc agbef | bad dfegcab bd cadfbge
cea efdbagc cdaf ca ebdfgc dbeac egbafc aebdg adcebf cfebd | dbegfc ceadb ac bfegadc
gfd fecbgda faebd gd fegdab adge ecfbg baefdc cbfadg gdfeb | cebfda gfd dg agdefcb
cade cbagf cfe cgdeabf ce dgfeab dabfe feacdb ebacf cdfebg | caed dcfageb gafcb efc
efbgda bdgc aebfc dc dfecga afdcb acd eadfgcb fabgdc dbafg | gdcb cdefag fdaegcb fegcad
efacbdg afeg adfcg bagcfd dcafe ebcda gdcbef fce ef dfcgae | afcgdb adfec efc aecbd
afbge adgbfc gef fbagd ecgdfb caedfbg eadg ebdafg ebcaf ge | edga afbge gdefcba dfbag
aegdc efbad cfbdea dcafe cfe cf gebfac fbdgae fcbd adcfgbe | ebadf dfcb dgaec abdef
acbfg adefcbg edagfb gadef dfb eabd bdfgce adgfec dgfba db | db fdegac aedcgf gfcade
aecgdb fdeba gfbc degafcb adbcg cf afdbgc cdfba fgdace cfd | gcbf fgcb cf fgecabd
cfaged badg bcfaeg dfg edbfg dabgfe dcfagbe dg debfc efbga | dfg agdb bfced cbdefag
ac dabeg cdfbe cbfa gedafc acd fadgecb bedcgf caedbf becda | egdcbf adc ac fgbaecd
fagbd ed cfegb gbcfde fdegb fbaegc fbcaedg efdbac dgec dfe | gafcbe bcefga edf gbdef
dfecg gb fcgbaed gcb dgbf cfgdbe dcbge cebad acegdf gafebc | cfebag eacdfg cgb fdgb
abegdfc fdbeag fdac gdefc cfgdeb gceafd ad cdaeg gad aecbg | dbefga da acdf gfbdec
cgaf gedab egafdc gcfbed cfdbae cge fabegcd gc caefd cgead | eacfdg fgdaebc ceg ecg
faceb fbaceg fgcdbe fdcga fbdca bfd db eafdcb aegfbcd dbae | cfadb cbaef gcadf ecbfa
adeb cbgdae cde fcegb ed fdgbac adgefbc bcegd fgcdae gadbc | gdfbca dce dgbac bacdgfe
geadc fegda cabe bcfgad gbecad ced ec fecgdab efcgbd adgcb | gdfae dbfcega edc ecd
acg eacdg ca dcegb fgdea abcdefg bcad adgbec gebdcf acgbfe | cga dgbec gac cfbeag
bdgfc dcg febcg gcedab dc bafdg deafbg fbdgac dgabefc cdfa | cd dc fbdcg fbgce
efdbag adgec fcgbda bcd dbecgfa ecfb fbdge bc dfegcb dcgeb | gdcfbe dbcaefg fdgeb bgefd
adebgcf dfbgac bdfeac dcega egdcfb febdc bad acbed ab beaf | befcgd afeb acedb dbefc
bead cdgfe bdacfg eb abfcd badgfec abcdfe feagbc cbe ebcfd | dcfbe bec baecfg eb
dcgafe dafcbg cfage df gedcb cfedg dfg fead bfeacg gafcbde | egdfca fgdec edfa aedf
ac cfbgaed adcbg bgcde eadc abdgf cebgfa becdga dgebcf bac | adcbg dbacgef egdbfc ca
ac acd cdage acfbgde dbacge degfcb gdefa bgca cbedg fcdeab | decga ac cda efbcdag
facbg gdaceb cg gecf agefb acdgfeb cafegb cag dcafb adbfge | agbedf fgec gc adcbf
bfdc cbged egcbdf egfdb bgf bdagce fgead bf ecgbaf cgebfad | fb bgf dbecgf bfdc
fbeacgd dfabcg ebdgc bafdg cfdabe ac abc acfg dgbac fbgade | bacgd gbdec adbfeg cdagb
dbfga dcf debgcaf egbdaf gdfac cf efabdc fbdgac gbcf cgeda | fc cgbf cf gfeabd
fdebac gcdb efabg aegbdc bcega gac dfgebca cadeb gc egcfad | gc acg edcba ecabd
cgfeb gcebfd fge gedb bgcfa fcdgea acdebf ge bcgafde defcb | gfdceb bcafed cbfgead fdecga
gdcfba cadge ac aefc decfga dgbafe fgeda dfceabg cebdg cag | acbfdge edgabf dagec gac
befacdg bdeg acefdg gdebac bcadgf ecadg fbeac cdb cdbea db | bd db gfdace bd
gcabf edfb dbacf acefgd daceb ebacfgd fd feabcd gbdcea fda | dacgfe gcebda cgaefd fgcbdae
bcadg fbcgd caeg ecbda gaebcd fbaced gedafb ag gad bgfcaed | aebfcd bfceda dga ga
egfad efab gdcba afcgde gfb cgdfeab gdafeb ebcgfd fadbg bf | dacbg dbafg fbg eafgd
efda bfdcg aecfgb becfda adb bcdfa caefb da baecgd dfcageb | cfabe dab bgfcdae gfedacb
bead dgabc begcfd cgfbea cabeg edbgca dacfg dfgabce dbc bd | fgdebc bdc egabcd adbe
be gaecdfb fcgdb dbfec dbfcag ebc cfade dcbefg gefcab egbd | fgabdc aegfcb degcbf bcdgfe
acedf agdb efdbag bdgfec afdcgbe dbfae cfagbe fba ab egfbd | acefgbd bedfa gdebaf ab
dab efcbd edafbcg cbfeda abgedf fbca bfcedg ab cabde dgcea | ebcgfd bagdfce dgcafeb ecbdf
fgdeac gda gabedfc gcefa da adec gedbf abdgcf egdfa cagbfe | gbcfae eacbgfd eacfbg ad
bdacegf ef begcda ecgf cefda fea bcdaf dfgeac gbfeda cdaeg | ef ecfg fcge dgcea
fabdge cgfad defga fc bdecgfa cbgad gcfe fdacbe fac cgeafd | cfa fegadb eadfbc gcef
dbcefa agbced acfbg efbagd dc bdacg ebdcagf decg gdaeb acd | cd badfce dc dabfec
adgb cgdaf ga eacgdfb cgdbf agfbcd gaf gcfdeb fbegac eafcd | abdcgf cgedfb bdcfg gbcfd
fgcadb dag da cbafg abgedc afcd gebfd dfagb efcbga adgbfce | cbfdage acdf afdc fcad
badce efbcda bcd dbfecg ebdaf bfcaedg afcd baecg cd dbfaeg | cd ecdbfa aebdc gbdfec
bfgdac cafbeg abgecd cgebdfa bgacf dc gfcd dcafb cbd dfbea | dc fdgc fabgc dc
bafgd egaf eadfgb dfa bcgefda cfdgb gaebd fa degacb dcabfe | agfedbc badecg cebadg adf
bc egdca cbdfae fcbgea aegcb fcgb egfbda cba fbega bacdgfe | cb egdbcfa abgce cb
bagef fgb cagfbe bg acegf adfgbc fedabgc dfeagc efbda gecb | bgf aebcfdg bg fgb
fdabeg egbfc fcgbda cgda gfbdc cedabf cdf abdgf bdcgefa dc | fbgdeca dbcafg fbgce cfegb
bacedf bdgca fdgbaec dbc dc bgacf bfdgca dfcg becagf gadeb | caedfb bgacf cbgafe acfebg
dfcebag def edbc aedgbf dacef bdcaf fcage ed cbfgda becfda | bdec edafbg caedfb dcafb
cfabg ebfagc aecdgb edgbcaf cgf fg fdabc ebgfdc gcbae fega | adbegc gcf bgeac bfacg
eadf beagdc dag bfdcg gafbec feadbg bgdaf gbaef bdgfcea ad | egfbcad ecadgb bgdaf acgfdbe
dabgce ecfdg fg fega dcefb bcedfga edgca fcbdag cgdeaf fcg | gaef cgfdaeb agfedc aegdc
gebd de dfega efbacd ead gdcfbae afcgd fabecg ebfag gfadbe | cgeabf bfcega cfabeg bgceaf
cfd gcedb egdbfc egdfca df ecgdbaf bcaef dacgeb fcedb bfgd | gdbec fd fd gbecd
agdec dbfcea gfcead gfbaecd cab gdba cgfeb ceabg eacbgd ab | eacdgf bfecad gedcfa abdg
cbgae baedcg dbcg bdefacg cb acb fadbce eacfg aebfdg egdab | aegbc dgcb cbdg ebfagd
fgbcd dfa gdab fecga afedbc dcfga fgacdb da bfcedg begcadf | facedb aedgbcf gdbcfa cdgfbe
aedfb fda bagfdc ad efbga bdfce gebdfc egfcbad ecbfad ecad | dgebcf dcae edca bdfgca
acfb cebgf daebgf gacfe fecdag gbf bcdeg gbadcfe bf gcbfea | aecgfd badfge gfcaed gcedb
abdg cfdeb ceafgd ebgdaf edg fcagbe dg bdegcfa gedfb aefgb | egbadcf bedgf abefcdg gcbaef
cgaefbd dfa cdag da fdecg dcgfae caefd ceafb gdefba cdbfge | dgcbef ad gefdc edbcgaf
bc dfcage geafb cgefdb bcefad fcaeb cfb agecbdf dcab acfed | ecfda bcda bdac cgdabfe
adbef bdcae dgeabf gfad gaedbfc ecbfdg fcebga af abf dfbeg | afebgc dfabe fba efbcga
agbc fcbea cfged bdfgaec cafbeg gacef ag dfbage aefdbc agf | cbga ga fegdc ag
gadbe acfg cgaed fcgbde edcfbag dca degcf ac dcafeg cdfabe | fgecad cgaf cgdefb ebadg
dcbegf gdaeb dcgfea dcbag fcba afgcd gbdcefa cb bcafdg dbc | cdbfeg bdc cbfa dcaebgf
cagbe fcdbg cfdgabe fdagec ecd agfbcd ebdf bfgdec de gbecd | bfcegd eagbc defb gcfbad
fbgdeca cged bgd bcdef cfgabd gd eagbf fdcbea gfbde ecbdfg | gd gdb cdeg dcfaeb
df bgacfd fagbc facbged bgfad fdecab dcgf bgdea afgceb adf | bdafgc gbfac acebgdf gafbec
adc abecfdg acbgfd ebafgd cfged ca acgdf adfbg dbcage bfac | bfgacd ca dafbg cebdga
debfa fgdbec edgcaf cbdfe dfa af bafdceg afcebd edabg bafc | edafb fa cefdab fa
bacfg gadecfb fda cgfeba cebda bfdg df fadcb fedcga acgfbd | abcfeg cdbaf daebc fecdgba
fdgec fegbc cb degacf cbg bgdcea fabge dbeafgc edbgfc bfdc | fdcb gbacfed cb bgc
afcb ab bad adcgf bgcfaed gabcd dgfacb faegcd gbafed gbedc | acbdg cbaf bcfa acgdb
fdebag ecdfga fagdb dcfaegb fb egfad fbcged efab cdgba dfb | efab cfbedg dbgafce aebf
edbag bgaecd fcagd bedafgc fagdb dbef badfeg bf gfacbe gbf | fdbe bf caefgb bdfage
adgfe dfec fd degacfb agedb afd cfdabg cfaeg eafcgb cafgde | afbgcd adf eabgfc cfde
dba cfgdbea da gbcae gaebfd aefd agdcbf defgb bedcgf agbde | egbda fedbgac edaf ad
ae fbgecda gabdc bagde defbg ebca fbdgac ecgbda cfadeg aeg | agdbe bdeag ecba age
dg gdc agcfe ebcgdf dfabcg fbcda fdgca gfeabcd bfaecd dbga | fbdegc dg fcbdga gd
fdebga bdgfa dabecfg bcedg cda ac acgbd abdgfc eacfgd acbf | cda cfgdba fedcag ca
dbceag fcgab efgab agc cdbgfe bgcfd ca bfacgd bdcafeg fadc | ca fcbgd adcf abcfg
fdacbg bcde acfdge acdfb abcef gabfe befagdc fce ec aedfbc | adcfeg agdcbef acdfge ecadgbf
aegf bgcdeaf gbfcad dcaefg gdefc bfced fg ecabgd gcf gadec | cgaed cgf ecgfd cfg
becdg gbfca cebgfa faebcd ae fcadgb fgeadcb ecgab gfea aeb | gebca cagbe ecfgabd agef
dbc agcd eagbdcf bfecda caegfb gcabf dfbcg dc bdgcaf gbdef | cd bdc fbcgae bcgfea
dbcaeg cfag gfedcab bcefa beagf egfacb defbg ag gba dfacbe | faecbgd acfedb eacdfb baefg
ebcgafd badce abfeg efgabc adfeb adf begafd dfeg gfdcab fd | ecbda aebfcg dcfagb ebcad
cdgbef ecgdab fgacdeb cdbfg eadfbg dcg cabgf cd fecd bgdfe | abdecg edgbfa fdbge gadbef
bfagd dfge acgbe bde dacbfg ebgda cefdbga fbcade de dafbeg | dafebc gbdeafc dfeg fedg
gadefbc gdfabc db dgbaf acbd agdfe bcgfa decbfg cfgeba dbf | gdbecfa bcda agfbec adgef
deg fgcd gd ecabg eagcd edcfa fgcdbea efbdca edacfg gdbfea | edagcf eacbdf daceg gde
cdfegb gcae dfabge begcf ae cabfeg abecf bea dcafb adfcgbe | ae cebgdfa gcbafe ae
acgfb da dga bcgad dcfa ebdafg bdcefag debgc bfdgca bgfeac | dga dag da abgfc
dgace gfc gfab gdcfa dbfegc bafcgd gdefbac fcadb bdfcae gf | gcf gecad fgacbde bdagfc
gabfce aegfd ac cgafe abcg eca debcfa fegcbda cgdfbe gecfb | ca aec ca gaefdcb
bcefa efdbc ed fde gdcfb dgabfec dbae adfebc ecgfab gcadfe | fed bafec baefc ebfacd
dagbf fb fgdeac bcdag fecbdg aefdg efab bfg gbaefd bfgcaed | efab afdbge gdbfa dcabg
gcdefa cdbgea gd efdab edg fegbdac fagce bgface dgcf eadgf | cdeafbg cagfde faebcgd gd
dfbgce edfgc deacgf fdbgeac befd acfbgd gbd gcdbe eabgc bd | dgecb egcdf eagcb fdgec
fdcbeg agbcd facdge aefb gedba dea fgbead dgacfeb ebdgf ea | dgeab afebgd dbgfec dae
dbfeag bfgca dfac abdfcg baf gceba bgdcf fa cfebgd afgbcde | afdc adfc bgdefac abf
fedbga cgdbef afdcgeb bdefg ae bdacg gdbea adef eba egabcf | ea dgecbaf gdabe fadbecg
daegcb bgc ecfdb adgfb egfc dcfgb cg efacdb cegfdb gcafedb | gadbcfe cgdbf dbgcfe gdbaf
dagf bcdfg aebcdf afbdecg fgdcab ebacg afb egcfbd af bgfac | afdg af fcgdb afb
dafbe dbfg agedf fg fabgce fga egbfad gcead ecadbf dfcgabe | cgaedfb edbgacf bcgfae bceadfg
cadbfeg abfde fabdc dfebag gdaecb afe daefgc gbdea efgb fe | eaf bgead gadbe befdga
gb cgba decbg abefdc egb bagecfd egcadb ecgfd begfad adcbe | bcga feabcd bg bg
dcbg bfdega dg gfceba cgeba caedf edgac ebgadc dag bceafdg | fceabg dg gdeac dg
ecadgb cdebgfa edabg fe efgd eagbfc bagdef dcfab fbade eaf | gbdaec efdab fae afe
begac fdbeag dce cdaf fdgceba fcdbea dgfecb cbeda cd dbeaf | ecfbda edcfab baecg edbagf
cdabef cbaef gedcafb dbfce fbgcde af fdab gfcead aef gcabe | edcafg fedbc cfdagbe cefdb
gadecb adecb fda fa abfc aefdc badegf fgcedab feacdb gdfec | fecdba fcba bdcea aefgdb
ec adgcb cdagfb abecdg fgbed aecd ecbgd ecfgdab gecfab bce | bec aecd ce adce
bgfdace fb fdabge faedgc cbfed dfcbea fbe gcbde fcab cfade | gbcedfa becafdg faecbgd bfcdea
aegd gbedacf fbgcda gebdc dcabeg gcbef gd baedc bdg cefabd | edbcfag baefcgd agbfdce gaed
efgcb ecbdag gdac ca gbade geabdfc bagec aebcdf ace gbaefd | eca egbcad aec cagd
ecabg abdgce efdbag bg cbdg cedfabg fbeca dgaec beg gafdec | gbdc aegdc baefc cgeba
ecfdgab fde afgd bcfeg ecdga gbdace fedabc cdefg df dgcafe | fd efdcg dfe cgfdae
bcdfeg gdfaec badcg db bdc gdfaebc egbac fdba cafgd dfcgba | cfadg afbd edcgaf agedcf
bcaeg beadg gbdaec dbe ebgcaf ed aedc dfbaecg gfdab gbefcd | bdegfc abgced gecba bagedcf
gdabf bfage bef bceg abgedfc cbgefa eb defbac acegdf eacgf | eb begaf becg fgecba
egdc dgfcb fdg fbegad bgecaf cafegbd gd fbcdeg dcabf cegbf | gd gacebf fcdab gbfcaed
cged dcbfae badgef gdcbefa aecgf aefcd aeg cadfeg eg gafbc | fedca degc faecd edgc
eadcfg ad acegf cefad eagdcfb adc ceagbf gfdabc fecdb geda | cfedbag cafegb gcfdeab ad
gfcdbe edbac cgbfae df cdbgfa dgecafb cfbge cefdb dcf egfd | cbfgeda df cdebf bgecf
facdb bfcadeg edafcb fbdcag bdgc edfcag cg eagfb gcf gcafb | fdceab gfeab dafcb gcdbaf
ebafc fecdbg fed fcbdea dfeab afdc gedba efagbc df gacedbf | cbfaeg fegdbac cgebfa df
fcd feabd cdge fdagcb fdaebcg cd egfbc fbedc dcgfeb afcgbe | dfebc fdbce dcf cefdb
gfaed degcfa cf caf eadcb aecbfg fbedagc cfdg afdec gbdeaf | dcfg cf afc caf
ebgcdf cfgeb dbcfa fgdebac gfea cea bfcaeg ae egadcb befac | efga ebcgaf cgedbf ecfab
efcbg afcbe afedbc fgedc bge bgac deabgf gb cabgedf fecgba | bgcaef egb gb cdabefg
dbeafg dg degbacf gdac becdf cebdg dgb gebcad cbeag fbcgea | egcbd cgaeb dg decgbfa
dcb dfcga adbfgc fdbeac acefdg bgfcd bgefd bacdfge cagb bc | adecfbg bc gabc fcgdab
bdf fdea df fdcegb befgcad cabef gabcfe facdb abecfd gabdc | eadgcbf fd fdea dafe
afgeb fdgcb gbedafc ecafbg ceafdb ad fdgab dgae fad eadgbf | afgdb gdae da ad
dbcfg gfed cebgf gbcae ebf cebdaf bdefcg afebgcd fe acbfdg | fe bdafegc ef bgeca
cebga ceg cgfbde aegdcb gc febgcad fgeab dcag ebcad cebdfa | gc gcda fgabe ecbdfa
gcdefa egbac defbcga da cegdba fdebg bcda fcgaeb eda dageb | gbaced begdf dae da
bgfd bcaeg bd ecadfb aebdfg gcaedbf dcefga geadb edb adfeg | db efdgba gfdb egdba
dgc defc dcagfeb gcfbd gdbef abfgc bagcde bgaefd dc egdbfc | febdg adgcbe dc agbfc
gbdafe fgb cebg ecbfad bcfae acfbg cfagdeb facegb gb cagdf | gfb bfceadg gb agbcf
edfab db acegfdb dbf fbaedg dgbe cgabef geabf dcafe abgcdf | db gdcfba bfadge fbgaed
eagbfc efca abefdgc dbgfce cgabd abfedg cf efgba cgafb cbf | cebgfa fc gbcad bcf
dfgbc fgc dbgca begdf gecadf geacbdf cdaegb fdagbc cbaf fc | cfab bdcgea cf gdcab
fdcbg cegfd dbafc gebf bg bfcdge gcdfae acgedb acfebdg bgd | gdfce efcdg eagdcf dbg
dg fcbdae bgad faecbdg agedf ecdbfg efabdg aegcf fabed egd | fgaedcb dafgbe cbedaf efdag
bcea gdaefb abd ecbdf fgcbed efdcab dfacb cafgd ba fdegcba | cgafd daegbf cbea ba
cfedbga dgce gafbe fdeag cdafbe dg gecafd dgf facbdg cdaef | dcefga fadeg bgeaf abfeg
cegdaf gbaec dceafb gfdb febag dafgebc fg abefd gef fbagde | gf dfegcba bdgaef efgdbca
fagcd aecdbf ebcadg edcbf gdbfc bg bgc dfcgeb bfeg afecbgd | dgcabe cfaedb gbc fdgbc
geafdb gacfed eadbg gce eagbdc begca ce bced cagfb gdaecfb | fcedag gecbad gcbfa ec
begcfa gaebfd fd cdgeb gafce afcd agfdce fdgec dgf gfacdbe | efgabd adfc efacg eacfgd
gefca de adefg eacd gbdfec ecfbag geabfcd efd dfegca bgadf | afcgbed agefcd bfdeagc dcea
ebadc cagedb adcfbeg eadfcb fcdbg cfea faebdg fe decbf fed | gcbfd cabedfg bcadef bedcf
ce cbfdg cbfge gacbef aceg cbgadfe gbefa feagbd cef dbfcae | egcbf dafegb ec cef
eagfd gc fgdace gcea dfacgb dgeafb fgc dbfce gdfbeac degfc | cagdfe cg aegc gc
bcade efgbda fgbed edcbagf cfegbd cbegd cgef gdc gc afbcdg | bdace befdg dcg cg
fdabce gcdebaf dceg bdgfae dag gacbf dg cbaed bcdag begcda | aegbdc gd cabedg gda
gf gfdbac fgecb gfc gdceab fbagce ecdbf efdgcba aefg eagbc | gbeca cefbg gcabe acedgb
dbcaf bcedf gdbfa cedafbg edbgfa agfecb afc bcafgd ca cagd | cafbd caf gcfadb ac
dbeafc dcaegb cd dcb cgdba dabecgf egadb dbagfe gbfca gced | cabgf gced dc dgce
bgce aebcfg agecdf bface efc febda gaebdcf ec bgcaf bacgfd | abfde ec efc cebg
gdeafc ce defcgb gbec fce fdbea cbdef fcdbga gbaecfd bdgcf | ce ecfgdb cfe fedcb
aegb fedag egf dabfeg dbcafeg defba fcbdae ge adgfc fgedcb | efg dcegfb edcafbg ge
debagf febgcd gfc dcfgba fbegc fc ebfgd cfed bagfedc aecbg | cfg cgf fgc gebac
gcfdbe bagfecd fcdbe eb fbeg bgadec cadbf gefdc dbe cafgde | deb eb egfbdac dcgaeb
dgfea bdefa efb eb edcb gcbafe efdcab dacfb bgcfad adcgbef | dbce ebdc dbgfca bafgec
becfgd eagcdb dceafb bdf cdbfa edbac efgdcba dcagf fb feba | ebfa fabe bcedaf bf
bdcaf cgfdeb cfeadg cgd gd dega dfabcge fcagd ecfgab fcaeg | gdc dcbfa eafcbdg gdcaf
cgedba edagf gb gfcdae gcfeabd dbg aedfbg gdebf edfbc abgf | dfabeg dfgea gdb fcabdge
agbefcd fgadb dfbage efbd gcbda adfge cbfgea fb fab dcfega | bf ebgdacf fedag fegda
efdbg fbdcge agdfe cabfdg ebf eb dfgcb dagcfeb ebcd cafegb | eb dfbegac bgfde fcdbg
ecfba egf dgbeaf dbgcfea gf egfcb egabcd cfgd gefcdb egdcb | gbfec dcfg fg fbgce`
