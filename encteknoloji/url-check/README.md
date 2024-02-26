
# URL Checker Container

## Description

This containerized application periodically checks a specified URL and verifies it returns the expected status code. It reads configuration parameters from environment variables, allowing customization of the monitored URL, check interval, expected status code, and certificate verification behavior.

## Environment Variables
  
| Environment Variable | Describtion |
|--|--|
| CHECK_URL | The URL to be checked. (Required) |
| CHECK_INTERVAL | The interval between checks in seconds (default: 3). |
| CHECK_STATUS_CODE | The expected status code (default: 200, OK). |
| INSECURE_SKIP_VERIFY | Set to true to skip TLS certificate verification (WARNING: Use with caution!). Default is false. |

## Building and Running

Build the image:

    docker build -t encteknoloji/url-checker .

Run the container:

    docker run -d \
      -e CHECK_URL=https://example.com \
      -e CHECK_INTERVAL=10 \
      --name url-checker \
    encteknoloji/url-checker

Replace `https://example.com` with the URL you want to monitor and adjust `CHECK_INTERVAL` as needed.

## Disclaimer

Skipping TLS certificate verification can compromise security and should only be used in controlled environments where trust is well established.

## Additional Notes

This container can be integrated with monitoring tools or used in conjunction with other scripts for custom actions based on the check results.

Consider incorporating logging mechanisms for detailed information about the checks and potential errors.
