package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoUserSellerGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        用户
    */
    User  domain.TaobaoUserSellerGetUser `json:"user,omitempty" `
}
