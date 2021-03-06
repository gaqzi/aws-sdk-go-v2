package external

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/defaults"
	"github.com/aws/aws-sdk-go-v2/aws/ec2metadata"
	"github.com/aws/aws-sdk-go-v2/aws/ec2rolecreds"
	"github.com/aws/aws-sdk-go-v2/aws/endpointcreds"
	"github.com/aws/aws-sdk-go-v2/aws/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// ResolveDefaultAWSConfig will write default configuration values into the cfg
// value. It will write the default values, overwriting any previous value.
//
// This should be used as the first resolver in the slice of resolvers when
// resolving external configuration.
func ResolveDefaultAWSConfig(cfg *aws.Config, configs Configs) error {
	*cfg = defaults.Config()
	return nil
}

// ResolveCustomCABundle extracts the first instance of a custom CA bundle filename
// from the external configurations. It will update the HTTP Client's builder
// to be configured with the custom CA bundle.
//
// Config provider used:
// * CustomCABundleProvider
func ResolveCustomCABundle(cfg *aws.Config, configs Configs) error {
	pemCerts, found, err := GetCustomCABundle(configs)
	if err != nil {
		// TODO error handling, What is the best way to handle this?
		// capture previous errors continue. error out if all errors
		return err
	}
	if !found {
		return nil
	}

	type withTransportOptions interface {
		WithTransportOptions(...func(*http.Transport)) aws.HTTPClient
	}

	trOpts, ok := cfg.HTTPClient.(withTransportOptions)
	if !ok {
		return fmt.Errorf("unable to add custom RootCAs HTTPClient, "+
			"has no WithTransportOptions, %T", cfg.HTTPClient)
	}

	var appendErr error
	client := trOpts.WithTransportOptions(func(tr *http.Transport) {
		if tr.TLSClientConfig == nil {
			tr.TLSClientConfig = &tls.Config{}
		}
		if tr.TLSClientConfig.RootCAs == nil {
			tr.TLSClientConfig.RootCAs = x509.NewCertPool()
		}
		if !tr.TLSClientConfig.RootCAs.AppendCertsFromPEM(pemCerts) {
			appendErr = awserr.New("LoadCustomCABundleError",
				"failed to load custom CA bundle PEM file", nil)
		}
	})
	if appendErr != nil {
		return appendErr
	}

	cfg.HTTPClient = client
	return err
}

// ResolveRegion extracts the first instance of a Region from the Configs slice.
//
// Config providers used:
// * RegionProvider
func ResolveRegion(cfg *aws.Config, configs Configs) error {
	v, found, err := GetRegion(configs)
	if err != nil {
		// TODO error handling, What is the best way to handle this?
		// capture previous errors continue. error out if all errors
		return err
	}
	if !found {
		return nil
	}

	cfg.Region = v
	return nil
}

// ResolveCredentialsValue extracts the first instance of Credentials from the
// config slices.
//
// Config providers used:
// * CredentialsValueProvider
func ResolveCredentialsValue(cfg *aws.Config, configs Configs) error {
	v, found, err := GetCredentialsValue(configs)
	if err != nil {
		// TODO error handling, What is the best way to handle this?
		// capture previous errors continue. error out if all errors
		return err
	}
	if !found {
		return nil
	}

	cfg.Credentials = aws.StaticCredentialsProvider{Value: v}

	return nil
}

// ResolveEndpointCredentials will extract the credentials endpoint from the config
// slice. Using the endpoint, provided, to create a endpoint credential provider.
//
// Config providers used:
// * CredentialsEndpointProvider
func ResolveEndpointCredentials(cfg *aws.Config, configs Configs) error {
	v, found, err := GetCredentialsEndpoint(configs)
	if err != nil {
		// TODO error handling, What is the best way to handle this?
		// capture previous errors continue. error out if all errors
		return err
	}
	if !found {
		return nil
	}

	if err := validateLocalURL(v); err != nil {
		return err
	}

	cfgCp := cfg.Copy()
	cfgCp.EndpointResolver = aws.ResolveWithEndpointURL(v)

	provider := endpointcreds.New(cfgCp)
	provider.ExpiryWindow = 5 * time.Minute

	cfg.Credentials = provider

	return nil
}

const containerCredentialsEndpoint = "http://169.254.170.2"

// ResolveContainerEndpointPathCredentials will extract the container credentials
// endpoint from the config slice. Using the endpoint provided, to create a
// endpoint credential provider.
//
// Config providers used:
// * ContainerCredentialsEndpointPathProvider
func ResolveContainerEndpointPathCredentials(cfg *aws.Config, configs Configs) error {
	v, found, err := GetContainerCredentialsEndpointPath(configs)
	if err != nil {
		// TODO error handling, What is the best way to handle this?
		// capture previous errors continue. error out if all errors
		return err
	}
	if !found {
		return nil
	}

	cfgCp := cfg.Copy()

	v = containerCredentialsEndpoint + v
	cfgCp.EndpointResolver = aws.ResolveWithEndpointURL(v)

	provider := endpointcreds.New(cfgCp)
	provider.ExpiryWindow = 5 * time.Minute

	cfg.Credentials = provider

	return nil
}

// ResolveAssumeRoleCredentials extracts the assume role configuration from
// the external configurations.
//
// Config providers used:
func ResolveAssumeRoleCredentials(cfg *aws.Config, configs Configs) error {
	v, found, err := GetAssumeRoleConfig(configs)
	if err != nil {
		// TODO error handling, What is the best way to handle this?
		// capture previous errors continue. error out if all errors
		return err
	}
	if !found {
		return nil
	}

	cfgCp := cfg.Copy()
	// TODO support additional credential providers that are already set?
	cfgCp.Credentials = aws.StaticCredentialsProvider{Value: v.Source.Credentials}

	provider := stscreds.NewAssumeRoleProvider(
		sts.New(cfgCp), v.RoleARN,
	)
	provider.RoleSessionName = v.RoleSessionName

	if id := v.ExternalID; len(id) > 0 {
		provider.ExternalID = aws.String(id)
	}
	if len(v.MFASerial) > 0 {
		tp, foundTP, err := GetMFATokenFunc(configs)
		if err != nil {
			return err
		}
		if !foundTP {
			return fmt.Errorf("token provider required for AssumeRole with MFA")
		}
		provider.SerialNumber = aws.String(v.MFASerial)
		provider.TokenProvider = tp
	}

	cfg.Credentials = provider

	return nil
}

// ResolveFallbackEC2Credentials will configure the AWS config credentials to
// use EC2 Instance Role always.
func ResolveFallbackEC2Credentials(cfg *aws.Config, configs Configs) error {
	cfgCp := cfg.Copy()

	provider := ec2rolecreds.NewProvider(ec2metadata.New(cfgCp))
	provider.ExpiryWindow = 5 * time.Minute

	cfg.Credentials = provider

	return nil
}

// ResolveEnableEndpointDiscovery will configure the AWS config for Endpoint Discovery
// based on the first value discovered from the provided slice of configs.
func ResolveEnableEndpointDiscovery(cfg *aws.Config, configs Configs) error {
	endpointDiscovery, found, err := GetEnableEndpointDiscovery(configs)
	if err != nil {
		return err
	}

	if !found {
		return nil
	}

	cfg.EnableEndpointDiscovery = endpointDiscovery

	return nil
}
