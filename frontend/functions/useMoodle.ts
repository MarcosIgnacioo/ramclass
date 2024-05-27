import { useQuery } from '@tanstack/react-query';
import UserData from '../classes/UserData';
import moodle from './fetchs/moodle';

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
