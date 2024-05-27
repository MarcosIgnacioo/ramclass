import { useQuery } from '@tanstack/react-query';
import { getTerms } from './fetchs/terms';

const useTerms = () => {
 return useQuery(
  {
   queryKey: ["terms", "terms"],
   queryFn: getTerms,
  }
 )
}

export default useTerms
