# Quick Deployment Guide

## 🚀 Deploy in 3 Steps

### Step 1: Deploy API (Railway) - 5 minutes

1. Go to https://railway.app → Sign up with GitHub
2. New Project → Deploy from GitHub repo
3. Set root directory: `api`
4. Add environment variables:
   ```
   SUPABASE_URL=your-supabase-url
   SUPABASE_KEY=your-supabase-key
   PORT=8080
   ```
5. Deploy → Copy API URL (e.g., `https://your-api.railway.app`)

**Test:** `curl https://your-api.railway.app/api/health`

---

### Step 2: Deploy Web UI (Vercel) - 3 minutes

1. Go to https://vercel.com → Sign up with GitHub
2. Add New Project → Import GitHub repo
3. Set root directory: `web`
4. Add environment variable:
   ```
   NEXT_PUBLIC_API_URL=https://your-api.railway.app
   ```
5. Deploy → Copy Vercel URL (e.g., `https://your-app.vercel.app`)

---

### Step 3: Update CORS - 1 minute

1. Go back to Railway → Your API → Variables
2. Add/Update:
   ```
   ALLOWED_ORIGIN=https://your-app.vercel.app
   ```
3. API auto-redeploys

**Done!** 🎉 Open your Vercel URL and test!

---

## 📋 What You Need

- GitHub repository (code pushed)
- Supabase credentials (already have)
- Railway account (free)
- Vercel account (free)

---

## 🔍 Verify Everything Works

1. **API Health:**
   ```bash
   curl https://your-api.railway.app/api/health
   ```

2. **Web UI:**
   - Open Vercel URL
   - Add a todo
   - Mark complete
   - Delete todo

3. **Check Logs:**
   - Railway → Deployments → View logs
   - Vercel → Deployments → View logs

---

## 🆘 Common Issues

| Problem | Fix |
|---------|-----|
| CORS error | Update `ALLOWED_ORIGIN` in Railway |
| Can't connect | Check `NEXT_PUBLIC_API_URL` in Vercel |
| Database error | Verify Supabase credentials in Railway |

---

## 📚 Full Guide

See `DEPLOYMENT_GUIDE.md` for detailed instructions.

