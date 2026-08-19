import { getCourseById, updateCourse } from "$lib/server/api/admin/courses";
import type { PageServerLoad } from "./$types";
import { superValidate } from "sveltekit-superforms";
import { editCourseFormSchema } from "./schema";
import { zod4 } from "sveltekit-superforms/adapters";
import { fail, redirect } from "@sveltejs/kit";
import { AuthorizationError } from "$lib/server/api/errors";

export const load: PageServerLoad = async ({ params }) => {
  const breadcrumbs = [
    {
      name: 'Kelola Kelas',
      href: '/admin/courses'
    },
    {
      name: 'Kelola Detail Kelas',
      href: `/admin/courses/${params.id}`
    },
  ]

  try {
    const course = await getCourseById(parseInt(params.id));
    if (!course.data || course.error) {
      return {
        breadcrumbs,
        course: null,
        form: await superValidate(null, zod4(editCourseFormSchema)),
        message: 'Course not found'
      };
    }
    const form = await superValidate(course.data ? {
      name: course.data.name,
      year: course.data.year.toString(),
      term: course.data.term,
      url: course.data.url
    } : null, zod4(editCourseFormSchema))
    return {
      breadcrumbs,
      course: course.data,
      form: form
    };
  } catch (error) {
    if (error instanceof AuthorizationError) {
      throw redirect(307, '/logout');
    }
    return {
      breadcrumbs,
      course: null,
      form: await superValidate(null, zod4(editCourseFormSchema)),
      message: 'An error occurred while fetching the course'
    };
  }
};

export const actions = {
  updateCourse: async ({ request, params }) => {
    const form = await superValidate(request, zod4(editCourseFormSchema));
    if (!form.valid) {
      return fail(400, {
        form,
        message: 'Invalid form data'
      });
    }

    const reqBody = {
      ...form.data,
      year: parseInt(form.data.year),
    }

    try {
      const res = await updateCourse(parseInt(params.id), reqBody)
      if (res.error) {
        return fail(400, {
          form,
          message: `Request failed to update course: ${res.error}`
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
        message: 'An error occurred while updating the course'
      })
    }
  },
}