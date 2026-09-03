import { logout } from '$lib/server/api/auth';
import { AuthorizationError } from '$lib/server/errors';

export const handle = async ({ event, resolve }) => {
  try {
    if (event.route.id?.startsWith('/(authed)')) {
      const authToken = event.cookies.get('authToken');
      if (!authToken) throw new AuthorizationError("No auth token found in cookies");
      const userInfo = event.cookies.get('userInfo');
      event.locals.userInfo = JSON.parse(userInfo ?? '{}');
    }
  } catch (error) {
    if (error instanceof AuthorizationError) {
      await logout(`${event.route.id ?? undefined} (hooks)`);
      event.locals.userInfo = undefined
    }
  }

  return resolve(event);
};