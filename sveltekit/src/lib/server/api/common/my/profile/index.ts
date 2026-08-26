import { typedFetch } from "$lib/server/utils"
import * as Schema from "./schema"

const BASE_URL = '/common/my/profile'

export const getMyProfile = async () => {
  return await typedFetch(
    `${BASE_URL}`,
    Schema.myProfileResponseSchema,
    {
      method: 'GET',
      requireAuthentication: true
    })
}