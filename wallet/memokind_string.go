// Code generated by "stringer -type=MemoKind -linecomment"; DO NOT EDIT.

package wallet

import "strconv"

const _MemoKind_name = "textidhashreturn"

var _MemoKind_index = [...]uint8{0, 4, 6, 10, 16}

func (i MemoKind) String() string {
	i -= 1
	if i < 0 || i >= MemoKind(len(_MemoKind_index)-1) {
		return "MemoKind(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _MemoKind_name[_MemoKind_index[i]:_MemoKind_index[i+1]]
}
