# Supabase Setup Guide

## Step 1: Create Supabase Project

1. Go to https://supabase.com and sign up/login
2. Click "New Project"
3. Fill in:
   - **Name**: listy (or any name you prefer)
   - **Database Password**: Choose a strong password (save it!)
   - **Region**: Choose closest to you
4. Click "Create new project" (takes 1-2 minutes)

## Step 2: Create the Todos Table

1. In your Supabase project, go to **Table Editor** (left sidebar)
2. Click **"New table"**
3. Name it: `todos`
4. Add these columns:

| Column Name | Type | Default Value | Nullable | Primary Key |
|------------|------|--------------|---------|-------------|
| id | int8 | (auto-increment) | No | Yes |
| item | text | - | No | No |
| done | bool | false | No | No |

5. Click **"Save"**

## Step 3: Enable Row Level Security (RLS) for Public Access

Since we're not using authentication yet, we need to allow public access:

1. Go to **Authentication** > **Policies** (or click on the `todos` table > **Policies**)
2. Click **"New Policy"**
3. Choose **"For full customization"**
4. Policy name: `Allow public access`
5. Allowed operation: Check **ALL** (SELECT, INSERT, UPDATE, DELETE)
6. Policy definition: `true` (allows everything)
7. Click **"Review"** then **"Save policy"**

## Step 4: Get Your Credentials

1. Go to **Settings** (gear icon) > **API**
2. You'll see:
   - **Project URL**: `https://xxxxx.supabase.co`
   - **anon public** key: `eyJhbGc...` (long string)

## Step 5: Configure Your Local Environment

1. Copy `.env.example` to `.env`:
   ```bash
   cp .env.example .env
   ```

2. Edit `.env` and add your credentials:
   ```
   SUPABASE_URL=https://your-project-id.supabase.co
   SUPABASE_KEY=your-anon-key-here
   ```

3. **Important**: Never commit `.env` to Git (it's already in `.gitignore`)

## Step 6: Test It!

```bash
go run todolist.go add "Test todo"
go run todolist.go list
```

You should see your todo! Check your Supabase dashboard > Table Editor > todos to see it in the database.

## Troubleshooting

**Error: "Supabase client not initialized"**
- Check that `.env` file exists and has correct values
- Make sure no extra spaces in `.env` file

**Error: "permission denied"**
- Make sure RLS policies are set up correctly
- Check that you're using the `anon` key, not the `service_role` key

**Error: "relation 'todos' does not exist"**
- Make sure the table name is exactly `todos` (lowercase)
- Check that table was created successfully

## Next Steps

Once this works, you can:
- Add user authentication later
- Add more features (categories, due dates, etc.)
- Build a web interface that uses the same database

