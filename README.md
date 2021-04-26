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


### Documentation

- [Application Front](documentation/APPLICATION.md)
- [Platform](documentation/PLATFORM.md)
