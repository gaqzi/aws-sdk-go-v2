// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartBuildInput struct {
	_ struct{} `type:"structure"`

	// Build output artifact settings that override, for this build only, the latest
	// ones already defined in the build project.
	ArtifactsOverride *ProjectArtifacts `locationName:"artifactsOverride" type:"structure"`

	// A buildspec file declaration that overrides, for this build only, the latest
	// one already defined in the build project.
	//
	// If this value is set, it can be either an inline buildspec definition, the
	// path to an alternate buildspec file relative to the value of the built-in
	// CODEBUILD_SRC_DIR environment variable, or the path to an S3 bucket. The
	// bucket must be in the same AWS Region as the build project. Specify the buildspec
	// file using its ARN (for example, arn:aws:s3:::my-codebuild-sample2/buildspec.yml).
	// If this value is not provided or is set to an empty string, the source code
	// must contain a buildspec file in its root directory. For more information,
	// see Buildspec File Name and Storage Location (https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec-ref-name-storage).
	BuildspecOverride *string `locationName:"buildspecOverride" type:"string"`

	// A ProjectCache object specified for this build that overrides the one defined
	// in the build project.
	CacheOverride *ProjectCache `locationName:"cacheOverride" type:"structure"`

	// The name of a certificate for this build that overrides the one specified
	// in the build project.
	CertificateOverride *string `locationName:"certificateOverride" type:"string"`

	// The name of a compute type for this build that overrides the one specified
	// in the build project.
	ComputeTypeOverride ComputeType `locationName:"computeTypeOverride" type:"string" enum:"true"`

	// The AWS Key Management Service (AWS KMS) customer master key (CMK) that overrides
	// the one specified in the build project. The CMK key encrypts the build output
	// artifacts.
	//
	// You can use a cross-account KMS key to encrypt the build output artifacts
	// if your service role has permission to that key.
	//
	// You can specify either the Amazon Resource Name (ARN) of the CMK or, if available,
	// the CMK's alias (using the format alias/alias-name ).
	EncryptionKeyOverride *string `locationName:"encryptionKeyOverride" min:"1" type:"string"`

	// A container type for this build that overrides the one specified in the build
	// project.
	EnvironmentTypeOverride EnvironmentType `locationName:"environmentTypeOverride" type:"string" enum:"true"`

	// A set of environment variables that overrides, for this build only, the latest
	// ones already defined in the build project.
	EnvironmentVariablesOverride []EnvironmentVariable `locationName:"environmentVariablesOverride" type:"list"`

	// The user-defined depth of history, with a minimum value of 0, that overrides,
	// for this build only, any previous depth of history defined in the build project.
	GitCloneDepthOverride *int64 `locationName:"gitCloneDepthOverride" type:"integer"`

	// Information about the Git submodules configuration for this build of an AWS
	// CodeBuild build project.
	GitSubmodulesConfigOverride *GitSubmodulesConfig `locationName:"gitSubmodulesConfigOverride" type:"structure"`

	// A unique, case sensitive identifier you provide to ensure the idempotency
	// of the StartBuild request. The token is included in the StartBuild request
	// and is valid for 12 hours. If you repeat the StartBuild request with the
	// same token, but change a parameter, AWS CodeBuild returns a parameter mismatch
	// error.
	IdempotencyToken *string `locationName:"idempotencyToken" type:"string"`

	// The name of an image for this build that overrides the one specified in the
	// build project.
	ImageOverride *string `locationName:"imageOverride" min:"1" type:"string"`

	// The type of credentials AWS CodeBuild uses to pull images in your build.
	// There are two valid values:
	//
	//    * CODEBUILD specifies that AWS CodeBuild uses its own credentials. This
	//    requires that you modify your ECR repository policy to trust AWS CodeBuild's
	//    service principal.
	//
	//    * SERVICE_ROLE specifies that AWS CodeBuild uses your build project's
	//    service role.
	//
	// When using a cross-account or private registry image, you must use SERVICE_ROLE
	// credentials. When using an AWS CodeBuild curated image, you must use CODEBUILD
	// credentials.
	ImagePullCredentialsTypeOverride ImagePullCredentialsType `locationName:"imagePullCredentialsTypeOverride" type:"string" enum:"true"`

	// Enable this flag to override the insecure SSL setting that is specified in
	// the build project. The insecure SSL setting determines whether to ignore
	// SSL warnings while connecting to the project source code. This override applies
	// only if the build's source is GitHub Enterprise.
	InsecureSslOverride *bool `locationName:"insecureSslOverride" type:"boolean"`

	// Log settings for this build that override the log settings defined in the
	// build project.
	LogsConfigOverride *LogsConfig `locationName:"logsConfigOverride" type:"structure"`

	// Enable this flag to override privileged mode in the build project.
	PrivilegedModeOverride *bool `locationName:"privilegedModeOverride" type:"boolean"`

	// The name of the AWS CodeBuild build project to start running a build.
	//
	// ProjectName is a required field
	ProjectName *string `locationName:"projectName" min:"1" type:"string" required:"true"`

	// The number of minutes a build is allowed to be queued before it times out.
	QueuedTimeoutInMinutesOverride *int64 `locationName:"queuedTimeoutInMinutesOverride" min:"5" type:"integer"`

	// The credentials for access to a private registry.
	RegistryCredentialOverride *RegistryCredential `locationName:"registryCredentialOverride" type:"structure"`

	// Set to true to report to your source provider the status of a build's start
	// and completion. If you use this option with a source provider other than
	// GitHub, GitHub Enterprise, or Bitbucket, an invalidInputException is thrown.
	//
	// The status of a build triggered by a webhook is always reported to your source
	// provider.
	ReportBuildStatusOverride *bool `locationName:"reportBuildStatusOverride" type:"boolean"`

	// An array of ProjectArtifacts objects.
	SecondaryArtifactsOverride []ProjectArtifacts `locationName:"secondaryArtifactsOverride" type:"list"`

	// An array of ProjectSource objects.
	SecondarySourcesOverride []ProjectSource `locationName:"secondarySourcesOverride" type:"list"`

	// An array of ProjectSourceVersion objects that specify one or more versions
	// of the project's secondary sources to be used for this build only.
	SecondarySourcesVersionOverride []ProjectSourceVersion `locationName:"secondarySourcesVersionOverride" type:"list"`

	// The name of a service role for this build that overrides the one specified
	// in the build project.
	ServiceRoleOverride *string `locationName:"serviceRoleOverride" min:"1" type:"string"`

	// An authorization type for this build that overrides the one defined in the
	// build project. This override applies only if the build project's source is
	// BitBucket or GitHub.
	SourceAuthOverride *SourceAuth `locationName:"sourceAuthOverride" type:"structure"`

	// A location that overrides, for this build, the source location for the one
	// defined in the build project.
	SourceLocationOverride *string `locationName:"sourceLocationOverride" type:"string"`

	// A source input type, for this build, that overrides the source input defined
	// in the build project.
	SourceTypeOverride SourceType `locationName:"sourceTypeOverride" type:"string" enum:"true"`

	// A version of the build input to be built, for this build only. If not specified,
	// the latest version is used. If specified, must be one of:
	//
	//    * For AWS CodeCommit: the commit ID, branch, or Git tag to use.
	//
	//    * For GitHub: the commit ID, pull request ID, branch name, or tag name
	//    that corresponds to the version of the source code you want to build.
	//    If a pull request ID is specified, it must use the format pr/pull-request-ID
	//    (for example pr/25). If a branch name is specified, the branch's HEAD
	//    commit ID is used. If not specified, the default branch's HEAD commit
	//    ID is used.
	//
	//    * For Bitbucket: the commit ID, branch name, or tag name that corresponds
	//    to the version of the source code you want to build. If a branch name
	//    is specified, the branch's HEAD commit ID is used. If not specified, the
	//    default branch's HEAD commit ID is used.
	//
	//    * For Amazon Simple Storage Service (Amazon S3): the version ID of the
	//    object that represents the build input ZIP file to use.
	//
	// If sourceVersion is specified at the project level, then this sourceVersion
	// (at the build level) takes precedence.
	//
	// For more information, see Source Version Sample with CodeBuild (https://docs.aws.amazon.com/codebuild/latest/userguide/sample-source-version.html)
	// in the AWS CodeBuild User Guide.
	SourceVersion *string `locationName:"sourceVersion" type:"string"`

	// The number of build timeout minutes, from 5 to 480 (8 hours), that overrides,
	// for this build only, the latest setting already defined in the build project.
	TimeoutInMinutesOverride *int64 `locationName:"timeoutInMinutesOverride" min:"5" type:"integer"`
}

// String returns the string representation
func (s StartBuildInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartBuildInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartBuildInput"}
	if s.EncryptionKeyOverride != nil && len(*s.EncryptionKeyOverride) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EncryptionKeyOverride", 1))
	}
	if s.ImageOverride != nil && len(*s.ImageOverride) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ImageOverride", 1))
	}

	if s.ProjectName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectName"))
	}
	if s.ProjectName != nil && len(*s.ProjectName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectName", 1))
	}
	if s.QueuedTimeoutInMinutesOverride != nil && *s.QueuedTimeoutInMinutesOverride < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("QueuedTimeoutInMinutesOverride", 5))
	}
	if s.ServiceRoleOverride != nil && len(*s.ServiceRoleOverride) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServiceRoleOverride", 1))
	}
	if s.TimeoutInMinutesOverride != nil && *s.TimeoutInMinutesOverride < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("TimeoutInMinutesOverride", 5))
	}
	if s.ArtifactsOverride != nil {
		if err := s.ArtifactsOverride.Validate(); err != nil {
			invalidParams.AddNested("ArtifactsOverride", err.(aws.ErrInvalidParams))
		}
	}
	if s.CacheOverride != nil {
		if err := s.CacheOverride.Validate(); err != nil {
			invalidParams.AddNested("CacheOverride", err.(aws.ErrInvalidParams))
		}
	}
	if s.EnvironmentVariablesOverride != nil {
		for i, v := range s.EnvironmentVariablesOverride {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "EnvironmentVariablesOverride", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.GitSubmodulesConfigOverride != nil {
		if err := s.GitSubmodulesConfigOverride.Validate(); err != nil {
			invalidParams.AddNested("GitSubmodulesConfigOverride", err.(aws.ErrInvalidParams))
		}
	}
	if s.LogsConfigOverride != nil {
		if err := s.LogsConfigOverride.Validate(); err != nil {
			invalidParams.AddNested("LogsConfigOverride", err.(aws.ErrInvalidParams))
		}
	}
	if s.RegistryCredentialOverride != nil {
		if err := s.RegistryCredentialOverride.Validate(); err != nil {
			invalidParams.AddNested("RegistryCredentialOverride", err.(aws.ErrInvalidParams))
		}
	}
	if s.SecondaryArtifactsOverride != nil {
		for i, v := range s.SecondaryArtifactsOverride {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SecondaryArtifactsOverride", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SecondarySourcesOverride != nil {
		for i, v := range s.SecondarySourcesOverride {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SecondarySourcesOverride", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SecondarySourcesVersionOverride != nil {
		for i, v := range s.SecondarySourcesVersionOverride {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SecondarySourcesVersionOverride", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SourceAuthOverride != nil {
		if err := s.SourceAuthOverride.Validate(); err != nil {
			invalidParams.AddNested("SourceAuthOverride", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartBuildOutput struct {
	_ struct{} `type:"structure"`

	// Information about the build to be run.
	Build *Build `locationName:"build" type:"structure"`
}

// String returns the string representation
func (s StartBuildOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartBuild = "StartBuild"

// StartBuildRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Starts running a build.
//
//    // Example sending a request using StartBuildRequest.
//    req := client.StartBuildRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/StartBuild
func (c *Client) StartBuildRequest(input *StartBuildInput) StartBuildRequest {
	op := &aws.Operation{
		Name:       opStartBuild,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartBuildInput{}
	}

	req := c.newRequest(op, input, &StartBuildOutput{})
	return StartBuildRequest{Request: req, Input: input, Copy: c.StartBuildRequest}
}

// StartBuildRequest is the request type for the
// StartBuild API operation.
type StartBuildRequest struct {
	*aws.Request
	Input *StartBuildInput
	Copy  func(*StartBuildInput) StartBuildRequest
}

// Send marshals and sends the StartBuild API request.
func (r StartBuildRequest) Send(ctx context.Context) (*StartBuildResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartBuildResponse{
		StartBuildOutput: r.Request.Data.(*StartBuildOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartBuildResponse is the response type for the
// StartBuild API operation.
type StartBuildResponse struct {
	*StartBuildOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartBuild request.
func (r *StartBuildResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
