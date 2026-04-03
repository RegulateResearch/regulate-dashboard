import type { Actions } from './$types';

export const actions: Actions = {
  default: ({ cookies }) => {
    cookies.delete('authToken', {
      path: '/'
    })
  }
} satisfies Actions;