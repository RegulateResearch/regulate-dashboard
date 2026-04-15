import { z } from "zod";

const passwordValidation = z
  .string()
  .min(8, { message: "Kata sandi harus minimal 8 karakter" })
  .regex(/[A-Z]/, { message: "Kata sandi harus mengandung setidaknya satu huruf kapital" })
  .regex(/[a-z]/, { message: "Kata sandi harus mengandung setidaknya satu huruf kecil" })
  .regex(/[0-9]/, { message: "Kata sandi harus mengandung setidaknya satu angka" })
  .regex(/[^a-zA-Z0-9]/, { message: "Kata sandi harus mengandung setidaknya satu karakter khusus" })

export const formSchema = z.object({
  email: z.email(),
  username: z
    .string()
    .regex(/^[a-zA-Z0-9._-]+$/, "Username hanya boleh mengandung huruf, angka, titik (.), underscore (_), dan tanda hubung (-)")
    .min(8, { message: "Username harus minimal 8 karakter" })
    .max(20, { message: "Username tidak boleh lebih dari 20 karakter" }),
  displayName: z.string(),
  password: passwordValidation,
  passwordConfirmation: passwordValidation,
}).refine((data) => data.password === data.passwordConfirmation, {
  message: "Periksa kembali kata sandi Anda",
});

export type FormSchema = typeof formSchema;