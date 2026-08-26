import { z } from "zod";

export const editUserRoleSchema = z.object({
  id: z.number(),
  role: z.enum(['admin', 'user']),
  academicRole: z.enum(['student', 'lecturer'])
});

export type EditUserRoleData = typeof editUserRoleSchema;