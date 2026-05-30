import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = ({ locals }) => {
  const role = locals.userInfo?.role
  switch (role) {
    case 'admin':
      throw redirect(307, '/admin');
    case 'teacher':
      throw redirect(307, '/teacher');
    case 'student':
      throw redirect(307, '/student');
    default:
      throw redirect(307, '/student');
  }
};
