# c

to learn https://www.sigbus.info/compilerbook

```
docker build -t ubuntugo .

# test
docker run --rm -it -v /$(pwd):/home/user ubuntugo go test

# bash
docker run --rm -it -v /$(pwd):/home/user ubuntugo bash
```