import { DataTableSortableHeader, DataTableLongText, renderComponent, renderSnippet } from "$lib/components/ui/data-table";
import type { ColumnDef } from "@tanstack/table-core";
import DataTableActions from "./data-table-actions.svelte";
import { createRawSnippet } from "svelte";

export type Course = {
  id: number;
  name: string;
  year: number;
  term: "odd" | "even" | "short";
  url?: string;
};

export const columnsLabel = [
  { id: 'id', label: 'ID' },
  { id: 'name', label: 'Nama Kelas' },
  { id: 'year', label: 'Tahun Akademik' },
  { id: 'term', label: 'Semester' },
  { id: 'action', label: 'Aksi Kelola' },
]

export const columns: ColumnDef<Course>[] = [
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
    accessorKey: "name",
    header: ({ column }) =>
      renderComponent(DataTableSortableHeader, {
        label: "Nama Kelas",
        onclick: column.getToggleSortingHandler(),
        class: "flex space-x-2 w-full"
      }),
    cell: ({ row }) =>
      renderComponent(DataTableLongText, {
        label: row.original.name,
        class: "min-w-125"
      })
  },
  {
    accessorKey: "year",
    header: ({ column }) =>
      renderComponent(DataTableSortableHeader, {
        label: "Tahun",
        onclick: column.getToggleSortingHandler(),
        class: "text-end"
      }),
    cell: ({ row }) => {
      const yearSnippet = createRawSnippet<[{ year: number }]>(
        (getYear) => {
          const { year } = getYear();
          return {
            render: () => `<div class="text-end">${year}/${year + 1}</div>`
          };
        }
      );
      return renderSnippet(yearSnippet, {
        year: row.original.year
      });
    }
  },
  {
    accessorKey: "term",
    header: ({ column }) =>
      renderComponent(DataTableSortableHeader, {
        label: "Semester",
        onclick: column.getToggleSortingHandler(),
        class: "text-center"
      }),
    cell: ({ row }) => {
      const termSnippet = createRawSnippet<[{ term: "odd" | "even" | "short" }]>(
        (getTerm) => {
          const { term } = getTerm();
          let termString: 'Ganjil' | 'Genap' | 'Semester Pendek' = 'Ganjil';
          let termStyle;
          switch (term) {
            case 'odd':
              termString = 'Ganjil';
              termStyle = 'bg-blue-200 text-blue-500'
              break;
            case 'even':
              termString = 'Genap';
              termStyle = 'bg-purple-200 text-purple-500'
              break;
            case 'short':
              termString = 'Semester Pendek';
              termStyle = 'bg-yellow-200 text-yellow-500'
              break;
          }
          return {
            render: () =>
              `<div class="flex justify-center">
                <div class="px-2 py-1 text-center rounded-lg text-xs font-semibold ${termStyle}">${termString}<div>
              </div>`
          };
        }
      );
      return renderSnippet(termSnippet, {
        term: row.original.term
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