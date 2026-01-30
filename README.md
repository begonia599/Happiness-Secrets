# 🎭 Happiness Secrets

> A philosophical metaphor about pursuit and loss

[中文文档](./README_CN.md) | English

A beautifully designed error page API service that explores the concept of "unattainability". Each error page is a visual narrative about absence, interruption, and unavailability.

## 🎨 Design Philosophy

This is a collection of error pages, each with a unique design style:

- **404** - Page Not Found (Dark red theme)
- **502** - Bad Gateway (Warm orange theme)
- **503** - Service Unavailable (Cool blue theme)

Each page features:
- Unique color schemes
- Elegant typography
- Smooth animations
- Noise textures and glow effects
- Real-time visit statistics

## 🚀 Quick Start

### Using Docker Compose (Recommended)

```bash
docker-compose up -d
```

The service will start on `http://localhost:3000`

Access different error pages:
- `http://localhost:3000/` - Homepage
- `http://localhost:3000/gallery` - Gallery page
- `http://localhost:3000/404?style=dark&token=YOUR_TOKEN` - 404 page
- `http://localhost:3000/502?style=warm&token=YOUR_TOKEN` - 502 page
- `http://localhost:3000/503?style=cool&token=YOUR_TOKEN` - 503 page

### Using Docker

```bash
# Build image
docker build -t happiness-secrets .

# Run container
docker run -d -p 3000:3000 \
  -e DATABASE_URL=postgres://postgres:postgres@db:5432/happiness_secrets?sslmode=disable \
  happiness-secrets
```

### Local Development

```bash
go run main.go
```

## � Features

- ✨ Multiple beautifully designed error pages
- 🎨 Unique color scheme and style for each page
- 📈 Unified visit statistics
- 💾 Persistent data storage (PostgreSQL)
- 🐳 Docker containerized deployment
- 📱 Fully responsive design
- 🔌 RESTful API support
- 🌐 CORS cross-origin support
- 🔗 Easy integration into any project

## 🛠️ Tech Stack

- **Frontend**: HTML5 + CSS3 (Native animations)
- **Backend**: Go (Standard library)
- **Database**: PostgreSQL 16
- **Deployment**: Docker + Docker Compose

## 📁 Project Structure

```
happiness-secrets/
├── main.go              # Go backend service
├── go.mod               # Go module definition
├── pages/               # Error pages directory
│   ├── index.html      # Homepage
│   ├── gallery.html    # Gallery page
│   ├── 404/            # 404 error pages
│   │   └── dark.html   # Dark style
│   ├── 502/            # 502 error pages
│   │   └── warm.html   # Warm style
│   └── 503/            # 503 error pages
│       └── cool.html   # Cool style
├── Dockerfile           # Docker image definition
├── docker-compose.yml   # Docker Compose configuration
└── README.md            # Project documentation
```

## 🔌 API Usage

### Token Generation

Visit the homepage and click "Generate Token" to create your unique token. Each token independently tracks visit statistics.

### Direct Access

Access error pages directly in the browser:
```
https://your-domain.com/404?style=dark&token=YOUR_TOKEN
https://your-domain.com/502?style=warm&token=YOUR_TOKEN
https://your-domain.com/503?style=cool&token=YOUR_TOKEN
```

### API Endpoint

Get error page HTML via API:
```
GET /api/error?code=404&style=dark&token=YOUR_TOKEN
GET /api/error?code=502&style=warm&token=YOUR_TOKEN
GET /api/error?code=503&style=cool&token=YOUR_TOKEN
```

Response headers include visit count:
```
X-Visit-Count: 123
```

### Nginx Integration

Use custom error pages in Nginx configuration:

```nginx
server {
    listen 80;
    server_name your-domain.com;

    # Custom error pages (with Token)
    error_page 404 https://happiness-secrets.example.com/404?style=dark&token=YOUR_TOKEN;
    error_page 502 https://happiness-secrets.example.com/502?style=warm&token=YOUR_TOKEN;
    error_page 503 https://happiness-secrets.example.com/503?style=cool&token=YOUR_TOKEN;

    location / {
        proxy_pass http://localhost:8080;
    }
}
```

### JavaScript Integration

```javascript
// Get error page and read visit statistics
async function loadErrorPage(code, style, token) {
    const response = await fetch(
        `https://happiness-secrets.example.com/api/error?code=${code}&style=${style}&token=${token}`
    );
    
    // Get visit count from response headers
    const visitCount = response.headers.get('X-Visit-Count');
    console.log('Visit count:', visitCount);
    
    // Get HTML content
    const html = await response.text();
    document.open();
    document.write(html);
    document.close();
}

// Usage example
const myToken = 'YOUR_TOKEN';
loadErrorPage('404', 'dark', myToken);
```

### React Integration Example

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
            <div>Visit count: {visitCount}</div>
            <div dangerouslySetInnerHTML={{ __html: html }} />
        </div>
    );
}
```

## 💡 Use Cases

- 🌐 Custom error pages for websites (via Nginx configuration)
- 🎨 Art projects or conceptual exhibitions
- 💻 Creative elements for personal websites
- 🤔 Visual expression of philosophical thinking
- 🔌 Error page API service for other projects
- 📱 Error handling interface for frontend applications

**Inspiration:** The initial idea came from the concept of "Happiness Secrets" - placing a seemingly hopeful link in a personal signature, but upon access, it's a 404 error, forming a metaphor about expectation and disappointment.

## 📝 Production Deployment

### Using Nginx Reverse Proxy

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

### Configure HTTPS Certificate (using Let's Encrypt)

```bash
certbot --nginx -d your-domain.com
```

### Manage with Docker Compose

```bash
# Start
docker-compose up -d

# View logs
docker-compose logs -f

# Stop
docker-compose down

# Restart
docker-compose restart
```

## 🎭 Philosophical Reflection

**Happiness Secrets** is not just a collection of error pages, it's a metaphor about pursuit and loss:

- 404 - The happiness you seek doesn't exist
- 502 - The path to happiness is interrupted
- 503 - Happiness is temporarily unavailable

Sometimes, happiness is like these error pages - we keep searching, only to find it may never have existed.
Or perhaps, this search itself is a form of happiness?

## 📄 License

MIT License

## 🤝 Contributing

Issues and Pull Requests are welcome!

---

Made with 🎭 and philosophical humor
