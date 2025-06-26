package response

import (
    "topsdk/ability316/domain"
)

type AlibabaInteractAopdataRegisterResponse struct {

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
    Result  domain.AlibabaInteractAopdataRegisterResult `json:"result,omitempty" `
}
