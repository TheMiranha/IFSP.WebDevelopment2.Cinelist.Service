# Arquitetura Clean Architecture

Este documento descreve como a aplicação **Cinelist Service** segue os princípios da Clean Architecture.

## Visão Geral

A aplicação está organizada em camadas bem definidas, seguindo os princípios da Clean Architecture proposta por Robert C. Martin (Uncle Bob). A arquitetura garante que as dependências apontem sempre para o centro (domínio), mantendo o código de negócio independente de frameworks, bancos de dados e detalhes de implementação.

## Estrutura de Camadas

### 1. **Domain (Domínio)** - Camada Central

A camada mais interna, contém as regras de negócio puras e não depende de nenhuma outra camada.

#### `domain/entities/`
Contém as entidades do domínio que representam os conceitos principais da aplicação:
- `User` - Usuário do sistema
- `Movie` - Filme
- `Actor` - Ator
- `Favorite` - Filme favoritado
- `ToWatch` - Filme para assistir
- `Watched` - Filme assistido
- `Cast` - Elenco

#### `domain/dtos/`
Contém os Data Transfer Objects (DTOs) usados para comunicação entre camadas:
- DTOs de requisição e resposta
- DTOs de transferência de dados entre camadas

#### `domain/repositories/`
**Interfaces** que definem os contratos para acesso a dados:
- `UserRepository` - Interface para operações de usuário
- `MovieRepository` - Interface para operações de filme
- `ActorRepository` - Interface para operações de ator
- `MovieInteractionRepository` - Interface para interações do usuário com filmes

**Importante**: Estas são apenas interfaces, sem implementações. As implementações ficam na camada de infraestrutura.

#### `domain/services/`
Interfaces para serviços externos necessários ao domínio:
- `AuthService` - Interface para serviços de autenticação (hash de senha, JWT)

### 2. **Application (Aplicação)** - Casos de Uso

Camada que contém a lógica de aplicação e orquestra as operações do domínio.

#### `application/usecases/`
Contém os casos de uso da aplicação:
- `AuthenticationUseCase` - Autenticação e registro de usuários
- `UserUseCase` - Operações relacionadas a usuários
- `MovieUseCase` - Operações relacionadas a filmes
- `ActorUseCase` - Operações relacionadas a atores
- `MovieInteractionUseCase` - Interações do usuário com filmes (favoritar, assistir, etc.)

**Características importantes**:
- Os use cases dependem apenas de **interfaces** do domínio (`domain/repositories`, `domain/services`)
- Não conhecem detalhes de implementação (banco de dados, frameworks HTTP, etc.)
- Contêm a lógica de negócio da aplicação

### 3. **Infrastructure (Infraestrutura)** - Detalhes de Implementação

Camada externa que implementa os contratos definidos no domínio.

#### `infrastructure/database/repositories/`
Implementações concretas das interfaces de repositório:
- `UserRepository` - Implementação usando SQL
- `MovieRepository` - Implementação usando SQL
- `ActorRepository` - Implementação usando SQL
- `MovieInteractionRepository` - Implementação usando SQL

**Características**:
- Implementam as interfaces definidas em `domain/repositories/`
- Contêm detalhes específicos de banco de dados (SQL queries)
- Podem ser substituídas sem afetar o domínio ou casos de uso

#### `infrastructure/services/`
Implementações concretas dos serviços do domínio:
- `AuthService` - Implementação usando bcrypt e JWT

#### `infrastructure/http/`
Camada de apresentação HTTP:
- `controllers/` - Controladores HTTP que recebem requisições e chamam os use cases
- `middlewares/` - Middlewares HTTP (autenticação, CORS)
- `http.go` - Configuração do servidor HTTP (Gin)

**Características**:
- Depende dos use cases da camada de aplicação
- Não contém lógica de negócio
- Responsável apenas por receber requisições HTTP e retornar respostas

### 4. **CMD (Comandos)** - Pontos de Entrada

#### `cmd/api/`
Ponto de entrada principal da aplicação API:
- `main.go` - Inicializa o banco de dados e o servidor HTTP

#### `cmd/migrate/`
Ponto de entrada para migrações de banco de dados

## Princípios da Clean Architecture Aplicados

### 1. **Inversão de Dependências (Dependency Inversion)**

As dependências sempre apontam para dentro (domínio):
- ✅ Use cases dependem de **interfaces** do domínio, não de implementações
- ✅ Repositórios implementam interfaces do domínio
- ✅ Serviços implementam interfaces do domínio
- ✅ Controllers dependem de use cases, não de repositórios

**Exemplo**:
```go
// ❌ ANTES (violação)
type MovieUseCase struct {
    repo repositories.MovieRepository  // Dependência direta da infraestrutura
}

// ✅ DEPOIS (correto)
type MovieUseCase struct {
    repo domain_repositories.MovieRepository  // Dependência da interface do domínio
}
```

### 2. **Independência de Frameworks**

O domínio e os casos de uso não dependem de frameworks externos:
- ✅ Nenhuma importação de Gin, SQL, etc. no domínio
- ✅ Nenhuma importação de frameworks nos use cases
- ✅ Frameworks são usados apenas na camada de infraestrutura

### 3. **Testabilidade**

A arquitetura facilita testes:
- ✅ Use cases podem ser testados com mocks das interfaces
- ✅ Repositórios podem ser testados independentemente
- ✅ Lógica de negócio isolada e testável

### 4. **Independência de UI**

A lógica de negócio não depende de como a aplicação é apresentada:
- ✅ Use cases não conhecem HTTP
- ✅ Pode-se trocar Gin por outro framework sem afetar o domínio

### 5. **Independência de Banco de Dados**

A lógica de negócio não depende de detalhes de persistência:
- ✅ Use cases não conhecem SQL
- ✅ Pode-se trocar PostgreSQL por MongoDB sem afetar o domínio
- ✅ Interfaces de repositório abstraem os detalhes de persistência

## Fluxo de Dados

```
HTTP Request
    ↓
Controller (infrastructure/http/controllers)
    ↓
Use Case (application/usecases)
    ↓
Repository Interface (domain/repositories)
    ↓
Repository Implementation (infrastructure/database/repositories)
    ↓
Database
```

## Injeção de Dependências

A injeção de dependências é feita no ponto de entrada (`infrastructure/http/http.go`):

```go
// Criação das implementações concretas
userRepository := repositories.NewUserRepository(database)
authService := infrastructure_services.NewAuthService()

// Injeção nas interfaces do domínio
authenticationUseCase := usecases.NewAuthenticationUseCase(userRepository, authService)
```

## Benefícios da Arquitetura

1. **Manutenibilidade**: Código organizado e fácil de entender
2. **Testabilidade**: Cada camada pode ser testada independentemente
3. **Flexibilidade**: Fácil trocar implementações (banco de dados, frameworks)
4. **Escalabilidade**: Fácil adicionar novos recursos sem quebrar código existente
5. **Desacoplamento**: Mudanças em uma camada não afetam outras

## Exemplo de Extensibilidade

Para adicionar um novo repositório (ex: MongoDB):

1. Criar implementação em `infrastructure/database/repositories/mongodb/`
2. Implementar a interface de `domain/repositories/`
3. Injetar no `http.go`
4. **Nenhuma mudança** necessária nos use cases ou domínio

## Conclusão

A aplicação segue rigorosamente os princípios da Clean Architecture, garantindo:
- Separação clara de responsabilidades
- Independência entre camadas
- Facilidade de manutenção e testes
- Flexibilidade para evoluir e escalar

Todas as dependências apontam para o centro (domínio), mantendo o código de negócio puro e independente de detalhes de implementação.

