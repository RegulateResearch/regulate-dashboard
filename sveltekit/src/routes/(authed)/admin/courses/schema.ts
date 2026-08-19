import { z } from "zod";

export const newCourseFormSchema = z.object({
  name: z.string(),
  year: z.string().regex(/^\d+$/), // Ensure it's a string that represents a number
  term: z.enum(['odd', 'even', 'short']),
  url: z.url().optional()
});

export type NewCourseFormSchema = typeof newCourseFormSchema;