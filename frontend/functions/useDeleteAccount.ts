import { useQuery } from '@tanstack/react-query';
import deleteAccount from './fetchs/deleteAccount';

const useDeleteAccount = (identifier: string | null) => {
 return useQuery(
  {
   queryKey: ["delete-account", identifier],
   queryFn: deleteAccount,
   enabled: (identifier !== ""),
  }
 )
}

export default useDeleteAccount
