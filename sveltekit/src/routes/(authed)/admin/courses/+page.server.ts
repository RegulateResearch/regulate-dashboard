import { createCourse, deleteCourse, getCourses } from "$lib/server/api/admin/courses";
import { superValidate } from "sveltekit-superforms";
import { newCourseFormSchema } from "./schema";
import { zod4 } from "sveltekit-superforms/adapters";
import { fail } from "@sveltejs/kit";

export const load = async () => {
  const breadcrumbs = [
    {
      name: 'Kelola Kelas',
      href: '/admin/courses'
    }
  ]
  const courses = await getCourses();

  return {
    breadcrumbs,
    courses: courses.data,
    form: await superValidate(zod4(newCourseFormSchema)),
  };
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
    } catch {
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
    } catch {
      return fail(400, {
        message: 'An error occurred while deleting the course'
      })
    }
  }
}