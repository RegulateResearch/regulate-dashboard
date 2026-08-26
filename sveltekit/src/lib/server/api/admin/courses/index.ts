import { typedFetch } from "$lib/server/utils"
import * as Schema from "./schema"

const BASE_URL = '/admin/courses'

export const getCourses = async () => {
  return await typedFetch(
    `${BASE_URL}`,
    Schema.getCoursesResponseSchema,
    {
      method: 'GET',
      requireAuthentication: true
    })
}

export const getCourseById = async (id: number) => {
  return await typedFetch(
    `${BASE_URL}/${id}`,
    Schema.getCourseByIdResponseSchema,
    {
      method: 'GET',
      requireAuthentication: true
    })
}

export const createCourse = async (body: Schema.CreateCourseRequest) => {
  return await typedFetch(
    `${BASE_URL}`,
    Schema.createCourseResponseSchema,
    {
      method: 'POST',
      body,
      bodySchema: Schema.createCourseRequestSchema,
      requireAuthentication: true
    })
}

export const updateCourse = async (id: number, body: Schema.UpdateCourseRequest) => {
  return await typedFetch(
    `${BASE_URL}/${id}`,
    Schema.updateCourseResponseSchema,
    {
      method: 'PUT',
      body,
      bodySchema: Schema.updateCourseRequestSchema,
      requireAuthentication: true
    })
}

export const deleteCourse = async (id: number) => {
  return await typedFetch(
    `${BASE_URL}/${id}`,
    Schema.deleteCourseResponseSchema,
    {
      method: 'DELETE',
      requireAuthentication: true
    })
}