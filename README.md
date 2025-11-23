# VICIdial Non-Agent API - Go Implementation

A comprehensive RESTful API implementation in Go for VICIdial's non-agent functions. This API provides programmatic access to all administrative, management, and reporting functions for VICIdial call center systems.

## Features

- **RESTful API Design**: Clean, modern REST API with JSON responses
- **Comprehensive Coverage**: All non-agent API functions from the PHP implementation
- **Database Connection Pooling**: Efficient MySQL connection management
- **Authentication Middleware**: Secure request authentication
- **CORS Support**: Cross-origin resource sharing enabled
- **Structured Logging**: Request/response logging for debugging
- **Error Handling**: Consistent error responses across all endpoints

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Running the API](#running-the-api)
- [API Endpoints](#api-endpoints)
- [Authentication](#authentication)
- [Examples](#examples)
- [Development](#development)

## Installation

### Prerequisites

- Go 1.21 or higher
- MySQL 5.7+ or MariaDB 10.3+
- VICIdial database schema installed

### Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd go-api
```

2. Install dependencies:
```bash
go mod download
```

3. Copy the environment template:
```bash
cp .env.example .env
```

4. Edit `.env` with your database credentials:
```bash
DB_HOST=localhost
DB_PORT=3306
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=asterisk
API_PORT=8080
TIMEZONE=America/New_York
```

## Configuration

All configuration is managed through environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | Database host | localhost |
| `DB_PORT` | Database port | 3306 |
| `DB_USER` | Database username | root |
| `DB_PASSWORD` | Database password | |
| `DB_NAME` | Database name | asterisk |
| `API_PORT` | API server port | 8080 |
| `API_KEY` | Shared API key required for all requests | _(none)_ |
| `TIMEZONE` | System timezone | America/New_York |
| `LOG_LEVEL` | Logging level | info |

Timezone values are automatically URL-encoded for the DSN; supply a valid IANA TZ name (e.g., `America/New_York`, `Europe/London`).

## Running the API

### Development Mode

```bash
go run main.go
```

### Production Build

```bash
# Build binary
go build -o non-agent-api

# Run binary
./non-agent-api
```

### Docker (Optional)

```bash
# Build Docker image
docker build -t vicidial-api .

# Run container
docker run -p 8080:8080 --env-file .env vicidial-api
```

## API Endpoints

### Authentication

All API requests require the shared API key set in your environment (`API_KEY`).

Send the key in either location:
- HTTP header `X-API-Key: <your-api-key>`
- Query/Form parameter `api_key=<your-api-key>`

### Base URL

```
http://localhost:8080/api/v1
```

### Health Check

```
GET /health
```

### Version Information

```
GET /api/v1/version
```

---

## API Categories

### 1. Lead Management

#### Add Lead
```http
POST /api/v1/leads
Content-Type: application/json

{
  "list_id": 101,
  "phone_number": "5551234567",
  "first_name": "John",
  "last_name": "Doe",
  "email": "john@example.com",
  "status": "NEW"
}
```

#### Update Lead
```http
PUT /api/v1/leads/{lead_id}
```

#### Batch Update Leads
```http
PUT /api/v1/leads/batch
{
  "lead_ids": [1, 2, 3],
  "status": "CALLBACK",
  "owner": "agent1"
}
```

#### Search Leads
```http
GET /api/v1/leads/search?phone_number=555&first_name=John
```

#### Get Lead Information
```http
GET /api/v1/leads/{lead_id}/info
```

#### Get Lead Field
```http
GET /api/v1/leads/{lead_id}/field-info?field=email
```

#### Search by Status
```http
GET /api/v1/leads/status-search?status=CALLBACK&list_id=101
```

#### Get Callback Info
```http
GET /api/v1/leads/{lead_id}/callback-info
```

#### Dearchive Lead
```http
POST /api/v1/leads/{lead_id}/dearchive
```

#### Check Phone Number
```http
GET /api/v1/phone/check?phone_number=5551234567
```

---

### 2. List Management

#### Add List
```http
POST /api/v1/lists
{
  "list_name": "New Campaign List",
  "campaign_id": "TESTCAMP",
  "active": "Y",
  "list_description": "Test campaign leads"
}
```

#### Update List
```http
PUT /api/v1/lists/{list_id}
```

#### Get List Info
```http
GET /api/v1/lists/{list_id}/info
```

#### Manage Custom Fields
```http
GET /api/v1/lists/{list_id}/custom-fields
POST /api/v1/lists/{list_id}/custom-fields
PUT /api/v1/lists/{list_id}/custom-fields
```

---

### 3. User/Agent Management

#### Add User
```http
POST /api/v1/users
{
  "user": "agent1",
  "pass": "password123",
  "full_name": "John Smith",
  "user_level": 1,
  "user_group": "AGENTS",
  "active": "Y"
}
```

#### Update User
```http
PUT /api/v1/users/{user_id}
```

#### Copy User
```http
POST /api/v1/users/{user_id}/copy
{
  "new_user": "agent2",
  "new_pass": "password123"
}
```

#### Get User Details
```http
GET /api/v1/users/{user_id}/details
```

#### Get Logged-In Agents
```http
GET /api/v1/users/logged-in
```

#### Get Agent Status
```http
GET /api/v1/agents/status?campaign_id=TESTCAMP
```

#### Get Agent Ingroups
```http
GET /api/v1/agents/{agent_id}/ingroup-info
```

#### Get Agent Campaigns
```http
GET /api/v1/agents/{agent_id}/campaigns
```

#### Update Remote Agent
```http
PUT /api/v1/remote-agents/{agent_id}
```

---

### 4. Campaign Management

#### Update Campaign
```http
PUT /api/v1/campaigns/{campaign_id}
{
  "campaign_name": "Updated Campaign",
  "active": "Y",
  "dial_method": "RATIO"
}
```

#### List Campaigns
```http
GET /api/v1/campaigns?active=Y
```

#### Get Hopper List
```http
GET /api/v1/campaigns/{campaign_id}/hopper
```

#### Bulk Insert to Hopper
```http
POST /api/v1/campaigns/{campaign_id}/hopper/bulk
{
  "lead_ids": [1, 2, 3, 4, 5],
  "priority": 50,
  "source": "API"
}
```

#### Get Campaigns with Lists
```http
GET /api/v1/campaigns/with-lists?active=Y&campaign_id=TESTCAMP
```

Returns campaigns with all associated lists and lead counts in JSON format.

**Response:**
```json
{
  "success": true,
  "message": "Campaigns with lists retrieved successfully",
  "data": {
    "count": 1,
    "campaigns": [
      {
        "campaign_id": "TESTCAMP",
        "campaign_name": "Test Campaign",
        "active": "Y",
        "dial_status": "AUTO",
        "dial_method": "RATIO",
        "auto_dial_level": "1.0",
        "lead_order": "DOWN",
        "local_call_time": "9am-5pm",
        "lists": [
          {
            "list_id": 101,
            "list_name": "January Leads",
            "active": "Y",
            "list_description": "Leads for January",
            "lead_count": 1250
          },
          {
            "list_id": 102,
            "list_name": "February Leads",
            "active": "Y",
            "list_description": "Leads for February",
            "lead_count": 980
          }
        ]
      }
    ]
  }
}
```

---

### 5. SIP/Carrier Logs

#### Get SIP/Carrier Log
```http
GET /api/v1/sip/carrier-log?start_date=2025-01-01&end_date=2025-01-31&limit=100
```

**Query Parameters:**
- `start_date` (optional): Filter by start date
- `end_date` (optional): Filter by end date
- `lead_id` (optional): Filter by lead ID
- `server_ip` (optional): Filter by server IP
- `dialstatus` (optional): Filter by dial status
- `sip_hangup_cause` (optional): Filter by SIP hangup cause
- `limit` (optional): Limit results (default: 100)

**Response:**
```json
{
  "success": true,
  "message": "SIP/Carrier logs retrieved successfully",
  "data": {
    "count": 50,
    "logs": [
      {
        "uniqueid": "1234567890.123",
        "call_date": "2025-01-08T10:30:00Z",
        "server_ip": "192.168.1.10",
        "lead_id": 12345,
        "hangup_cause": 16,
        "dialstatus": "ANSWER",
        "channel": "SIP/carrier-00000001",
        "dial_time": 5,
        "answered_time": 120,
        "sip_hangup_cause": 200,
        "sip_hangup_reason": "Normal Clearing",
        "caller_code": "1"
      }
    ]
  }
}
```

#### Get SIP Event Log
```http
GET /api/v1/sip/event-log?start_date=2025-01-01&sip_event=INVITE
```

**Query Parameters:**
- `start_date` (optional): Filter by start date
- `end_date` (optional): Filter by end date
- `sip_call_id` (optional): Filter by SIP call ID
- `sip_event` (optional): Filter by SIP event type
- `limit` (optional): Limit results (default: 100)

#### Get Live SIP Channels
```http
GET /api/v1/sip/live-channels?server_ip=192.168.1.10
```

**Query Parameters:**
- `server_ip` (optional): Filter by server IP
- `channel_group` (optional): Filter by channel group

**Response:**
```json
{
  "success": true,
  "message": "Live SIP channels retrieved successfully",
  "data": {
    "count": 5,
    "channels": [
      {
        "channel": "SIP/8001-00000001",
        "server_ip": "192.168.1.10",
        "channel_group": "AGENTS",
        "extension": "8001",
        "context": "default",
        "caller_id_number": "5551234567",
        "caller_id_name": "John Doe",
        "application": "Dial",
        "app_data": "SIP/carrier/18005551234"
      }
    ]
  }
}
```

---

### 6. KPI & Analytics

#### Get Disposition KPIs
```http
GET /api/v1/kpi/dispositions?group_by=campaign&campaign_id=TESTCAMP&start_date=2025-01-01
```

**Query Parameters:**
- `group_by` (required): Grouping method - "list", "campaign", or "both"
- `list_id` (optional): Filter by list ID
- `campaign_id` (optional): Filter by campaign ID
- `start_date` (optional): Filter by start date
- `end_date` (optional): Filter by end date

**Example 1: Group by Campaign**
```http
GET /api/v1/kpi/dispositions?group_by=campaign&campaign_id=TESTCAMP
```

**Response:**
```json
{
  "success": true,
  "message": "KPI dispositions by campaign retrieved successfully",
  "data": {
    "count": 1,
    "campaigns": [
      {
        "campaign_id": "TESTCAMP",
        "campaign_name": "Test Campaign",
        "total_calls": 1000,
        "dispositions": [
          {
            "status": "SALE",
            "count": 150,
            "percentage": 15.0,
            "avg_talk_time": 245.5,
            "total_talk_time": 36825
          },
          {
            "status": "NI",
            "count": 350,
            "percentage": 35.0,
            "avg_talk_time": 45.2,
            "total_talk_time": 15820
          },
          {
            "status": "CB",
            "count": 200,
            "percentage": 20.0,
            "avg_talk_time": 120.8,
            "total_talk_time": 24160
          }
        ]
      }
    ]
  }
}
```

**Example 2: Group by List**
```http
GET /api/v1/kpi/dispositions?group_by=list&list_id=101
```

**Response:**
```json
{
  "success": true,
  "message": "KPI dispositions by list retrieved successfully",
  "data": {
    "count": 1,
    "lists": [
      {
        "list_id": 101,
        "list_name": "January Leads",
        "total_calls": 500,
        "dispositions": [
          {
            "status": "SALE",
            "count": 75,
            "percentage": 15.0,
            "avg_talk_time": 250.0,
            "total_talk_time": 18750
          }
        ]
      }
    ]
  }
}
```

**Example 3: Group by Both Campaign and List**
```http
GET /api/v1/kpi/dispositions?group_by=both&campaign_id=TESTCAMP
```

**Response:**
```json
{
  "success": true,
  "message": "KPI dispositions by campaign and list retrieved successfully",
  "data": {
    "count": 2,
    "results": [
      {
        "campaign_id": "TESTCAMP",
        "campaign_name": "Test Campaign",
        "list_id": 101,
        "list_name": "January Leads",
        "total_calls": 500,
        "dispositions": [
          {
            "status": "SALE",
            "count": 75,
            "percentage": 15.0,
            "avg_talk_time": 250.0,
            "total_talk_time": 18750
          }
        ]
      },
      {
        "campaign_id": "TESTCAMP",
        "campaign_name": "Test Campaign",
        "list_id": 102,
        "list_name": "February Leads",
        "total_calls": 500,
        "dispositions": [
          {
            "status": "SALE",
            "count": 75,
            "percentage": 15.0,
            "avg_talk_time": 241.0,
            "total_talk_time": 18075
          }
        ]
      }
    ]
  }
}
```

---

### 7. Test Calls

#### Send Test Call
Places a campaign-based test call using campaign settings. This mimics the VICIdial admin.php test call functionality, using campaign dial rules, CID overrides, and dial prefixes.

```http
POST /api/v1/test-call/send
Content-Type: application/json

{
  "campaign_id": "TESTCAMP",
  "phone_number": "5551234567",
  "phone_code": "1",
  "user": "API",
  "vdad_exten": "8368",
  "server_ip": "10.0.0.5"
}
```

**Request Body:**
- `campaign_id` (required): Campaign ID to use for dial settings
- `phone_number` (required): Phone number to dial (minimum 6 digits)
- `phone_code` (optional): Phone/country code (default: "1")
- `user` (optional): User placing the call (default: "API")
- `vdad_exten` (optional): Override the VDAD/routing extension; defaults to campaign setting, then server answer_transfer_agent, then 8368
- `server_ip` (optional): Target a specific active dialer; defaults to the first active server

**Example 1: Basic Test Call**
```bash
curl -X POST http://localhost:8080/api/v1/test-call/send \
  -H "Content-Type: application/json" \
  -H "X-API-Key: $API_KEY" \
  -d '{
    "campaign_id": "TESTCAMP",
    "phone_number": "5551234567",
    "server_ip": "10.0.0.5"
  }'
```

**Example 2: International Test Call**
```bash
curl -X POST http://localhost:8080/api/v1/test-call/send \
  -H "Content-Type: application/json" \
  -H "X-API-Key: $API_KEY" \
  -d '{
    "campaign_id": "TESTCAMP",
    "phone_number": "2079460123",
    "phone_code": "44",
    "user": "admin",
    "vdad_exten": "8366"
  }'
```

**Response:**
```json
{
  "success": true,
  "message": "Test call placed successfully",
  "data": {
    "caller_code": "V01081530450000012345",
    "manager_id": 98765,
    "lead_id": 12345,
    "campaign_id": "TESTCAMP",
    "campaign_name": "Test Campaign",
    "phone_number": "5551234567",
    "phone_code": "1",
    "server_ip": "192.168.1.10",
    "server_id": "server1",
    "channel": "Local/95551234567@default",
    "extension": "8368",
    "dial_string": "95551234567",
    "caller_id": "\"V01081530450000012345\" <5559876543>",
    "call_date": "2025-01-08 15:30:45"
  }
}
```

#### Get Test Call Status
```http
GET /api/v1/test-call/status?caller_code=V01081530450000012345
```

**Query Parameters:**
- `caller_code` (required): Caller code from test call response

**Response:**
```json
{
  "success": true,
  "message": "Test call status retrieved",
  "data": {
    "caller_code": "V01081530450000012345",
    "manager_id": 98765,
    "entry_date": "2025-01-08T15:30:45Z",
    "status": "SENT",
    "response": "Y",
    "action": "Originate",
    "channel": "Local/95551234567@default",
    "context": "default",
    "extension": "8368",
    "call_date": "2025-01-08T15:30:45Z"
  }
}
```

#### List Test Calls
```http
GET /api/v1/test-call/list?limit=50
```

**Query Parameters:**
- `limit` (optional): Number of results (default: 50)
- `phone_login` (optional): Filter by phone login suffix

**Response:**
```json
{
  "success": true,
  "message": "Test calls retrieved",
  "data": {
    "count": 10,
    "calls": [
      {
        "caller_code": "V01081530450000012345",
        "call_date": "2025-01-08T15:30:45Z",
        "extension": "8368",
        "channel": "Local/95551234567@default",
        "server_ip": "192.168.1.10",
        "status": "SENT",
        "response": "Y"
      }
    ]
  }
}
```

**How It Works:**

The test call functionality uses campaign settings and the Asterisk Manager Interface (AMI):

1. **Validates campaign** - Checks campaign_id exists and gets campaign settings
2. **Creates test lead** - Inserts into vicidial_list with status 'CTCALL'
3. **Gets campaign settings** - Retrieves dial_prefix, campaign_cid, dial_timeout, omit_phone_code, campaign_vdad_exten
4. **Gets list CID overrides** - Checks for campaign_cid_override from manual_dial_list_id
5. **Gets server info** - Retrieves asterisk_version, routing_prefix, server_id
6. **Builds dial string** - Uses campaign dial rules and omit_phone_code setting
7. **Generates caller ID** - Format: VmddhhmmssLLLLLLLLLLL (V + date + lead_id)
8. **Inserts AMI Originate** - Queued in vicidial_manager table with proper campaign settings
9. **Logs to multiple tables** - vicidial_auto_calls, vicidial_dial_log, vicidial_dial_cid_log, vicidial_user_dial_log, user_call_log

**Campaign Settings Used:**
- `dial_prefix` - Prefix added to phone number (e.g., "9" for outside line)
- `campaign_cid` - Outbound caller ID for the campaign
- `dial_timeout` - Call timeout in seconds (converted to milliseconds)
- `omit_phone_code` - Whether to omit country code from dial string
- `campaign_vdad_exten` - VDAD extension for answering calls
- `manual_dial_list_id` - List ID for creating test leads
- `ext_context` - Dialplan context (default: "default")

**Asterisk 12+ Support:**
- Automatically adds `routing_prefix` to VDAD extension for Asterisk 12+

**Common Use Cases:**
- Test campaign dial settings before going live
- Verify outbound caller ID configuration
- Test carrier connections with campaign dial rules
- Validate dial prefix and phone code settings
- Troubleshoot campaign-specific dialing issues

---

### 8. Phone/DID Management

#### Add Phone
```http
POST /api/v1/phones
{
  "extension": "8001",
  "dialplan": "8001",
  "server_ip": "192.168.1.10",
  "active": "Y"
}
```

#### Update Phone
```http
PUT /api/v1/phones/{phone_id}
```

#### Add Phone Alias
```http
POST /api/v1/phone-aliases
```

#### Update Phone Alias
```http
PUT /api/v1/phone-aliases/{alias_id}
```

#### Add DID
```http
POST /api/v1/dids
{
  "did_pattern": "18005551234",
  "did_description": "Main Customer Line",
  "did_route": "INGROUP",
  "group_id": "SUPPORT"
}
```

#### Update DID
```http
PUT /api/v1/dids/{did_id}
```

#### Copy DID
```http
POST /api/v1/dids/{did_id}/copy
{
  "new_did_pattern": "18005555678"
}
```

---

### 9. DNC Management

#### Add to DNC
```http
POST /api/v1/dnc
{
  "phone_number": "5551234567",
  "campaign_id": "TESTCAMP"
}
```

#### Remove from DNC
```http
DELETE /api/v1/dnc/{phone_number}?campaign_id=TESTCAMP
```

#### Add to Filter Phone Group
```http
POST /api/v1/fpg
{
  "phone_number": "5551234567",
  "filter_phone_group_id": "BADNUMBERS"
}
```

#### Remove from Filter Phone Group
```http
DELETE /api/v1/fpg/{phone_number}?filter_phone_group_id=BADNUMBERS
```

---

### 10. Reporting & Monitoring

#### Recording Lookup
```http
GET /api/v1/recordings/lookup?lead_id=12345&start_date=2025-01-01
```

#### DID Log Export
```http
GET /api/v1/did-logs/export?start_date=2025-01-01&end_date=2025-01-31
```

#### Phone Number History
```http
GET /api/v1/phone-logs/{phone_number}
```

#### Agent Stats Export
```http
GET /api/v1/agent-stats/export?start_date=2025-01-01&api_key=YOUR_API_KEY
```

#### Call Status Statistics
```http
GET /api/v1/call-stats/status?campaign_id=TESTCAMP&start_date=2025-01-01
```

#### Call Disposition Report
```http
GET /api/v1/call-stats/dispo?campaign_id=TESTCAMP
```

#### Blind Monitor
```http
POST /api/v1/monitor/blind
{
  "user": "agent1",
  "extension": "8002",
  "server_ip": "192.168.1.10"
}
```

---

### 11. System Management

#### List Sounds
```http
GET /api/v1/system/sounds
```

#### List Music on Hold
```http
GET /api/v1/system/moh
```

#### List Voicemail
```http
GET /api/v1/system/voicemail
```

#### List Inbound Groups
```http
GET /api/v1/ingroups?active=Y
```

#### Inbound Group Status
```http
GET /api/v1/ingroups/status
```

#### List Call Menus
```http
GET /api/v1/callmenus
```

#### List Containers
```http
GET /api/v1/containers
```

#### Server Refresh
```http
POST /api/v1/system/refresh
{
  "server_ip": "192.168.1.10"
}
```

#### User Group Status
```http
GET /api/v1/user-groups/status
```

---

### 12. Advanced Features

#### Add Group Alias
```http
POST /api/v1/group-aliases
```

#### Update Log Entry
```http
PUT /api/v1/log-entries/{entry_id}
```

#### Update CID Group Entry
```http
PUT /api/v1/cid-groups/{entry_id}
```

#### Update Alt URL
```http
PUT /api/v1/alt-urls/{url_id}
```

#### Update Presets
```http
PUT /api/v1/presets/{preset_id}
```

#### Get Call ID Info
```http
GET /api/v1/calls/{call_id}/info
```

#### Get CCC Lead Info
```http
GET /api/v1/ccc/lead-info/{lead_id}
```

---

## Authentication

All requests must include the shared API key defined in your environment as `API_KEY`.

### Header (recommended)
```bash
curl http://localhost:8080/api/v1/version \
  -H "X-API-Key: $API_KEY"
```

### Query/Form parameter (alternate)
```bash
curl "http://localhost:8080/api/v1/version?api_key=$API_KEY"
```

Optional: you may still supply `user` in headers/query to tag requests for logging only; it is not used for authentication.

## Response Format

All responses follow this JSON structure:

### Success Response
```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": {
    ...
  }
}
```

### Error Response
```json
{
  "success": false,
  "error": "Error message description"
}
```

## Examples

### Complete Lead Management Workflow

```bash
# 1. Add a new lead
curl -X POST http://localhost:8080/api/v1/leads \
  -H "Content-Type: application/json" \
  -H "X-API-Key: $API_KEY" \
  -d '{
    "list_id": 101,
    "phone_number": "5551234567",
    "first_name": "John",
    "last_name": "Doe",
    "email": "john@example.com"
  }'

# 2. Search for the lead
curl "http://localhost:8080/api/v1/leads/search?phone_number=5551234567&api_key=$API_KEY"

# 3. Update the lead
curl -X PUT http://localhost:8080/api/v1/leads/12345 \
  -H "Content-Type: application/json" \
  -H "X-API-Key: $API_KEY" \
  -d '{
    "status": "CALLBACK",
    "comments": "Customer requested callback at 3pm"
  }'

# 4. Get callback information
curl "http://localhost:8080/api/v1/leads/12345/callback-info?api_key=$API_KEY"
```

## Development

### Project Structure

```
go-api/
├── main.go                 # Application entry point
├── config/
│   └── config.go          # Configuration management
├── database/
│   └── database.go        # Database connection
├── handlers/
│   ├── handler.go         # Base handler
│   ├── leads.go          # Lead management endpoints
│   ├── lists.go          # List management endpoints
│   ├── users.go          # User/agent management endpoints
│   ├── campaigns.go      # Campaign management endpoints
│   ├── phones.go         # Phone/DID management endpoints
│   ├── dnc.go            # DNC management endpoints
│   ├── reporting.go      # Reporting endpoints
│   ├── system.go         # System management endpoints
│   ├── advanced.go       # Advanced features endpoints
│   └── version.go        # Version endpoint
├── middleware/
│   ├── auth.go           # Authentication middleware
│   └── logging.go        # Logging middleware
├── models/
│   └── models.go         # Data models
├── .env.example          # Environment template
├── .gitignore
├── go.mod                # Go module definition
└── README.md             # This file
```

### Adding New Endpoints

1. Define the model in `models/models.go`
2. Create handler function in appropriate handler file
3. Register route in `main.go`
4. Update this README with endpoint documentation

### Running Tests

```bash
go test ./...
```

### Building for Production

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o non-agent-api-linux

# Windows
GOOS=windows GOARCH=amd64 go build -o non-agent-api.exe

# macOS
GOOS=darwin GOARCH=amd64 go build -o non-agent-api-macos
```

## License


This project implements the VICIdial Non-Agent API which is:
Copyright (C) 2025 Matt Florell <vicidial.com>
Licensed under AGPLv2

## Support

For issues and questions:
- VICIdial Documentation: http://www.vicidial.org/docs/
- VICIdial Forum: http://www.vicidial.org/vicidial_forum/

## Credits

- Original PHP Non-Agent API by Matt Florell
- Go Implementation: 2025
- VICIdial Project: http://www.vicidial.org/
