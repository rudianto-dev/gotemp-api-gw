package configuration

import (
	"log"
	"net/url"
)

func (cf ConfigurationSchema) ServiceList() map[string]*url.URL {
	var services = make(map[string]*url.URL)
	for _, service := range cf.Services {
		serviceURL, err := url.Parse(service.Host)
		if err != nil {
			log.Panic(err)
		}
		services[service.Name] = serviceURL
	}
	return services
}
