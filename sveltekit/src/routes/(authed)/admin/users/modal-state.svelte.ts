import { setContext, getContext } from 'svelte';
import type { User } from './columns';


class EditUserRoleModalState {
  isOpen = $state(false);
  rowData: User = $state({ id: 0, username: '', displayName: '', role: 'user', academicRole: 'student', email: '', civitasId: '' })

  open(userRoleData?: User) {
    this.rowData = userRoleData || { id: 0, username: '', displayName: '', role: 'user', academicRole: 'student', email: '', civitasId: '' };
    this.isOpen = true;
  }

  close() {
    this.rowData = { id: 0, username: '', displayName: '', role: 'user', academicRole: 'student', email: '', civitasId: '' };
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