package request


type TaobaoFenxiaoReturngoodsSupRefuseRequest struct {
    /*
        拒绝退货的原因     */
    RefuseReason  *string `json:"refuse_reason,omitempty" required:"false" `
    /*
        要操作的逆向退款ID     */
    RefundId  *int64 `json:"refund_id" required:"true" `
}

func (s *TaobaoFenxiaoReturngoodsSupRefuseRequest) SetRefuseReason(v string) *TaobaoFenxiaoReturngoodsSupRefuseRequest {
    s.RefuseReason = &v
    return s
}
func (s *TaobaoFenxiaoReturngoodsSupRefuseRequest) SetRefundId(v int64) *TaobaoFenxiaoReturngoodsSupRefuseRequest {
    s.RefundId = &v
    return s
}

func (req *TaobaoFenxiaoReturngoodsSupRefuseRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefuseReason != nil) {
        paramMap["refuse_reason"] = *req.RefuseReason
    }
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    return paramMap
}

func (req *TaobaoFenxiaoReturngoodsSupRefuseRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}