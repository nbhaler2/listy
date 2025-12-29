# Deployment Guide - Listy Todo App

This guide will help you deploy both the Go API server and Next.js frontend.

## Architecture

```
Vercel (Web UI) ──> Railway (Go API) ──> Supabase (Database)
```

## Prerequisites

1. **GitHub account** - For connecting repositories
2. **Railway account** - For API server (free tier available)
3. **Vercel account** - For Web UI (free tier available)
4. **Supabase project** - Already set up

---

## Part 1: Deploy Go API Server (Railway)

### Step 1: Create Railway Account

1. Go to https://railway.app
2. Sign up with GitHub
3. Create a new project

### Step 2: Deploy API

**Option A: Deploy from GitHub (Recommended)**

1. Push your code to GitHub:
   ```bash
   git add .
   git commit -m "Prepare for deployment"
   git push origin main
   ```

2. In Railway:
   - Click "New Project"
   - Select "Deploy from GitHub repo"
   - Choose your repository
   - Select the `api` folder as root directory

3. Configure environment variables:
   - Go to Variables tab
   - Add:
     ```
     SUPABASE_URL=your-supabase-url
     SUPABASE_KEY=your-supabase-key
     PORT=8080
     ALLOWED_ORIGIN=https://your-vercel-url.vercel.app
     ```

4. Railway will automatically:
   - Detect Go project
   - Build and deploy
   - Provide a URL (e.g., `https://your-api.railway.app`)

**Option B: Deploy with Dockerfile**

1. Railway will use the Dockerfile in `api/` directory
2. Same environment variables as above

### Step 3: Get API URL

After deployment, Railway will give you a URL like:
```
https://your-api-name.railway.app
```

**Save this URL** - you'll need it for the Web UI!

### Step 4: Test API

```bash
curl https://your-api-name.railway.app/api/health
```

Should return: `{"service":"listy-api","status":"healthy"}`

---

## Part 2: Deploy Next.js Web UI (Vercel)

### Step 1: Create Vercel Account

1. Go to https://vercel.com
2. Sign up with GitHub

### Step 2: Deploy Web UI

1. In Vercel dashboard:
   - Click "Add New Project"
   - Import your GitHub repository
   - Set root directory to `web`
   - Framework: Next.js (auto-detected)

2. Configure environment variables:
   - Go to Settings → Environment Variables
   - Add:
     ```
     NEXT_PUBLIC_API_URL=https://your-api-name.railway.app
     ```

3. Deploy:
   - Click "Deploy"
   - Vercel will build and deploy automatically

### Step 3: Update API CORS

After getting your Vercel URL, update Railway environment variables:

1. Go to Railway → Your API → Variables
2. Update `ALLOWED_ORIGIN`:
   ```
   ALLOWED_ORIGIN=https://your-app.vercel.app
   ```
3. Redeploy API (Railway will auto-redeploy)

### Step 4: Test Web UI

Open your Vercel URL (e.g., `https://your-app.vercel.app`)

You should see the todo list interface!

---

## Part 3: Update CLI to Use Deployed API

Update your CLI to use the deployed API:

```bash
# Set environment variable
export LISTY_API_URL=https://your-api-name.railway.app

# Or create a config file
echo "LISTY_API_URL=https://your-api-name.railway.app" > ~/.listy/config
```

Then use CLI normally:
```bash
cd cli
go run main.go api_client.go config.go list
```

---

## Environment Variables Summary

### Railway (API Server)
```
SUPABASE_URL=your-supabase-url
SUPABASE_KEY=your-supabase-key
PORT=8080
ALLOWED_ORIGIN=https://your-vercel-url.vercel.app
```

### Vercel (Web UI)
```
NEXT_PUBLIC_API_URL=https://your-railway-api.railway.app
```

---

## Testing After Deployment

1. **Test API:**
   ```bash
   curl https://your-api.railway.app/api/health
   curl https://your-api.railway.app/api/todos
   ```

2. **Test Web UI:**
   - Open Vercel URL
   - Add a todo
   - Check if it appears

3. **Test CLI:**
   ```bash
   LISTY_API_URL=https://your-api.railway.app go run main.go list
   ```

---

## Troubleshooting

### API not accessible
- Check Railway logs
- Verify environment variables
- Check if port is correct

### CORS errors in browser
- Verify `ALLOWED_ORIGIN` in Railway matches Vercel URL
- Check browser console for exact error

### Web UI can't connect to API
- Verify `NEXT_PUBLIC_API_URL` in Vercel
- Check API is running (health check)
- Check network tab in browser dev tools

### Database connection issues
- Verify Supabase credentials in Railway
- Check Supabase dashboard for connection status

---

## Alternative Deployment Options

### API Server Alternatives:
- **Fly.io** - Good for Go apps
- **Render** - Simple deployment
- **AWS/GCP** - More control

### Web UI Alternatives:
- **Netlify** - Similar to Vercel
- **Railway** - Can host both together

---

## Cost Estimate

**Free Tier:**
- Railway: Free tier available (limited usage)
- Vercel: Free tier (generous limits)
- Supabase: Free tier (500MB database)

**Total: $0/month** for small projects! 🎉

---

## Next Steps After Deployment

1. ✅ Test all functionality
2. ✅ Share with others
3. ✅ Monitor performance
4. 🎯 Add AI features
5. 🎯 Add authentication
6. 🎯 Scale as needed

