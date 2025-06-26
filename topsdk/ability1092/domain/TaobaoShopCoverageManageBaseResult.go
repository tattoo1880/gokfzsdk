package domain


type TaobaoShopCoverageManageBaseResult struct {
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

func (s *TaobaoShopCoverageManageBaseResult) SetSuccess(v bool) *TaobaoShopCoverageManageBaseResult {
    s.Success = &v
    return s
}
func (s *TaobaoShopCoverageManageBaseResult) SetErrorCode(v string) *TaobaoShopCoverageManageBaseResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoShopCoverageManageBaseResult) SetErrorMsg(v string) *TaobaoShopCoverageManageBaseResult {
    s.ErrorMsg = &v
    return s
}
