package response

import (
    "topsdk/ability229/domain"
)

type CainiaoCloudprintClientinfoPutResponse struct {

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
    Result  domain.CainiaoCloudprintClientinfoPutCloudPrintBaseResult `json:"result,omitempty" `
}
