
Dockerbuild

```bash
docker buildx build --platform=linux/amd64 . -t long-running  --build-arg path=./long-running


docker image inspect -f '{{.Config.Entrypoint}} {{.Config.Cmd}}'  ghcr.io/atlanhq/go-leak:latest
```