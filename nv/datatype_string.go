// Code generated by "stringer -type=dataType"; DO NOT EDIT

package nv

import "fmt"

const _dataType_name = "_UNKNOWN_BOOLEAN_BYTE_INT16_UINT16_INT32_UINT32_INT64_UINT64_STRING_BYTE_ARRAY_INT16_ARRAY_UINT16_ARRAY_INT32_ARRAY_UINT32_ARRAY_INT64_ARRAY_UINT64_ARRAY_STRING_ARRAY_HRTIME_NVLIST_NVLIST_ARRAY_BOOLEAN_VALUE_INT8_UINT8_BOOLEAN_ARRAY_INT8_ARRAY_UINT8_ARRAY_DOUBLE"

var _dataType_index = [...]uint16{0, 8, 16, 21, 27, 34, 40, 47, 53, 60, 67, 78, 90, 103, 115, 128, 140, 153, 166, 173, 180, 193, 207, 212, 218, 232, 243, 255, 262}

func (i dataType) String() string {
	if i >= dataType(len(_dataType_index)-1) {
		return fmt.Sprintf("dataType(%d)", i)
	}
	return _dataType_name[_dataType_index[i]:_dataType_index[i+1]]
}
