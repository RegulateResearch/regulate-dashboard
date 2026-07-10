import { DataTableSortableHeader, renderComponent } from "$lib/components/ui/data-table";
import type { ColumnDef } from "@tanstack/table-core";

export type Course = {
  id: number;
  name: string;
  year: number;
  term: "Ganjil" | "Genap" | "Semester Pendek";
};

export const columns: ColumnDef<Course>[] = [
  {
    accessorKey: "name",
    header: ({ column }) =>
      renderComponent(DataTableSortableHeader, {
        label: "Nama Kelas",
        onclick: column.getToggleSortingHandler(),
      }),
  },
  {
    accessorKey: "year",
    header: ({ column }) =>
      renderComponent(DataTableSortableHeader, {
        label: "Tahun",
        onclick: column.getToggleSortingHandler(),
      }),
  },
  {
    accessorKey: "term",
    header: ({ column }) =>
      renderComponent(DataTableSortableHeader, {
        label: "Semester",
        onclick: column.getToggleSortingHandler(),
      }),
  }
];