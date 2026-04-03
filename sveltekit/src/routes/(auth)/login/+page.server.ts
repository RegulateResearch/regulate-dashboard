import type { Actions, PageServerLoad } from './$types';
import { superValidate } from "sveltekit-superforms";
import { formSchema } from "./schema";
import { zod4 } from "sveltekit-superforms/adapters";
import { fail, redirect } from "@sveltejs/kit";
import { login } from '$lib/server/api/auth';

export const load: PageServerLoad = async ({ url, cookies }) => {
  const authToken = cookies.get('authToken')
  if (authToken) throw redirect(301, url.searchParams.get('redirectTo') || 'after-login')
  return {
    form: await superValidate(zod4(formSchema)),
  };
};

export const actions: Actions = {
  default: async ({ request, url }) => {
    const form = await superValidate(request, zod4(formSchema));
    if (!form.valid) {
      return fail(400, {
        form, message: 'failed'
      });
    }

    try {
      await login(form.data)
    } catch (error) {
      return fail(400, {
        form, message: 'failed'
      });
    }

    const redirectTo = url.searchParams.get('redirectTo') || 'after-login'
    throw redirect(301, redirectTo)
  },
} satisfies Actions;