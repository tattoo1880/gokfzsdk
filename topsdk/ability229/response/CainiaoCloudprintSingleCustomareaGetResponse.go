package response

import (
    "topsdk/ability229/domain"
)

type CainiaoCloudprintSingleCustomareaGetResponse struct {

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
    Result  domain.CainiaoCloudprintSingleCustomareaGetCloudPrintBaseResult `json:"result,omitempty" `
}
