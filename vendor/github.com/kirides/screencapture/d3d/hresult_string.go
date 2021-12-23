// Code generated by "stringer -type=HRESULT -output=hresult_string.go"; DO NOT EDIT.

package d3d

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[S_OK-0]
	_ = x[E_INVALIDARG-2147942487]
	_ = x[DXGI_STATUS_OCCLUDED-142213121]
	_ = x[DXGI_STATUS_CLIPPED-142213122]
	_ = x[DXGI_STATUS_NO_REDIRECTION-142213124]
	_ = x[DXGI_STATUS_NO_DESKTOP_ACCESS-142213125]
	_ = x[DXGI_STATUS_GRAPHICS_VIDPN_SOURCE_IN_USE-142213126]
	_ = x[DXGI_STATUS_MODE_CHANGED-142213127]
	_ = x[DXGI_STATUS_MODE_CHANGE_IN_PROGRESS-142213128]
	_ = x[DXGI_ERROR_INVALID_CALL-2289696769]
	_ = x[DXGI_ERROR_NOT_FOUND-2289696770]
	_ = x[DXGI_ERROR_MORE_DATA-2289696771]
	_ = x[DXGI_ERROR_UNSUPPORTED-2289696772]
	_ = x[DXGI_ERROR_DEVICE_REMOVED-2289696773]
	_ = x[DXGI_ERROR_DEVICE_HUNG-2289696774]
	_ = x[DXGI_ERROR_DEVICE_RESET-2289696775]
	_ = x[DXGI_ERROR_WAS_STILL_DRAWING-2289696778]
	_ = x[DXGI_ERROR_FRAME_STATISTICS_DISJOINT-2289696779]
	_ = x[DXGI_ERROR_GRAPHICS_VIDPN_SOURCE_IN_USE-2289696780]
	_ = x[DXGI_ERROR_DRIVER_INTERNAL_ERROR-2289696800]
	_ = x[DXGI_ERROR_NONEXCLUSIVE-2289696801]
	_ = x[DXGI_ERROR_NOT_CURRENTLY_AVAILABLE-2289696802]
	_ = x[DXGI_ERROR_REMOTE_CLIENT_DISCONNECTED-2289696803]
	_ = x[DXGI_ERROR_REMOTE_OUTOFMEMORY-2289696804]
	_ = x[DXGI_ERROR_ACCESS_LOST-2289696806]
	_ = x[DXGI_ERROR_WAIT_TIMEOUT-2289696807]
	_ = x[DXGI_ERROR_SESSION_DISCONNECTED-2289696808]
	_ = x[DXGI_ERROR_RESTRICT_TO_OUTPUT_STALE-2289696809]
	_ = x[DXGI_ERROR_CANNOT_PROTECT_CONTENT-2289696810]
	_ = x[DXGI_ERROR_ACCESS_DENIED-2289696811]
	_ = x[DXGI_ERROR_NAME_ALREADY_EXISTS-2289696812]
	_ = x[DXGI_ERROR_SDK_COMPONENT_MISSING-2289696813]
	_ = x[DXGI_ERROR_NOT_CURRENT-2289696814]
	_ = x[DXGI_ERROR_HW_PROTECTION_OUTOFMEMORY-2289696816]
	_ = x[DXGI_ERROR_DYNAMIC_CODE_POLICY_VIOLATION-2289696817]
	_ = x[DXGI_ERROR_NON_COMPOSITED_UI-2289696818]
	_ = x[DXGI_STATUS_UNOCCLUDED-142213129]
	_ = x[DXGI_STATUS_DDA_WAS_STILL_DRAWING-142213130]
	_ = x[DXGI_ERROR_MODE_CHANGE_IN_PROGRESS-2289696805]
	_ = x[DXGI_STATUS_PRESENT_REQUIRED-142213167]
	_ = x[DXGI_ERROR_CACHE_CORRUPT-2289696819]
	_ = x[DXGI_ERROR_CACHE_FULL-2289696820]
	_ = x[DXGI_ERROR_CACHE_HASH_COLLISION-2289696821]
	_ = x[DXGI_ERROR_ALREADY_EXISTS-2289696822]
	_ = x[DXGI_DDI_ERR_WASSTILLDRAWING-2289762305]
	_ = x[DXGI_DDI_ERR_UNSUPPORTED-2289762306]
	_ = x[DXGI_DDI_ERR_NONEXCLUSIVE-2289762307]
}

const (
	_HRESULT_name_0 = "S_OK"
	_HRESULT_name_1 = "DXGI_STATUS_OCCLUDEDDXGI_STATUS_CLIPPED"
	_HRESULT_name_2 = "DXGI_STATUS_NO_REDIRECTIONDXGI_STATUS_NO_DESKTOP_ACCESSDXGI_STATUS_GRAPHICS_VIDPN_SOURCE_IN_USEDXGI_STATUS_MODE_CHANGEDDXGI_STATUS_MODE_CHANGE_IN_PROGRESSDXGI_STATUS_UNOCCLUDEDDXGI_STATUS_DDA_WAS_STILL_DRAWING"
	_HRESULT_name_3 = "DXGI_STATUS_PRESENT_REQUIRED"
	_HRESULT_name_4 = "E_INVALIDARG"
	_HRESULT_name_5 = "DXGI_ERROR_INVALID_CALLDXGI_ERROR_NOT_FOUNDDXGI_ERROR_MORE_DATADXGI_ERROR_UNSUPPORTEDDXGI_ERROR_DEVICE_REMOVEDDXGI_ERROR_DEVICE_HUNGDXGI_ERROR_DEVICE_RESET"
	_HRESULT_name_6 = "DXGI_ERROR_WAS_STILL_DRAWINGDXGI_ERROR_FRAME_STATISTICS_DISJOINTDXGI_ERROR_GRAPHICS_VIDPN_SOURCE_IN_USE"
	_HRESULT_name_7 = "DXGI_ERROR_DRIVER_INTERNAL_ERRORDXGI_ERROR_NONEXCLUSIVEDXGI_ERROR_NOT_CURRENTLY_AVAILABLEDXGI_ERROR_REMOTE_CLIENT_DISCONNECTEDDXGI_ERROR_REMOTE_OUTOFMEMORYDXGI_ERROR_MODE_CHANGE_IN_PROGRESSDXGI_ERROR_ACCESS_LOSTDXGI_ERROR_WAIT_TIMEOUTDXGI_ERROR_SESSION_DISCONNECTEDDXGI_ERROR_RESTRICT_TO_OUTPUT_STALEDXGI_ERROR_CANNOT_PROTECT_CONTENTDXGI_ERROR_ACCESS_DENIEDDXGI_ERROR_NAME_ALREADY_EXISTSDXGI_ERROR_SDK_COMPONENT_MISSINGDXGI_ERROR_NOT_CURRENT"
	_HRESULT_name_8 = "DXGI_ERROR_HW_PROTECTION_OUTOFMEMORYDXGI_ERROR_DYNAMIC_CODE_POLICY_VIOLATIONDXGI_ERROR_NON_COMPOSITED_UIDXGI_ERROR_CACHE_CORRUPTDXGI_ERROR_CACHE_FULLDXGI_ERROR_CACHE_HASH_COLLISIONDXGI_ERROR_ALREADY_EXISTS"
	_HRESULT_name_9 = "DXGI_DDI_ERR_WASSTILLDRAWINGDXGI_DDI_ERR_UNSUPPORTEDDXGI_DDI_ERR_NONEXCLUSIVE"
)

var (
	_HRESULT_index_1 = [...]uint8{0, 20, 39}
	_HRESULT_index_2 = [...]uint8{0, 26, 55, 95, 119, 154, 176, 209}
	_HRESULT_index_5 = [...]uint8{0, 23, 43, 63, 85, 110, 132, 155}
	_HRESULT_index_6 = [...]uint8{0, 28, 64, 103}
	_HRESULT_index_7 = [...]uint16{0, 32, 55, 89, 126, 155, 189, 211, 234, 265, 300, 333, 357, 387, 419, 441}
	_HRESULT_index_8 = [...]uint8{0, 36, 76, 104, 128, 149, 180, 205}
	_HRESULT_index_9 = [...]uint8{0, 28, 52, 77}
)

func (i HRESULT) String() string {
	switch {
	case i == 0:
		return _HRESULT_name_0
	case 142213121 <= i && i <= 142213122:
		i -= 142213121
		return _HRESULT_name_1[_HRESULT_index_1[i]:_HRESULT_index_1[i+1]]
	case 142213124 <= i && i <= 142213130:
		i -= 142213124
		return _HRESULT_name_2[_HRESULT_index_2[i]:_HRESULT_index_2[i+1]]
	case i == 142213167:
		return _HRESULT_name_3
	case i == 2147942487:
		return _HRESULT_name_4
	case 2289696769 <= i && i <= 2289696775:
		i -= 2289696769
		return _HRESULT_name_5[_HRESULT_index_5[i]:_HRESULT_index_5[i+1]]
	case 2289696778 <= i && i <= 2289696780:
		i -= 2289696778
		return _HRESULT_name_6[_HRESULT_index_6[i]:_HRESULT_index_6[i+1]]
	case 2289696800 <= i && i <= 2289696814:
		i -= 2289696800
		return _HRESULT_name_7[_HRESULT_index_7[i]:_HRESULT_index_7[i+1]]
	case 2289696816 <= i && i <= 2289696822:
		i -= 2289696816
		return _HRESULT_name_8[_HRESULT_index_8[i]:_HRESULT_index_8[i+1]]
	case 2289762305 <= i && i <= 2289762307:
		i -= 2289762305
		return _HRESULT_name_9[_HRESULT_index_9[i]:_HRESULT_index_9[i+1]]
	default:
		return "HRESULT(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
