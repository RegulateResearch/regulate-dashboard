import { generalAPIResponseSchema, userSchemaWithId } from '$lib/schema';
import { z } from 'zod';

export const getUserByIdResponseSchema = z.object({
  ...generalAPIResponseSchema.shape,
  data: z.optional(
    userSchemaWithId
  ),
});
export type GetUserByIdResponse = z.infer<typeof getUserByIdResponseSchema>;

export const getUsersResponseSchema = z.object({
  ...generalAPIResponseSchema.shape,
  data: z.array(userSchemaWithId).optional()
});
export type GetUsersResponse = z.infer<typeof getUsersResponseSchema>;

export const editRoleSchema = z.object({
  id: z.number(),
  role: z.enum(['admin', 'user']),
  academicRole: z.enum(['student', 'lecturer']).optional()
})

export const bulkUpdateUserRoleRequestSchema = z.array(editRoleSchema);
export type BulkUpdateUsersRequest = z.infer<typeof bulkUpdateUserRoleRequestSchema>;

export const bulkUpdateUserRoleResponseSchema = z.object({
  ...generalAPIResponseSchema.shape,
  data: z.array(editRoleSchema).optional()
});
export type BulkUpdateUsersResponse = z.infer<typeof bulkUpdateUserRoleResponseSchema>;
