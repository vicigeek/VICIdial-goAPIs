# VICIdial Non-Agent API - Complete API Reference

## Table of Contents

1. [Authentication](#authentication)
2. [Response Format](#response-format)
3. [Error Codes](#error-codes)
4. [Rate Limiting](#rate-limiting)
5. [API Endpoints](#api-endpoints)

---

## Authentication

All API endpoints (except `/health`) require the shared API key configured in your environment as `API_KEY`.

- **Header (recommended):**
  - `X-API-Key: YOUR_API_KEY`
- **Query/Form parameter (alternate):**
  - `api_key=YOUR_API_KEY`

Optional: you may include `user` for request tagging/logging; it is not used for authentication.

---

## Response Format

### Standard Success Response
```json
{
  "success": true,
  "message": "Description of the operation",
  "data": {
    // Response data here
  }
}
```

### Standard Error Response
```json
{
  "success": false,
  "error": "Error description"
}
```

---

## Error Codes

| HTTP Code | Meaning |
|-----------|---------|
| 200 | Success |
| 400 | Bad Request - Invalid parameters |
| 401 | Unauthorized - Authentication failed |
| 404 | Not Found - Resource doesn't exist |
| 405 | Method Not Allowed |
| 500 | Internal Server Error |

---

## Rate Limiting

Currently no rate limiting is implemented. It's recommended to implement rate limiting at the reverse proxy level (nginx, Apache, etc.) for production deployments.

---

## API Endpoints

### System

#### Get Version
Returns API version and system information.

**Endpoint:** `GET /api/v1/version`

**Response:**
```json
{
  "success": true,
  "message": "Version information retrieved",
  "data": {
    "version": "1.0.0",
    "build": "20250108",
    "timezone": "America/New_York",
    "date": "2025-01-08 10:30:45"
  }
}
```

#### Health Check
Check if API is running (no authentication required).

**Endpoint:** `GET /health`

**Response:**
```
OK
```

---

### Lead Management

#### Add Lead

**Endpoint:** `POST /api/v1/leads`

**Request Body:**
```json
{
  "list_id": 101,
  "phone_number": "5551234567",
  "first_name": "John",
  "last_name": "Doe",
  "middle_initial": "A",
  "address1": "123 Main St",
  "address2": "Apt 4B",
  "address3": "",
  "city": "New York",
  "state": "NY",
  "province": "",
  "postal_code": "10001",
  "country_code": "1",
  "gender": "M",
  "date_of_birth": "1980-01-15",
  "alt_phone": "5559876543",
  "email": "john@example.com",
  "security": "",
  "comments": "VIP customer",
  "status": "NEW",
  "rank": 0,
  "owner": "agent1"
}
```

**Required Fields:**
- `list_id` (integer)
- `phone_number` (string)

**Default Values:**
- `status`: "NEW"
- `country_code`: "1"

**Response:**
```json
{
  "success": true,
  "message": "Lead added successfully",
  "data": {
    "lead_id": 12345,
    "list_id": 101,
    "phone_number": "5551234567",
    ...
  }
}
```

#### Update Lead

**Endpoint:** `PUT /api/v1/leads/{lead_id}`

**Parameters:**
- `lead_id` (path parameter): Lead ID to update

**Request Body:**
```json
{
  "first_name": "John",
  "last_name": "Smith",
  "status": "CALLBACK",
  "comments": "Customer requested callback",
  "email": "john.smith@example.com"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Lead updated successfully",
  "data": {
    "lead_id": 12345
  }
}
```

#### Batch Update Leads

**Endpoint:** `PUT /api/v1/leads/batch`

**Request Body:**
```json
{
  "lead_ids": [1, 2, 3, 4, 5],
  "status": "CALLBACK",
  "owner": "agent1"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Leads updated successfully",
  "data": {
    "updated_count": 5
  }
}
```

#### Search Leads

**Endpoint:** `GET /api/v1/leads/search`

**Query Parameters:**
- `phone_number` (string, optional): Phone number to search
- `first_name` (string, optional): First name to search
- `last_name` (string, optional): Last name to search
- `email` (string, optional): Email to search
- `list_id` (integer, optional): List ID filter
- `status` (string, optional): Status filter

**Example:**
```
GET /api/v1/leads/search?phone_number=555&status=NEW&api_key=YOUR_API_KEY
```

**Response:**
```json
{
  "success": true,
  "message": "Leads retrieved successfully",
  "data": [
    {
      "lead_id": 1,
      "list_id": 101,
      "phone_number": "5551234567",
      "first_name": "John",
      "last_name": "Doe",
      "email": "john@example.com",
      "status": "NEW",
      "entry_date": "2025-01-08T10:30:00Z"
    },
    ...
  ]
}
```

**Note:** Results limited to 100 records

#### Get Lead Information

**Endpoint:** `GET /api/v1/leads/{lead_id}/info`

**Parameters:**
- `lead_id` (path parameter): Lead ID

**Response:**
```json
{
  "success": true,
  "message": "Lead information retrieved",
  "data": {
    "lead_id": 12345,
    "list_id": 101,
    "phone_number": "5551234567",
    "first_name": "John",
    "last_name": "Doe",
    "middle_initial": "A",
    "address1": "123 Main St",
    "city": "New York",
    "state": "NY",
    "postal_code": "10001",
    "country_code": "1",
    "gender": "M",
    "date_of_birth": "1980-01-15",
    "alt_phone": "5559876543",
    "email": "john@example.com",
    "comments": "VIP customer",
    "status": "NEW",
    "entry_date": "2025-01-08T10:30:00Z",
    "modify_date": "2025-01-08T10:30:00Z",
    "called_count": 0,
    "rank": 0,
    "owner": "agent1"
  }
}
```

#### Get Lead Field Information

**Endpoint:** `GET /api/v1/leads/{lead_id}/field-info`

**Parameters:**
- `lead_id` (path parameter): Lead ID
- `field` (query parameter): Field name to retrieve

**Valid Field Names:**
- phone_number
- first_name
- last_name
- email
- status
- comments
- address1
- city
- state

**Example:**
```
GET /api/v1/leads/12345/field-info?field=email&api_key=YOUR_API_KEY
```

**Response:**
```json
{
  "success": true,
  "message": "Field retrieved",
  "data": {
    "email": "john@example.com"
  }
}
```

#### Search Leads by Status

**Endpoint:** `GET /api/v1/leads/status-search`

**Query Parameters:**
- `status` (string, required): Status code to search
- `list_id` (integer, optional): List ID filter

**Example:**
```
GET /api/v1/leads/status-search?status=CALLBACK&list_id=101&api_key=YOUR_API_KEY
```

**Response:**
```json
{
  "success": true,
  "message": "Leads retrieved",
  "data": [
    {
      "lead_id": 1,
      "list_id": 101,
      "phone_number": "5551234567",
      "first_name": "John",
      "last_name": "Doe",
      "status": "CALLBACK"
    },
    ...
  ]
}
```

**Note:** Results limited to 100 records

#### Get Lead Callback Information

**Endpoint:** `GET /api/v1/leads/{lead_id}/callback-info`

**Parameters:**
- `lead_id` (path parameter): Lead ID

**Response:**
```json
{
  "success": true,
  "message": "Callbacks retrieved",
  "data": [
    {
      "callback_id": 1,
      "lead_id": 12345,
      "list_id": 101,
      "campaign_id": "TESTCAMP",
      "status": "ACTIVE",
      "entry_time": "2025-01-08T10:00:00Z",
      "callback_time": "2025-01-08T15:00:00Z",
      "user": "agent1",
      "recipient": "ANYONE",
      "comments": "Customer requested 3pm callback"
    },
    ...
  ]
}
```

**Note:** Returns last 10 callbacks

#### Dearchive Lead

**Endpoint:** `POST /api/v1/leads/{lead_id}/dearchive`

**Parameters:**
- `lead_id` (path parameter): Lead ID to restore from archive

**Response:**
```json
{
  "success": true,
  "message": "Lead restored successfully",
  "data": {
    "lead_id": 12345
  }
}
```

#### Check Phone Number

**Endpoint:** `GET /api/v1/phone/check`

**Query Parameters:**
- `phone_number` (string, required): Phone number to check

**Example:**
```
GET /api/v1/phone/check?phone_number=5551234567&api_key=YOUR_API_KEY
```

**Response:**
```json
{
  "success": true,
  "message": "Phone check complete",
  "data": {
    "phone_number": "5551234567",
    "exists": true,
    "count": 3,
    "checked_by": "6666"
  }
}
```

---

### List Management

#### Add List

**Endpoint:** `POST /api/v1/lists`

**Request Body:**
```json
{
  "list_name": "January 2025 Leads",
  "campaign_id": "TESTCAMP",
  "active": "Y",
  "list_description": "Leads for January campaign",
  "script": "GENERIC",
  "web_form": "http://example.com/form"
}
```

**Required Fields:**
- `list_name` (string)

**Default Values:**
- `active`: "Y"

**Response:**
```json
{
  "success": true,
  "message": "List created successfully",
  "data": {
    "list_id": 102,
    "list_name": "January 2025 Leads",
    "campaign_id": "TESTCAMP",
    "active": "Y"
  }
}
```

#### Update List

**Endpoint:** `PUT /api/v1/lists/{list_id}`

**Parameters:**
- `list_id` (path parameter): List ID to update

**Request Body:**
```json
{
  "list_name": "Updated List Name",
  "campaign_id": "NEWCAMP",
  "active": "N",
  "list_description": "Updated description"
}
```

**Response:**
```json
{
  "success": true,
  "message": "List updated successfully",
  "data": {
    "list_id": "102"
  }
}
```

#### Get List Information

**Endpoint:** `GET /api/v1/lists/{list_id}/info`

**Parameters:**
- `list_id` (path parameter): List ID

**Response:**
```json
{
  "success": true,
  "message": "List retrieved",
  "data": {
    "list": {
      "list_id": 102,
      "list_name": "January 2025 Leads",
      "campaign_id": "TESTCAMP",
      "active": "Y",
      "list_description": "Leads for January campaign",
      "script": "GENERIC",
      "web_form": "http://example.com/form"
    },
    "lead_count": 1250
  }
}
```

#### Get List Custom Fields

**Endpoint:** `GET /api/v1/lists/{list_id}/custom-fields`

**Parameters:**
- `list_id` (path parameter): List ID

**Response:**
```json
{
  "success": true,
  "message": "Custom fields retrieved",
  "data": [
    {
      "field_id": 1,
      "field_label": "Customer Type",
      "field_name": "customer_type",
      "field_type": "SELECT",
      "field_options": "Residential,Commercial,Government",
      "field_size": 20,
      "field_max": 50,
      "field_default": "Residential",
      "field_required": "N"
    },
    ...
  ]
}
```

#### Add List Custom Field

**Endpoint:** `POST /api/v1/lists/{list_id}/custom-fields`

**Parameters:**
- `list_id` (path parameter): List ID

**Request Body:**
```json
{
  "field_label": "Customer Type",
  "field_name": "customer_type",
  "field_type": "SELECT",
  "field_options": "Residential,Commercial,Government",
  "field_size": 20,
  "field_max": 50,
  "field_default": "Residential",
  "field_required": "N"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Field added successfully",
  "data": {
    "field_id": 5
  }
}
```

#### Update List Custom Field

**Endpoint:** `PUT /api/v1/lists/{list_id}/custom-fields`

**Parameters:**
- `list_id` (path parameter): List ID

**Request Body:**
```json
{
  "field_id": 5,
  "field_label": "Updated Label",
  "field_type": "TEXT",
  "field_size": 30,
  "field_required": "Y"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Field updated successfully",
  "data": null
}
```

---

### Test Calls

#### Send Test Call

Place a test call using campaign settings (mirrors admin.php test call).

**Endpoint:** `POST /api/v1/test-call/send`

**Request Body:**
```json
{
  "campaign_id": "TESTCAMP",
  "phone_number": "5551234567",
  "phone_code": "1",
  "user": "API",
  "vdad_exten": "8366",
  "server_ip": "10.0.0.5"
}
```

- `campaign_id` (required): Campaign to use for dial rules.
- `phone_number` (required): Destination number (min 6 digits).
- `phone_code` (optional): Country/area code, default `"1"`.
- `user` (optional): For logging notes, default `"API"`.
- `vdad_exten` (optional): Force VDAD/routing extension; otherwise uses campaign setting, then server answer_transfer_agent, then 8368.
- `server_ip` (optional): Target a specific active server; otherwise first active server is used.

**Response (success):**
```json
{
  "success": true,
  "message": "Test call placed successfully",
  "data": {
    "caller_code": "V01081530450000012345",
    "manager_id": 123456,
    "lead_id": 999,
    "campaign_id": "TESTCAMP",
    "campaign_name": "Test Campaign",
    "phone_number": "5551234567",
    "phone_code": "1",
    "server_ip": "10.0.0.5",
    "server_id": "server1",
    "channel": "Local/95551234567@default",
    "extension": "138366",
    "dial_string": "95551234567",
    "caller_id": "\"V01081530450000012345\" <15559876543>",
    "call_date": "2025-01-08 15:30:45"
  }
}
```

#### Get Test Call Status

**Endpoint:** `GET /api/v1/test-call/status?caller_code=V01081530450000012345`

Returns latest status from `vicidial_manager` and `vicidial_dial_log`.

#### List Test Calls

**Endpoint:** `GET /api/v1/test-call/list?limit=50`

Lists recent test calls (defaults to 50); supports optional `phone_login` suffix filter.

---

## Complete Endpoint List

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/api/v1/version` | Get version info |
| POST | `/api/v1/leads` | Add lead |
| PUT | `/api/v1/leads/{lead_id}` | Update lead |
| PUT | `/api/v1/leads/batch` | Batch update leads |
| GET | `/api/v1/leads/search` | Search leads |
| GET | `/api/v1/leads/{lead_id}/info` | Get lead info |
| GET | `/api/v1/leads/{lead_id}/field-info` | Get lead field |
| GET | `/api/v1/leads/status-search` | Search by status |
| GET | `/api/v1/leads/{lead_id}/callback-info` | Get callbacks |
| POST | `/api/v1/leads/{lead_id}/dearchive` | Dearchive lead |
| GET | `/api/v1/phone/check` | Check phone number |
| POST | `/api/v1/lists` | Add list |
| PUT | `/api/v1/lists/{list_id}` | Update list |
| GET | `/api/v1/lists/{list_id}/info` | Get list info |
| GET | `/api/v1/lists/{list_id}/custom-fields` | Get custom fields |
| POST | `/api/v1/lists/{list_id}/custom-fields` | Add custom field |
| PUT | `/api/v1/lists/{list_id}/custom-fields` | Update custom field |
| POST | `/api/v1/users` | Add user |
| PUT | `/api/v1/users/{user_id}` | Update user |
| POST | `/api/v1/users/{user_id}/copy` | Copy user |
| GET | `/api/v1/users/{user_id}/details` | Get user details |
| GET | `/api/v1/users/logged-in` | Get logged-in agents |
| GET | `/api/v1/agents/status` | Get agent status |
| GET | `/api/v1/agents/{agent_id}/ingroup-info` | Get agent ingroups |
| GET | `/api/v1/agents/{agent_id}/campaigns` | Get agent campaigns |
| PUT | `/api/v1/remote-agents/{agent_id}` | Update remote agent |
| PUT | `/api/v1/campaigns/{campaign_id}` | Update campaign |
| GET | `/api/v1/campaigns` | List campaigns |
| GET | `/api/v1/campaigns/{campaign_id}/hopper` | Get hopper |
| POST | `/api/v1/campaigns/{campaign_id}/hopper/bulk` | Bulk insert hopper |
| POST | `/api/v1/phones` | Add phone |
| PUT | `/api/v1/phones/{phone_id}` | Update phone |
| POST | `/api/v1/phone-aliases` | Add phone alias |
| PUT | `/api/v1/phone-aliases/{alias_id}` | Update phone alias |
| POST | `/api/v1/dids` | Add DID |
| PUT | `/api/v1/dids/{did_id}` | Update DID |
| POST | `/api/v1/dids/{did_id}/copy` | Copy DID |
| POST | `/api/v1/dnc` | Add to DNC |
| DELETE | `/api/v1/dnc/{phone}` | Remove from DNC |
| POST | `/api/v1/fpg` | Add to filter group |
| DELETE | `/api/v1/fpg/{phone}` | Remove from filter group |
| GET | `/api/v1/recordings/lookup` | Search recordings |
| GET | `/api/v1/did-logs/export` | Export DID logs |
| GET | `/api/v1/phone-logs/{phone}` | Phone number history |
| GET | `/api/v1/agent-stats/export` | Export agent stats |
| GET | `/api/v1/call-stats/status` | Call status stats |
| GET | `/api/v1/call-stats/dispo` | Call dispo report |
| POST | `/api/v1/monitor/blind` | Blind monitor |
| POST | `/api/v1/test-call/send` | Send campaign test call |
| GET | `/api/v1/test-call/status` | Get test call status |
| GET | `/api/v1/test-call/list` | List recent test calls |
| GET | `/api/v1/system/sounds` | List sounds |
| GET | `/api/v1/system/moh` | List MOH |
| GET | `/api/v1/system/voicemail` | List voicemail |
| GET | `/api/v1/ingroups` | List inbound groups |
| GET | `/api/v1/ingroups/status` | Inbound group status |
| GET | `/api/v1/callmenus` | List call menus |
| GET | `/api/v1/containers` | List containers |
| POST | `/api/v1/system/refresh` | Server refresh |
| GET | `/api/v1/user-groups/status` | User group status |
| POST | `/api/v1/group-aliases` | Add group alias |
| PUT | `/api/v1/log-entries/{entry_id}` | Update log entry |
| PUT | `/api/v1/cid-groups/{entry_id}` | Update CID group |
| PUT | `/api/v1/alt-urls/{url_id}` | Update alt URL |
| PUT | `/api/v1/presets/{preset_id}` | Update presets |
| GET | `/api/v1/calls/{call_id}/info` | Get call info |
| GET | `/api/v1/ccc/lead-info/{lead_id}` | Get CCC lead info |
