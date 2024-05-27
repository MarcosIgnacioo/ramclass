import { useQuery } from '@tanstack/react-query';
import UserData from '../classes/UserData';
import classRoom from './fetchs/classroom';

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
