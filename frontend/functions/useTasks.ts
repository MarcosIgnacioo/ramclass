import { useQuery } from '@tanstack/react-query';
import saveTasks from './fetchs/tasks.ts';

const useTasks = (tasks: Object, identifier: string) => {
 return useQuery(
  {
   queryKey: ["tasks", { tasks, identifier }],
   queryFn: saveTasks,
   enabled: (tasks !== null && identifier !== null),
  }
 )
}

export default useTasks
