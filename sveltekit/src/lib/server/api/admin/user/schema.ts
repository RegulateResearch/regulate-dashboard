import { z } from 'zod';
import { generalAPIResponseSchema } from '../../schema';

export const userSchema = z.object({
  email: z.email().optional(),
  username: z.string(),
  displayName: z.string(),
  role: z.enum(['admin', 'user']),
  academicRole: z.enum(['student', 'lecturer']),
  civitasId: z.string().optional(),
})
export type User = z.infer<typeof userSchema>;

export const userSchemaWithId = z.object({
  ...userSchema.shape,
  id: z.number(),
})
export type UserWithId = z.infer<typeof userSchemaWithId>;

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
