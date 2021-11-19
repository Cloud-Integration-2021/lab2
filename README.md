# Cloud-Integration - Lab 2

Golang API Rest for lab3

## 🛠️ Installation Steps

### 🐳 Option 1: Run from Docker run

```bash
# Run the container
$ docker run \
  -v /etc/localtime:/etc/localtime:ro \
  -e "ENV=dev" \
  -e "DB_HOST=localhost" \
  -e "DB_PORT=5432" \
  -e "DB_USER=postgres" \
  -e "DB_PASSWORD=postgres" \
  -e "DB_NAME=postgres" \
  --restart always \
  --name lab2 \
  -p 8081:8081 \
  thomaslacaze/lab2
```

### 🐳 Option 2: Run from Docker-compose

**See [here](https://github.com/Cloud-Integration-2021/lab2/blob/main/docker-compose.yml)** 

### 💻 Option 3: Run from source
#### Prerequisites
* NodeJS, yarn.

1. Clone the repository

```bash
git clone https://github.com/Cloud-Integration-2021/lab2.git
```

2. Change the working directory

```bash
cd lab2
```

1. Setup environnement variables

| Environment Variable | Default     | Description                                       |
| -------------------- | ----------- | ------------------------------------------------- |
| `ENV`                | `dev`       | Whether to use development or production settings |
| `DB_HOST`            | `localhost` | Hostname of the postgres server                   |
| `DB_PORT`            | `5432`      | Port of the postgres server                       |
| `DB_USER`            | `postgres`  | Username to connect to the postgres server        |
| `DB_PASSWORD`        | `postgres`  | Password to connect to the postgres server        |
| `DB_NAME`            | `postgres`  | Name of the database                              |

1. Run the app

```bash
$ go build -o lab2 && ./lab2
```


🌟 You are all set!

## Dockerfile
<a href="https://github.com/Cloud-Integration-2021/lab2/blob/main/Dockerfile">Dockerfile</a>

## License
<a href="https://github.com/Cloud-Integration-2021/lab2/blob/main/LICENSE">MIT</a>
