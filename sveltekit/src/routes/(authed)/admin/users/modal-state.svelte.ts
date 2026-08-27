import { setContext, getContext } from 'svelte';
import type { UserWithId } from '$lib/schema';


class EditUserRoleModalState {
  isOpen = $state(false);
  rowData: UserWithId = $state({ id: 0, username: '', displayName: '', role: 'user', academicRole: 'student', email: '', civitasId: '' })

  open(userRoleData: UserWithId) {
    this.rowData = userRoleData;
    this.isOpen = true;
  }

  close() {
    this.isOpen = false;
  }
}

const EDIT_ROLE_MODAL_KEY = Symbol('edit-role-modal-state');

export function setEditRoleModalState() {
  return setContext(EDIT_ROLE_MODAL_KEY, new EditUserRoleModalState());
}

export function getEditRoleModalState(): EditUserRoleModalState {
  return getContext(EDIT_ROLE_MODAL_KEY);
}