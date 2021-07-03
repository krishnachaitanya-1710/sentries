# sentries

Sentries is a security and compliance scanner tool, built to scan cloud infrastructure and to run compliance and security rules to mark violations.


## Overview

Organizations leveraging the cloud at scale require company policies tied to security, compliance and operational best practices to be adhered to at all times. Sentries makes it possible for organizations to enforce policies at different levels.

The key features of Sentries are:

* Policy as a Code: Sentries is a policy as a code framework that makes it possible to define organizational policies in areas such as security, compliance, or operational best practices as code. These policies can then be evaluated as part of an automated infrastructure provisioning workflow and also provides an ability to get on demand resource compliance status at any given point of time.

	```
		If a compliance scan fails, a warning message is printed/logged along with the violation context
	```

* Pre provisioning infrastructure compliance scan: Sentries tool contains in built infrastructure compliance rules that supports infrastructure components that resides in major cloud platforms such as aws, azure, google and supports validating your infrastructure prior to provisioning ( easily integrated with infrastructure provisionining workflows)
	
* Post provisioning infrastructure compliance scan


Contents
-----------
* [Introduction](#introduction)
* [components](#components)
* [examples usage](#example_usage)


<a name="introduction"></a>
## Introduction

<a name="components"></a>
## Components

<a name="example_usage"></a>
## Usage

An Example usage as shown below

```
	sentries inspect 
	
	response:
		 pass   - lambda function with out vpc rule
		 pass   - lambda function public subnet rule
		 pass   - lambda kmsKey encryption rule
		 pass   - lambda environment variable encryption check rule
		 fail   - s3 versioning rule
		 fail   - s3 kmsKey encryption rule
		 fail   - s3 default encryption rule
		 
```

## Contributing

We work hard to provide a high-quality and useful SDK for our Sentries services, and we greatly value feedback and contributions from our community. Please review our [contributing guidelines](./CONTRIBUTING.md) before submitting any [issues](https://github.com/krishnachaitanya-1710/sentries/issues) or [pull_requests](https://github.com/krishnachaitanya-1710/sentries/pulls) to ensure we have all the necessary information to effectively respond to your bug report or contribution.


Developing Sentinel
--------------------

This repository contains only Sentries core, which includes the command line interface, and the in built resource level compliance rules.


Resources
----------

AWS
-------

[Developer guide](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/welcome.html) - This document
is a general introduction on how to configure and make requests with the SDK.
If this is your first time using the SDK, this documentation and the API
documentation will help you get started. This document focuses on the syntax
and behavior of the SDK. The [Service Developer Guide](https://aws.amazon.com/documentation/)
will help you get started using specific AWS services.

[SDK API Reference Documentation](https://docs.aws.amazon.com/sdk-for-go/api/) - Use this
document to look up all API operation input and output parameters for AWS
services supported by the SDK. The API reference also includes documentation of
the SDK, and examples how to using the SDK, service client API operations, and
API operation require parameters.

[Service Documentation](https://aws.amazon.com/documentation/) - Use this
documentation to learn how to interface with AWS services. These guides are
great for getting started with a service, or when looking for more
information about a service. While this document is not required for coding,
services may supply helpful samples to look out for.

[SDK Examples](https://github.com/aws/aws-sdk-go/tree/main/example) -
Included in the SDK's repo are several hand crafted examples using the SDK
features and AWS services.

Azure
----------

- SDK docs are at [godoc.org](https://godoc.org/github.com/Azure/azure-sdk-for-go/).
- SDK samples are at [Azure-Samples/azure-sdk-for-go-samples](https://github.com/Azure-Samples/azure-sdk-for-go-samples).
- SDK notifications are published via the [Azure update feed](https://azure.microsoft.com/updates/).
- Azure API docs are at [docs.microsoft.com/rest/api](https://docs.microsoft.com/rest/api/).
- General Azure docs are at [docs.microsoft.com/azure](https://docs.microsoft.com/azure).


Google
----------

- [Go on Google Cloud](https://cloud.google.com/go/home)
- [Getting started with Go on Google Cloud](https://cloud.google.com/go/getting-started)
- [App Engine Quickstart](https://cloud.google.com/appengine/docs/standard/go/quickstart)
- [Cloud Functions Quickstart](https://cloud.google.com/functions/docs/quickstart-go)
- [Cloud Run Quickstart](https://cloud.google.com/run/docs/quickstarts/build-and-deploy#go)


License
-----------
[Apache License 2.0](https://github.com/krishnachaitanya-1710/sentries/blob/main/LICENSE)
