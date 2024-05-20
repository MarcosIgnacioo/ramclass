import { useQuery } from '@tanstack/react-query';
import getTasks from './fetchs/getTasks.ts';

const useGetTasks = (identifier: string) => {
 return useQuery(
  {
   queryKey: ["get-tasks", identifier],
   queryFn: getTasks,
   enabled: (identifier !== null),
  }
 )
}

export default useGetTasks
