package response

import (
    "topsdk/ability316/domain"
)

type TaobaoWirelessVideoCheckResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        系统自动生成
    */
    Result  domain.TaobaoWirelessVideoCheckRopResultTo `json:"result,omitempty" `
}
