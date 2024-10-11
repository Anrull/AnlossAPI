# Anloss API

Anloss API - it`s api written on golang. 

### Struct .env:

`CONFIG_PATH="./config/local.yaml"`

`ADMIN_PASSWORD="YourAdminPassword"`

`TOKEN_BOT="YourTokenBot"`

`TG_CHAT_ID="YourTgChatID"`

### Struct config:

`env: "local"` (local/prod/dev)

`records_path: "./storage/records.db"` path for records.db

`students_path: "./storage/student.db"` path for student.db

`http_server:`

`address: "localhost:8082"` address server

`timeout: 4s` maximum time doing

`idle_timeout: 5m` time saving session

## Methods API

### Get method `/addRecord`. Params:

`name` required

`class` required

`olimp` required

`sub` required

`teacher` required

`stage` required

`date` optional

Example request: `http://localhost:8080/addRecord?name=Alex&class=10A&olimp=...&sub=Phisic&teacher=...&stage=...`

Example good output (200): `{"message": "ok"}`

Possible statuses: 400, 500

### Get method `/getRecords`. Params:

`name` required

`sub` optional

`olimp` optional

`teacher` optional

`stage` optional

Example request: `http://localhost:8080/getRecords?name=Petr&sub=...`

Example good output (200): `[]`

Possible statuses: 400, 500

### Get method `/getRecordsCount`. Params:

`name` required

`sub` optional

`olimp` optional

`teacher` optional

`stage` optional

Example request: `http://localhost:8080/getRecordsCount?name=Petr&sub=...`

Example good output (200): `{"message": "ok", "count": 6}`

Possible statuses: 400, 500

### Get method `/getAllRecords`. Params:

Empty params

Example request: `http://localhost:8080/getAllRecords`

Example good output (200): `[]`

Possible status: 500

### Get method `/deleteAllRecords`. Params:

`password` required

Example request: `http://localhost:8080/deleteAllRecords?password=YourPassword`

Example good output (200): `{"message": "ok"}`

Possible statuses: 400, 401, 500

### Get method `/checkSnils`. Params:

`snils` required (hashed)

Example request: `http://localhost:8080/checkSnils?snils=HashedSnils`

Example good output (200): `{"name": "Alex", "stage": "10A"}`

Example bad output (204): `{"message": "no student with this () hashing snils"}`

Possible status: 400

### Get method `/addStudent`. Params:

`name` required

`class` required

`snils` required (hashed)

Example request: `http://localhost:8080/addStudent?name=Alex&class=10A&snils=HashedSnils`

Example good output (201): `{"message": "student added"}`

Possible statuses: 400, 500

### Get method `/getJson`. Params:

`path` required

Example request: `http://localhost:8080/getJson?path=storage/jsons/weeks.json`

Example good output: choice json data

Possible statuses: 400, 500