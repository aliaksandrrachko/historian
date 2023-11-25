# Historical events

## Environment variables

* GIN_MODE=[debug|release|test]


# Development

## Dockerfile

ARG GIT_COMMIT_HASH
ARG GIT_COMMIT_AUTHOR
ARG APP_VERSION

```
docker build -t hist-evo-backend \
    --no-cache \
    --build-arg GIT_COMMIT_HASH=$(git rev-parse HEAD) \
    --build-arg GIT_COMMIT_AUTHOR=$(git log -1 --format='%an' HEAD) \
    --build-arg APP_VERSION=v0.0.1-SNAPSHOT \
    .

docker run -it -p 8080:8080 --rm --name hist-evo-backend hist-evo-backend
```

## Ld flags

```
-ldflags="-X 'github.com/aliaksandrrachko/historian/historical-events/internal/build.version=$APP_VERSION' \
    -X 'github.com/aliaksandrrachko/historian/historical-events/internal/build.gitCommit=$GIT_COMMIT_HASH' \
    -X 'github.com/aliaksandrrachko/historian/historical-events/internal/build.gitAuthor=$GIT_COMMIT_AUTHOR' \
    -X 'github.com/aliaksandrrachko/historian/historical-events/internal/build.time=$(date)'"
```

# Deploy

## Create docker-registry secret

```
kubectl delete secret historian-gitlab-registry --namespace "$NAMESPACE"

kubectl create secret docker-registry historian-gitlab-registry \
      --docker-server registry.gitlab.com \
      --docker-email "aliaksandrrachko@gitlab.com" \
      --docker-username "$DEPLOY_TOKEN_USERNAME" \
      --docker-password "$DEPLOY_TOKEN_VALUE" \
      --namespace "$NAMESPACE"
```