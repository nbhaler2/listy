# Deployment Checklist

Use this checklist to ensure everything is ready before deploying.

## Pre-Deployment

### API Server (Railway)
- [ ] Code is pushed to GitHub
- [ ] `api/Dockerfile` exists
- [ ] `api/go.mod` is up to date
- [ ] Environment variables documented:
  - [ ] `SUPABASE_URL`
  - [ ] `SUPABASE_KEY`
  - [ ] `PORT` (optional, defaults to 8080)
  - [ ] `ALLOWED_ORIGIN` (will set after Vercel deployment)

### Web UI (Vercel)
- [ ] `web/package.json` exists
- [ ] `web/next.config.js` or `next.config.ts` exists (if needed)
- [ ] Environment variable documented:
  - [ ] `NEXT_PUBLIC_API_URL` (will set after Railway deployment)

### Database (Supabase)
- [ ] Supabase project is active
- [ ] `todos` table exists
- [ ] RLS policies are set (public read/write for now)
- [ ] Credentials are saved securely

---

## Deployment Steps

### Step 1: Deploy API (Railway)
1. [ ] Create Railway account
2. [ ] Create new project
3. [ ] Connect GitHub repository
4. [ ] Set root directory to `api`
5. [ ] Add environment variables:
   - [ ] `SUPABASE_URL`
   - [ ] `SUPABASE_KEY`
   - [ ] `PORT=8080`
6. [ ] Deploy
7. [ ] Get API URL (e.g., `https://your-api.railway.app`)
8. [ ] Test health endpoint: `curl https://your-api.railway.app/api/health`

### Step 2: Deploy Web UI (Vercel)
1. [ ] Create Vercel account
2. [ ] Import GitHub repository
3. [ ] Set root directory to `web`
4. [ ] Add environment variable:
   - [ ] `NEXT_PUBLIC_API_URL=https://your-api.railway.app`
5. [ ] Deploy
6. [ ] Get Vercel URL (e.g., `https://your-app.vercel.app`)

### Step 3: Update API CORS
1. [ ] Go back to Railway
2. [ ] Add/Update environment variable:
   - [ ] `ALLOWED_ORIGIN=https://your-app.vercel.app`
3. [ ] Redeploy API (auto-redeploys on env change)

### Step 4: Final Testing
1. [ ] Open Vercel URL in browser
2. [ ] Test adding a todo
3. [ ] Test marking todo as complete
4. [ ] Test deleting a todo
5. [ ] Test filtering (pending/completed)
6. [ ] Check browser console for errors
7. [ ] Check Railway logs for API errors

---

## Post-Deployment

### Documentation
- [ ] Save all URLs:
  - [ ] API URL: `_________________`
  - [ ] Web UI URL: `_________________`
  - [ ] Supabase URL: `_________________`

### Monitoring
- [ ] Check Railway logs
- [ ] Check Vercel logs
- [ ] Monitor Supabase dashboard

### Next Steps
- [ ] Share with others
- [ ] Test on mobile devices
- [ ] Plan AI feature integration
- [ ] Consider adding authentication

---

## Troubleshooting Quick Reference

| Issue | Solution |
|-------|----------|
| API not accessible | Check Railway logs, verify env vars |
| CORS errors | Update `ALLOWED_ORIGIN` in Railway |
| Web UI can't connect | Verify `NEXT_PUBLIC_API_URL` in Vercel |
| Database errors | Check Supabase credentials in Railway |
| Build fails | Check logs, verify dependencies |

---

## Quick Commands

### Test API locally before deploying:
```bash
cd api
go run main.go
curl http://localhost:8080/api/health
```

### Test Web UI locally:
```bash
cd web
npm run dev
# Open http://localhost:3000
```

### Test deployed API:
```bash
curl https://your-api.railway.app/api/health
curl https://your-api.railway.app/api/todos
```

