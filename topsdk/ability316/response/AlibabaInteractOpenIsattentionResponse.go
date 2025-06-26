package response

import (
    "topsdk/ability316/domain"
)

type AlibabaInteractOpenIsattentionResponse struct {

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
    Result  domain.AlibabaInteractOpenIsattentionResultDo `json:"result,omitempty" `
}
