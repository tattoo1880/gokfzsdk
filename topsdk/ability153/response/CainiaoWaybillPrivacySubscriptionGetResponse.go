package response

import (
    "topsdk/ability153/domain"
)

type CainiaoWaybillPrivacySubscriptionGetResponse struct {

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
    Result  domain.CainiaoWaybillPrivacySubscriptionGetResult `json:"result,omitempty" `
}
