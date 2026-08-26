import { generalAPIResponseSchema } from '$lib/server/schema';
import { z } from 'zod';

export const sessionResponseSchema = z.object({
  ...generalAPIResponseSchema.shape,
  data: z.optional(z.object({
    id: z.number(),
    role: z.enum(['admin', 'user']),
  })),
});
export type SessionResponse = z.infer<typeof sessionResponseSchema>;

export const loginRequestSchema = z.object({
  email: z.email(),
  password: z.string(),
});
export type LoginRequest = z.infer<typeof loginRequestSchema>;

export const loginResponseSchema = z.object({
  ...generalAPIResponseSchema.shape,
  data: z.string(),
});
export type LoginResponse = z.infer<typeof loginResponseSchema>;

export const registerRequestSchema = z.object({
  email: z.email(),
  username: z.string(),
  displayName: z.string(),
  password: z.string(),
})
export type RegisterRequest = z.infer<typeof registerRequestSchema>;

export const registerResponseSchema = z.object({
  ...generalAPIResponseSchema.shape,
  data: z.any().optional(),
})
export type RegisterResponse = z.infer<typeof registerResponseSchema>;

export const ssoRequestSchema = z.object({
  ticket: z.string(),
  service: z.string(),
})
export type SSORequest = z.infer<typeof ssoRequestSchema>;

export const ssoResponseSchema = loginResponseSchema
export type SSOResponse = z.infer<typeof ssoResponseSchema>;