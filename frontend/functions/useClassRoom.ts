import { useQuery } from '@tanstack/react-query';
import UserData from '../classes/UserData.ts';
import classRoom from './fetchs/classroom.ts';

const useClassRoom = (loginParams: UserData | null | undefined) => {
 return useQuery(
  {
   queryKey: ["classroom", loginParams],
   queryFn: classRoom,
   enabled: (loginParams !== null),
  }
 )
}

export default useClassRoom
