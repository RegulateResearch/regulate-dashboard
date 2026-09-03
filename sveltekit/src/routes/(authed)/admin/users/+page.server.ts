import { getUsers, updateRoleUser } from "$lib/server/api/admin/user";
import type { BulkUpdateUsersRequest } from "$lib/server/api/admin/user/schema.js";
import { AuthorizationError } from "$lib/server/errors";
import { fail, redirect } from "@sveltejs/kit";
import { superValidate } from "sveltekit-superforms";
import { zod4 } from "sveltekit-superforms/adapters";
import type { PageServerLoad } from "./$types";
import { editBulkUserRoleSchema, editSingleUserRoleSchema } from "./schema";

export const load: PageServerLoad = async ({ route }) => {
  const breadcrumbs = [
    {
      name: 'Kelola Pengguna',
      href: '/admin/users'
    }
  ]
  const users = await getUsers(route.id ?? undefined);

  return {
    breadcrumbs,
    users: users.data,
    editSingleRoleForm: await superValidate(null, zod4(editSingleUserRoleSchema)),
    editBulkRoleForm: await superValidate(null, zod4(editBulkUserRoleSchema)),
  };
};

export const actions = {
  editSingleUserRole: async ({ request, route }) => {
    const form = await superValidate(request, zod4(editSingleUserRoleSchema));
    if (!form.valid) {
      return fail(400, {
        form,
        message: 'Invalid form data'
      });
    }

    const reqBody = [form.data]

    try {
      const res = await updateRoleUser(reqBody, route.id ?? undefined)
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
  editBulkUserRole: async ({ request, route }) => {
    const form = await superValidate(request, zod4(editBulkUserRoleSchema));
    if (!form.valid) {
      return fail(400, {
        form,
        message: 'Invalid form data'
      });
    }

    const data = form.data
    const reqBody: BulkUpdateUsersRequest = [];

    data.id.forEach(id => {
      reqBody.push({ id, role: data.role, academicRole: data.academicRole })
    });

    try {
      const res = await updateRoleUser(reqBody, route.id ?? undefined)
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