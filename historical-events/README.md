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

# Known issue

container 'historical-events-backend' image 'registry.gitlab.com/aliaksandrrachko/historian:main.1085016981' not found or it is private and 'imagePullSecrets' is not properly configured.

2023-11-25 19:53:25.00 UTChistorical-events-historical-events-backend-8484577f4f-t2z6l[pod-event]Successfully assigned ales-litvin/historical-events-historical-events-backend-8484577f4f-t2z6l to gke-cloud-dev-3-de78e4eb-pf05
2023-11-25 19:53:26.00 UTChistorical-events-historical-events-backend-8484577f4f-t2z6l[pod-event]Pulling image "registry.gitlab.com/aliaksandrrachko/historian:main.1085016981"
2023-11-25 19:53:27.00 UTChistorical-events-historical-events-backend-8484577f4f-t2z6l[pod-event]Failed to pull image "registry.gitlab.com/aliaksandrrachko/historian:main.1085016981": rpc error: code = Unknown desc = failed to pull and unpack image "registry.gitlab.com/aliaksandrrachko/historian:main.1085016981": failed to resolve reference "registry.gitlab.com/aliaksandrrachko/historian:main.1085016981": failed to authorize: failed to fetch anonymous token: unexpected status from GET request to https://gitlab.com/jwt/auth?scope=repository%3Aaliaksandrrachko%2Fhistorian%3Apull&service=container_registry: 403 Forbidden
2023-11-25 19:53:27.00 UTChistorical-events-historical-events-backend-8484577f4f-t2z6l[pod-event]Error: ErrImagePull
