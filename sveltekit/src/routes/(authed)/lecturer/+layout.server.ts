import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = ({ locals }) => {
  if (locals.userInfo?.role === 'admin') throw redirect(307, '/admin');
  if (locals.userInfo?.academicRole === 'student') throw redirect(307, '/student');
};
