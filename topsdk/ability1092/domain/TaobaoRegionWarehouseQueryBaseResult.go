package domain


type TaobaoRegionWarehouseQueryBaseResult struct {
    /*
        返回覆盖地址信息     */
    Data  *string `json:"data,omitempty" `

    /*
        errorCode     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        errorMsg     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        success     */
    Success  *bool `json:"success,omitempty" `

}

func (s *TaobaoRegionWarehouseQueryBaseResult) SetData(v string) *TaobaoRegionWarehouseQueryBaseResult {
    s.Data = &v
    return s
}
func (s *TaobaoRegionWarehouseQueryBaseResult) SetErrorCode(v string) *TaobaoRegionWarehouseQueryBaseResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoRegionWarehouseQueryBaseResult) SetErrorMsg(v string) *TaobaoRegionWarehouseQueryBaseResult {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoRegionWarehouseQueryBaseResult) SetSuccess(v bool) *TaobaoRegionWarehouseQueryBaseResult {
    s.Success = &v
    return s
}
