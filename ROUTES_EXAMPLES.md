# Exemplos de Payloads para Rotas de Interação com Filmes

## 1. Favoritar Filme

**Rota:** `POST /api/v1/movies/favorite`  
**Autenticação:** Requerida (Bearer Token)

### Payload de Exemplo:

```json
{
  "movieId": "123e4567-e89b-12d3-a456-426614174000"
}
```

### Resposta de Sucesso (200 OK):

```json
{
  "success": true,
  "message": "Movie favorited successfully"
}
```

---

## 2. Adicionar à Lista "Para Assistir" (To Watch)

**Rota:** `POST /api/v1/movies/to-watch`  
**Autenticação:** Requerida (Bearer Token)

### Payload de Exemplo:

```json
{
  "movieId": "123e4567-e89b-12d3-a456-426614174000"
}
```

### Resposta de Sucesso (200 OK):

```json
{
  "success": true,
  "message": "Movie added to watch list successfully"
}
```

---

## 3. Criar Avaliação (Watched/Rate)

**Rota:** `POST /api/v1/movies/watched`  
**Autenticação:** Requerida (Bearer Token)

### Payload de Exemplo:

```json
{
  "movieId": "123e4567-e89b-12d3-a456-426614174000",
  "rate": 8.5,
  "description": "Excelente filme! A atuação foi incrível e a história muito envolvente."
}
```

**Nota:** O campo `description` é opcional. O campo `rate` é obrigatório e deve ser um número decimal.

### Resposta de Sucesso (200 OK):

```json
{
  "success": true,
  "message": "Watched entry created successfully"
}
```

### Exemplo com apenas rate (sem description):

```json
{
  "movieId": "123e4567-e89b-12d3-a456-426614174000",
  "rate": 7.0
}
```

---

## Exemplos de Requisições cURL

### 1. Favoritar Filme:

```bash
curl -X POST http://localhost:8000/api/v1/movies/favorite \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer SEU_TOKEN_AQUI" \
  -d '{
    "movieId": "123e4567-e89b-12d3-a456-426614174000"
  }'
```

### 2. Adicionar à Lista "Para Assistir":

```bash
curl -X POST http://localhost:8000/api/v1/movies/to-watch \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer SEU_TOKEN_AQUI" \
  -d '{
    "movieId": "123e4567-e89b-12d3-a456-426614174000"
  }'
```

### 3. Criar Avaliação:

```bash
curl -X POST http://localhost:8000/api/v1/movies/watched \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer SEU_TOKEN_AQUI" \
  -d '{
    "movieId": "123e4567-e89b-12d3-a456-426614174000",
    "rate": 8.5,
    "description": "Excelente filme!"
  }'
```

---

## Observações Importantes

1. **Autenticação:** Todas as rotas requerem autenticação via Bearer Token no header `Authorization`
2. **movieId:** Deve ser um UUID válido
3. **rate:** Deve ser um número decimal (float64) para a rota de watched
4. **description:** Campo opcional apenas para a rota de watched
5. **Comportamento:** Se o registro já existir (favorite/to-watch), ele será atualizado. Para watched, se já existir uma avaliação do mesmo usuário para o mesmo filme, ela será atualizada.

