package response

import (
    "topsdk/ability1092/domain"
)

type TaobaoShopCoverageQueryResponse struct {

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
    Result  domain.TaobaoShopCoverageQueryBaseResult `json:"result,omitempty" `
}
