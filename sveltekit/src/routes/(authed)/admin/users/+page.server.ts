import type { PageServerLoad } from "./$types";
import { getUsers, updateRoleUser } from "$lib/server/api/admin/user";
import { AuthorizationError } from "$lib/server/errors";
import { fail, redirect } from "@sveltejs/kit";
import { superValidate } from "sveltekit-superforms";
import { editUserRoleSchema } from "./schema";
import { zod4 } from "sveltekit-superforms/adapters";

export const load: PageServerLoad = async () => {
  const breadcrumbs = [
    {
      name: 'Kelola Pengguna',
      href: '/admin/users'
    }
  ]
  const users = await getUsers();

  return {
    breadcrumbs,
    users: users.data,
    form: await superValidate(null, zod4(editUserRoleSchema)),
  };
};

export const actions = {
  editSingleUserRole: async ({ request }) => {
    const form = await superValidate(request, zod4(editUserRoleSchema));
    if (!form.valid) {
      return fail(400, {
        form,
        message: 'Invalid form data'
      });
    }

    const reqBody = [form.data]

    try {
      const res = await updateRoleUser(reqBody)
      if (res.error) {
        return fail(400, {
          form,
          message: `Request failed to update user role: ${res.error}`
        })
      }
      return {
        form,
        message: 'success'
      }
    } catch (error) {
      if (error instanceof AuthorizationError) {
        throw redirect(307, '/logout');
      }
      return fail(400, {
        form,
        message: 'An error occurred while updating the user role'
      })
    }
  },
}