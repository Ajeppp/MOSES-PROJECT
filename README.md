# 📖 Moses System API Documentation (Official)

> Automated Church Scheduling System
> Admin-based availability management & auto scheduling engine

Base URL:
```
http://localhost:8080/api
```

---

# 🔐 System Concept
**No user login system**
- Semua data diinput oleh admin
- Player tidak input availability sendiri
- Admin input siapa yang **TIDAK tersedia (unavailable)**
- Sistem otomatis generate jadwal dari player yang tersedia

---

# 📌 1. Players Management

## ➕ Create Player
**POST** `/players`

Body (JSON):
```json
{
  "name": "John",
  "role": "SINGER"
}
```

Response:
```json
{
  "message": "player created successfully"
}
```

---

## 📥 Get Players
**GET** `/players`

Response:
```json
[
  {
    "id": 1,
    "name": "John",
    "role": "SINGER"
  }
]
```

---

# 📌 2. Unavailability (Admin Input)
> Data player yang **TIDAK bisa pelayanan di tanggal tertentu**

## ➕ Create Unavailability
**POST** `/unavailable`

Body (JSON):
```json
{
  "player_id": 1,
  "service_date": "2026-02-08",
  "month": 2,
  "year": 2026,
  "reason": "Out of town"
}
```

Response:
```json
{
  "message": "unavailability created successfully"
}
```

---

## 📥 Get Unavailability
**GET** `/unavailable`

Query Params:
```
month=2
year=2026
```

Response:
```json
[
  {
    "id": 10,
    "player_id": 1,
    "service_date": "2026-02-08",
    "month": 2,
    "year": 2026,
    "reason": "Out of town"
  }
]
```

---

# 📌 3. Scheduler Engine


## ⚙️ Generate Monthly Schedule
**POST** `/generate-schedule`

Body (JSON):
```json
{
  "month": 2,
  "year": 2026
}
```

Process:
- Ambil semua player
- Exclude player yang unavailable
- Generate team per Sabtu
- Auto assign berdasarkan role

Response:
```json
{
  "message": "schedule generated successfully"
}
```

---

## 📥 Get Raw Schedule Data
**GET** `/schedules`

Query Params:
```
month=2
year=2026
```

Response:
```json
[
  {
    "id": 1,
    "service_date": "2026-02-07",
    "month": 2,
    "year": 2026,
    "role": "SINGER",
    "player_id": 3
  }
]
```

---

# 📌 4. Schedule View (Formatted Output)

## 📊 Get Schedule View
**GET** `/schedule-view`

Query Params:
```
month=2
year=2026
```

Response:
```json
[
  {
    "service_date": "2026-02-07",
    "teams": {
      "WL": ["Andi", "Budi"],
      "SINGER": ["Maria", "Sinta"],
      "BASS": ["Riko"],
      "KEYS": ["Daniel"],
      "DRUM": ["Kevin"],
      "GUITAR": ["Rama"]
    }
  }
]
```

---

# 🧱 Role Constants (SYSTEM ENUM)
```txt
WL
SINGER
BASS
KEYS
DRUM
GUITAR
```

---

# 🧠 System Flow
1. Admin input players
2. Admin input unavailable players
3. Admin generate monthly schedule
4. System auto assign per Saturday
5. Admin view formatted schedule
6. Admin share manually to group

---

# 🧬 Architecture
```
Next.js (Admin Panel)
        ↓
Gin API (Backend)
        ↓
Service Layer (Business Logic)
        ↓
Repository Layer (DB Access)
        ↓
PostgreSQL
```

---

# 🚀 Roadmap

## Phase 1
- Schedule filter
- Export JSON
- Admin UI

## Phase 2
- PDF Export
- Excel Export
- WhatsApp Message Generator

## Phase 3
- Fairness Rotation Engine
- No-repeat Algorithm
- Priority Weighting
- Fatigue Management

## Phase 4 (SaaS Mode)
- Multi-church
- Multi-admin
- Auth system
- Tenant system
- Billing system
- Cloud deployment

---

# 🐐 Moses System
> Automated Church Scheduling Engine
> Admin Controlled • Fair Assignment • Scalable Architecture
> Built for real ministry operations

