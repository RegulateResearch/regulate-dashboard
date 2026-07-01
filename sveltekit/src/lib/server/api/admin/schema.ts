import { z } from 'zod';

export const getCoursesResponseSchema = z.object({
  message: z.string(),
  data: z.optional(
    z.array(
      z.object({
        id: z.number(),
        name: z.string(),
        year: z.number(),
        term: z.enum(['odd', 'even', 'short']),
        url: z.string().optional()
      }
      )
    )
  ),
});

export type GetCoursesResponse = z.infer<typeof getCoursesResponseSchema>;

export const createCoursesRequestSchema = z.object({
  name: z.string(),
  year: z.number(),
  term: z.enum(['odd', 'even', 'short']),
  url: z.url().optional()
});

export type CreateCoursesRequest = z.infer<typeof createCoursesRequestSchema>;

export const createCoursesResponseSchema = z.object({
  message: z.string(),
  data: z.optional(
    z.object({
      id: z.number(),
      name: z.string(),
      year: z.number(),
      term: z.enum(['odd', 'even', 'short']),
      url: z.string().optional()
    }
    )
  ),
});

export type CreateCoursesResponse = z.infer<typeof createCoursesResponseSchema>;