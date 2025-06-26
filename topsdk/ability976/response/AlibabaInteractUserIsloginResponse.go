package response

import (
    "topsdk/ability976/domain"
)

type AlibabaInteractUserIsloginResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        result
    */
    Result  domain.AlibabaInteractUserIsloginMtopResult `json:"result,omitempty" `
}
