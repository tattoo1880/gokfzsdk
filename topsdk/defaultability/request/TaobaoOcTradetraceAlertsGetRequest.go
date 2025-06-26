package request


type TaobaoOcTradetraceAlertsGetRequest struct {
    /*
        异常类型代码：发货回写淘宝异常=1，商家系统漏单提醒=2，发货超时提醒=3，物流寄送超时=4，买家签收超时=5，物流中转异常=6，退货寄回异常=7，订单追回提醒=8"。     */
    AbnormalType  *int64 `json:"abnormal_type" required:"true" `
    /*
        返回数据的页码     */
    PageNo  *int64 `json:"page_no" required:"true" `
    /*
        返回数据每页包含的记录数目 defalutValue��100    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
}

func (s *TaobaoOcTradetraceAlertsGetRequest) SetAbnormalType(v int64) *TaobaoOcTradetraceAlertsGetRequest {
    s.AbnormalType = &v
    return s
}
func (s *TaobaoOcTradetraceAlertsGetRequest) SetPageNo(v int64) *TaobaoOcTradetraceAlertsGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoOcTradetraceAlertsGetRequest) SetPageSize(v int64) *TaobaoOcTradetraceAlertsGetRequest {
    s.PageSize = &v
    return s
}

func (req *TaobaoOcTradetraceAlertsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AbnormalType != nil) {
        paramMap["abnormal_type"] = *req.AbnormalType
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *TaobaoOcTradetraceAlertsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}