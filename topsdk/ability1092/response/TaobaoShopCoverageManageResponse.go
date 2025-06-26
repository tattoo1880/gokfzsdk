package response

import (
    "topsdk/ability1092/domain"
)

type TaobaoShopCoverageManageResponse struct {

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
    Result  domain.TaobaoShopCoverageManageBaseResult `json:"result,omitempty" `
}
