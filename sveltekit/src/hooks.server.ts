import { getSession } from '$lib/server/api/auth';

export const handle = async ({ event, resolve }) => {
  try {
    const res = await getSession();
    if (res.data && res.message === 'token is valid') {
      event.locals.userInfo = { ...res.data }
    }
  } catch {
    event.locals.userInfo = undefined
  }

  return resolve(event);
};