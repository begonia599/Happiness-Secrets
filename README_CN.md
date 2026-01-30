# 🎭 Happiness Secrets

> 一个关于追寻与失落的哲学隐喻

中文文档 | [English](./README.md)

精心设计的错误页面集合，探索"不可及"的概念。每个错误页面都是一次关于缺失、中断和不可用的视觉叙事。

## 🎨 设计理念

这是一个错误页面集合项目，每个错误页面都有独特的设计风格：

- **404** - 页面未找到（暗黑红色主题）
- **502** - 网关错误（暖色橙色主题）
- **503** - 服务不可用（冷色蓝色主题）

每个页面都采用：
- 独特的配色方案
- 优雅的排版设计
- 细腻的动画效果
- 噪点纹理和光晕效果
- 实时访问统计

## 🚀 快速开始

### 使用 Docker Compose（推荐）

```bash
docker-compose up -d
```

服务将在 `http://localhost:3000` 启动

访问不同的页面：
- `http://localhost:3000/` - 首页
- `http://localhost:3000/gallery` - 展示页面
- `http://localhost:3000/404?style=dark&token=YOUR_TOKEN` - 404 页面
- `http://localhost:3000/502?style=warm&token=YOUR_TOKEN` - 502 页面
- `http://localhost:3000/503?style=cool&token=YOUR_TOKEN` - 503 页面

### 使用 Docker

```bash
# 构建镜像
docker build -t happiness-secrets .

# 运行容器
docker run -d -p 3000:3000 \
  -e DATABASE_URL=postgres://postgres:postgres@db:5432/happiness_secrets?sslmode=disable \
  happiness-secrets
```

### 本地运行

```bash
go run main.go
```

## 📊 功能特性

- ✨ 多个精美的错误页面设计
- 🎨 每个页面独特的配色和风格
- 📈 统一的访问量统计
- 💾 访问数据持久化存储（PostgreSQL）
- 🐳 Docker 容器化部署
- 📱 完全响应式设计
- 🔌 RESTful API 支持
- 🌐 CORS 跨域支持
- 🔗 易于集成到任何项目

## 🛠️ 技术栈

- **前端**: HTML5 + CSS3 (原生动画)
- **后端**: Go (标准库)
- **数据库**: PostgreSQL 16
- **部署**: Docker + Docker Compose

## 📁 项目结构

```
happiness-secrets/
├── main.go              # Go 后端服务
├── go.mod               # Go 模块定义
├── pages/               # 错误页面目录
│   ├── index.html      # 首页
│   ├── gallery.html    # 展示页面
│   ├── 404/            # 404 错误页面
│   │   └── dark.html   # 暗黑风格
│   ├── 502/            # 502 错误页面
│   │   └── warm.html   # 暖色风格
│   └── 503/            # 503 错误页面
│       └── cool.html   # 冷色风格
├── Dockerfile           # Docker 镜像定义
├── docker-compose.yml   # Docker Compose 配置
└── README.md            # 项目说明
```

## 🔌 API 使用

### Token 生成

访问首页，点击"生成新 Token"按钮创建你的专属 token。每个 token 独立追踪访问统计。

### 直接访问

直接在浏览器访问错误页面：
```
https://your-domain.com/404?style=dark&token=YOUR_TOKEN
https://your-domain.com/502?style=warm&token=YOUR_TOKEN
https://your-domain.com/503?style=cool&token=YOUR_TOKEN
```

### API 端点

通过 API 获取错误页面 HTML：
```
GET /api/error?code=404&style=dark&token=YOUR_TOKEN
GET /api/error?code=502&style=warm&token=YOUR_TOKEN
GET /api/error?code=503&style=cool&token=YOUR_TOKEN
```

响应头包含访问统计：
```
X-Visit-Count: 123
```

### Nginx 集成

在 Nginx 配置中使用自定义错误页面：

```nginx
server {
    listen 80;
    server_name your-domain.com;

    # 自定义错误页面（带 Token）
    error_page 404 https://happiness-secrets.example.com/404?style=dark&token=YOUR_TOKEN;
    error_page 502 https://happiness-secrets.example.com/502?style=warm&token=YOUR_TOKEN;
    error_page 503 https://happiness-secrets.example.com/503?style=cool&token=YOUR_TOKEN;

    location / {
        proxy_pass http://localhost:8080;
    }
}
```

### JavaScript 集成

```javascript
// 获取错误页面并读取访问统计
async function loadErrorPage(code, style, token) {
    const response = await fetch(
        `https://happiness-secrets.example.com/api/error?code=${code}&style=${style}&token=${token}`
    );
    
    // 从响应头获取访问统计
    const visitCount = response.headers.get('X-Visit-Count');
    console.log('访问次数:', visitCount);
    
    // 获取 HTML 内容
    const html = await response.text();
    document.open();
    document.write(html);
    document.close();
}

// 使用示例
const myToken = 'YOUR_TOKEN';
loadErrorPage('404', 'dark', myToken);
```

### React 集成示例

```jsx
import { useEffect, useState } from 'react';

function ErrorPage({ code = '404', style = 'dark', token }) {
    const [html, setHtml] = useState('');
    const [visitCount, setVisitCount] = useState(0);

    useEffect(() => {
        fetch(`https://happiness-secrets.example.com/api/error?code=${code}&style=${style}&token=${token}`)
            .then(res => {
                setVisitCount(res.headers.get('X-Visit-Count'));
                return res.text();
            })
            .then(setHtml)
            .catch(console.error);
    }, [code, style, token]);

    return (
        <div>
            <div>访问次数: {visitCount}</div>
            <div dangerouslySetInnerHTML={{ __html: html }} />
        </div>
    );
}
```

## 💡 使用场景

- 🌐 作为网站的自定义错误页面（通过 Nginx 配置）
- 🎨 艺术项目或概念展示
- 💻 个人网站的创意元素
- 🤔 哲学思考的视觉表达
- 🔌 作为错误页面 API 服务供其他项目调用
- 📱 前端应用的错误处理界面

**灵感来源：** 项目最初的想法来自于"幸福的秘籍"这个概念——在个人签名中放置一个看似充满希望的链接，但访问后却是404错误，形成一种关于期待与落空的隐喻。

## 📝 生产环境部署

### 使用 Nginx 反向代理

```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

### 配置 HTTPS 证书（使用 Let's Encrypt）

```bash
certbot --nginx -d your-domain.com
```

### 使用 Docker Compose 管理

```bash
# 启动
docker-compose up -d

# 查看日志
docker-compose logs -f

# 停止
docker-compose down

# 重启
docker-compose restart
```

## 🎭 哲学思考

**Happiness Secrets** 不仅仅是一个错误页面集合，它是一个关于追寻与失落的隐喻：

- 404 - 你寻找的幸福不存在
- 502 - 通往幸福的路径中断了
- 503 - 幸福暂时不可用

有时候，幸福就像这些错误页面——我们一直在寻找，却发现它可能从未存在过。
或者，这种寻找本身，就是一种幸福？

## 📄 开源协议

MIT License

## 🤝 贡献指南

我们欢迎贡献！如果你设计了一个漂亮的错误页面并想与社区分享：

### 如何贡献你的设计

1. **Fork 这个仓库**
2. **创建新分支**用于你的设计：
   ```bash
   git checkout -b feature/add-404-minimal-style
   ```
3. **添加你的错误页面**：
   - 将你的 HTML 文件放在相应的目录中（例如：`pages/404/your-style.html`）
   - 遵循现有的命名规范：`pages/{错误码}/{样式名称}.html`
   - 确保你的设计是自包含的（推荐使用内联 CSS/JS）
4. **更新后端代码** (`main.go`)：
   - 在相应的 switch case 中添加你的样式
5. **更新展示页面** (`pages/gallery.html`)：
   - 为你的设计添加预览卡片
6. **测试你的设计**：
   ```bash
   docker-compose up -d
   # 访问 http://localhost:3000/gallery 预览
   ```
7. **提交并推送**：
   ```bash
   git add .
   git commit -m "Add minimal style for 404 page"
   git push origin feature/add-404-minimal-style
   ```
8. **创建 Pull Request** 到 `main` 分支

### 设计指南

- **自包含**：优先使用内联 CSS 和 JavaScript
- **响应式**：确保你的设计在移动设备上正常工作
- **可访问性**：遵循无障碍最佳实践
- **独特性**：带来你自己的创意视角
- **轻量级**：保持文件大小合理

### 我们期待的贡献

- 现有错误码的新样式（404、502、503）
- 新的错误码（500、403 等）
- 对"不可及"主题的创意诠释
- 独特的字体、配色方案和动画效果

欢迎提交 Issue 和 Pull Request！

---

Made with 🎭 and philosophical humor
