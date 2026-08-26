import { DataTableSortableHeader, DataTableLongText, renderComponent, renderSnippet } from "$lib/components/ui/data-table";
import type { ColumnDef } from "@tanstack/table-core";
import DataTableActions from "./data-table-actions.svelte";
import { createRawSnippet } from "svelte";

export type User = {
  id: number;
  username: string;
  displayName: string;
  role: "admin" | "user";
  academicRole: "student" | "lecturer";
  email?: string;
  civitasId?: string;
};

export const columnsLabel = [
  { id: 'id', label: 'id' },
  { id: 'email', label: 'Email' },
  { id: 'username', label: 'Username' },
  { id: 'displayName', label: 'Nama Lengkap' },
  { id: 'role', label: 'Peran (Sistem)' },
  { id: 'academicRole', label: 'Peran (Akademik)' },
  { id: 'civitasId', label: 'ID Civitas' },
]

export const columns: ColumnDef<User>[] = [
  {
    accessorKey: "id",
    header: ({ column }) =>
      renderComponent(DataTableSortableHeader, {
        label: "ID",
        onclick: column.getToggleSortingHandler(),
        class: "text-center w-8"
      }),
    cell: ({ row }) => {
      const idSnippet = createRawSnippet<[{ id: number }]>(
        (getId) => {
          const { id } = getId();
          return {
            render: () => `<div class="text-center text-stone-500">${id}</div>`
          };
        }
      );
      return renderSnippet(idSnippet, {
        id: row.original.id
      });
    }
  },
  {
    accessorKey: "username",
    header: ({ column }) =>
      renderComponent(DataTableSortableHeader, {
        label: "Username",
        onclick: column.getToggleSortingHandler(),
        class: "flex space-x-2 w-full"
      })
  },
  {
    accessorKey: "email",
    header: ({ column }) =>
      renderComponent(DataTableSortableHeader, {
        label: "Email",
        onclick: column.getToggleSortingHandler(),
        class: "flex space-x-2 w-full"
      })
  },
  {
    accessorKey: "civitasId",
    header: ({ column }) =>
      renderComponent(DataTableSortableHeader, {
        label: "ID Civitas",
        onclick: column.getToggleSortingHandler(),
        class: "flex space-x-2 w-full ju"
      }),
    cell: ({ row }) => {
      const civitasIdSnippet = createRawSnippet<[{ civitasId: string }]>(
        (getCivitasId) => {
          const { civitasId } = getCivitasId();
          return {
            render: () => `<div class="text-center text-stone-500">${civitasId}</div>`
          };
        }
      );
      return renderSnippet(civitasIdSnippet, {
        civitasId: row.original.civitasId || "-"
      });
    }
  },
  {
    accessorKey: "displayName",
    header: ({ column }) =>
      renderComponent(DataTableSortableHeader, {
        label: "Nama Lengkap",
        onclick: column.getToggleSortingHandler(),
        class: "flex space-x-2 w-full"
      }),
    cell: ({ row }) =>
      renderComponent(DataTableLongText, {
        label: row.original.displayName,
        class: "min-w-100"
      })
  },
  {
    accessorKey: "role",
    header: ({ column }) =>
      renderComponent(DataTableSortableHeader, {
        label: "Peran (Sistem)",
        onclick: column.getToggleSortingHandler(),
        class: "text-center"
      }),
    cell: ({ row }) => {
      const roleSnippet = createRawSnippet<[{ role: "admin" | "user" }]>(
        (getRole) => {
          const { role } = getRole();
          let roleString: 'Admin' | 'User' = 'User';
          let roleStyle;
          switch (role) {
            case 'admin':
              roleString = 'Admin';
              roleStyle = 'bg-red-200 text-red-500';
              break;
            case 'user':
              roleString = 'User';
              roleStyle = 'bg-blue-200 text-blue-500';
              break;
          }
          return {
            render: () =>
              `<div class="flex justify-center">
                <div class="px-2 py-1 text-center rounded-lg text-xs font-semibold ${roleStyle}">${roleString}<div>
              </div>`
          };
        }
      );
      return renderSnippet(roleSnippet, {
        role: row.original.role
      });
    },
    filterFn: (row, id, value) => {
      return value.includes(row.getValue(id));
    },
  },
  {
    accessorKey: "academicRole",
    header: ({ column }) =>
      renderComponent(DataTableSortableHeader, {
        label: "Peran (Akademik)",
        onclick: column.getToggleSortingHandler(),
        class: "text-center"
      }),
    cell: ({ row }) => {
      const academicRoleSnippet = createRawSnippet<[{ academicRole: "student" | "lecturer" }]>(
        (getAcademicRole) => {
          const { academicRole } = getAcademicRole();
          let academicRoleString: 'Mahasiswa' | 'Dosen' = 'Mahasiswa';
          let academicRoleStyle;
          switch (academicRole) {
            case 'student':
              academicRoleString = 'Mahasiswa';
              academicRoleStyle = 'bg-yellow-200 text-yellow-500';
              break;
            case 'lecturer':
              academicRoleString = 'Dosen';
              academicRoleStyle = 'bg-purple-200 text-purple-500';
              break;
          }
          return {
            render: () =>
              `<div class="flex justify-center">
                <div class="px-2 py-1 text-center rounded-lg text-xs font-semibold ${academicRoleStyle}">${academicRoleString}<div>
              </div>`
          };
        }
      );
      return renderSnippet(academicRoleSnippet, {
        academicRole: row.original.academicRole
      });
    },
    filterFn: (row, id, value) => {
      return value.includes(row.getValue(id));
    },
  },
  {
    accessorKey: "action",
    header: () => {
      const actionSnippet = createRawSnippet<[{ label: string }]>(
        (getLabel) => {
          const { label } = getLabel();
          return {
            render: () => `<div class="text-center text-xs">${label}</div>`
          };
        }
      );
      return renderSnippet(actionSnippet, {
        label: "Aksi Kelola"
      });
    },
    cell: ({ row }) => {
      return renderComponent(DataTableActions, { data: row.original });
    },
  }
];