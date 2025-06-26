package request


type TaobaoFenxiaoRefundSupAgreeRequest struct {
    /*
        同意退款的原因     */
    AgreeReason  *string `json:"agree_reason,omitempty" required:"false" `
    /*
        要操作的逆向退款ID     */
    RefundId  *int64 `json:"refund_id" required:"true" `
}

func (s *TaobaoFenxiaoRefundSupAgreeRequest) SetAgreeReason(v string) *TaobaoFenxiaoRefundSupAgreeRequest {
    s.AgreeReason = &v
    return s
}
func (s *TaobaoFenxiaoRefundSupAgreeRequest) SetRefundId(v int64) *TaobaoFenxiaoRefundSupAgreeRequest {
    s.RefundId = &v
    return s
}

func (req *TaobaoFenxiaoRefundSupAgreeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AgreeReason != nil) {
        paramMap["agree_reason"] = *req.AgreeReason
    }
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    return paramMap
}

func (req *TaobaoFenxiaoRefundSupAgreeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}