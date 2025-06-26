package domain


type TaobaoContentMediaUploadSecretHzGetStsUploadTokenResponse struct {
    /*
        oss接入域名     */
    Endpoint  *string `json:"endpoint,omitempty" `

    /*
        文件上传OSS时需要的access_key_id     */
    AccessId  *string `json:"access_id,omitempty" `

    /*
        文件上传OSS时需要的access_key_secret     */
    AccessSecret  *string `json:"access_secret,omitempty" `

    /*
        security_token的过期UTC时间，未过期前token可以重复使用     */
    Expiration  *string `json:"expiration,omitempty" `

    /*
        文件上传的sts token     */
    SecurityToken  *string `json:"security_token,omitempty" `

    /*
        oss存储 bucket     */
    Bucket  *string `json:"bucket,omitempty" `

    /*
        有权限的存储路径     */
    Path  *string `json:"path,omitempty" `

}

func (s *TaobaoContentMediaUploadSecretHzGetStsUploadTokenResponse) SetEndpoint(v string) *TaobaoContentMediaUploadSecretHzGetStsUploadTokenResponse {
    s.Endpoint = &v
    return s
}
func (s *TaobaoContentMediaUploadSecretHzGetStsUploadTokenResponse) SetAccessId(v string) *TaobaoContentMediaUploadSecretHzGetStsUploadTokenResponse {
    s.AccessId = &v
    return s
}
func (s *TaobaoContentMediaUploadSecretHzGetStsUploadTokenResponse) SetAccessSecret(v string) *TaobaoContentMediaUploadSecretHzGetStsUploadTokenResponse {
    s.AccessSecret = &v
    return s
}
func (s *TaobaoContentMediaUploadSecretHzGetStsUploadTokenResponse) SetExpiration(v string) *TaobaoContentMediaUploadSecretHzGetStsUploadTokenResponse {
    s.Expiration = &v
    return s
}
func (s *TaobaoContentMediaUploadSecretHzGetStsUploadTokenResponse) SetSecurityToken(v string) *TaobaoContentMediaUploadSecretHzGetStsUploadTokenResponse {
    s.SecurityToken = &v
    return s
}
func (s *TaobaoContentMediaUploadSecretHzGetStsUploadTokenResponse) SetBucket(v string) *TaobaoContentMediaUploadSecretHzGetStsUploadTokenResponse {
    s.Bucket = &v
    return s
}
func (s *TaobaoContentMediaUploadSecretHzGetStsUploadTokenResponse) SetPath(v string) *TaobaoContentMediaUploadSecretHzGetStsUploadTokenResponse {
    s.Path = &v
    return s
}
