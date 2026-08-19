import { z } from 'zod';

export const generalAPIResponseSchema = z.object({
  message: z.string(),
  error: z.optional(z.string())
});