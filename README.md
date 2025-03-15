# Golang Electerm sync server

A golang server to sync electerm data.

## Installation

### Docker deployment

Pull image:

```bash
docker pull giterhub/electerm-sync-server:latest
```

Create the [`compose.yaml`](./compose.yaml) file and run:

```bash
docker compose up -d
```

### Local compile

Install [git](https://git-scm.com) and [Go](https://go.dev) locally and run:

```bash
# Clone project code
git clone https://github.com/dragonish/electerm-sync-server-go.git

# Enter project directory
cd electerm-sync-server-go

# Run
ELECTERM_JWT_SECRET=abcdef123 ELECTERM_JWT_USERS=user1 go run main.go
```

## Usage

See: [Custom sync server](https://github.com/electerm/electerm/wiki/Custom-sync-server).

API URL: `http://<domain>:7837/api/sync`.

## Environment variables

| Name | Type | Default | Description |
| --- | --- | --- | --- |
| `ELECTERM_JWT_SECRET` | `string` | `""` | JWT secret |
| `ELECTERM_JWT_USERS` | `string` | `""` | JWT users. Multiple users separated by comma(`,`) |
| `ELECTERM_PORT` | `int` | `7837` | Web service running port |

## Credits

- [electerm/electerm](https://github.com/electerm/electerm)
- [electerm/electerm-sync-server-node](https://github.com/electerm/electerm-sync-server-node)

## License

[MIT](./LICENSE)
