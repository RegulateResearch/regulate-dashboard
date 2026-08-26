import { z } from 'zod';

export const generalAPIResponseSchema = z.object({
  message: z.string(),
  error: z.optional(z.string())
});

export const userSchema = z.object({
  email: z.email().optional(),
  username: z.string(),
  displayName: z.string(),
  role: z.enum(['admin', 'user']),
  academicRole: z.enum(['student', 'lecturer']),
  civitasId: z.string().optional(),
})
export type User = z.infer<typeof userSchema>;

export const userSchemaWithId = z.object({
  ...userSchema.shape,
  id: z.number(),
})
export type UserWithId = z.infer<typeof userSchemaWithId>;

export const courseSchema = z.object({
  name: z.string(),
  year: z.number(),
  term: z.enum(['odd', 'even', 'short']),
  url: z.string().optional()
})
export type Course = z.infer<typeof courseSchema>;

export const courseSchemaWithId = z.object({
  ...courseSchema.shape,
  id: z.number(),
})
export type CourseWithId = z.infer<typeof courseSchemaWithId>;