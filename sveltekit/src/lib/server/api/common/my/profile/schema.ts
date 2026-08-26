import { generalAPIResponseSchema, userSchemaWithId } from '$lib/server/schema';
import { z } from 'zod';

export const myProfileResponseSchema = z.object({
  ...generalAPIResponseSchema.shape,
  data: z.optional(
    userSchemaWithId
  ),
});
export type MyProfileResponse = z.infer<typeof myProfileResponseSchema>;