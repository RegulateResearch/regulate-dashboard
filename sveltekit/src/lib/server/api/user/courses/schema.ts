import { courseSchema, courseSchemaWithId, generalAPIResponseSchema } from '$lib/server/schema';
import { z } from 'zod';

export const getCourseByIdResponseSchema = z.object({
  ...generalAPIResponseSchema.shape,
  data: z.optional(
    courseSchemaWithId
  ),
});
export type GetCourseByIdResponse = z.infer<typeof getCourseByIdResponseSchema>;

export const getCoursesResponseSchema = z.object({
  ...generalAPIResponseSchema.shape,
  data: z.optional(
    z.array(courseSchemaWithId)
  ),
});
export type GetCoursesResponse = z.infer<typeof getCoursesResponseSchema>;

export const createCourseRequestSchema = courseSchema;
export type CreateCourseRequest = z.infer<typeof createCourseRequestSchema>;

export const createCourseResponseSchema = z.object({
  ...generalAPIResponseSchema.shape,
  data: z.optional(
    courseSchemaWithId
  ),
});
export type CreateCourseResponse = z.infer<typeof createCourseResponseSchema>;

export const updateCourseRequestSchema = courseSchema;
export type UpdateCourseRequest = z.infer<typeof updateCourseRequestSchema>;

export const updateCourseResponseSchema = z.object({
  ...generalAPIResponseSchema.shape,
  data: z.optional(
    courseSchemaWithId
  ),
});
export type UpdateCourseResponse = z.infer<typeof updateCourseResponseSchema>;

export const deleteCourseResponseSchema = z.object({
  ...generalAPIResponseSchema.shape,
  data: z.optional(
    z.any()
  ),
});
export type DeleteCourseResponse = z.infer<typeof deleteCourseResponseSchema>;