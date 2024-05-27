import { useQuery } from '@tanstack/react-query';
import getTasks from './fetchs/getTasks';

const useGetTasks = (identifier: string, cache: Object | null) => {
 return useQuery(
  {
   queryKey: ["get-tasks", identifier],
   queryFn: getTasks,
   enabled: (identifier !== null && cache === null),
  }
 )
}

export default useGetTasks
