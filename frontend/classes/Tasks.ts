export default interface Tasks {
 identifier: string
 tasks: TaskClass[]
}

export class TaskClass {
 task_description: string | null | undefined
 is_done: boolean
 is_deleted: boolean
 constructor(task_description: string | null | undefined, is_done: boolean) {
  this.task_description = task_description;
  this.is_done = is_done;
  this.is_deleted = false;
 }
}

export interface TaskDrag {
 day: string
 index: number
}
