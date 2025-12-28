# Documentação do workspace.json

Guia completo sobre todas as possibilidades de configuração do arquivo `workspace.json` no DCM (Docker Compose Manager).

## 📋 Índice

- [Estrutura Básica](#estrutura-básica)
- [Propriedades Principais](#propriedades-principais)
  - [version](#version)
  - [projects](#projects)
  - [groups](#groups)
- [Exemplos Práticos](#exemplos-práticos)
- [Casos de Uso Avançados](#casos-de-uso-avançados)
- [Validação](#validação)
- [Melhores Práticas](#melhores-práticas)

---

## Estrutura Básica

O `workspace.json` é um arquivo JSON que define a configuração de todos os seus projetos Docker Compose e como eles podem ser agrupados e gerenciados.

```json
{
  "version": "1.0",
  "projects": { ... },
  "groups": { ... }
}
```

---

## Propriedades Principais

### version

**Tipo:** `string`  
**Obrigatório:** Sim  
**Valores aceitos:** `"1.0"`

Define a versão do schema do workspace.

```json
{
  "version": "1.0"
}
```

---

### projects

**Tipo:** `object`  
**Obrigatório:** Sim  
**Descrição:** Mapa de projetos Docker Compose disponíveis.

Cada projeto é identificado por uma **chave única** (nome do projeto) e contém as seguintes propriedades:

#### Propriedades de um Project

| Propriedade | Tipo | Obrigatório | Descrição |
|-------------|------|-------------|-----------|
| `path` | `string` | ✅ Sim | Caminho relativo ou absoluto para a pasta contendo o `docker-compose.yml` |
| `description` | `string` | ❌ Não | Descrição do projeto (exibida no comando `dcm list`) |

#### Exemplo de projects

```json
{
  "projects": {
    "database": {
      "path": "./infra/database",
      "description": "PostgreSQL database"
    },
    "api": {
      "path": "./services/api",
      "description": "REST API backend"
    },
    "frontend": {
      "path": "./services/frontend",
      "description": "React frontend application"
    }
  }
}
```

---

### groups

**Tipo:** `object`  
**Obrigatório:** Não  
**Descrição:** Mapa de grupos que combinam múltiplos projetos.

Cada grupo é identificado por uma **chave única** (nome do grupo) e contém as seguintes propriedades:

#### Propriedades de um Group

| Propriedade | Tipo | Obrigatório | Descrição |
|-------------|------|-------------|-----------|
| `services` | `array<string>` | ✅ Sim | Lista de nomes de projetos ou especificações de serviços |
| `extends` | `string` | ❌ Não | Nome de outro grupo para herdar serviços |
| `parallel` | `boolean` | ❌ Não | Se `true`, inicia serviços em paralelo. Se `false`, inicia sequencialmente. Padrão: `true` |

#### Especificação de Serviços

Os serviços podem ser especificados de duas formas:

1. **Nome do projeto completo:** `"database"` - Inicia todos os serviços do projeto
2. **Projeto:Serviço específico:** `"api:web"` - Inicia apenas o serviço `web` do projeto `api`

#### Exemplo de groups

```json
{
  "groups": {
    "backend": {
      "services": ["database", "api"],
      "parallel": false
    },
    "frontend": {
      "services": ["frontend"]
    },
    "full": {
      "extends": "backend",
      "services": ["frontend"],
      "parallel": true
    }
  }
}
```

---

## Exemplos Práticos

### 1. Configuração Simples

Ideal para projetos pequenos com poucos serviços.

```json
{
  "version": "1.0",
  "projects": {
    "app": {
      "path": "./app",
      "description": "Aplicação principal"
    },
    "db": {
      "path": "./database",
      "description": "Banco de dados"
    }
  },
  "groups": {
    "dev": {
      "services": ["db", "app"],
      "parallel": false
    }
  }
}
```

**Uso:**
```bash
dcm up dev  # Inicia db primeiro, depois app
```

---

### 2. Microserviços com Dependências

Quando você tem múltiplos serviços com ordem de inicialização importante.

```json
{
  "version": "1.0",
  "projects": {
    "postgres": {
      "path": "./infra/postgres",
      "description": "PostgreSQL database"
    },
    "redis": {
      "path": "./infra/redis",
      "description": "Redis cache"
    },
    "auth-service": {
      "path": "./services/auth",
      "description": "Authentication service"
    },
    "user-service": {
      "path": "./services/users",
      "description": "User management service"
    },
    "api-gateway": {
      "path": "./services/gateway",
      "description": "API Gateway"
    }
  },
  "groups": {
    "infra": {
      "services": ["postgres", "redis"],
      "parallel": true,
      "description": "Infraestrutura básica"
    },
    "backend": {
      "extends": "infra",
      "services": ["auth-service", "user-service"],
      "parallel": false
    },
    "full": {
      "extends": "backend",
      "services": ["api-gateway"]
    }
  }
}
```

**Uso:**
```bash
dcm up infra     # Apenas infraestrutura
dcm up backend   # Infra + serviços backend
dcm up full      # Tudo
```

---

### 3. Serviços Específicos

Controle fino sobre quais serviços de um projeto iniciar.

```json
{
  "version": "1.0",
  "projects": {
    "monitoring": {
      "path": "./monitoring",
      "description": "Prometheus, Grafana, AlertManager"
    },
    "app": {
      "path": "./app",
      "description": "Aplicação com múltiplos containers"
    }
  },
  "groups": {
    "metrics-only": {
      "services": ["monitoring:prometheus", "monitoring:grafana"],
      "parallel": true
    },
    "dev": {
      "services": ["app:web", "app:worker"],
      "parallel": false
    },
    "full-monitoring": {
      "services": ["monitoring"],
      "parallel": true
    }
  }
}
```

**Uso:**
```bash
dcm up metrics-only      # Apenas Prometheus e Grafana
dcm up dev               # Apenas web e worker do app
dcm up full-monitoring   # Todos os serviços de monitoring
```

---

### 4. Ambientes Diferentes

Configurações para desenvolvimento, teste e produção.

```json
{
  "version": "1.0",
  "projects": {
    "db-dev": {
      "path": "./infra/db-dev",
      "description": "Database para desenvolvimento"
    },
    "db-test": {
      "path": "./infra/db-test",
      "description": "Database para testes"
    },
    "api": {
      "path": "./services/api",
      "description": "API backend"
    },
    "frontend": {
      "path": "./services/frontend",
      "description": "Frontend"
    },
    "test-runner": {
      "path": "./tests",
      "description": "Container de testes E2E"
    }
  },
  "groups": {
    "dev": {
      "services": ["db-dev", "api", "frontend"],
      "parallel": false
    },
    "test": {
      "services": ["db-test", "api:test", "test-runner"],
      "parallel": false
    },
    "api-only": {
      "services": ["db-dev", "api"],
      "parallel": false
    }
  }
}
```

**Uso:**
```bash
dcm up dev       # Ambiente de desenvolvimento
dcm up test      # Ambiente de testes
dcm up api-only  # Apenas API para desenvolvimento frontend
```

---

### 5. Herança de Grupos Complexa

Grupos que estendem outros grupos para máxima reutilização.

```json
{
  "version": "1.0",
  "projects": {
    "postgres": { "path": "./db/postgres" },
    "redis": { "path": "./db/redis" },
    "rabbitmq": { "path": "./messaging/rabbitmq" },
    "auth": { "path": "./services/auth" },
    "users": { "path": "./services/users" },
    "orders": { "path": "./services/orders" },
    "notifications": { "path": "./services/notifications" },
    "web": { "path": "./web" }
  },
  "groups": {
    "databases": {
      "services": ["postgres", "redis"],
      "parallel": true
    },
    "messaging": {
      "services": ["rabbitmq"]
    },
    "core-services": {
      "extends": "databases",
      "services": ["auth", "users"],
      "parallel": false
    },
    "business-services": {
      "extends": "core-services",
      "services": ["orders", "notifications"],
      "parallel": false
    },
    "full-stack": {
      "extends": "business-services",
      "services": ["messaging", "web"]
    }
  }
}
```

**Ordem de inicialização do grupo `full-stack`:**
1. `postgres` e `redis` (paralelo)
2. `auth` (sequencial)
3. `users` (sequencial)
4. `orders` (sequencial)
5. `notifications` (sequencial)
6. `messaging` (sequencial)
7. `web` (sequencial)

---

### 6. Configuração Mínima

O menor workspace.json válido possível.

```json
{
  "version": "1.0",
  "projects": {
    "app": { "path": "./app" }
  }
}
```

**Uso:**
```bash
dcm up app  # Inicia o projeto diretamente
```

---

## Casos de Uso Avançados

### Inicialização Sequencial vs Paralela

#### Paralela (padrão)
Mais rápido, mas sem garantia de ordem.

```json
{
  "groups": {
    "fast": {
      "services": ["service1", "service2", "service3"],
      "parallel": true
    }
  }
}
```

#### Sequencial
Mais lento, mas garante ordem de inicialização.

```json
{
  "groups": {
    "ordered": {
      "services": ["database", "api", "frontend"],
      "parallel": false
    }
  }
}
```

> [!TIP]
> Use `parallel: false` quando:
> - Serviços têm dependências entre si
> - Banco de dados precisa estar pronto antes da API
> - Migrations precisam rodar antes da aplicação

---

### Herança em Cadeia

Grupos podem estender outros grupos, criando uma hierarquia.

```json
{
  "groups": {
    "base": {
      "services": ["db"]
    },
    "backend": {
      "extends": "base",
      "services": ["api"]
    },
    "full": {
      "extends": "backend",
      "services": ["web"]
    }
  }
}
```

**Resultado do grupo `full`:** `["db", "api", "web"]`

> [!WARNING]
> **Ciclos de herança são detectados e causam erro!**
> ```json
> {
>   "groups": {
>     "a": { "extends": "b", "services": [] },
>     "b": { "extends": "a", "services": [] }
>   }
> }
> ```
> ❌ Erro: "ciclo de herança detectado"

---

### Serviços Específicos em Grupos

Você pode especificar serviços individuais de um projeto multi-container.

**docker-compose.yml do projeto `app`:**
```yaml
services:
  web:
    image: nginx
  worker:
    image: myapp-worker
  scheduler:
    image: myapp-scheduler
```

**workspace.json:**
```json
{
  "projects": {
    "app": { "path": "./app" }
  },
  "groups": {
    "web-only": {
      "services": ["app:web"]
    },
    "background-only": {
      "services": ["app:worker", "app:scheduler"],
      "parallel": true
    }
  }
}
```

---

### Combinando Projetos Completos e Serviços Específicos

```json
{
  "projects": {
    "infra": { "path": "./infra" },
    "app": { "path": "./app" }
  },
  "groups": {
    "dev": {
      "services": [
        "infra",           // Todos os serviços de infra
        "app:web",         // Apenas o serviço web do app
        "app:worker"       // Apenas o serviço worker do app
      ]
    }
  }
}
```

---

## Validação

O DCM valida automaticamente o `workspace.json` ao carregar. Use o comando:

```bash
dcm validate
```

### Validações Realizadas

#### ✅ Projetos

- [ ] **Caminho existe:** Verifica se `path` aponta para um diretório válido
- [ ] **docker-compose.yml existe:** Verifica se há um arquivo docker-compose no caminho

**Exemplo de erro:**
```
❌ Projeto 'api': caminho não encontrado: ./services/api
```

#### ✅ Grupos

- [ ] **Projetos referenciados existem:** Todos os projetos em `services` devem estar definidos em `projects`
- [ ] **Grupo estendido existe:** Se usar `extends`, o grupo deve existir
- [ ] **Sem ciclos de herança:** Detecta referências circulares

**Exemplos de erros:**
```
❌ Grupo 'dev': projeto 'database' não definido
❌ Grupo 'full': estende grupo inexistente 'backend'
❌ Grupo 'a': ciclo de herança detectado
```

---

## Melhores Práticas

### 1. Organize por Camadas

```json
{
  "projects": {
    "postgres": { "path": "./infra/postgres" },
    "redis": { "path": "./infra/redis" },
    "api": { "path": "./services/api" },
    "web": { "path": "./services/web" }
  },
  "groups": {
    "infra": {
      "services": ["postgres", "redis"],
      "parallel": true
    },
    "app": {
      "extends": "infra",
      "services": ["api", "web"],
      "parallel": false
    }
  }
}
```

### 2. Use Descrições Claras

```json
{
  "projects": {
    "auth": {
      "path": "./services/auth",
      "description": "Serviço de autenticação JWT"
    }
  }
}
```

### 3. Crie Grupos para Diferentes Cenários

```json
{
  "groups": {
    "dev": { ... },           // Desenvolvimento local
    "test": { ... },          // Testes automatizados
    "debug": { ... },         // Debugging específico
    "minimal": { ... }        // Mínimo necessário
  }
}
```

### 4. Use `parallel: false` com Sabedoria

```json
{
  "groups": {
    "backend": {
      "services": ["db", "migrations", "api"],
      "parallel": false  // DB → Migrations → API
    }
  }
}
```

### 5. Evite Herança Muito Profunda

❌ **Evite:**
```json
{
  "groups": {
    "a": { "services": ["s1"] },
    "b": { "extends": "a", "services": ["s2"] },
    "c": { "extends": "b", "services": ["s3"] },
    "d": { "extends": "c", "services": ["s4"] }
  }
}
```

✅ **Prefira:**
```json
{
  "groups": {
    "base": { "services": ["s1", "s2"] },
    "extended": { "extends": "base", "services": ["s3", "s4"] }
  }
}
```

### 6. Nomeie Projetos de Forma Consistente

```json
{
  "projects": {
    "db-postgres": { ... },     // Prefixo por tipo
    "db-redis": { ... },
    "svc-auth": { ... },        // Prefixo por camada
    "svc-users": { ... }
  }
}
```

---

## Comandos Relacionados

### Inicialização
```bash
dcm init              # Cria workspace.json inicial
dcm validate          # Valida configuração
dcm list              # Lista projetos e grupos
dcm inspect <grupo>   # Inspeciona configuração de um grupo
```

### Gerenciamento
```bash
dcm up <grupo>        # Inicia grupo
dcm up <projeto>      # Inicia projeto individual
dcm down              # Para todos
dcm down <grupo>      # Para grupo específico
dcm down -v           # Para e remove volumes
dcm down <grupo> -v   # Para grupo e remove volumes
```

### Monitoramento
```bash
dcm status            # Status de todos os containers
dcm logs              # Logs de todos os serviços
dcm restart           # Reinicia todos os serviços
```

---

## Referência Rápida

### Estrutura Completa

```json
{
  "version": "1.0",
  "projects": {
    "<nome-do-projeto>": {
      "path": "<caminho-relativo-ou-absoluto>",
      "description": "<descrição-opcional>"
    }
  },
  "groups": {
    "<nome-do-grupo>": {
      "services": ["<projeto>", "<projeto:servico>"],
      "extends": "<nome-de-outro-grupo>",
      "parallel": true | false
    }
  }
}
```

### Tipos de Dados

```typescript
interface Workspace {
  version: string;
  projects: Record<string, Project>;
  groups?: Record<string, Group>;
}

interface Project {
  path: string;
  description?: string;
}

interface Group {
  services: string[];
  extends?: string;
  parallel?: boolean;
}
```

---

## Solução de Problemas

### Erro: "workspace.json não encontrado"
**Solução:** Execute `dcm init` na raiz do projeto.

### Erro: "projeto 'X' não encontrado"
**Solução:** Verifique se o projeto está definido em `projects`.

### Erro: "caminho não encontrado"
**Solução:** Verifique se o `path` está correto e o diretório existe.

### Erro: "ciclo de herança detectado"
**Solução:** Remova referências circulares em `extends`.

### Serviços não iniciam na ordem esperada
**Solução:** Use `"parallel": false` no grupo.

---

## Exemplos Completos

### E-commerce Completo

```json
{
  "version": "1.0",
  "projects": {
    "postgres": { "path": "./infra/postgres", "description": "PostgreSQL 15" },
    "redis": { "path": "./infra/redis", "description": "Redis cache" },
    "elasticsearch": { "path": "./infra/elasticsearch", "description": "Search engine" },
    "rabbitmq": { "path": "./infra/rabbitmq", "description": "Message broker" },
    "auth": { "path": "./services/auth", "description": "Authentication" },
    "users": { "path": "./services/users", "description": "User management" },
    "products": { "path": "./services/products", "description": "Product catalog" },
    "orders": { "path": "./services/orders", "description": "Order processing" },
    "payments": { "path": "./services/payments", "description": "Payment gateway" },
    "notifications": { "path": "./services/notifications", "description": "Email/SMS" },
    "admin": { "path": "./web/admin", "description": "Admin panel" },
    "storefront": { "path": "./web/storefront", "description": "Customer facing" }
  },
  "groups": {
    "infra": {
      "services": ["postgres", "redis", "elasticsearch", "rabbitmq"],
      "parallel": true
    },
    "core": {
      "extends": "infra",
      "services": ["auth", "users"],
      "parallel": false
    },
    "catalog": {
      "extends": "core",
      "services": ["products"]
    },
    "checkout": {
      "extends": "catalog",
      "services": ["orders", "payments", "notifications"],
      "parallel": false
    },
    "dev": {
      "extends": "checkout",
      "services": ["storefront"]
    },
    "full": {
      "extends": "checkout",
      "services": ["admin", "storefront"]
    }
  }
}
```

---

## Conclusão

O `workspace.json` oferece flexibilidade total para gerenciar projetos Docker Compose de qualquer tamanho e complexidade. Use as combinações apresentadas nesta documentação para criar configurações que atendam às necessidades específicas do seu projeto.

Para mais informações, consulte o [README.md](file:///c:/Users/diney/Projects/dcm-cli/README.md) do projeto.
