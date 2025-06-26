package response

import (
    "topsdk/ability1092/domain"
)

type TaobaoRegionWarehouseQueryResponse struct {

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
    Result  domain.TaobaoRegionWarehouseQueryBaseResult `json:"result,omitempty" `
}
