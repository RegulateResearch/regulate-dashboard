import { setContext, getContext } from 'svelte';

type CourseData = {
	name: string,
	year: string,
	term: "odd" | "even" | "short",
	url: string | undefined
}

class CreateCourseModalState {
	isOpen = $state(false);
	rowData: CourseData = $state({ name: '', year: '2026', term: 'odd', url: undefined })

	open(courseData?: CourseData) {
		this.rowData = courseData || { name: '', year: '2026', term: 'odd', url: undefined };
		this.isOpen = true;
	}

	close() {
		this.rowData = { name: '', year: '2026', term: 'odd', url: undefined };
		this.isOpen = false;
	}
}

const CREATE_COURSE_MODAL_KEY = Symbol('create-course-modal-state');

export function setCreateCourseModalState() {
	return setContext(CREATE_COURSE_MODAL_KEY, new CreateCourseModalState());
}

export function getCreateCourseModalState(): CreateCourseModalState {
	return getContext(CREATE_COURSE_MODAL_KEY);
}

class DeleteCourseModalState {
	isOpen = $state(false);
	rowData: { id: number } | null = $state(null);
	isConfirmed = $state(false);

	open(id: number) {
		this.rowData = { id };
		this.isOpen = true;
		this.isConfirmed = false;
	}

	close() {
		this.isOpen = false;
		this.rowData = null;
		this.isConfirmed = false;
	}
}

const DELETE_COURSE_MODAL_KEY = Symbol('delete-course-modal-state');

export function setDeleteCourseModalState() {
	return setContext(DELETE_COURSE_MODAL_KEY, new DeleteCourseModalState());
}

export function getDeleteCourseModalState(): DeleteCourseModalState {
	return getContext(DELETE_COURSE_MODAL_KEY);
}
