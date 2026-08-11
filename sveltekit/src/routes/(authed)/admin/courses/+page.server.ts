import { createCourse, getCourses } from "$lib/server/api/admin";
import type { Actions, PageServerLoad } from "./$types";
import { superValidate } from "sveltekit-superforms";
import { formSchema } from "./schema";
import { zod4 } from "sveltekit-superforms/adapters";

export const load: PageServerLoad = async () => {
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
    form: await superValidate(zod4(formSchema)),
  };
};

export const actions = {
  default: async ({ request }) => {
    console.log("Attempting to create course with form data:");
    const form = await superValidate(request, zod4(formSchema));
    if (!form.valid) {
      return {
        form,
        message: 'failed'
      }
    }

    const reqBody = {
      name: form.data.name,
      year: parseInt(form.data.year),
      term: form.data.term
    }

    try {
      const res = await createCourse(reqBody)
      if (!res.data) {
        return {
          form,
          message: 'failed'
        }
      }
      return {
        form,
        message: 'success'
      }
    } catch {
      console.log("FAILED to create course:");
      return {
        form,
        message: 'failed'
      }
    }
  }
} satisfies Actions;