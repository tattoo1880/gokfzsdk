package domain


type TaobaoRegionWarehouseManageBaseResult struct {
    /*
        true/false     */
    Data  *bool `json:"data,omitempty" `

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

func (s *TaobaoRegionWarehouseManageBaseResult) SetData(v bool) *TaobaoRegionWarehouseManageBaseResult {
    s.Data = &v
    return s
}
func (s *TaobaoRegionWarehouseManageBaseResult) SetErrorCode(v string) *TaobaoRegionWarehouseManageBaseResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoRegionWarehouseManageBaseResult) SetErrorMsg(v string) *TaobaoRegionWarehouseManageBaseResult {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoRegionWarehouseManageBaseResult) SetSuccess(v bool) *TaobaoRegionWarehouseManageBaseResult {
    s.Success = &v
    return s
}
