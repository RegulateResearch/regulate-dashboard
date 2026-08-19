import { typedFetch } from "../../utils"
import * as Schema from "./schema"

const BASE_URL = '/admin/users'

export const getUsers = async () => {
  return await typedFetch(
    `${BASE_URL}`,
    Schema.getUsersResponseSchema,
    {
      method: 'GET',
      requireAuthentication: true
    })
}

export const getUserById = async (id: number) => {
  return await typedFetch(
    `${BASE_URL}/${id}`,
    Schema.getUserByIdResponseSchema,
    {
      method: 'GET',
      requireAuthentication: true
    })
}

export const updateRoleUser = async (id: number, body: Schema.BulkUpdateUsersRequest) => {
  return await typedFetch(
    `${BASE_URL}/${id}`,
    Schema.bulkUpdateUserRoleResponseSchema,
    {
      method: 'PUT',
      body,
      bodySchema: Schema.bulkUpdateUserRoleRequestSchema,
      requireAuthentication: true
    })
}