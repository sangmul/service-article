# 📘 Article API (Golang Clean Architecture)

## 🚀 Setup Project

### 1. Request Environment

Silakan request file environment ke PM melalui:

* Email
* WhatsApp

---

### 2. Setup File `.env` dan  `certs/ca.pem `

Buat file `.env` di root project, lalu isi dengan parameter berikut:

```
DB_USER=
DB_PASS=
DB_HOST=
DB_PORT=
DB_NAME=
DB_TLS=
```

Buat file `ca.pem` di folder certs

```
-----BEGIN CERTIFICATE-----
XXXXX
-----END CERTIFICATE-----
```


Isi value sesuai dengan yang diberikan oleh PM.

---

### 3. Postman Collection

Jika membutuhkan Postman Collection, silakan request ke PM melalui:

* Email
* WhatsApp

---

## 🌐 API Endpoint

### 🔹 Create Article

* **POST** `/article`

Body:

```
{
  "title": "string (min 20 char)",
  "content": "string (min 200 char)",
  "category": "string (min 3 char)",
  "status": "publish | draft | thrash"
}
```

---

### 🔹 Get All Article

* **GET** `/article`
* Optional query:

  * `limit`
  * `offset`

Contoh:

```
/article?limit=10&offset=0
```

---

### 🔹 Get Article By ID

* **GET** `/article/:id`

---

### 🔹 Update Article

* **PUT / PATCH** `/article/:id`

Body:

```
{
  "title": "",
  "content": "",
  "category": "",
  "status": ""
}
```

---

### 🔹 Delete Article

* **DELETE** `/article/:id`

---

## ▶️ Run Application

```bash
go run cmd/main.go
```

---

## 🎉 Selamat Mencoba!

Jika ada kendala, silakan hubungi PM.
