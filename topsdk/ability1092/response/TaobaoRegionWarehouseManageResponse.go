package response

import (
    "topsdk/ability1092/domain"
)

type TaobaoRegionWarehouseManageResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回结果
    */
    Result  domain.TaobaoRegionWarehouseManageBaseResult `json:"result,omitempty" `
}
