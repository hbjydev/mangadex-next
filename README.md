# mangadex-next

A Go port of the leaked Mangadex source code, as a proof of concept for a better
language implementation for the codebase. It also includes some Kubernetes-based
DevOps implementation, with GitHub Actions providing automatic updates.

API documentation in the form of an OpenAPI spec can be found in the `docs`
directory.

## Developing

This repository comes with a Docker Compose configuration used to bring up a dev
stack and the service components too. To start it, you can simply run...

```sh
$ docker-compose up -d
```

## Deployment

The `deploy` directory contains the Kubernetes manifests and other such files.

## Database Connection

The MySQL driver used in this stack requires a differently-formatted connection
URI from normal.

```env
DATABASE_URI="<username>:<password>@tcp(<host>:<port>)/<database>"
```

## Migrations

This repository uses the `migrate` CLI to run its migrations.

See https://github.com/golang-migrate/migrate/tree/master/cmd/migrate for an
installation guide.
