package response

import (
    "topsdk/ability316/domain"
)

type AlibabaInteractLotteryactivityRegisterResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        接口返回model
    */
    Result  domain.AlibabaInteractLotteryactivityRegisterResult `json:"result,omitempty" `
}
