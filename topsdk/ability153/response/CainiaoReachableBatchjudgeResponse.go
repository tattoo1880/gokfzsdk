package response

import (
    "topsdk/ability153/domain"
)

type CainiaoReachableBatchjudgeResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        结果
    */
    Result  domain.CainiaoReachableBatchjudgeBaseResultDto `json:"result,omitempty" `
}
