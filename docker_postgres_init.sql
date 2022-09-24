SELECT 'CREATE DATABASE auth_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'auth_db')\gexec

SELECT 'CREATE DATABASE journal_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'journal_db')\gexec

SELECT 'CREATE DATABASE discussions_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'discussions_db')\gexec