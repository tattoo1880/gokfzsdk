package response

import (
    "topsdk/ability229/domain"
)

type CainiaoWaybillCloudprintNetprintVerifycodeResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回值
    */
    Result  domain.CainiaoWaybillCloudprintNetprintVerifycodeIsvResult `json:"result,omitempty" `
}
