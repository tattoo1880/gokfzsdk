package domain


type TaobaoShopCoverageQueryBaseResult struct {
    /*
        门店覆盖范围地址列表     */
    Data  *[]string `json:"data,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误信息     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

}

func (s *TaobaoShopCoverageQueryBaseResult) SetData(v []string) *TaobaoShopCoverageQueryBaseResult {
    s.Data = &v
    return s
}
func (s *TaobaoShopCoverageQueryBaseResult) SetSuccess(v bool) *TaobaoShopCoverageQueryBaseResult {
    s.Success = &v
    return s
}
func (s *TaobaoShopCoverageQueryBaseResult) SetErrorCode(v string) *TaobaoShopCoverageQueryBaseResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoShopCoverageQueryBaseResult) SetErrorMsg(v string) *TaobaoShopCoverageQueryBaseResult {
    s.ErrorMsg = &v
    return s
}
