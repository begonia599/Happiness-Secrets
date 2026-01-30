package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// 连接 PostgreSQL
	var err error
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@db:5432/happiness_secrets?sslmode=disable"
	}

	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("无法连接数据库:", err)
	}
	defer db.Close()

	// 等待数据库就绪
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Printf("等待数据库连接... (%d/30)\n", i+1)
		time.Sleep(time.Second)
	}

	if err != nil {
		log.Fatal("数据库连接超时:", err)
	}

	// 初始化数据库表
	initDB()

	// 设置路由
	http.HandleFunc("/", serveErrorPage)
	http.HandleFunc("/404", serve404)
	http.HandleFunc("/502", serve502)
	http.HandleFunc("/503", serve503)
	http.HandleFunc("/api/token", handleTokenGeneration)
	http.HandleFunc("/api/error", handleErrorAPI)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🎭 Happiness Secrets 服务器启动在端口 %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, corsMiddleware(http.DefaultServeMux)))
}

// 初始化数据库表
func initDB() {
	schema := `
	CREATE TABLE IF NOT EXISTS tokens (
		id SERIAL PRIMARY KEY,
		token VARCHAR(255) UNIQUE NOT NULL,
		visit_count INTEGER DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		last_visit_at TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_tokens_token ON tokens(token);
	`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal("初始化数据库失败:", err)
	}

	log.Println("✓ 数据库初始化成功")
}

// CORS 中间件
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func serveErrorPage(w http.ResponseWriter, r *http.Request) {
	// favicon 请求
	if r.URL.Path == "/favicon.ico" {
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write([]byte(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100"><text y="75" font-size="75">🎭</text></svg>`))
		return
	}

	// 主页显示介绍页面
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "pages/index.html")
		return
	}

	// 展示页面路由
	if r.URL.Path == "/gallery" {
		http.ServeFile(w, r, "pages/gallery.html")
		return
	}

	// 其他路径显示 404
	serve404(w, r)
}

func serve404(w http.ResponseWriter, r *http.Request) {
	// 检查是否有 style 参数
	style := r.URL.Query().Get("style")
	token := r.URL.Query().Get("token")

	if token != "" {
		incrementVisit(w, token)
	}

	// 根据 style 参数选择不同的 404 页面
	var filePath string
	switch style {
	case "dark":
		filePath = "pages/404/dark.html"
	case "minimal":
		filePath = "pages/404/minimal.html"
	case "creative":
		filePath = "pages/404/creative.html"
	default:
		// 默认样式
		filePath = "pages/404/dark.html"
	}

	http.ServeFile(w, r, filePath)
}

func serve502(w http.ResponseWriter, r *http.Request) {
	style := r.URL.Query().Get("style")
	token := r.URL.Query().Get("token")

	if token != "" {
		incrementVisit(w, token)
	}

	var filePath string
	switch style {
	case "warm":
		filePath = "pages/502/warm.html"
	case "minimal":
		filePath = "pages/502/minimal.html"
	case "tech":
		filePath = "pages/502/tech.html"
	default:
		filePath = "pages/502/warm.html"
	}

	http.ServeFile(w, r, filePath)
}

func serve503(w http.ResponseWriter, r *http.Request) {
	style := r.URL.Query().Get("style")
	token := r.URL.Query().Get("token")

	if token != "" {
		incrementVisit(w, token)
	}

	var filePath string
	switch style {
	case "cool":
		filePath = "pages/503/cool.html"
	case "minimal":
		filePath = "pages/503/minimal.html"
	case "modern":
		filePath = "pages/503/modern.html"
	default:
		filePath = "pages/503/cool.html"
	}

	http.ServeFile(w, r, filePath)
}

// 生成新 Token
func handleTokenGeneration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 生成 UUID 作为 token
	token := uuid.New().String()

	// 插入数据库
	_, err := db.Exec(
		"INSERT INTO tokens (token, visit_count, created_at) VALUES ($1, 0, $2)",
		token,
		time.Now(),
	)

	if err != nil {
		log.Printf("创建 token 失败: %v\n", err)
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	log.Printf("🔑 新 Token 生成: %s\n", token)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

// 增加访问计数
func incrementVisit(w http.ResponseWriter, token string) int {
	var count int

	err := db.QueryRow(`
		UPDATE tokens 
		SET visit_count = visit_count + 1, last_visit_at = $1
		WHERE token = $2
		RETURNING visit_count
	`, time.Now(), token).Scan(&count)

	if err != nil {
		log.Printf("更新访问计数失败: %v\n", err)
		return 0
	}

	// 在响应头中返回访问计数
	w.Header().Set("X-Visit-Count", string(rune(count)))

	log.Printf("📊 Token %s 访问计数: %d\n", token[:8]+"...", count)

	return count
}

// API 端点：返回错误页面的 HTML
func handleErrorAPI(w http.ResponseWriter, r *http.Request) {
	errorCode := r.URL.Query().Get("code")
	token := r.URL.Query().Get("token")

	if errorCode == "" {
		errorCode = "404"
	}

	var filePath string
	switch errorCode {
	case "404":
		filePath = "pages/404.html"
	case "502":
		filePath = "pages/502.html"
	case "503":
		filePath = "pages/503.html"
	default:
		http.Error(w, "Unsupported error code", http.StatusBadRequest)
		return
	}

	// 如果提供了 token，增加访问计数
	if token != "" {
		incrementVisit(w, token)
	}

	// 读取文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Failed to read error page", http.StatusInternalServerError)
		return
	}

	// 返回 HTML 内容
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(content)

	log.Printf("📄 API 请求: %s (token: %s)\n", errorCode, token[:8]+"...")
}
