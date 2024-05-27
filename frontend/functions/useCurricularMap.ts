import { useQuery } from '@tanstack/react-query';
import UserData from '../classes/UserData';
import curricularMap from './fetchs/curricularMap';

const useCurricularMap = (loginParams: UserData | null | undefined) => {
 return useQuery(
  {
   queryKey: ["curricular_map", loginParams],
   queryFn: curricularMap,
   enabled: (loginParams !== null),
  }
 )
}

export default useCurricularMap
