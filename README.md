# Simple K6

Grafana Dashboard: https://grafana.com/grafana/dashboards/4411-k6-load-testing-results/

![](doc/dashboard.jpg)

Ports:
- 8080: UI
- 3000: Grafana
- 5665: K6 Status, only work when k6 script is running

Note:
- K6 Report will save under `reports` directory. 

## Dependency

- Git
- Docker
- Docker Compose

## Usage

```bash
git clone https://github.com/timfanda35/simple-k6.git 
cd simple-k6

make docker-build
docker compose up -d
```

## Note

The container user is k6

```
uid=12345(k6) gid=12345(k6) groups=12345(k6)
```

We can grant permission on Ubuntu:

```bash
sudo apt-get install -y acl

setfacl -m u:12345:rwx reports
```
