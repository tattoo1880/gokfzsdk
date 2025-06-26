package response

import (
    "topsdk/ability153/domain"
)

type CainiaoWaybillIiConfirmResponse struct {

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
    ConfirmResponse  []domain.CainiaoWaybillIiConfirmWaybillOrderConfirmResponse `json:"confirm_response,omitempty" `
}
