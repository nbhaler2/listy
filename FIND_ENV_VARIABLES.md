# How to Find and Set Environment Variables

This guide shows you exactly where to find each value and how to add them in Railway.

---

## Step 1: Find Your Supabase Credentials

### Option A: If you already have a `.env` file locally

1. Open your `.env` file in the project root:
   ```bash
   cat .env
   ```
   
   You'll see:
   ```
   SUPABASE_URL=https://xxxxx.supabase.co
   SUPABASE_KEY=eyJhbGc...
   ```
   
   **Copy these values** - you'll need them for Railway!

### Option B: Get from Supabase Dashboard

1. **Go to Supabase Dashboard:**
   - Visit https://supabase.com
   - Sign in to your account
   - Select your project (or create one if you don't have it)

2. **Get Your Credentials:**
   - Click **Settings** (gear icon) in the left sidebar
   - Click **API** in the settings menu
   - You'll see two important values:

   **a) Project URL (SUPABASE_URL):**
   ```
   https://xxxxxxxxxxxxx.supabase.co
   ```
   - This is your `SUPABASE_URL`
   - Copy the entire URL

   **b) anon public key (SUPABASE_KEY):**
   ```
   eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Inh4eHh4eHh4eHgiLCJyb2xlIjoiYW5vbiIsImlhdCI6MTY0NjE2ODAwMCwiZXhwIjoxOTYxNzQ0MDAwfQ.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
   ```
   - This is your `SUPABASE_KEY`
   - Make sure to use the **"anon public"** key (NOT the "service_role" key)
   - Copy the entire long string

---

## Step 2: Add Environment Variables in Railway

### Detailed Steps:

1. **Go to Railway Dashboard:**
   - Visit https://railway.app
   - Sign in (or create account with GitHub)
   - You should see your project (or create a new one)

2. **Navigate to Your API Service:**
   - Click on your project
   - Click on the service that runs your API (or create one if deploying for first time)

3. **Open Variables Tab:**
   - In the service view, click on the **"Variables"** tab (or look for "Environment" or "Env" in the menu)
   - You'll see a list of environment variables (might be empty initially)

4. **Add Each Variable:**

   **a) Add SUPABASE_URL:**
   - Click **"+ New Variable"** or **"Add Variable"** button
   - **Name:** `SUPABASE_URL`
   - **Value:** Paste your Supabase URL (e.g., `https://xxxxx.supabase.co`)
   - Click **"Add"** or **"Save"**

   **b) Add SUPABASE_KEY:**
   - Click **"+ New Variable"** again
   - **Name:** `SUPABASE_KEY`
   - **Value:** Paste your anon public key (the long string)
   - Click **"Add"** or **"Save"**

   **c) Add PORT (Optional):**
   - Click **"+ New Variable"** again
   - **Name:** `PORT`
   - **Value:** `8080`
   - Click **"Add"** or **"Save"**
   - *Note: Railway might set this automatically, but it's good to set it explicitly*

5. **Verify Variables:**
   - You should now see all three variables listed:
     ```
     SUPABASE_URL = https://xxxxx.supabase.co
     SUPABASE_KEY = eyJhbGc...
     PORT = 8080
     ```

6. **Deploy:**
   - Railway will automatically redeploy when you add variables
   - Or click **"Deploy"** button if needed
   - Wait for deployment to complete (usually 1-2 minutes)

---

## Step 3: Get Your API URL (After Deployment)

1. **In Railway Dashboard:**
   - Go to your API service
   - Look for a section called **"Settings"** or **"Networking"**
   - Or look at the top of the service page

2. **Find the Domain/URL:**
   - You'll see something like:
     ```
     https://your-api-name.railway.app
     ```
   - Or it might show as:
     ```
     https://your-project-production.up.railway.app
     ```
   - **This is your API URL!**

3. **Test It:**
   ```bash
   curl https://your-api-name.railway.app/api/health
   ```
   
   Should return: `{"service":"listy-api","status":"healthy"}`

4. **Save This URL:**
   - You'll need it for the Web UI deployment (as `NEXT_PUBLIC_API_URL`)

---

## Step 4: Get Your Vercel URL (After Web UI Deployment)

1. **In Vercel Dashboard:**
   - Visit https://vercel.com
   - Sign in and go to your project

2. **Find the URL:**
   - After deployment, Vercel shows your URL at the top
   - It looks like:
     ```
     https://your-app-name.vercel.app
     ```
   - Or a custom domain if you set one up

3. **Save This URL:**
   - You'll need it to update Railway's `ALLOWED_ORIGIN` variable

---

## Step 5: Update ALLOWED_ORIGIN in Railway

After you get your Vercel URL:

1. **Go back to Railway:**
   - Open your API service
   - Go to **"Variables"** tab

2. **Add ALLOWED_ORIGIN:**
   - Click **"+ New Variable"**
   - **Name:** `ALLOWED_ORIGIN`
   - **Value:** Your Vercel URL (e.g., `https://your-app-name.vercel.app`)
   - Click **"Add"** or **"Save"**

3. **Railway will auto-redeploy** with the new CORS setting

---

## Complete Environment Variables Summary

### For Railway (API Server):
```
SUPABASE_URL=https://xxxxx.supabase.co          ← From Supabase Settings > API
SUPABASE_KEY=eyJhbGc...                         ← From Supabase Settings > API (anon key)
PORT=8080                                        ← Set to 8080 (default)
ALLOWED_ORIGIN=https://your-app.vercel.app      ← From Vercel (add after Web UI deploys)
```

### For Vercel (Web UI):
```
NEXT_PUBLIC_API_URL=https://your-api.railway.app  ← From Railway (after API deploys)
```

---

## Visual Guide (What You'll See)

### Railway Variables Tab:
```
┌─────────────────────────────────────┐
│  Environment Variables              │
├─────────────────────────────────────┤
│  SUPABASE_URL                       │
│  https://xxxxx.supabase.co          │
│  [Edit] [Delete]                    │
├─────────────────────────────────────┤
│  SUPABASE_KEY                       │
│  eyJhbGciOiJIUzI1NiIsInR5cCI6...   │
│  [Edit] [Delete]                    │
├─────────────────────────────────────┤
│  PORT                               │
│  8080                               │
│  [Edit] [Delete]                    │
├─────────────────────────────────────┤
│  [+ New Variable]                   │
└─────────────────────────────────────┘
```

### Supabase API Settings:
```
┌─────────────────────────────────────┐
│  API Settings                       │
├─────────────────────────────────────┤
│  Project URL                        │
│  https://xxxxx.supabase.co          │
│  [Copy]                             │
├─────────────────────────────────────┤
│  anon public                        │
│  eyJhbGciOiJIUzI1NiIsInR5cCI6...   │
│  [Copy]                             │
├─────────────────────────────────────┤
│  service_role (secret)              │
│  ⚠️ Don't use this one!             │
└─────────────────────────────────────┘
```

---

## Quick Checklist

- [ ] Found SUPABASE_URL from Supabase dashboard
- [ ] Found SUPABASE_KEY (anon public) from Supabase dashboard
- [ ] Added SUPABASE_URL to Railway
- [ ] Added SUPABASE_KEY to Railway
- [ ] Added PORT=8080 to Railway
- [ ] Deployed API and got Railway URL
- [ ] Added NEXT_PUBLIC_API_URL to Vercel (after API deploys)
- [ ] Deployed Web UI and got Vercel URL
- [ ] Added ALLOWED_ORIGIN to Railway (after Web UI deploys)

---

## Troubleshooting

**"Can't find Variables tab in Railway":**
- Look for "Environment" or "Env" instead
- Or check the service settings

**"Supabase credentials not working":**
- Make sure you're using the **anon public** key (not service_role)
- Check for extra spaces when copying
- Verify the URL is correct (should end with `.supabase.co`)

**"API URL not showing in Railway":**
- Wait for deployment to complete
- Check the "Settings" or "Networking" section
- Look for "Domains" or "Public URL"

**"Variables not saving":**
- Make sure you click "Add" or "Save" after entering values
- Check if there are any validation errors

---

## Need Help?

If you're stuck:
1. Check Railway logs (in the "Deployments" tab)
2. Check Supabase logs (in the "Logs" section)
3. Verify all variables are set correctly
4. Test API health endpoint: `curl https://your-api.railway.app/api/health`

