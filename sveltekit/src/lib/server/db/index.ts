import { drizzle } from 'drizzle-orm/postgres-js';
import postgres from 'postgres';
import * as schema from './schema';
import { env } from '$env/dynamic/private';


const databaseUrl = `postgres://${env.DB_USERNAME}:${env.DB_PASSWORD}@${env.DB_HOST}:5432/${env.DB_NAME}`

if (!env.DATABASE_URL) throw new Error('DATABASE_URL is not set');

const client = postgres(env.DATABASE_URL);

export const db = drizzle(client, { schema });
