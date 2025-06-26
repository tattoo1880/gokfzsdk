package request


type TaobaoFenxiaoReturngoodsSupRefuseconfirmRequest struct {
    /*
        拒绝原因     */
    RefuseReason  *string `json:"refuse_reason,omitempty" required:"false" `
    /*
        要操作的逆向退款ID     */
    RefundId  *int64 `json:"refund_id" required:"true" `
}

func (s *TaobaoFenxiaoReturngoodsSupRefuseconfirmRequest) SetRefuseReason(v string) *TaobaoFenxiaoReturngoodsSupRefuseconfirmRequest {
    s.RefuseReason = &v
    return s
}
func (s *TaobaoFenxiaoReturngoodsSupRefuseconfirmRequest) SetRefundId(v int64) *TaobaoFenxiaoReturngoodsSupRefuseconfirmRequest {
    s.RefundId = &v
    return s
}

func (req *TaobaoFenxiaoReturngoodsSupRefuseconfirmRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefuseReason != nil) {
        paramMap["refuse_reason"] = *req.RefuseReason
    }
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    return paramMap
}

func (req *TaobaoFenxiaoReturngoodsSupRefuseconfirmRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}