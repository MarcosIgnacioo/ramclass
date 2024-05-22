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

export type Task = {
 task_description: string;
 is_done: boolean;
 is_deleted: boolean;
};

// Define the type for the state
export type TaskList = {
 [day: string]: Task[];
};
