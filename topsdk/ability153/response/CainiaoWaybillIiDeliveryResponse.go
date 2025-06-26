package response

import (
    "topsdk/ability153/domain"
)

type CainiaoWaybillIiDeliveryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        确认结果
    */
    ConfirmResponse  domain.CainiaoWaybillIiDeliveryWaybillOrderConfirmResponse `json:"confirm_response,omitempty" `
}
