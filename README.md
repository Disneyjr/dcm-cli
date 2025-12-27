# DCM - Docker Compose Manager

Gerencie múltiplos serviços Docker com **um único comando**.

## O Problema

```bash
# ❌ Sem DCM: Múltiplos terminais
Terminal 1: cd services/service-a && docker-compose up
Terminal 2: cd services/service-b && docker-compose up
Terminal 3: cd services/service-c && docker-compose up
# ... e assim vai
```

## A Solução

```bash
# ✅ Com DCM: Um comando
dcm up dev
```

## Características

- 🚀 **Comando único** - Inicie todos os serviços de uma vez
- ⚡ **Concorrência** - Inicialização paralela de serviços (super rápido!)
- 📦 **Sem dependências** - Binário standalone
- 🎯 **Grupos** - Organize serviços em combinações
- 🖥️ **Cross-platform** - Linux, macOS, Windows

## Quick Start

### 1. Instalar

```bash
# Baixe o binário em: https://github.com/Disneyjr/dcm/releases

# Linux/macOS
chmod +x install
./install

# Windows
double-click install.exe
```

### 2. Configurar

Navegue até a pasta raiz do seu projeto e execute:

```bash
dcm init
```

Isso criará um `workspace.json`. Veja como é simples organizar:

```json
{
  "version": "1.0",
  "projects": {
    "db": { "path": "./infra/db" },
    "api": { "path": "./services/api" }
  },
  "groups": {
    "dev": { 
      "services": ["db", "api"],
      "parallel": false 
    },
    "full": {
      "extends": "dev",
      "services": ["web"]
    }
  }
}
```

> [!TIP]
> Use `"parallel": false` quando a ordem de inicialização importar (ex: subir o banco antes da API).

### 3. Usar

```bash
dcm init      # Cria configuração inicial
dcm list      # Ver projetos e grupos
dcm up dev    # Iniciar grupo completo
dcm down      # Parar tudo
```

## Exemplos Rápidos

**Iniciar um grupo:**
```bash
dcm up dev          # Todos os serviços do grupo 'dev'
dcm up dev --build  # Força o rebuild das imagens
```

**Gerenciar serviços:**
```bash
dcm logs            # Ver logs de tudo
dcm status          # Status dos containers
dcm restart         # Reiniciar tudo
```

## Contribuindo

1. Fork o repositório
2. Crie uma branch: `git checkout -b feature/minha-feature`
3. Faça seus commits: `git commit -m "feat: descrição"`
4. Push: `git push origin feature/minha-feature`
5. Abra um Pull Request

## Licença

MIT - Use livremente!

## Suporte

Dúvidas ou problemas? Abra uma [issue](https://github.com/Disneyjr/dcm/issues) 🚀