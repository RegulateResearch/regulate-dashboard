import { z } from "zod";

export const formSchema = z.object({
  email: z.email(),
  password: z.string()
});

export type FormSchema = typeof formSchema;