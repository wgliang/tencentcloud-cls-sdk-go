package cls

// ClSCleint is a define of a TecentCloud Cloud Log Service cliet.
type ClSCleint struct {
	SecretId  string
	SecretKey string
	Host      string
}

// NewCLSClient
func NewCLSClient(secretId, secretKey, host string) *ClSCleint {
	return &ClSCleint{
		SecretId:  secretId,
		SecretKey: secretKey,
		Host:      host,
	}
}
