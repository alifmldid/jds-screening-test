# JDS-Screening-Test

Screening test Jabar Digital Service: Authentication App & Fetch App

## 1. Authentication App API Specification

**Base URL:**
http://localhost:8000/api

**METHOD:**
|Endpoint |Fungsi |
|---------|---------------|
|register |Registrasi user|
|login |Login user |
|verify |Validasi token |

## 1.1 Register Endpoint

| <!-- --> | <!-- -->  |
| -------- | --------- |
| Type     | POST      |
| Path     | /register |

**Format Request Body Payload**

```javascript
{
    "nik": "3574526883729839",
    "role": "admin",
    "password": "password"
}
```

**Format Response Body**

```javascript
{
    "success": true,
    "message": "Data user berhasil ditambahkan!",
    "data": {
        "nik": "3574526883729839",
        "role": "admin",
        "updated_at": "2023-02-14T20:06:57.000000Z",
        "created_at": "2023-02-14T20:06:57.000000Z",
        "id": 3
    }
}
```

## 1.2 Login Endpoint

| <!-- --> | <!-- --> |
| -------- | -------- |
| Type     | POST     |
| Path     | /login   |

**Format Request Body Payload**

```javascript
{
    "nik": "3574526883729838",
    "password": "password"
}
```

**Format Response Body**

```javascript
{
    "success": true,
    "user": {
        "id": 2,
        "nik": "3574526883729838",
        "role": "admin",
        "created_at": "2023-02-12T20:18:42.000000Z",
        "updated_at": "2023-02-12T20:18:42.000000Z"
    },
    "token": "xxxxxxxxxxxxxxxxxxxxxxx.xxxxxxxxxxxxxxxxxxxxxxx.xxxxxxxxxxxxxxxxxxxxxxx"
}
```

## 1.3 Verify Endpoint

| <!-- --> | <!-- --> |
| -------- | -------- |
| Type     | GET      |
| Path     | /verify  |

**Request Header**
Authorization: "Bearer "+token

**Format Response Body**

```javascript
{
    "success": true,
    "message": "Token valid!",
    "user": {
        "id": 2,
        "nik": "3574526883729838",
        "role": "admin",
        "created_at": "2023-02-12T20:18:42.000000Z",
        "updated_at": "2023-02-12T20:18:42.000000Z"
    }
}
```

## 2. Fetch App API Specification

**Base URL:**
http://localhost:9000/api

**METHOD:**
|Endpoint |Fungsi |
|----------|---------------------------------------------------------------------------------------------------------|
|fetch |Menarik data dari API https://60c18de74f7e880017dbfd51.mockapi.io/api/v1/jabar-digital-services/product |
|aggregate |Agregasi data dari API https://60c18de74f7e880017dbfd51.mockapi.io/api/v1/jabar-digital-services/product |
|verify |Validasi token |

## 2.1 Fetch Endpoint

| <!-- --> | <!-- --> |
| -------- | -------- |
| Type     | GET      |
| Path     | /fetch   |

**Request Header**
Authorization: "Bearer "+token

**Format Response Body**

```javascript
{
    "data": [
        {
            "id": 1,
            "createdAt": "2021-06-09T09:37:05.527Z",
            "price": 218,
            "department": "Outdoors",
            "product": "Salad",
            "IDR": 3317529
        },
        {
            "id": 2,
            "createdAt": "2021-06-09T10:11:46.024Z",
            "price": 784,
            "department": "Games",
            "product": "Chips",
            "IDR": 11930930
        },
        {
            ....
        }
    ],
    "message": "success"
}
```

## 2.2 Aggregate Endpoint

| <!-- --> | <!-- -->   |
| -------- | ---------- |
| Type     | GET        |
| Path     | /aggregate |

**Request Header**
Authorization: "Bearer "+token

**Format Response Body**

```javascript
{
    "data": [
        {
            "department": "Games",
            "product": "Computer",
            "price": 60872.094
        },
        {
            "department": "Computers",
            "product": "Fish",
            "price": 76090.12
        },
        {
            ...
        }
    ],
    "message": "success"
}
```

## 2.3 Verify Endpoint

| <!-- --> | <!-- --> |
| -------- | -------- |
| Type     | GET      |
| Path     | /verify  |

**Request Header**
Authorization: "Bearer "+token

**Format Response Body**

```javascript
{
    "data": {
        "id": 2,
        "nik": "3574526883729838",
        "role": "admin"
    },
    "message": "success"
}
```
