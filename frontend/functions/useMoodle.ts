import { useQuery } from '@tanstack/react-query';
import UserData from '../classes/UserData.ts';
import moodle from './fetchs/moodle.ts';

const useMoodle = (loginParams: UserData | null | undefined) => {
 return useQuery(
  {
   queryKey: ["moodle", loginParams],
   queryFn: moodle,
   enabled: (loginParams !== null),
  }
 )
}

export default useMoodle
