import { typedFetch } from "$lib/server/utils"
import * as Schema from "./schema"

const BASE_URL = '/admin/users'

export const getUsers = async (clientUrl?: string) => {
  return await typedFetch(
    `${BASE_URL}`,
    Schema.getUsersResponseSchema,
    {
      method: 'GET',
      requireAuthentication: true,
      clientUrl
    })
}

export const getUserById = async (id: number, clientUrl?: string) => {
  return await typedFetch(
    `${BASE_URL}/${id}`,
    Schema.getUserByIdResponseSchema,
    {
      method: 'GET',
      requireAuthentication: true,
      clientUrl
    })
}

export const updateRoleUser = async (body: Schema.BulkUpdateUsersRequest, clientUrl?: string) => {
  return await typedFetch(
    `${BASE_URL}`,
    Schema.bulkUpdateUserRoleResponseSchema,
    {
      method: 'PUT',
      body,
      bodySchema: Schema.bulkUpdateUserRoleRequestSchema,
      requireAuthentication: true,
      clientUrl
    })
}