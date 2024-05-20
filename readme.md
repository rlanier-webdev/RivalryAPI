# API Documentation

## Introduction
This API provides endpoints to manage and retrieve information about sports games.

## Endpoints

### 1. Get All Games
- **URL**: `/api/games/all`
- **Method**: `GET`
- **Description**: Returns all games in JSON format.
- **Response**: 
  - Status Code: `200 OK`
  - Body: JSON array of games.

### 2. Get Game by ID
- **URL**: `/api/games/:id`
- **Method**: `GET`
- **Description**: Returns the game with the specified ID.
- **Parameters**:
  - `id` (int): The ID of the game to retrieve.
- **Response**:
  - Status Code: 
    - `200 OK` if the game is found.
    - `404 Not Found` if the game with the specified ID is not found.
  - Body: JSON object representing the game.

### 3. Get Games by Team
- **URL**: `/api/games/team/:name`
- **Method**: `GET`
- **Description**: Returns all games involving the specified team.
- **Parameters**:
  - `name` (string): The name of the team to retrieve games for.
- **Response**:
  - Status Code:
    - `200 OK` if games involving the team are found.
    - `404 Not Found` if no games involving the team are found.
  - Body: JSON array of games involving the specified team.

### 4. Get Games by Year
- **URL**: `/api/games/year/:year`
- **Method**: `GET`
- **Description**: Returns all games played in the specified year.
- **Parameters**:
  - `year` (int): The year to retrieve games for.
- **Response**:
  - Status Code:
    - `200 OK` if games for the specified year are found.
    - `404 Not Found` if no games are found for the specified year.
  - Body: JSON array of games played in the specified year.

## Data Structures

### Team
Represents a sports team.
- `Name` (string): The name of the team.

### Game
Represents a sports game.
- `ID` (int): The unique identifier of the game.
- `HomeTeam` (Team): The home team participating in the game.
- `AwayTeam` (Team): The away team participating in the game.
- `Date` (string): The date of the game in the format "YYYY-MM-DD".
- `Score` (Score): The score of the game.
- `Notes` (string): Additional notes about the game.

### Score
Represents the score of a sports game.
- `HomeTeamScore` (int): The score of the home team.
- `AwayTeamScore` (int): The score of the away team.
