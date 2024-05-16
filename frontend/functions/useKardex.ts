import { useQuery } from '@tanstack/react-query';
import UserData from '../classes/UserData.ts';
import kardex from './fetchs/kardex.ts';

const useKardex = (loginParams: UserData | null | undefined) => {
 return useQuery(
  {
   queryKey: ["kardex", loginParams],
   queryFn: kardex,
   enabled: (loginParams !== null),
  }
 )
}

export default useKardex
