import { useQuery } from '@tanstack/react-query';
import saveTasks from '../functions/fetchs/tasks';

const useTasks = (tasks: Object | undefined, identifier: string | null) => {
 return useQuery(
  {
   queryKey: ["tasks", { tasks, identifier }],
   queryFn: saveTasks,
   enabled: (tasks !== undefined && identifier !== null),
  }
 )
}

export default useTasks
