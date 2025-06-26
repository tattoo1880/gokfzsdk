package request


type TaobaoFenxiaoReturngoodsSupAgreeRequest struct {
    /*
        收货街道名称     */
    ReceiverStreetName  *string `json:"receiver_street_name,omitempty" required:"false" `
    /*
        收货人姓名     */
    ReceiverName  *string `json:"receiver_name,omitempty" required:"false" `
    /*
        收货市名称     */
    ReceiverCityName  *string `json:"receiver_city_name,omitempty" required:"false" `
    /*
        收货区名称     */
    ReceiverAreaName  *string `json:"receiver_area_name,omitempty" required:"false" `
    /*
        收货人手机号     */
    ReceiverMobilePhone  *string `json:"receiver_mobile_phone,omitempty" required:"false" `
    /*
        收货详细地址     */
    ReceiverDetailAddress  *string `json:"receiver_detail_address,omitempty" required:"false" `
    /*
        收货省名称     */
    ReceiverProvinceName  *string `json:"receiver_province_name,omitempty" required:"false" `
    /*
        要操作的逆向退款ID     */
    RefundId  *int64 `json:"refund_id" required:"true" `
}

func (s *TaobaoFenxiaoReturngoodsSupAgreeRequest) SetReceiverStreetName(v string) *TaobaoFenxiaoReturngoodsSupAgreeRequest {
    s.ReceiverStreetName = &v
    return s
}
func (s *TaobaoFenxiaoReturngoodsSupAgreeRequest) SetReceiverName(v string) *TaobaoFenxiaoReturngoodsSupAgreeRequest {
    s.ReceiverName = &v
    return s
}
func (s *TaobaoFenxiaoReturngoodsSupAgreeRequest) SetReceiverCityName(v string) *TaobaoFenxiaoReturngoodsSupAgreeRequest {
    s.ReceiverCityName = &v
    return s
}
func (s *TaobaoFenxiaoReturngoodsSupAgreeRequest) SetReceiverAreaName(v string) *TaobaoFenxiaoReturngoodsSupAgreeRequest {
    s.ReceiverAreaName = &v
    return s
}
func (s *TaobaoFenxiaoReturngoodsSupAgreeRequest) SetReceiverMobilePhone(v string) *TaobaoFenxiaoReturngoodsSupAgreeRequest {
    s.ReceiverMobilePhone = &v
    return s
}
func (s *TaobaoFenxiaoReturngoodsSupAgreeRequest) SetReceiverDetailAddress(v string) *TaobaoFenxiaoReturngoodsSupAgreeRequest {
    s.ReceiverDetailAddress = &v
    return s
}
func (s *TaobaoFenxiaoReturngoodsSupAgreeRequest) SetReceiverProvinceName(v string) *TaobaoFenxiaoReturngoodsSupAgreeRequest {
    s.ReceiverProvinceName = &v
    return s
}
func (s *TaobaoFenxiaoReturngoodsSupAgreeRequest) SetRefundId(v int64) *TaobaoFenxiaoReturngoodsSupAgreeRequest {
    s.RefundId = &v
    return s
}

func (req *TaobaoFenxiaoReturngoodsSupAgreeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ReceiverStreetName != nil) {
        paramMap["receiver_street_name"] = *req.ReceiverStreetName
    }
    if(req.ReceiverName != nil) {
        paramMap["receiver_name"] = *req.ReceiverName
    }
    if(req.ReceiverCityName != nil) {
        paramMap["receiver_city_name"] = *req.ReceiverCityName
    }
    if(req.ReceiverAreaName != nil) {
        paramMap["receiver_area_name"] = *req.ReceiverAreaName
    }
    if(req.ReceiverMobilePhone != nil) {
        paramMap["receiver_mobile_phone"] = *req.ReceiverMobilePhone
    }
    if(req.ReceiverDetailAddress != nil) {
        paramMap["receiver_detail_address"] = *req.ReceiverDetailAddress
    }
    if(req.ReceiverProvinceName != nil) {
        paramMap["receiver_province_name"] = *req.ReceiverProvinceName
    }
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    return paramMap
}

func (req *TaobaoFenxiaoReturngoodsSupAgreeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}