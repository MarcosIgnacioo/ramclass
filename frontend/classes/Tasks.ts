export default interface Tasks {
 identifier: string
 tasks: TaskClass[]
}

export class TaskClass {
 day: string | null | undefined
 task_description: string | null | undefined
 is_done: boolean
 constructor(day: string | null | undefined, task_description: string | null | undefined, is_done: boolean) {
  this.day = day;
  this.task_description = task_description;
  this.is_done = is_done;
 }
}
