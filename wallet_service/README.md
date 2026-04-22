# Wallet Service

## 執行方式

### Docker

啟動完整服務（PostgreSQL + migration + 3 個 app replica + nginx）：

```bash
docker compose up --build
```

服務啟動後監聽 `http://localhost:8080`。

### 本機開發

啟動資料庫與執行 migration：

```bash
docker compose up -d postgres
make migrate-up
```

啟動 server：

```bash
make run
```

DB 預設連線 `postgres://wallet:wallet@localhost:55432/wallet`。

---

## API

| Method | Path            | 說明                                              |
| ------ | --------------- | ------------------------------------------------- |
| `POST` | `/accounts`     | 建立帳戶，預設餘額 10000                          |
| `GET`  | `/accounts/:id` | 查詢帳戶餘額（支援 `Accept: application/x-yaml`） |
| `POST` | `/transfer`     | 轉帳                                              |

### 建立帳戶

```http
POST /accounts
Content-Type: application/json

{"name": "Alice"}
```

### 查詢餘額

```http
GET /accounts/{id}
```

### 轉帳

```http
POST /transfer
Content-Type: application/json

{"from_id": "...", "to_id": "...", "amount": "100"}
```

`amount` 必須為字串，最多 8 位小數，不可為負數或零。
