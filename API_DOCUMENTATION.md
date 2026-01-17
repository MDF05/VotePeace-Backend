# API Documentation

Base URL: `http://localhost:3000`

## 🔐 Authentication

### Register
Create a new user account.
- **URL**: `/register`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "nik": "1234567890123456",
    "name": "John Doe",
    "password": "securepassword",
    "role": "USER" // Optional, defaults to USER
  }
  ```
- **Response**: `201 Created`

### Login
Authenticate user and receive access credentials.
- **URL**: `/login`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "nik": "1234567890123456",
    "password": "securepassword"
  }
  ```
- **Response**: `200 OK`
  ```json
  {
    "message": "Login success",
    "user": { ... }
  }
  ```

### Check Session
Verify if the current user is authenticated.
- **URL**: `/check`
- **Method**: `GET`
- **Response**: `200 OK` (User details) or `401 Unauthorized`

---

## 🗳️ Campaigns

### Get All Campaigns
Retrieve a list of all election campaigns.
- **URL**: `/campaigns`
- **Method**: `GET`
- **Response**: `200 OK` `[CampaignObject]`

### Get Campaign Detail
Retrieve details for a specific campaign.
- **URL**: `/campaigns/:id`
- **Method**: `GET`
- **Response**: `200 OK` `CampaignObject`

### Create Campaign
Create a new election campaign (Admin only).
- **URL**: `/campaigns`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "title": "Election 2024",
    "description": "General Election",
    "startDate": "2024-01-01T00:00:00Z",
    "endDate": "2024-01-02T00:00:00Z"
  }
  ```
- **Response**: `201 Created`

### Delete Campaign
Remove a campaign.
- **URL**: `/campaigns/:id`
- **Method**: `DELETE`
- **Response**: `200 OK`

### Get Campaign Summary
Get high-level statistics for a campaign.
- **URL**: `/campaigns/:id/summary`
- **Method**: `GET`
- **Response**: `200 OK`

### Get Campaign Votes
Get detailed vote logs/stats for a campaign.
- **URL**: `/campaigns/:id/votes`
- **Method**: `GET`
- **Response**: `200 OK`

---

## 👤 Candidates

### Get Candidates
List all candidates (optionally filtered by campaign).
- **URL**: `/candidates`
- **Method**: `GET`
- **Response**: `200 OK`

### Create Candidate
Add a candidate to a campaign.
- **URL**: `/candidates`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "campaignId": 1,
    "number": 1,
    "name": "Jane Doe",
    "vision": "A better future",
    "mission": "To serve the people",
    "photo": "http://image.url"
  }
  ```
- **Response**: `201 Created`

---

## 📊 General Stats

### Get System Stats
Retrieve global dashboard statistics.
- **URL**: `/stats`
- **Method**: `GET`
- **Response**: `200 OK`
