# Webzila Test Website

A simple Go Fiber web application for testing Webzila's auto-deployment capabilities.

## Features

- **Go Fiber**: Fast HTTP framework
- **Docker**: Containerized deployment  
- **Auto-deployment**: Triggered by GitHub webhooks
- **Health checks**: `/health` endpoint for monitoring

## Endpoints

- `/` - Main website page
- `/health` - Health check endpoint
- `/api/info` - API information

## Technology Stack

- Go 1.21
- Fiber web framework
- Docker multi-stage build
- Webzila auto-deployment

## Deployment

This website is automatically deployed via Webzila when pushed to GitHub:

1. Push to `main` branch
2. GitHub webhook triggers Webzila
3. Webzila builds Docker image
4. Container deployed with SSL via Caddy
5. Available at https://web1.zfus.net

## Version

- **Current**: v1.0.0
- **Last Update**: Initial deployment test# Test webhook trigger - Tue Aug 19 13:10:00 CST 2025
# Auto-deployment test Tue Aug 19 13:11:57 CST 2025
# Final auto-deployment test Tue Aug 19 13:16:05 CST 2025
