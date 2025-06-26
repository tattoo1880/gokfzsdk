package domain


type TaobaoWirelessVideoCheckRopResultTo struct {
    /*
        系统自动生成     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        系统自动生成     */
    Code  *int64 `json:"code,omitempty" `

    /*
        系统自动生成     */
    Success  *bool `json:"success,omitempty" `

}

func (s *TaobaoWirelessVideoCheckRopResultTo) SetErrorMsg(v string) *TaobaoWirelessVideoCheckRopResultTo {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoWirelessVideoCheckRopResultTo) SetCode(v int64) *TaobaoWirelessVideoCheckRopResultTo {
    s.Code = &v
    return s
}
func (s *TaobaoWirelessVideoCheckRopResultTo) SetSuccess(v bool) *TaobaoWirelessVideoCheckRopResultTo {
    s.Success = &v
    return s
}
