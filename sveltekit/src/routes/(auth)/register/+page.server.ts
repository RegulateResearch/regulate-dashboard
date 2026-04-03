import type { Actions, PageServerLoad } from './$types';
import { superValidate } from "sveltekit-superforms";
import { formSchema } from "./schema";
import { zod4 } from "sveltekit-superforms/adapters";
import { fail } from "@sveltejs/kit";

export const load: PageServerLoad = async () => {
  return {
    form: await superValidate(zod4(formSchema)),
  };
};

export const actions: Actions = {
  default: async (event) => {
    const form = await superValidate(event, zod4(formSchema));
    if (!form.valid) {
      return fail(400, {
        form,
      });
    }

    return {
      form,
    };
  },
} satisfies Actions;