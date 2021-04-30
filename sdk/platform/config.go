package platform

// PlatformConfig provides configuration to a service client instance.
type PlatformConfig struct {
	CompanyID   string
	APIKey      string
	APISecret   string
	Domain      string
	OAuthClient *OAuthClient
}

//NewPlatformConfig provides platform configuration
func NewPlatformConfig(companyID, apiKey, apiSecret, domain string) *PlatformConfig {
	return &PlatformConfig{companyID, apiKey, apiSecret, domain, &OAuthClient{}}
}

//SetOAuthClient sets OAuthClient into platform configuration
func (p *PlatformConfig) SetOAuthClient() *OAuthClient {
	p.OAuthClient = NewOAuthClient(p)
	return p.OAuthClient
}

//GetAccessToken returns the access token
func (p *PlatformConfig) GetAccessToken() string {
	return p.OAuthClient.GetAccessToken()
}
