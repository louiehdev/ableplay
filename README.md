# Ableplay

Ableplay is a curated video game database that focuses on accessibility features. It allows users to browse games, view details, and explore which accessibility options each game supports. The project includes a Go backend API and a Go-powered frontend that renders UI using HTML templates and HTMX.

## Features

- **Game catalog** with metadata such as developer, release year, and platforms
- **Accessibility feature tagging** for each game
- **Server-rendered UI** using the Go `html/template` package
- **HTMX** for dynamic updates without a full SPA
- **PostgreSQL backend** for relational data storage
- Clear separation of concerns between **API service** and **frontend UI**

---

## Public Demo

A live demo of Ableplay is available here:

**[https://ableplay-frontend-latest.onrender.com](https://ableplay-frontend-latest.onrender.com)**

This demo showcases the core browsing experience and accessibility feature listings.  
**Note:** Since this is a free-tier deployment, cold starts or slight loading delays may occur. If no data populates for the frontend, use the link below to wake up the backend as it may have spun down due to inactivity.

The demo uses the API base of:

**[https://ableplay-api-latest.onrender.com](https://ableplay-api-latest.onrender.com)**

---

## Quick Start

### Requirements

- Go 1.22+
- Docker + Docker Compose

### 1. Start the Database and Servers

```
docker compose up --build
```

This runs:

    PostgreSQL database

    Backend API server

    Frontend Web server


### 2. Visit the App

Open your browser to:

http://localhost:9634

**Note:** No games or features are loaded in by default

### Environment Variables

Create a .env file in the project root:

    DB_PASSWORD=mysecretpassword
    DB_USER=ableplay
    DB_NAME=ableplay
    DB_URL=postgres://${DB_USER}:${DB_PASSWORD}@db:5432/${DB_NAME}?sslmode=disable
    API_BASE=http://api:9633

Docker Compose will load this automatically.

### How the Services Work

| Service   | Description | Default Port |
|----------|-------------|--------------|
| Frontend | Renders HTML pages and calls the API server-side | `9634` |
| API      | Exposes JSON endpoints for games and features    | `9633` |
| Database | PostgreSQL instance                              | `5432` |

The frontend server communicates with the API server internally, so calls do **not** require browser-side CORS handling.

---

## Usage

### API Endpoints
|Method|Endpoint|Description|
|------|--------|-----------|
|GET|/api/games/{gameID}|Returns details for a single game|
|GET|/api/games|Returns all games|
|GET|/api/games/search|Returns all games that wildcard match the provided query|
|POST|/api/games|Creates a new game|
|PUT|/api/games/{gameID}|Updates an existing game|
|DELETE|/api/games/{gameID}|Removes a game|
|GET|/api/features/{featureID}|Returns details for one feature|
|GET|/api/features|Returns all features|
|GET|/api/features/search|Returns all features that wildcard match the provided query|
|POST|/api/features|Creates a new feature|
|PUT|/api/features/{featureID}|Updates a feature|
|DELETE|/api/features/{featureID}|Removes a feature|
|GET|/api/games/{gameID}/features/{featureID}|Returns details for one game-feature relationship|
|GET|/api/games/features|Returns all games with their features|
|GET|/api/games/{gameID}/features|Returns features for one game|
|GET|/api/features/{featureID}/games|Returns games with a specific feature|
|POST|/api/games/{gameID}/features|Associates a feature to a game|
|PUT|/api/games/{gameID}/features/{featureID}|Updates relationship notes / verification|
|DELETE|/api/games/{gameID}/features/{featureID}|Removes a feature from a game|

### API Example Request

```
GET http://localhost:9633/api/games/features
```
Example response:
```
{
   [
      {
         "id":"1234",
         "title":"Super Mario Sunshine",
         "developer":"Nintendo EAD",
         "publisher":"Nintendo",
         "release_year":2002,
         "platforms":[
            "Nintendo GameCube",
            "Nintendo Switch"
         ],
         "description":"A 2002 platforming game where Mario is framed for a crime on the tropical vacation island of 
          Isle Delfino. He must use his new companion, FLUDD (a water-squirting backpack), to clean up a mysterious 
          goop that is covering the island while also pursuing the real culprit.",
         "game_features":[
            {
                "id": "4321",
                "name": "Colorblind Modes",
                "notes": "Game provides multiple colorblind modes for visual assistance",
                "verified": true
            },
            {
                "..."
            }
         ]
      },
      {
         "..."
      }
   ]
}
```

---

## Roadmap

The following features and improvements are planned for future development:

### 1. User Authentication & Roles
Currently, all API write operations are open. The goal is to introduce user accounts and role-based permissions:
- User registration and login (session or token-based auth)
- Roles such as **User**, **Moderator**, and **Admin**
- Users can submit new game entries or accessibility features
- Moderator/Admin approval workflow before new data becomes publicly visible
- Audit logs to track submissions and approvals

### 2. Community Mod Support
A database table for community-created game mods already exists, but no API or UI integration is in place yet. Planned work includes:
- Create API endpoints for listing and submitting mods
- Display mods on individual game detail pages
- Allow filtering/sorting mods by popularity, date, or game compatibility

### 3. Frontend Pagination
As the game catalog grows, the frontend will need incremental loading rather than rendering the full list. Planned improvements:
- Pagination controls on game listings
- API support for `limit` and `offset` query parameters
- HTMX-based partial page updates for a smooth browsing experience

### 4. Improved Query Parameter Handling
Search and filtering will benefit from structured and validated query parameters. Planned additions:
- Strongly-typed parameter parsing in handlers
- URL-based filters (e.g. `?platform=PC&accessibility=text-to-speech`)
- Search results that are shareable/bookmarkable

### 5. UI / UX Improvements
- Add game cover art or accessibility icons for quicker visual scanning
- Refine template structure to reduce duplication
- Explore lightweight CSS frameworks or tailwind-utility selective usage

### 6. (Optional) Consolidation of Servers
Currently the project runs separate frontend and backend servers. A future refactor may unify them into a single Go service for simpler deployment, configuration, and routing.

---



## Technologies Used

    Go (net/http, html/template)

    HTMX for progressive enhancement in the frontend

    PostgreSQL as the primary datastore

    sqlc for type-safe query generation

    Docker Compose for local orchestration

## Contributing

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.
