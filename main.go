package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Webzila Test Website v1.0",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Webzila Test Site</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
            margin: 0;
            padding: 40px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        .container {
            background: white;
            padding: 60px;
            border-radius: 20px;
            box-shadow: 0 20px 40px rgba(0,0,0,0.1);
            text-align: center;
            max-width: 600px;
        }
        h1 {
            color: #333;
            margin-bottom: 20px;
            font-size: 2.5em;
        }
        .badge {
            display: inline-block;
            background: #4CAF50;
            color: white;
            padding: 8px 16px;
            border-radius: 20px;
            font-size: 0.9em;
            margin: 10px 0;
        }
        .info {
            color: #666;
            margin: 20px 0;
            line-height: 1.6;
        }
        .tech-stack {
            margin-top: 30px;
            padding: 20px;
            background: #f8f9fa;
            border-radius: 10px;
        }
        .tech-item {
            display: inline-block;
            background: #007bff;
            color: white;
            padding: 5px 12px;
            margin: 5px;
            border-radius: 15px;
            font-size: 0.8em;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🚀 Webzila Test Website</h1>
        <div class="badge">Auto-Deployed via GitHub Webhook</div>
        
        <div class="info">
            <p>This is a simple Go Fiber web application deployed automatically by Webzila.</p>
            <p><strong>Domain:</strong> web1.zfus.net</p>
            <p><strong>Deployment:</strong> Automatic via GitHub push</p>
            <p><strong>Version:</strong> 1.0.0</p>
        </div>

        <div class="tech-stack">
            <h3>Technology Stack</h3>
            <span class="tech-item">Go</span>
            <span class="tech-item">Fiber</span>
            <span class="tech-item">Docker</span>
            <span class="tech-item">Webzila</span>
            <span class="tech-item">Caddy</span>
            <span class="tech-item">GitHub</span>
        </div>
    </div>
</body>
</html>
		`
		return c.Type("html").SendString(html)
	})

	// Health endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "healthy",
			"service": "webzila-test-website",
			"version": "1.0.0",
		})
	})

	// API endpoint
	app.Get("/api/info", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":        "Webzila Test Website",
			"version":     "1.0.0",
			"description": "Simple Go Fiber app for testing auto-deployment",
			"technology":  []string{"Go", "Fiber", "Docker", "Webzila"},
		})
	})

	// Get port from environment or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🌐 Test website starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}