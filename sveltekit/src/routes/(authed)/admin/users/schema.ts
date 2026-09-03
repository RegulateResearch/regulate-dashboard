import { z } from "zod";

export const editSingleUserRoleSchema = z.object({
  id: z.number(),
  role: z.enum(['admin', 'user']),
  academicRole: z.enum(['student', 'lecturer'])
});

export type EditSingleUserRoleData = typeof editSingleUserRoleSchema;

export const editBulkUserRoleSchema = z.object({
  id: z.array(z.number()),
  role: z.enum(['admin', 'user']),
  academicRole: z.enum(['student', 'lecturer'])
});

export type EditBulkUserRoleData = typeof editBulkUserRoleSchema;