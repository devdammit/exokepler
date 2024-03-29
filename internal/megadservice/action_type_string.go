// Code generated by "stringer -type=ActionType -output=action_type_string.go"; DO NOT EDIT.

package megadservice

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TOGGLE-1]
	_ = x[ON-2]
	_ = x[OFF-3]
	_ = x[PAUSE-4]
	_ = x[REPEAT-5]
	_ = x[GLOBAL-6]
	_ = x[DIMMER-7]
	_ = x[TDIMMER-8]
}

const _ActionType_name = "TOGGLEONOFFPAUSEREPEATGLOBALDIMMERTDIMMER"

var _ActionType_index = [...]uint8{0, 6, 8, 11, 16, 22, 28, 34, 41}

func (i ActionType) String() string {
	i -= 1
	if i < 0 || i >= ActionType(len(_ActionType_index)-1) {
		return "ActionType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ActionType_name[_ActionType_index[i]:_ActionType_index[i+1]]
}
