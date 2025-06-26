package response

import (
    "topsdk/ability316/domain"
)

type AlibabaInteractLotterydrawDodrawResponse struct {

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
    Result  domain.AlibabaInteractLotterydrawDodrawResultDto `json:"result,omitempty" `
}
