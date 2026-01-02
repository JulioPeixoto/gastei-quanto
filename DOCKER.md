# Docker Setup

## Arquitetura

Este projeto utiliza:
- 2 instancias da API (api1 e api2) rodando na porta 8080 internamente
- Nginx como load balancer distribuindo as requisicoes entre as duas instancias
- As APIs sao expostas individualmente nas portas 8081 e 8082
- O Nginx e exposto na porta 80

## Comandos

### Iniciar todos os servicos

```bash
docker-compose up -d
```

### Ver logs de todos os servicos

```bash
docker-compose logs -f
```

### Ver logs de um servico especifico

```bash
docker-compose logs -f api1
docker-compose logs -f api2
docker-compose logs -f nginx
```

### Parar todos os servicos

```bash
docker-compose down
```

### Reconstruir as imagens

```bash
docker-compose up -d --build
```

### Verificar status dos servicos

```bash
docker-compose ps
```

## Endpoints

### Atraves do Load Balancer (Nginx)
- API: http://localhost/api/v1
- Health: http://localhost/api/v1/health
- Swagger: http://localhost/swagger/index.html

### APIs Diretas
- API 1: http://localhost:8081/api/v1
- API 2: http://localhost:8082/api/v1

## Variaveis de Ambiente

Crie um arquivo `.env` na raiz do projeto:

```env
JWT_SECRET=seu-secret-key-super-seguro
```

## Health Checks

Os health checks estao configurados para verificar automaticamente se as APIs estao respondendo corretamente.

## Load Balancing

O Nginx distribui as requisicoes de forma round-robin entre as duas instancias da API.
