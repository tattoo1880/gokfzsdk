package request

import (
        "topsdk/util"
    )

type TaobaoUserSellerGetRequest struct {
    /*
        需要返回的字段列表，可选值为返回示例值中的可以看到的字段     */
    Fields  *[]string `json:"fields" required:"true" `
}

func (s *TaobaoUserSellerGetRequest) SetFields(v []string) *TaobaoUserSellerGetRequest {
    s.Fields = &v
    return s
}

func (req *TaobaoUserSellerGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    return paramMap
}

func (req *TaobaoUserSellerGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}