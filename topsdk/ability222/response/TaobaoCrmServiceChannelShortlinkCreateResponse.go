package response

import (
)

type TaobaoCrmServiceChannelShortlinkCreateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回的淘短链。
    */
    ShortLink  string `json:"short_link,omitempty" `
}
