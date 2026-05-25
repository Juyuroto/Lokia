<div align="center">

<!-- Logo -->
<img src="./assets/logo.png" alt="Lokia Logo" width="140" />

> *Logo en cours d'arrivée — à remplacer par `./assets/logo.png`*

# Lokia

**Ton assistant IA local, illimité et privé.**

Lokia est une interface de chat IA construite sur [Ollama](https://ollama.com), avec historique de conversation, support d'images, analyse de code et intégration d'APIs externes. Tout tourne en local sur ta machine — zéro abonnement, zéro limite.

</div>

---

## Aperçu

> *Screenshots en cours d'arrivée — ils seront ajoutés ici prochainement.*

<!-- Une fois les screens dispo, remplace ce bloc par :
<img src="./assets/screen-chat.png" alt="Interface de chat" width="100%" />
<img src="./assets/screen-history.png" alt="Historique des conversations" width="100%" />
<img src="./assets/screen-code.png" alt="Aide au code" width="100%" />
-->

---

## Fonctionnalités

- **Chat avec historique** — toutes tes conversations sont sauvegardées en base PostgreSQL
- **Support d'images** — montre des screenshots, des erreurs, des maquettes à ton IA
- **Aide au code** — coloration syntaxique, debug, génération de code
- **APIs externes** — météo, news, recherche web, et bien d'autres
- **100% local** — aucune donnée ne quitte ta machine
- **Streaming** — les réponses s'affichent en temps réel
- **Interface moderne** — construite avec React + Vite + Tailwind

---

## Stack technique

| Couche | Technologie |
|--------|-------------|
| Backend | Go + Gin |
| IA | Ollama (`llama3.2-vision:11b`) |
| Base de données | PostgreSQL + GORM |
| Frontend | React + Vite + CSS pure |
| Streaming | WebSocket (gorilla/websocket) |
| Config | godotenv |

---

## Installation

### Prérequis

- [Go 1.22+](https://golang.org/dl/)
- [Node.js 20+](https://nodejs.org)
- [PostgreSQL 15+](https://www.postgresql.org/download/)
- [Ollama](https://ollama.com/download)
- [Docker](https://www.docker.com/)

### 1. Clone le projet

```bash
git clone https://github.com/Juyuroto/lokia.git
cd lokia
```

### 2. Configure les variables d'environnement

```bash
cp .env.example .env
```

Édite le fichier `.env` :

```env
# Base de données
DB_HOST=localhost
DB_PORT=5432
DB_USER=lokia
DB_PASSWORD=tonmotdepasse
DB_NAME=lokia

# Ollama
OLLAMA_HOST=http://localhost:11434
OLLAMA_MODEL=llama3.2-vision:11b

# APIs externes (optionnel)
BRAVE_SEARCH_API_KEY=ta_cle
OPENWEATHER_API_KEY=ta_cle
NEWS_API_KEY=ta_cle
```

### 3. Lance avec Docker

```bash
docker-compose up -d
```

### 3b. Ou lance manuellement (Pas recommandé)

**Ollama**
```bash
ollama pull llama3.2-vision:11b
```

**Backend :**
```bash
cd backend
go mod tidy
go run main.go
```

**Frontend :**
```bash
cd frontend
npm install
npm run dev
```

L'interface est accessible sur **http://localhost:3000**

---

## Structure du projet

```
lokia/
├── backend/
│   ├── controllers
│   ├── middlewares/
│   ├── models/
│   ├── routes/
│   ├── services/
│   ├── dockerfile
│   ├── go.mod/
│   └── main.go
├── frontend/
│   ├── public/
│   ├── src/
│   │   ├── assets/
│   │   ├── components/
│   │   ├── pages/
│   │   └── services/
│   ├── dockerfile
│   ├── eslint.config.js
│   ├── index.html
│   ├── package.json
│   ├── package-lock.json
│   ├── node_module
│   └── vite.config.js
├── docker-compose.yml
├── .env.example
└── README.md
```

---

## APIs intégrées

| API | Usage | Gratuit |
|-----|-------|---------|
| Brave Search | Recherche web en temps réel | Oui (2000 req/mois) |
| OpenWeatherMap | Météo | Oui |
| NewsAPI | Actualités | Oui |
| NASA API | Photo du jour, données spatiales | Oui |
| JokeAPI | Blagues | Oui |

---

## Contribution

Les contributions sont les bienvenues. N'hésite pas à ouvrir une issue ou une pull request.

---

<div align="center">
  Fait avec soin — propulsé par <a href="https://ollama.com">Ollama</a>
</div>
