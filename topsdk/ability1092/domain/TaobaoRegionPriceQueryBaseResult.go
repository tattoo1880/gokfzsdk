package domain


type TaobaoRegionPriceQueryBaseResult struct {
    /*
        data     */
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

func (s *TaobaoRegionPriceQueryBaseResult) SetData(v string) *TaobaoRegionPriceQueryBaseResult {
    s.Data = &v
    return s
}
func (s *TaobaoRegionPriceQueryBaseResult) SetErrorCode(v string) *TaobaoRegionPriceQueryBaseResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoRegionPriceQueryBaseResult) SetErrorMsg(v string) *TaobaoRegionPriceQueryBaseResult {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoRegionPriceQueryBaseResult) SetSuccess(v bool) *TaobaoRegionPriceQueryBaseResult {
    s.Success = &v
    return s
}
