package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoOcTradetraceAlertsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        异常订单数据
    */
    ResultList  []domain.TaobaoOcTradetraceAlertsGetSimpleAbnormalOrderDetail `json:"result_list,omitempty" `
}
