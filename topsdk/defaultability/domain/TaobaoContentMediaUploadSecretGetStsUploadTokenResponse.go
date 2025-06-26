package domain


type TaobaoContentMediaUploadSecretGetStsUploadTokenResponse struct {
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
        oss存储目标上传bucket     */
    Bucket  *string `json:"bucket,omitempty" `

}

func (s *TaobaoContentMediaUploadSecretGetStsUploadTokenResponse) SetAccessId(v string) *TaobaoContentMediaUploadSecretGetStsUploadTokenResponse {
    s.AccessId = &v
    return s
}
func (s *TaobaoContentMediaUploadSecretGetStsUploadTokenResponse) SetAccessSecret(v string) *TaobaoContentMediaUploadSecretGetStsUploadTokenResponse {
    s.AccessSecret = &v
    return s
}
func (s *TaobaoContentMediaUploadSecretGetStsUploadTokenResponse) SetExpiration(v string) *TaobaoContentMediaUploadSecretGetStsUploadTokenResponse {
    s.Expiration = &v
    return s
}
func (s *TaobaoContentMediaUploadSecretGetStsUploadTokenResponse) SetSecurityToken(v string) *TaobaoContentMediaUploadSecretGetStsUploadTokenResponse {
    s.SecurityToken = &v
    return s
}
func (s *TaobaoContentMediaUploadSecretGetStsUploadTokenResponse) SetBucket(v string) *TaobaoContentMediaUploadSecretGetStsUploadTokenResponse {
    s.Bucket = &v
    return s
}
