package pagarmeapisdk

import (
    "github.com/apimatic/go-core-runtime/https"
    "os"
)

type ConfigurationOptions func(*Configuration)

type Configuration struct {
    environment       Environment
    httpConfiguration https.HttpConfiguration
    basicAuthUserName string
    basicAuthPassword string
}

func newConfiguration(options ...ConfigurationOptions) Configuration {
    config := Configuration{}
    
    for _, option := range options {
        option(&config)
    }
    return config
}

func WithEnvironment(environment Environment) ConfigurationOptions {
    return func(c *Configuration) {
        c.environment = environment
    }
}

func WithHttpConfiguration(httpConfiguration https.HttpConfiguration) ConfigurationOptions {
    return func(c *Configuration) {
        c.httpConfiguration = httpConfiguration
    }
}

func WithBasicAuthUserName(basicAuthUserName string) ConfigurationOptions {
    return func(c *Configuration) {
        c.basicAuthUserName = basicAuthUserName
    }
}

func WithBasicAuthPassword(basicAuthPassword string) ConfigurationOptions {
    return func(c *Configuration) {
        c.basicAuthPassword = basicAuthPassword
    }
}

func (c *Configuration) Environment() Environment {
    return c.environment
}

func (c *Configuration) HttpConfiguration() https.HttpConfiguration {
    return c.httpConfiguration
}

func (c *Configuration) BasicAuthUserName() string {
    return c.basicAuthUserName
}

func (c *Configuration) BasicAuthPassword() string {
    return c.basicAuthPassword
}

func (c *Configuration) LoadFromEnvironment() {
    environment := os.Getenv("PAGARMEAPISDK_ENVIRONMENT")
    if environment != "" {
        c.environment = Environment(environment)
    }
    basicAuthUserName := os.Getenv("PAGARMEAPISDK_BASIC_AUTH_USER_NAME")
    if basicAuthUserName != "" {
        c.basicAuthUserName = basicAuthUserName
    }
    basicAuthPassword := os.Getenv("PAGARMEAPISDK_BASIC_AUTH_PASSWORD")
    if basicAuthPassword != "" {
        c.basicAuthPassword = basicAuthPassword
    }
}

// Available Servers
type Server string

const (
    ENUMDEFAULT Server = "default"
)

// Available Environments
type Environment string

const (
    PRODUCTION Environment = "production"
)

func RetryConfigurationFactory(options ...https.RetryConfigurationOptions) https.RetryConfiguration {
    config := DefaultRetryConfiguration()
    
    for _, option := range options {
        option(&config)
    }
    return config
}

func HttpConfigurationFactory(options ...https.HttpConfigurationOptions) https.HttpConfiguration {
    config := DefaultHttpConfiguration()
    
    for _, option := range options {
        option(&config)
    }
    return config
}

func ConfigurationFactory(options ...ConfigurationOptions) Configuration {
    config := DefaultConfiguration()
    
    for _, option := range options {
        option(&config)
    }
    return config
}
