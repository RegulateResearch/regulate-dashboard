import { logout } from '$lib/server/api/auth';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ locals }) => {
  try {
    const userInfo = locals.userInfo
    if (userInfo) {
      return { userInfo }
    }
    return logout()
  } catch {
    return logout()
  }
}