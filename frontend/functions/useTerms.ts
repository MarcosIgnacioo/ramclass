import { useQuery } from '@tanstack/react-query';
import { getTerms } from './fetchs/terms.ts';

const useTerms = () => {
 return useQuery(
  {
   queryKey: ["terms"],
   queryFn: getTerms,
  }
 )
}

export default useTerms
