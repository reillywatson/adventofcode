package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	binary := ""
	for _, c := range input {
		n, _ := strconv.ParseInt(string(c), 16, 64)
		dig := strconv.FormatInt(n, 2)
		dig = strings.Repeat("0", 4-len(dig)) + dig
		binary += dig
	}
	_, val := parsePacket(binary)
	fmt.Println("VAL:", val)
}

const (
	SUM     = "000"
	PRODUCT = "001"
	MIN     = "010"
	MAX     = "011"
	LITERAL = "100"
	GT      = "101"
	LT      = "110"
	EQ      = "111"
)

var versionSum = 0

func parsePacket(binary string) (string, int) {
	packetVersion := binary[:3]
	binary = binary[3:]
	versionSum += bin2dec(packetVersion)
	typeId := binary[:3]
	binary = binary[3:]
	if typeId == LITERAL {
		res := ""
		for {
			more := binary[:1]
			binary = binary[1:]
			num := binary[:4]
			binary = binary[4:]
			res += num
			if more == "0" {
				break
			}
		}
		return binary, bin2dec(res)
	}
	var fn func([]int) int
	switch typeId {
	case SUM:
		fn = func(ns []int) int {
			s := 0
			for _, n := range ns {
				s += n
			}
			return s
		}
	case PRODUCT:
		fn = func(ns []int) int {
			p := 1
			for _, n := range ns {
				p *= n
			}
			return p
		}
	case MIN:
		fn = func(ns []int) int {
			min := ns[0]
			for _, n := range ns {
				if n < min {
					min = n
				}
			}
			return min
		}
	case MAX:
		fn = func(ns []int) int {
			max := ns[0]
			for _, n := range ns {
				if n > max {
					max = n
				}
			}
			return max
		}
	case GT:
		fn = func(ns []int) int {
			if ns[0] > ns[1] {
				return 1
			}
			return 0
		}
	case LT:
		fn = func(ns []int) int {
			if ns[0] < ns[1] {
				return 1
			}
			return 0
		}
	case EQ:
		fn = func(ns []int) int {
			if ns[0] == ns[1] {
				return 1
			}
			return 0
		}
	}

	lengthType := string(binary[:1])
	binary = binary[1:]
	var subvals []int
	switch lengthType {
	case "0":
		length := binary[:15]
		binary = binary[15:]
		sub := binary[:bin2dec(length)]
		for len(sub) > 0 {
			var val int
			sub, val = parsePacket(sub)
			subvals = append(subvals, val)
		}
		binary = binary[bin2dec(length):]
	case "1":
		numSubpackets := bin2dec(binary[:11])
		binary = binary[11:]
		for i := 0; i < numSubpackets; i++ {
			var val int
			binary, val = parsePacket(binary)
			subvals = append(subvals, val)
		}
	}
	return binary, fn(subvals)
}

func bin2dec(b string) int {
	n, _ := strconv.ParseInt(b, 2, 64)
	return int(n)
}

const input = `220D4B80491FE6FBDCDA61F23F1D9B763004A7C128012F9DA88CE27B000B30F4804D49CD515380352100763DC5E8EC000844338B10B667A1E60094B7BE8D600ACE774DF39DD364979F67A9AC0D1802B2A41401354F6BF1DC0627B15EC5CCC01694F5BABFC00964E93C95CF080263F0046741A740A76B704300824926693274BE7CC880267D00464852484A5F74520005D65A1EAD2334A700BA4EA41256E4BBBD8DC0999FC3A97286C20164B4FF14A93FD2947494E683E752E49B2737DF7C4080181973496509A5B9A8D37B7C300434016920D9EAEF16AEC0A4AB7DF5B1C01C933B9AAF19E1818027A00A80021F1FA0E43400043E174638572B984B066401D3E802735A4A9ECE371789685AB3E0E800725333EFFBB4B8D131A9F39ED413A1720058F339EE32052D48EC4E5EC3A6006CC2B4BE6FF3F40017A0E4D522226009CA676A7600980021F1921446700042A23C368B713CC015E007324A38DF30BB30533D001200F3E7AC33A00A4F73149558E7B98A4AACC402660803D1EA1045C1006E2CC668EC200F4568A5104802B7D004A53819327531FE607E118803B260F371D02CAEA3486050004EE3006A1E463858600F46D8531E08010987B1BE251002013445345C600B4F67617400D14F61867B39AA38018F8C05E430163C6004980126005B801CC0417080106005000CB4002D7A801AA0062007BC0019608018A004A002B880057CEF5604016827238DFDCC8048B9AF135802400087C32893120401C8D90463E280513D62991EE5CA543A6B75892CB639D503004F00353100662FC498AA00084C6485B1D25044C0139975D004A5EB5E52AC7233294006867F9EE6BA2115E47D7867458401424E354B36CDAFCAB34CBC2008BF2F2BA5CC646E57D4C62E41279E7F37961ACC015B005A5EFF884CBDFF10F9BFF438C014A007D67AE0529DED3901D9CD50B5C0108B13BAFD6070`
