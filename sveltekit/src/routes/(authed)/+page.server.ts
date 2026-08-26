import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = ({ locals }) => {
  if (locals.userInfo?.role === 'admin') throw redirect(307, '/admin');
  if (locals.userInfo?.academicRole === 'lecturer') throw redirect(307, '/lecturer');
  throw redirect(307, '/student');
};
