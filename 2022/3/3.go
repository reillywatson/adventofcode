package main

import (
	"fmt"
	"strings"
)

func main() {
	total := 0
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i += 3 {
		group := []string{lines[i], lines[i+1], lines[i+2]}
		r := common(group)
		total += priority(r)
	}
	fmt.Println(total)
}

func main_partone() {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		compartments := []string{line[:len(line)/2], line[len(line)/2:]}
		seen := map[rune]bool{}
		for _, r := range compartments[0] {
			seen[r] = true
		}
		for _, r := range compartments[1] {
			if seen[r] {
				total += priority(r)
				break
			}
		}
	}
	fmt.Println(total)
}

func common(s []string) rune {
	seen := map[rune]bool{}
	for _, r := range s[0] {
		seen[r] = true
	}
	for r := range seen {
		seenAll := true
		for _, l := range s[1:] {
			seen := false
			for _, x := range l {
				if x == r {
					seen = true
					break
				}
			}
			if !seen {
				seenAll = false
				break
			}
		}
		if seenAll {
			return r
		}
	}
	panic("no common")
}

func priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r) - int('a') + 1
	}
	return int(r) - int('A') + 27
}

var testinput = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

var input = `GGVGlqWFgVfFqqVZGFlblJPMsDbbMrDMpDsJRn
LwzHtwdLHHwDrzPZzzsJbJ
wdLTBvSvHvZVGCjhfN
HsSSnZVHjjssZnJpSJjBHHWgQGcgqqQLQdQFqNgWgqGNDg
rmmRwrtfThtTrbCrGGGcLBDTqDBNQLdL
mwPrrbzPfwvbzhwMMnnjHnBjZlnzMM
gjjdMBgdqdTpJpBcjgRRRlqnvrqfnZtrtZDw
zHShWLhCszCWHVbVzQWCPtQvNZRwtfftfNnrnftlfR
PzPSssHbVhCLFMJFcMFJJMjdJw
ZqJdtpfpjmpJjpnwWdttTCDLLTQFNTzTzrcqrQqc
MsSlBGvBsSGGSlbGsCgggNTgzNLczFQNrNQVQcFzFF
sGHHSGllhvMGhSRGCjtjtjnjnnmHWpWWtJ
tMdjQlHPHsGjsCtsCpwwqfhfnnFMDMrpfD
SbNvWvBRJRWwFSgppgSrfg
RNcNbvzJRcVLRVzTRFLjdHCQttdZdPlHstPl
QWqgpdBflpHNCNWNHHPm
VVMbbJsLFVMhrMJMmRjFNHwHjjCTGSSRFj
mbMsZzsLmVhJZrcLcJhLMtnqvBnZdggplDffvlnlvnDn
prnNnsFnZpnBNdNtFrNnzNQQwTTQZqTHTQJQMwHDMDlZ
jgfgcSmbLmhmcPShghRdmwJTQjTlqGlJQJHqQqGHqQ
hRVhPfbCgbVggLVRSSmRhRPhrrrnCzzsvCvrnvFnNppsvBtd
QJLNDWSWQdLFFFhLdt
npHhHMsfsjpZjznRtmrMCdBwFBFrBdmt
HsjHqRRfnnHRsgfHffZspgzqDGQSWbQTDNGhQhSqNPhDWWbT
bsCmFDsGZCNsDmLDLZBSHSJTHnrZQMQSSQ
jqRpwvfqnnRQrftdBMHddB
phpchwpzjpvwRzwcsnlFsssPCCGzDlsD
rMqzVQfrfVZWZhTdRTQL
cgmtFtjFFJDDtFvSFRZdLlhpHZddmwTZWh
FbcSTtctcvFTJNgtJDGNPnCqMPMfMBfznGVsrMCq
wLJfGJJPZLBfwSLGHbqmhhDHHhFDzfhv
FsnpFjVjplTQCspNlCDbzhMMbqvMvsgmHDqb
lRdlTdTddllpRQFRltVVdFRcwrtrcWWcPrrWPrPSrZWLPc
VGVZhTppGTfPnJVJrFqbsmbSSshHqWqRHF
llzDCzlBLdNcCddlMMNBdCCtWHbFqFRRRsHjWtRwSWqbmjWm
NbcMBBvzzMQLCDBVTQQPVrPQPPZVPp
cdcgfmQdqlqhzzPjzfwpwf
GLBGBDvbvRzGwtnnmPpp
ZRCZBRFSRvLRLFvvbLLFQdHMTHTlQlNqNmqFlWdH
vzjzvHtcHvJcDStLLGSShCbbfF
MWFFTVZRMmMgdQdSQLwQrQwbGw
gFTmgmVZssRsWZRNzJlBHHnnJDvzNPlP
rHrvHpmHZfdGDDGGZd
cTlMsNhllMhGchNPCBlhMQgVDdgDSSWVbWVwRQwRSgbV
lnBjnNNTTMnCTcChPNhMvtzvFGLtJrjFtHHHzHJm
lgpdZZMmGVVzVZzt
HfHLrHqbPbzJJzRJJPTl
HsLWWbDqFrqlqfbsbDqDBncpgFmmvpnmmgpvdvjcdM
GpNVbTpJJNvMBMVvJTGvhnWQQScllnhhWlhVSznV
ZjswwHHLZzGnjWjSjl
sHdftLLtgLfwdtPmHtMbNpMTpNqGRbPvTqPv
sHSNNhNwsllGSGGlWSGWSsFrrVbQrdmFrVrrmnrHmrHr
QQMRZDDRcrcnmRcV
fJfCPMJCzTMZSGsWwsWBwqQf
HwQZZJsHdqqsdJQGRgCgVGgSqgpcGG
ljWWbnPhjBlGpCRCnScSGg
hrrztWlbPjltjMPSdJDZSsHttwsZwD
VzzbmzvpvNhvBDqc
QHSJSQGCwJCGrGQjjcgcBFhdgqdqFdDNNw
rCGJtZrHhhtLRsth
TMWwCLPpMTThrvtMRJjbjRvmJs
fDzcHFfSfFQfZzZRJbdmmqqssqtbSW
WWgGZglcllgPBBCBNVGPNr
wrwwhpTpbqhqrshrrfrFfwfzRJGdNJHNmcFzCCzCRJGzGR
vMggvjQvgPvQjVLMMPSZqWNJGCzcNGdcdzHPPzcmCzPz
qDZWvBZVfDhbTtrp
LpDvHdjVghnjbGrn
FBBBPwwlBlwSfFTWPHPWWhmgngmmnPnmbsngngbGrb
FwftBSCSfWCtwfVQDvHHCMVvdQLQ
ZrQpQlSpNlqQCVnQBmdDqmWDqmWWBDBB
HsZMsJvZzLMHTRwWhgDwmfDBgdhWdf
RZvTzJGzRjFrVNVjlQrS
mqjMwfqlSSPmSrlPhwhVpGRcppWcpcGRcGWv
ssJDJJNgZNDWrRWcRpvr
ZTsTnnsLJQgPnfMwmnMrfm
qsVBvZqWLdfbfvLj
mPNRgmHBBGQrCbSbrdfCCSbC
PlQTGcTTcgGFQQGPTGllpqMzwzpVJZwBMssZ
FWGcNRLRLhwJJQfV
nzbzlDBHSpTDbpDpzHwCqhqwJJghQqQMCCBw
JnzndzpmJFmNsrrFsc
gmRwwDwfnRDJgwZLFQFFNGNQrFBmFbbm
CCVHVWWThSrjVGvbNj
WpdqpplppHCWzlClMMTTZJcJsdscJLLdbDDfZDgL
VNtCCMDllpBqDvtdCczTSgjHlzGSHSGZTZ
hPFPsQhhFhLnbsRnLFssdzcHdsSHSSHgjzHG
QPWPQrPPmbdnbWLFPrrBVrVDBqqNMVwttDtBvD
PPNNRggwgRRgHBhDtwhTwbDs
SFGSFSMCJFMrcrCMSSsbtrTbbZhBvtHhrTHD
MFfSMpflQLQflfLjnLmddsLdddqq
RcgbcrrFscVrwZVCgVGGmHppNNndWnGdNqddqqNqND
jTlSTBSTjLTvlvjjPtvMLlhHnftphtDFNFqDnDHWNddn
QBMQvzzjzvJPjFQMmwZJrgmCCJVRVbbc
CzPJsWCpvsNszsJsNsHlDhMrrJGjhrRVhRGgDDjG
tFFdbqFLFdwctQdfVhjRRghTcrjVRTDW
bwQtFLdLBdFmwnHnWHPBNnHCnp
CNTstGNslRRRstlmNmmTZZqfFwtqgwqgfBPSwSWwqgWq
hpDbcHbpSrcgqqzhhWVfqg
DDcLDjbMjCSsZRNlMv
MhHMCMNbzbMHlcqmGmrmWc
tnPggdZPBPgdtttJpdnwVBnmqQcvlQrQlfGqmfWffBcqWD
VPPwPPLPwLGFGLzCbG
rqBcBmjHTGfPbcVgPG
dlDpsdshzlldlDvsWlWvLQbQBbfLFLbPvbBGQBgG
BlBznnRWzlphphBnhZjZtNNCNmrNqjCqHwHN
mQBvmvBmmLJvvrLtttQrfhGlcRGfRhVGWJVChlRG
MzPswTsbTPPsNgMNszgzMpbMfcRcGflVGRfWSpFRlWWWFhcC
bcPsTbgbbTTwNZzTZzvdjdjtQQndZvdrvdmZ
hQzTQJFFZJrcdcdZFFrGFSVWVRWRwRgRHVMWDCDSWc
lPmpNBnnnsNBnLnfbfnCDWMvwRvDCCMPwwHWvM
HpjmffNlnqqhddTddFZjGJ
BwsLFFbHLbVCSCSFbsbFLsJbqnTtZrRMHTZtrTrZTcRRRRTq
lGhNhpPmmhpztZTBrcpjRqpB
QPzdfzBQNgFJSCwsdLbS
ZsZsSBTgffSCqSSfrMnnMwjqmqmnnnqwMm
bbPPbzVclcPzGNlvzVtmnDBnQmtnQLBjJVLn
zPFGplGGvdPbHplcbzzvdlNBTThgRpCTCTfhfsCCsSRZhR
CVLSVCLVZRsHcnzSRpdZZRCdPlmcMWDDlPNqMqtDMmqPMlDt
TBnGjfQrQJjhfWlPPmPQDNlmlP
fjhhGvjvvrTTBhvTBTbvGVRLzVnbSRZpHddspHRzLs
DDtWjfRfftWMLzSQjzzhWjjwRVPHqFbBbZHVwZBFvFwZvZ
JGllgCJlJsrCGPrCNTPdslvZVVNVbvBqNbbpbbFHpBwZ
CcPdnCdmCJjfcftWhtSL
pgpfddDGHWzDNGNGpRCQjCTFHZZQFQjcRT
bJlhqmMvnlrRQFtTthPVhZ
lvbJrlJMBwfzGNTddB
wpbJGGZpsjvtdWvGWF
HqqhBhBqhhNQHTSHqqNzRHVPvTvddWrjtrjFvrvdTdVP
NRLCRzlqHQtNRBLzQllhhZcgbggwmLDZpsgssDpwwD
pDzFzJFcVMcWJFJFzpLBsqWLZssshsGLGbsS
wqHqfvnfrRwQtdQRthhBbBbZLhPLnBTGsh
CfQqlqvtfHNvMVmzmmMCFDMc
GcgpNHvcSNvpSLphdhsLdQTsdWThhQ
wwzttPrrhQswdhnT
tjJjMJRbRbjztmjtjbgcRsNlgglHpDFSlSvg
VVLvLqqPVlvcqLLdwLbHpzcHSsbRJppHbHpF
CfjjCNGmMWhWjhWHWb
ZmGZffggrDqZvZtlbTqL
TTmmhvBvvHWzHpsPpstpLVdwwsLb
qflfFgNctNcCnCCNDnfFFNDwrslwZbPswwZbJLJZbrlPLL
FgQDDcncStCgtqccjSDTHWMThvhTQMBQhWvWRG
SqhVghPccSBhgSBqWBFNQNsHQHMjCCQQWCwQHN
fLZftnlttcbbtZbZlpZtttQjwsCQjRwwRDQspMRMNNQs
TfLtvbJtZmlbTTTtlJbFvVqPSgBdPqPhFSGBcd
pPPNDptcqtpcDztLDhhngnnJgJTmJwNnwm
HVVCsSClHGBCHslWHbSCGGVngTrJwnnJnQRRBrhQhgJhdm
WTWWWsvVlvGbWCFvjDftPpjqZLLtDz
wWclwtDwRvflvffB
sMMsGdsSTMrJZNqczfdvhvnzCnfv
rspppMjMspSTSMpgLjcPFmwPLmPHwb
tCdSMHtHtRFHdWSSJQSgrrrnghTNJN
BGfcvDsfvsqcvqfGvfGnNLhggBNQJNJQmpgQJm
sGfQDPDZzfDZzclwDzwsDlfjtbdHClFRCMWjbMFMRFWbdj
pJNCcvqCccsVvFCpqsmvWJfCBWgSzBBRrrBRDDgDrSbbgQbQ
TMLnLjjffwfwGdjQjDDBjBrDtztRSb
MdPLGhHnMZhlPHHTFfZvVCpmmmcFcVFC
SwFMfCMRCdQDdMbmdFfdbbnlcVncVCcgLqWcNNnCcWlW
hPjQzzhGzhpPrtPJPpPHrVgnqVVncVVnNHlqVnncNB
ptjGrptztpthtrtJJhTsGwFDZZDQmSdfZSwsRZSwfZ
rSSWWCWrdllHWpjcnFNnRCNjQp
bGwwJqGVGbGJVVhgbBgttGmBQjFsMjpMcMnnMBcQFNnsssvv
bfthwmfJfgwwmmwZqVJPHNHSZHWzSlDPrdDdSH
nmJccvclcbwmlbbvVbvsHwJJPCPNCNPnLBhrBPPLBhLhBgBP
MdRGtdDRTqWDMqtMDtQDRWSdLLBsrhLgBCgrgCgNNLPBfNMf
dRZQdDdRRSQWGsjZmwzjmlzsZH
PBGGMrTQTrTBpPQpLpSlwjwfjtlnfwbmGttw
fCsJCWJcvRCtwwjbCl
NsqcsfcvDHFVDJvdLQTrpdTTzTPpHr
rltrwsBTlrfGZggGBLGGNN
jhMnRQJVphMnbhvQjDZNqqZDNTNHZVHGHq
MRvbhQRQQChpvbjvMSvQnMcsFsfwwmlCwFwWcTWwrmPc
DDvLLLBnvrzvbvbmtv
TMwRjTRMGCwGGwrjQQnmrQrrQdhZ
MJPFHFTwgCGqGqgJMGDfSWcsnBSccgVDlnpW
flzVzNrdLNLJzrGlfdlzjrQDgFTpDgDmmmgFmqFDQjQh
CbnBcsZnPZVSnwvVsZbRhhBDpgFphgmgDgTppq
ZWnsWSnncSZsntZCbsswwJMzdLzlMdNMLtMVfrllMt
ZffSgNfgJgGCHZcZrpHrNJTLhqvSLTqQnvVTLvzvLTjV
tWFtHMwlBlDqjjzjnqvvlV
DRMPDtWHPFDBFFwWMFBmFRPgZpJfsffNGJNrGcsprrsmfg
wRZRmpZmlPqZjzGrdrGq
bBhQQFPQbPDVNzVNzdGWNdrf
QFbcDcDbLHgHBPDFRsSSMtmvRttMpCLS
MpWJVVJMcWvpRShcwpLGflmqzSfNdfNLdQzN
CDBTtCgtbjgCRrBrPBTQqzflNqjGdLzzmqffzq
rFgnnBbttDTPtHCDPrPMnpwVJhJvMZvpMvppRZ
sWTTmpsfsWppPTTsTVZWHVVZNvVcdcJvdN
DjjBzjhRHvvvSzdc
rMBjjrjbjrGDlgMlMrGjBgRLPTTwHMsfnFwFQPMPMnmFFm
QRRbDjjmPzNQwFDNmrQmzCbVHrMhBVrJLJJfMGGLtfJBHh
dsWcsqqWSWvnWnWcWGPJLBqhLBqGhBJHHH
ZWnPWgWgPnlbCDDwmmDbRZ
nfPqqfLqQnfHBSqnzztQjVmjfGRWJNGRWsJWJfmJ
TTMlMMlFDMGVGRsVJH
CbDbFDbvgTFFwgTDlDprhlPSqBzSnPdLPtPPHgznQqBQ
fJmWVfHqjfjhZCQZ
NcNzBNvgszQmzjnthZQC
LsLsgBNFmFgTFgGBBgcdMdvPDPDJWrlpVbGpJWqHDlHJHD
SllDdvzgdFDdlPJvbFDDSzFScPTRTNcwfZRwRhcwwNnRZTtf
WBpWBCLGVpLjHrHGGVhZNwcTVcNhVnRcNZ
LHLQLspHWQGpWCHnBvdzDJFlqvdsqgSgqS
GcTctDMjMhpMDRjLsMMsfDWFfdPCFNbnCPnvCPgW
JmvwqlBwnmfdFPFP
SvZqHSZqqHZZZBlllBwSwsRsMHpLjpLsMGtsMspGRT
ClLnCLfClLVllfLLcQjLBCfCmSHVsttsmtsVStDNVdppdsSp
PFrRMbWqMRwFRqRSqwqvMvMsGtgsdmssrgNtdmpNdDGgdt
bwJbPWPwFFPFSczCZzZZCcfjBJ
MwmBmzwJQTcTmfPVfZPwhhwHPH
jlnrglFLvbrGRFGnvFZdNNFfPZddPThVhdPH
RjbjpgbnLGvpLgzBqBpmWmmzqTMS
FnsSpttPnPbNCFDtsPnFHQZTQZgcwgDDTfrfTHMZ
mRjzRzlvBvhjZrQmMMwfZZNN
ldzddlzLRlRWdhjdRLjhRWtJbJbNtJJpJPbCbGCWNG
wBwmNZBTmzzcVcmpzZqdMgPjnLSVlPgDPdbMdg
flJvGtHffHDddddbHjnb
RstrhfrhhRGFQtRhtftvQhvFZpsmpWwNlWqcWTccBNWswqNp
DPWhbzDlQLLlQbLDlLhPhLFNNJqCFGqnNJCCSCnGPnGN
wvwjtvtdwfssvSJgFFvGGSCFcp
mtdrZwwJsrtddrHRtZWbVThLlBzVTzhHQWhB
TsRRWctsTJMQZllggc
zDvhpbprgGvpvVlVQlZpQMJVlQ
rrrvFvGCDhDSrrrvChCgSstBNTSftWBjTdfWBN
JJdssBcLVGrgbBHWrH
QZTptvmvmlZpRDlMMMZCQvnjjFnrjWGFbjnrnFGWgZrz
TMRplDMggtwlppTlvhsJJqdcqwVPSSNcLd
JjTCCrcRvccPLmMP
NfGFPZlNnwBfPlbbbQZGqLHgzLghSmMBzSgvzmDMhv
ZfbnNQpQnZGlGlGpWTddjdTJdpRPTrCj
gWLblMMggdWsdRJlblMRMMqWDvPvcPPPccJPJVTZVZThmcDP
rQFfGfrCHrjnrtNTcPShTSPvvVLtmm
fQrCfLrpLHnCHwslqzqsslswzqRW
zpJtGlJPMPTlTjGJCDGCDljpdnvhhWnZnZnDwwmvnWDWWvdd
sHrVrSrRRRLNgLVBqSsZmWwvwcvwZjmwngmdbn
QsQQBrrLHTjPpTQzzP
JDlzHHzzptRDmbTMrrVQ
dRRNqnCBnmrQsVQQ
wFPCBNFgwjPwhgFNztftpJRPpzRvvHtZ
DlBhrDBPPwMWwhWchW
ntSqbbSJFJNqzVzjCfMvfSlSRWccRL
mVlHtNVtqldbJVmNHmdTTBBgrQQgGsPQdrDgsP
HWHNbBgvNLdcvQMnSf
wqqqVPDPhqwszFwrrszFfMdWthLcMdfhthSQfJSt
qVPVwTzFwFDpDrPDzDPFDPlCHjBGjTmZGjbWWGZBRTNjjH
GVgdWjllSqgjdgHqqlfmhwcpwCzhvZwMcScv
nsJQbLRQsNnzQDQQPPBbRBRhfZwpZcvwpvvmLCcvpcmfMM
DRJtnnRbBBnPztsrPzRBPbsFFHqqVrqggjFWqrgWjTGgFq
hhZJQPJFHGGlcWWslpNN
VwwwJjvwMtrCnwjDNDzlfscWszWW
nVStCrMqbVwqVqSqwnLPhTJFdRgJHZSFRLTP
vPgMbbRhhvMvNjjLWsWQsHQmHwBrmmBzww
tFctDnVFpppHVBTdzdTQwl
FtSFqSptfJCqqJStZCqDpDJMhvLLgLMgQgjgGZgGgMPLZg
zwsWgSGWLSVhPWhtLgLWhPVNQTmDrDQttZpdQtdpQDDQZQ
fjCHcvvjMDrppCQpVV
VMqRnVJMVLPzbRWhGh
mjRmzQlzDzNHWwDZ
FBfJBGqnnpfSVGnpJbJVfNtCsJHWZvrsNJCZrCNsvN
fZPBnfPqSBqdfpFbVnVSjgdcLLgRLjmgRhLLghlR
FSFnTcppdQtnnDhtzDfg
ZLGVmBLBVwZCVjjGqGhVwVVgzzbMDtNNvszMmMffNDDvtM
VZPJjBZVqBZZBjqwVqllpSTphhQFPShWSQcW
hTRdcLrCLgplLvBFGvlL
nZDZqzbDbDzRZtVNDzDWGwslsllBFpnlpGvJssFG
zbqjNWQVmVPrrRjRdRhS
VpNCbVHlHHZfflVfmchctqFcqQQjZmZM
WDSRGgsSvgJSRrnWgqQhmjBqmhqrtLqmQm
znSGTgDJnsDGzgwCwlpbCNwHzVtl
sTTTrpHFFFqTnQbbvfJdDzHHDLVV
CjMtgMgRvbPfjjvB
mhMvlhhWClvqshNTQQqsNN
tWFtFBzbwdFrpmdhdm
qTqDjJjJQQqMjTDLJjNqNqPNdmpcSmhdmhhmcrWZpdPGddcc
RjNQLJNTTJDDJRHHjQqnMWtlvvVvbtBvRVzgzgwgVg
CGdQjwdJrbBmpmZZZlRWcb
NgtMPVstgSzBLzhgzgLgDRlcmDWRmlZvcSmDSvvp
LhNsgPPLFPPsNzMhhVzPsGJBFqwQGfnqfQjdGdGfwr
CNbNdbzjCZpPNzjmzjsCMRJvnnMRGnsvJGRs
wrtdwTLWFcFWdFgwRRsnJGnGfTGJfMsq
FttcwgBtgVLgPldQSNZBzBpz
DjRZrrRmttRFDvDrFTZsnWnHVSTSSJVZJH
dNNhLqlLLqdCzfMMlCfSncTVVWcHdcVsVdSVnT
QqppMfzMfqWCwbRQrwFrrttQ
dwGjHrtjsdhfCHnPSpfMfDPpPDWS
lmNzzlLbFqcqNgzpWMSvbbvDQDGWDp
LBmglgmqBqmrwCGhCjVtBC
tvHgWZCCprlgpWglCtjPhLmPmhVdJFSzVzdJVmmQ
fBnTTnNNBnwfnNqcBbBBTbGJQQJhSSdQJJsmdJFSQGSmVV
cMcDwFbRfFRlHCRCZrrp
ZFWmgghzBgwgjWBzjzmRWWMmsVwnVrsdVdwNrrpnnVrPCnCP
GLLbtGqllctqvGJvSlQbJGsPnVdsdpsTPLsVppBCTVss
tJBStGSvctvDDfczmRgRZjzDjZmgzH
FMrLmsQQSWzCZBhpQJTQQZ
dPPVncVvPBJDCPhwJD
fvHbbVHvqnvvvBzgLbbGGmrbMr
mrZzrzqDrhZqDddSFrCGLLLPQPQBJPJJBnQq
TgbpGblWlMsjgWlgMfpNRgbRHHBnHHHtLpCJPCPBnBLJtQQL
sbTlblTlvRbbGblbFcdDzccVcDVvzzzd
zMzfzlGwSBMMSCMzhsPgfcPcfcbhjQPt
FHHqJVdJmFmdVrJdJppthscjGtqRPRcccgcQbR
rvNJJpLrvvLnJvNFFvZZZBWznBWGSDCMnCwz`
