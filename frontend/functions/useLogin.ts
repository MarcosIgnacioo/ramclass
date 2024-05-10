import { useQuery } from '@tanstack/react-query';
import login from '../functions/login.ts'

const useLogin = (loginParams) =>
    useQuery(
        {
            queryKey: ["user", loginParams],
            queryFn: login,
            enabled: (loginParams?.username !== undefined && loginParams?.password !== undefined)
        }
    )

export default useLogin
