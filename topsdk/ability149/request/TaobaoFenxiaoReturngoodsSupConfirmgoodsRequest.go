package request


type TaobaoFenxiaoReturngoodsSupConfirmgoodsRequest struct {
    /*
        确认收货的原因     */
    ConfirmGoodsReason  *string `json:"confirm_goods_reason,omitempty" required:"false" `
    /*
        要操作的逆向退款ID     */
    RefundId  *int64 `json:"refund_id" required:"true" `
}

func (s *TaobaoFenxiaoReturngoodsSupConfirmgoodsRequest) SetConfirmGoodsReason(v string) *TaobaoFenxiaoReturngoodsSupConfirmgoodsRequest {
    s.ConfirmGoodsReason = &v
    return s
}
func (s *TaobaoFenxiaoReturngoodsSupConfirmgoodsRequest) SetRefundId(v int64) *TaobaoFenxiaoReturngoodsSupConfirmgoodsRequest {
    s.RefundId = &v
    return s
}

func (req *TaobaoFenxiaoReturngoodsSupConfirmgoodsRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ConfirmGoodsReason != nil) {
        paramMap["confirm_goods_reason"] = *req.ConfirmGoodsReason
    }
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    return paramMap
}

func (req *TaobaoFenxiaoReturngoodsSupConfirmgoodsRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}