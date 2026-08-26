import { createCourse, deleteCourse, getCourses } from "$lib/server/api/admin/courses";
import { AuthorizationError } from "$lib/server/errors";
import { fail, redirect } from "@sveltejs/kit";
import { superValidate } from "sveltekit-superforms";
import { zod4 } from "sveltekit-superforms/adapters";
import { newCourseFormSchema } from "./schema";


export const load = async () => {
  const breadcrumbs = [
    {
      name: 'Kelola Kelas',
      href: '/admin/courses'
    }
  ]
  try {
    const courses = await getCourses();
    if (!courses.data || courses.error) {
      return {
        breadcrumbs,
        courses: null,
        form: await superValidate(null, zod4(newCourseFormSchema)),
        message: 'Courses not found'
      };
    }
    return {
      breadcrumbs,
      courses: courses.data,
      form: await superValidate(zod4(newCourseFormSchema)),
    };
  } catch (error) {
    if (error instanceof AuthorizationError) {
      throw redirect(307, '/logout');
    }
    return {
      breadcrumbs,
      courses: null,
      form: await superValidate(null, zod4(newCourseFormSchema)),
      message: 'An error occurred while fetching the courses'
    };
  }
};

export const actions = {
  createCourse: async ({ request }) => {
    const form = await superValidate(request, zod4(newCourseFormSchema));
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
      const res = await createCourse(reqBody)
      if (res.error) {
        return fail(400, {
          form,
          message: `Request failed to create course: ${res.error}`
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
        message: 'An error occurred while creating the course'
      })
    }
  },
  deleteCourse: async ({ request }) => {
    const form = await request.formData();
    const id = await form.get('id');

    if (id === null) {
      return fail(404, {
        message: 'Course not found'
      })
    }

    const parsedId = parseInt(id.toString());

    try {
      const res = await deleteCourse(parsedId)
      if (res.error) {
        return fail(400, {
          message: `Failed to delete course: ${res.error}`
        })
      }
      return {
        message: 'Delete course success'
      }
    } catch (error) {
      if (error instanceof AuthorizationError) {
        throw redirect(307, '/logout');
      }
      return fail(400, {
        message: 'An error occurred while deleting the course'
      })
    }
  }
}