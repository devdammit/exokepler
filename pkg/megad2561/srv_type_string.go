// Code generated by "stringer -type=SrvType -output=srv_type_string.go"; DO NOT EDIT.

package megad2561

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[HTTP-0]
	_ = x[MQTT-1]
}

const _SrvType_name = "HTTPMQTT"

var _SrvType_index = [...]uint8{0, 4, 8}

func (i SrvType) String() string {
	if i < 0 || i >= SrvType(len(_SrvType_index)-1) {
		return "SrvType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SrvType_name[_SrvType_index[i]:_SrvType_index[i+1]]
}
