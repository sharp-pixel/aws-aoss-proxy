# AWS AOSS Proxy

The AWS AOSS Proxy will sign incoming HTTP requests and forward them to the host specified in the `Host` header.

You can strip out arbitrary headers from the incoming request by using the -s option.

## Getting Started

Build and run the Proxy

The proxy uses the default AWS SDK for Go credential search path:

* Environment variables.
* Shared credentials file.
* IAM role for Amazon EC2 or ECS task role

More information can be found in the [developer guide](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html)

```bash
docker build -t aws-sigv4-proxy .

# Env vars
docker run --rm -ti \
  -e 'AWS_ACCESS_KEY_ID=<YOUR ACCESS KEY ID>' \
  -e 'AWS_SECRET_ACCESS_KEY=<YOUR SECRET ACCESS KEY>' \
  -p 8080:8080 \
  aws-sigv4-proxy -v

# Shared Credentials
docker run --rm -ti \
  -v ~/.aws:/root/.aws \
  -p 8080:8080 \
  -e 'AWS_SDK_LOAD_CONFIG=true' \
  -e 'AWS_PROFILE=<SOME PROFILE>' \
  aws-sigv4-proxy -v
```

### Configuration

When running the Proxy, the following flags can be used (none are required) :

| Flag (or short form)          | Type     | Description                                              | Default |
|-------------------------------|----------|----------------------------------------------------------|---------|
| `verbose` or `v`              | Boolean  | Enable additional logging, implies all the log-* options | `False` |
| `log-failed-requests`         | Boolean  | Log 4xx and 5xx response body                            | `False` |
| `log-signing-process`         | Boolean  | Log sigv4 signing process                                | `False` |
| `port`                        | String   | Port to serve http on                                    | `8080`  |
| `strip` or `s`                | String   | Headers to strip from incoming request                   | None    |
| `role-arn`                    | String   | Amazon Resource Name (ARN) of the role to assume         | None    |
| `name`                        | String   | AWS Service to sign for                                  | None    |
| `host`                        | String   | Host to proxy to                                         | None    |
| `region`                      | String   | AWS region to sign for                                   | None    |
| `no-verify-ssl`               | Boolean  | Disable peer SSL certificate validation                  | `False` |
| `transport.idle-conn-timeout` | Duration | Idle timeout to the upstream service                     | `40s`   |

## Examples

Amazon OpenSearch Service (Serverless)

```sh
curl -H 'host: <REST_API_ID>.aoss.<AWS_REGION>.amazonaws.com' http://localhost:9200/<PATH>
```

Running the service with Assume Role to use temporary credentials

```sh
docker run --rm -ti \
  -v ~/.aws:/root/.aws \
  -p 8080:8080 \
  -e 'AWS_SDK_LOAD_CONFIG=true' \
  -e 'AWS_PROFILE=<SOME PROFILE>' \
  aws-aoss-proxy -v --role-arn <ARN OF ROLE TO ASSUME>
```

Include service name & region overrides when you notice errors like `unable to determine service from host` for API gateway, for example.

```sh
docker run --rm -ti \
  -v ~/.aws:/root/.aws \
  -p 8080:8080 \
  -e 'AWS_SDK_LOAD_CONFIG=true' \
  -e 'AWS_PROFILE=<SOME PROFILE>' \
  aws-aoss-proxy -v --name execute-api --region us-east-1
```

## Reference

- [AWS SigV4 Signing Docs ](https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html)
- [AWS SigV4 Admission Controller](https://github.com/aws-observability/aws-sigv4-proxy-admission-controller) - Used to install the AWS SigV4 Proxy as a sidecar

## License

This library is licensed under the Apache 2.0 License.
