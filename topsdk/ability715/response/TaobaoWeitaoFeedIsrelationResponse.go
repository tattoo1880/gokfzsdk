package response

import (
)

type TaobaoWeitaoFeedIsrelationResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        是否关注
    */
    Result  int64 `json:"result,omitempty" `
}
