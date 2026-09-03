import { dev } from "$app/environment"
import { getRequestEvent } from "$app/server"
import { typedFetch } from "$lib/server/utils"
import * as Schema from "./schema"

const BASE_URL = '/common/my/profile'

export const getMyProfile = async (clientUrl?: string) => {
  const { cookies } = getRequestEvent()
  const response = await typedFetch(
    `${BASE_URL}`,
    Schema.myProfileResponseSchema,
    {
      method: 'GET',
      requireAuthentication: true,
      clientUrl
    })

  if (response.data) {
    cookies.set("userInfo", JSON.stringify(response.data), {
      path: '/',
      httpOnly: true,
      secure: !dev,
      sameSite: 'strict',
      maxAge: 60 * 60 * 24
    })
  }

  return response
}