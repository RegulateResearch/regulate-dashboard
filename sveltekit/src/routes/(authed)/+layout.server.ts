import { logout } from '$lib/server/api/auth';
import { getMyProfile } from '$lib/server/api/common/my/profile';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ locals }) => {
  try {
    const session = locals.session
    if (session) {
      if (!locals.userInfo) {
        const userInfoReq = await getMyProfile()
        if (userInfoReq.data) {
          locals.userInfo = userInfoReq.data
        }
      }
      return { userInfo: locals.userInfo }
    }
    return logout()
  } catch {
    return logout()
  }
}