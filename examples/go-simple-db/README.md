---

### **📌 Go Simple Database**  

A lightweight and simple in-memory database built with **Go** and **Gin** for learning purposes. Supports **CRUD operations** with concurrent access handling using `sync.RWMutex`.  

---

## **🚀 Features**  

✅ Insert and Update operations  
✅ Retrieve records by ID  
✅ JSON-based storage format  
✅ Integrated with **Gin** for API access  
✅ Thread-safe with **RWMutex**  
✅ Simple file-based persistence  

---


## **🛠️ Installation & Setup**  

### **1️⃣ Clone the Repository**  
```sh
git clone https://github.com/yourusername/go-simple-db.git
cd go-simple-db
```

### **2️⃣ Install Dependencies**  
```sh
go mod tidy
```

### **3️⃣ Run the Application**  
```sh
go run main.go
```

The server will start at **`http://localhost:8080`** 🚀  

---

## **📡 API Endpoints**  

| Method | Endpoint            | Description            |
|--------|---------------------|------------------------|
| `POST` | `/insert`           | Insert a new record   |
| `PUT`  | `/update/:id`       | Update a record by ID |
| `GET`  | `/get/:id`          | Retrieve a record     |
| `DELETE`  | `/delete/:id`    | Delete a record       |
| `GET`  | `/all`              | Get all records       |

---

## **📝 Example Record Format**  

```json
{
    "id": 1,
    "attributes": {
        "name": "Go Database",
        "language": "Go"
    }
}
```

---

Let me know if you want to **modify** anything! 🚀
