# KMS-Decrypter
Decrypt all files in folder then output to destination folder by AWS KMS.

## Usage
```
Usage of decrypter:
  -f string
        path to encrypted folder
  -m string
        aws or sops, default is 'aws' (default "aws")
  -n int
        number of worker, default is number of cpu cores
  -q    quiet mode, default is false
  -t string
        path where decrypted file will be writen to
```

## Example
We have this list
```
encrypted/.env
encrypted/secret/file.json
```

Run this command will decrypt all files:
```bash
./decrypter -f encrypted -t .
```

There file will be created
```
.env
secret/file.json
```

## Example when use with docker
In your entrypoint.sh
```bash
#!/bin/sh
set -e

decrypter -f ".encrypted/" -t .

exec "$@"
```

In your Dockerfile
```Dockerfile
FROM ghcr.io/yuyuvn/kms-decypter:v2.0.0-alpine AS decrypter
FROM mozilla/sops:v3-alpine AS sops
# main image
FROM ...

COPY --from=decrypter /decrypter /usr/local/bin/decrypter
COPY --from=sops /usr/local/bin/sops /usr/local/bin/sops

RUN chmod +x /usr/local/bin/*
```
