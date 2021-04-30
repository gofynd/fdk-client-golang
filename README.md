# FDK Golang

FDK client for Golang

## Getting Started

Get started with the Golang Development SDK for Fynd Platform

### Usage

```
import "https://github.com/gofynd/fdk-client-golang/"
```


### Sample Usage (ApplicationClient):

```golang
config, err := application.NewAppConfig("YOUR_APPLICATION_ID", "YOUR_APPLICATION_TOKEN", "DOMAIN", &application.Options{})
if err != nil {
	log.Println(err)
	return
}
client := application.NewAppClient(config);
res, err := client.Content.GetLandingPage()
if err != nil {
	log.Println(err)
	return
}

```

### Sample Usage (PlatformClient):

```golang
platformConfig := platform.NewPlatformConfig("YOUR_COMPANY_ID", "API_KEY", "API_SECRET", "DOMAIN")

platformConfig.SetOAuthClient()
//Set token using OAuthClient

platformClient := platform.NewPlatformClient(platformConfig)
res, err := platformClient.Catalog.GetProducts(platform.PlatformGetProductsXQuery{})
if err != nil {
	log.Println(err)
	return
}

```


### Documentation

- [Application Front](documentation/APPLICATION.md)
- [Platform](documentation/PLATFORM.md)
