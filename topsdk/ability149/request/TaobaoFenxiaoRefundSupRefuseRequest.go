package request


type TaobaoFenxiaoRefundSupRefuseRequest struct {
    /*
        拒绝退款的原因     */
    RefuseReason  *string `json:"refuse_reason,omitempty" required:"false" `
    /*
        要操作的逆向退款ID     */
    RefundId  *int64 `json:"refund_id" required:"true" `
}

func (s *TaobaoFenxiaoRefundSupRefuseRequest) SetRefuseReason(v string) *TaobaoFenxiaoRefundSupRefuseRequest {
    s.RefuseReason = &v
    return s
}
func (s *TaobaoFenxiaoRefundSupRefuseRequest) SetRefundId(v int64) *TaobaoFenxiaoRefundSupRefuseRequest {
    s.RefundId = &v
    return s
}

func (req *TaobaoFenxiaoRefundSupRefuseRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefuseReason != nil) {
        paramMap["refuse_reason"] = *req.RefuseReason
    }
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    return paramMap
}

func (req *TaobaoFenxiaoRefundSupRefuseRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}