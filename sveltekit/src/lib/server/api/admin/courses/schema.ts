import { z } from 'zod';

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

export const getCourseByIdResponseSchema = z.object({
  message: z.string(),
  data: z.optional(
    courseSchemaWithId
  ),
});
export type GetCourseByIdResponse = z.infer<typeof getCourseByIdResponseSchema>;

export const getCoursesResponseSchema = z.object({
  message: z.string(),
  data: z.optional(
    z.array(courseSchemaWithId)
  ),
});
export type GetCoursesResponse = z.infer<typeof getCoursesResponseSchema>;

export const createCourseRequestSchema = courseSchema;
export type CreateCourseRequest = z.infer<typeof createCourseRequestSchema>;

export const createCourseResponseSchema = z.object({
  message: z.string(),
  data: z.optional(
    courseSchemaWithId
  ),
});
export type CreateCourseResponse = z.infer<typeof createCourseResponseSchema>;

export const updateCourseRequestSchema = courseSchema;
export type UpdateCourseRequest = z.infer<typeof updateCourseRequestSchema>;

export const updateCourseResponseSchema = z.object({
  message: z.string(),
  data: z.optional(
    courseSchemaWithId
  ),
});
export type UpdateCourseResponse = z.infer<typeof updateCourseResponseSchema>;

export const deleteCourseResponseSchema = z.object({
  message: z.string(),
  data: z.optional(
    z.any()
  ),
});
export type DeleteCourseResponse = z.infer<typeof deleteCourseResponseSchema>;