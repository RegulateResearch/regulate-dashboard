import { typedFetch } from "../utils"
import * as Schema from "./schema"

export const getCourses = async () => {
  return await typedFetch(
    '/admin/courses',
    Schema.getCoursesResponseSchema,
    {
      method: 'GET',
      requireAuthentication: true
    })
}

export const createCourse = async (body: Schema.CreateCoursesRequest) => {
  return await typedFetch(
    '/admin/courses',
    Schema.createCoursesResponseSchema,
    {
      method: 'POST',
      body,
      bodySchema: Schema.createCoursesRequestSchema,
      requireAuthentication: true
    })
}