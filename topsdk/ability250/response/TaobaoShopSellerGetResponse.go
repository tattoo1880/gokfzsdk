package response

import (
    "topsdk/ability250/domain"
)

type TaobaoShopSellerGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        店铺信息
    */
    Shop  domain.TaobaoShopSellerGetShop `json:"shop,omitempty" `
}
