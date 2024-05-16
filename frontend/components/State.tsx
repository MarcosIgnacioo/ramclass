import { UseQueryResult } from '@tanstack/react-query'
import React from 'react'
import Loading from './Loading'
import Error from './Error'
import getUser, { storeInLocal } from '../functions/store'
import Empty from './Empty'
import UserData from '../classes/UserData'
import { useUser } from './UserContext'

interface Props {
 fetchedData: UseQueryResult<any, Error>;
 cache: Object[];
 nameSpace: string;
 Container: React.ComponentType<any>;
 returnArray: boolean
}

const State = ({ fetchedData, cache, nameSpace, Container, returnArray }) => {

 let userLocal: UserData | null
 userLocal = (useUser().username == "") ? getUser() : useUser()

 if (!userLocal) return (<h1 className='alert'>No has iniciado sesión</h1>)

 if (fetchedData.isLoading) {

  return <Loading />;
 }

 if (fetchedData.isError) {
  return <Error />;
 }

 if (fetchedData.isSuccess) {
  storeInLocal(fetchedData.data[nameSpace], nameSpace);
  cache = fetchedData.data[nameSpace];
 }

 if (!cache) {
  return <Error />;
 }

 if (cache.length === 0) return <Empty />

 const responseData = cache.map((data, index) => (
  <Container key={index} {...data} />
 ));

 if (!returnArray) return <>{responseData}</>;

 return cache
};

export default State;
