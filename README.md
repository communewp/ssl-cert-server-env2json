**ssl-cert-server-env2json** generates a configuration file compatible with [jxskiss/ssl-cert-server](https://github.com/jxskiss/ssl-cert-server).

## Configuration

The following environmental variables are available:

| Environmental Variable   | Configuration Key           | Default Value              | Required | Notes                                                                                                                     |
|--------------------------|-----------------------------|----------------------------|----------|---------------------------------------------------------------------------------------------------------------------------|
| `SSL_LISTEN`             | `listen`                    | `0.0.0.0:8999`             | no       |                                                                                                                           |
| `SSL_PID_FILE`           | `pid_file`                  | `/tmp/ssl-cert-server.pid` | no       |                                                                                                                           |
| `SSL_STORAGE_TYPE`       | `storage.type`              | `dir_cache`                | no       | May be either `dir_cache` or `redis`                                                                                      |
| `SSL_STORAGE_DIR_CACHE`  | `storage.dir_cache`         | `/data`                    | no       | Only used if `dir_cache` option is selected for `SSL_STORAGE_TYPE`                                                        |
| `SSL_STORAGE_REDIS_ADDR` | `storage-redis-addr`        | `redis:6379`               | no       | Only used if `redis` option is selected for `SSL_STORAGE_TYPE`                                                            |
| `SSL_LE_STAGING`         | `lets_encrypt.staging`      | `false`                    | no       |                                                                                                                           |
| `SSL_LE_FORCE_RSA`       | `lets_encrypt.force_rsa`    | `false`                    | no       |                                                                                                                           |
| `SSL_LE_RENEW_BEFORE`    | `lets_encrypt.renew_before` | `30`                       | no       |                                                                                                                           |
| `SSL_LE_EMAIL`           | `lets_encrypt.email`        | N/A                        | yes      | Email for Lets Encrypt notifications                                                                                      |
| `SSL_LE_DOMAINS`         | `lets_encrypt.domains`      | `[]`                       | no       | Array of allowed domains, separated by only commas, e.g. `example.com,example.org`                                        |
| `SSL_LE_REGEX_PATTERNS`  | `lets_encrypt.re_patterns`  | `[]`                       | no       | Array of allowed regex patterns, separated by only commas, e.g. `api1-(\\w+)\\.example\\.com,api2-(\\w+)\\.example\\.com` |

Currently, **ssl-cert-server-env2json** only supports certificate management configuration components for **Lets Encrypt**; **managed** and **self signed** certificate configuration components are not yet available.

## Example Usage

```bash
SSL_LE_EMAIL=demo@example.org ./ssl-cert-server-env2json > ./config.yaml
```
