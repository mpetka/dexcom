// Code generated by "stringer -type Trend"; DO NOT EDIT

package dexcom

import "fmt"

const _Trend_name = "UpUpUpUp45FlatDown45DownDownDownNotComputableOutOfRange"

var _Trend_index = [...]uint8{0, 4, 6, 10, 14, 20, 24, 32, 45, 55}

func (i Trend) String() string {
	i -= 1
	if i >= Trend(len(_Trend_index)-1) {
		return fmt.Sprintf("Trend(%d)", i+1)
	}
	return _Trend_name[_Trend_index[i]:_Trend_index[i+1]]
}
